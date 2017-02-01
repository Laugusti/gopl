package ftpsession

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/Laugusti/gopl/ch8/exercise_8_2/ftpsession/ftpdirectory"
)

type FTPSession struct {
	controlConn   net.Conn
	dataConn      *dataConnection
	authenticated bool
	cwd           ftpdirectory.Dir
}

func NewFTPSession(conn net.Conn) *FTPSession {
	return &FTPSession{controlConn: conn}
}

type dataConnection struct {
	family, addr, port string
	conn               net.Conn
}

var quit = errors.New("close connection")

// HandleConn handles the FTPSession
func (s *FTPSession) HandleConn() {
	// closes the control and data connection when HandleConn completes
	defer func() {
		s.controlConn.Close()
		if s.dataConn.conn != nil {
			s.dataConn.conn.Close()
		}
	}()

	err := s.writeControlResponse(readyForUser)
	if err != nil {
		return // e.g., client disconnected
	}

	// process commands from the data connection until the QUIT command
	// is received
	for {
		switch err := s.processAll(); err {
		case nil:
			continue
		case quit:
			return
		default:
			// log error and kill connection
			log.Println(err)
			return
		}
	}
}

// processAll reads from the control connection and handles each ftp command
// appropriately
func (s *FTPSession) processAll() error {
	command, args, err := s.getCommandWithArgs()
	if err != nil {
		if err == io.EOF {
			return quit
		}
		return err
	}

	// handle each ftp command
	switch command {
	case "USER":
		return s.processUser(args)
	case "SYST":
		return s.writeControlResponse(systemType)
	case "TYPE":
		return s.writeControlResponse(commandOk)
	case "EPRT":
		return s.processEprt(args)
	case "CWD":
		return s.processCwd(args)
	case "LIST":
		return s.processList(args)
	case "RETR":
		return s.processRetr(args)
	case "STOR":
		return s.processStor(args)
	case "QUIT":
		return s.processQuit()
	default:
		return s.writeControlResponse(notImplemented)
	}
	return nil
}

// checkAuthentication checks if the session is authenticated. If the session is not
// authenticated, it writes response 530 to the control connection
func (s *FTPSession) checkAuthentication() (isAuthenticated bool, err error) {
	if s.authenticated {
		isAuthenticated = true
	} else {
		err = s.writeControlResponse(userNotLoggedIn)
	}
	return
}

// processUser performs the USER ftp command
func (s *FTPSession) processUser(user string) error {
	// mark the session as unauthenticated
	s.authenticated = false

	// handle each user, default is unauthenticated
	switch user {
	case "anonymous":
		err := s.writeControlResponse(userLoggedIn)
		if err != nil {
			return err
		}
		s.authenticated = true
	default:
		err := s.writeControlResponse(userNotLoggedIn)
		if err != nil {
			return err
		}
	}
	return nil
}

// processEprt performs the EPRT ftp command
func (s *FTPSession) processEprt(args string) error {
	// check authentication
	authorized, err := s.checkAuthentication()
	if err != nil || !authorized {
		return err
	}

	var family, addr, port string
	args = strings.Replace(args, "|", " ", -1)
	fmt.Sscanf(args, "%s%s%s", &family, &addr, &port)

	// only supports tcp4 or tcp6
	if family != "1" && family != "2" {
		err := s.writeControlResponse(invalidProtocol)
		return err
	}
	laddr := net.JoinHostPort(addr, port)
	conn, err := net.Dial("tcp", laddr)
	if err != nil {
		return err
	}
	err = s.writeControlResponse(commandOk)
	if err != nil {
		return err
	}
	s.dataConn = &dataConnection{family, addr, port, conn}
	return nil
}

// processCwd performs the CWD ftp command
func (s *FTPSession) processCwd(args string) error {
	// check authentication
	authorized, err := s.checkAuthentication()
	if err != nil || !authorized {
		return err
	}

	// try to change the working directory
	err = s.cwd.ChangeDirectory(args)
	// cwd failed
	if err != nil {
		err := s.writeControlResponse(actionNotTaken)
		return err
	}

	// cwd successful
	return s.writeControlResponse(cwdSuccess + s.cwd.PrintWorkingDirectory())
}

// processList performs the LIST ftp command
func (s *FTPSession) processList(args string) error {
	// check authentication
	authorized, err := s.checkAuthentication()
	if err != nil || !authorized {
		return err
	}

	// try to get ls args
	files, err := s.cwd.ListFiles(args)
	// ls failed
	if err != nil {
		err := s.writeControlResponse(actionNotTaken)
		return err
	}

	return s.writeToDataConn(func(c net.Conn) error {
		// write filenames to data connection
		for _, f := range files {
			_, err := io.WriteString(c, f+"\r\n")
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// processRetr performs the RETR ftp command
func (s *FTPSession) processRetr(args string) error {
	// check authentication
	authorized, err := s.checkAuthentication()
	if err != nil || !authorized {
		return err
	}

	b, err := s.cwd.GetFile(args)
	if err != nil {
		err := s.writeControlResponse(actionNotTaken)
		return err
	}

	return s.writeToDataConn(func(c net.Conn) error {
		// write bytes to data connection
		_, err = c.Write(b)
		return err
	})
}

//processStor performs the STOR ftp command
func (s *FTPSession) processStor(args string) error {
	// check authentication
	authorized, err := s.checkAuthentication()
	if err != nil || !authorized {
		return err
	}

	// TODO: check for invalid directory error
	return s.readFromDataConn(func(c net.Conn) error {
		return s.cwd.StoreFile(args, c)
	})
}

// processQuit writes the quit message to control connection and returns the quit error
func (s *FTPSession) processQuit() error {
	err := s.writeControlResponse(closingCtrlConn)
	if err != nil {
		return err
	}
	return quit
}

// writeControlResponse writes the response to the control connection
func (s *FTPSession) writeControlResponse(response string) error {
	_, err := io.WriteString(s.controlConn, response+"\r\n")
	return err
}

// writeToDataConn calls the write function on the data connection
func (s *FTPSession) writeToDataConn(write func(net.Conn) error) error {
	// open the connection, if it is closed
	err := s.dataConn.createConnection()
	if err != nil {
		return s.writeControlResponse(cantOpenDataConn)
	}

	// start data transfer
	err = s.writeControlResponse(startingTransfer)
	if err != nil {
		return err
	}

	// write to data connection
	err = write(s.dataConn.conn)
	if err != nil {
		return err
	}

	// close data connection
	s.dataConn.conn.Close()
	s.dataConn.conn = nil
	err = s.writeControlResponse(closingDataConn)
	return err
}

// readFromDataConn calls the read function on the data connection
func (s *FTPSession) readFromDataConn(read func(net.Conn) error) error {
	// open the connection, if it is closed
	err := s.dataConn.createConnection()
	if err != nil {
		return s.writeControlResponse(cantOpenDataConn)
	}

	// start data transfer
	err = s.writeControlResponse(startingTransfer)
	if err != nil {
		return err
	}

	//read from data connection
	err = read(s.dataConn.conn)
	if err != nil {
		return err
	}

	// close data connection
	s.dataConn.conn.Close()
	s.dataConn.conn = nil
	err = s.writeControlResponse(closingDataConn)
	return err
}

// createConnection creates the data connection
func (dc *dataConnection) createConnection() error {
	if dc.conn == nil {
		conn, err := net.Dial("tcp", net.JoinHostPort(dc.addr, dc.port))
		if err != nil {
			return err
		}
		dc.conn = conn
	}
	return nil
}

// getCommandWithArgs reads the command from the control connection and splits
// it into it's command and argument components
func (s *FTPSession) getCommandWithArgs() (string, string, error) {
	r := bufio.NewReader(s.controlConn)
	str, err := r.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	// debug - output command to stdout
	log.Print(str)

	str = strings.TrimSpace(str)
	index := strings.Index(str, " ")
	if index < 0 {
		return str, "", nil
	}
	return str[:index], str[index+1:], nil
}
