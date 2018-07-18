package params

import (
	"net/http"
	"net/url"
	"testing"
)

func createRequestWithQuery(value string) *http.Request {
	u, _ := url.Parse("http://example.go")
	query := u.Query()
	query.Add("value", value)
	u.RawQuery = query.Encode()
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	return req
}

func TestValidNumber(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{"", true},
		{"a", true},
		{"-1", true},
		{"0", false},
		{"1", false},
	}

	for _, test := range tests {
		var v struct {
			Value string `http:"value,number"`
		}
		err := Unpack(createRequestWithQuery(test.input), &v)
		if err != nil && !test.expectError {
			t.Errorf("TestValidNumber -> failed to unpack: %q", test.input)
		}
		if test.expectError && err == nil {
			t.Errorf("TestValidNumber -> expected error, got nil: %q", test.input)
		}
	}
}

func TestValidVisa(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{"", true},
		{"a", true},
		{"0", true},
		{"1234-1234-1234-1234", false},
	}

	for _, test := range tests {
		var v struct {
			Value string `http:"value,visa"`
		}
		err := Unpack(createRequestWithQuery(test.input), &v)
		if err != nil && !test.expectError {
			t.Errorf("TestValidVisa -> failed to unpack: %q", test.input)
		}
		if test.expectError && err == nil {
			t.Errorf("TestValidVisa -> expected error, got nil: %q", test.input)
		}
	}
}

func TestValidEmailAddress(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{"", true},
		{"a", true},
		{"0", true},
		{"test@test.com", false},
	}

	for _, test := range tests {
		var v struct {
			Value string `http:"value,email"`
		}
		err := Unpack(createRequestWithQuery(test.input), &v)
		if err != nil && !test.expectError {
			t.Errorf("TestValidEmailAddress -> failed to unpack: %q", test.input)
		}
		if test.expectError && err == nil {
			t.Errorf("TestValidEmailAddress -> expected error, got nil: %q", test.input)
		}
	}
}

func TestValidZipCode(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{"", true},
		{"a", true},
		{"0", true},
		{"12345", false},
		{"12345-6789", false},
		{"123456789", true},
		{"12345-67890", true},
	}

	for _, test := range tests {
		var v struct {
			Value string `http:"value,zipcode"`
		}
		err := Unpack(createRequestWithQuery(test.input), &v)
		if err != nil && !test.expectError {
			t.Errorf("TestValidZipCode -> failed to unpack: %q", test.input)
		}
		if test.expectError && err == nil {
			t.Errorf("TestValidZipCode -> expected error, got nil: %q", test.input)
		}
	}
}
