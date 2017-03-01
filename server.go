package jata

import (
	"bufio"
	"fmt"
	"net"
)

// TCPServer represents a TCP Server
type TCPServer struct {
	Listener *net.Listener
}

// StartListeningOnPort starts the server
func (s *TCPServer) StartListeningOnPort(port int) (err error) {
	defer s.StopListening()
	portString := fmt.Sprintf(":%d", port)
	ln, err1 := net.Listen("tcp", portString)
	if err1 != nil {
		Error.Printf("StartListeningOnPort(%d) terminated due to %v", port, err1)
		return err1
	}
	s.Listener = &ln
	for {
		conn, err2 := (*s.Listener).Accept()
		if err2 == nil {
			go HandleConnection(s, conn)
		} else {
			Error.Printf("StartListeningOnPort(%d) terminated due to %v", port, err2)
			return err2
		}
	}

}

// StopListening closes the Listerner
func (s *TCPServer) StopListening() (err error) {
	return (*s.Listener).Close()
}

// HandleConnection handles all connection logic
func HandleConnection(s *TCPServer, c net.Conn) {
	defer c.Close()
	reader := bufio.NewReaderSize(c, 8192)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			Error.Printf("HandleConnection() terminated due to %v", err)
			s.StopListening()
			return
		}
		fmt.Print(line)
	}
}
