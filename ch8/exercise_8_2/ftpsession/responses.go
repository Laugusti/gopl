package ftpsession

const (
	startingTransfer = "125 Transfer starting"
	commandOk        = "200 Command okay"
	systemType       = "215 UNIX Type: L8"
	readyForUser     = "220 Service ready for new user"
	closingCtrlConn  = "221 Service closing TELNET connection"
	closingDataConn  = "226 Closing data connection - requested file action successful"
	userLoggedIn     = "230 User logged in, proceed"
	cwdSuccess       = "250 The current directory has been changed to "
	cantOpenDataConn = "Can't open data connection"
	notImplemented   = "502 Command not implemented"
	invalidProtocol  = "522 Extended Port Failure - unknown network protocol"
	userNotLoggedIn  = "530 Not logged in"
	actionNotTaken   = "550 Requested action not taken"
)
