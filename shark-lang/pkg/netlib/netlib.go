package netlib

import (
	"net"
	"time"
)

type TCPServer struct {
	listener net.Listener
}

type UDPServer struct {
	conn *net.UDPConn
}

func NewTCPServer(port string) (*TCPServer, error) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	return &TCPServer{listener: listener}, nil
}

func (s *TCPServer) Accept() (net.Conn, error) {
	return s.listener.Accept()
}

func (s *TCPServer) Close() error {
	return s.listener.Close()
}

func NewUDPServer(port string) (*UDPServer, error) {
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	return &UDPServer{conn: conn}, nil
}

func (s *UDPServer) Read(b []byte) (int, net.Addr, error) {
	n, addr, err := s.conn.ReadFromUDP(b)
	return n, addr, err
}

func (s *UDPServer) Close() error {
	return s.conn.Close()
}