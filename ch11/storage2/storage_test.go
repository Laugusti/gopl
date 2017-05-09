package storage

import "strings"
import "testing"

func TestCheckQuotaNotifiesUser(t *testing.T) {
	// Save and restore original notifyUser.
	saved := notifyUser
	defer func() { notifyUser = saved }()

	//Install the test's fake notifyUser.
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	const user = "joe@example.org"
	usage[user] = 980000000 // simulate a 980MB-used condition

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s",
			notifyUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification mesage <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}
