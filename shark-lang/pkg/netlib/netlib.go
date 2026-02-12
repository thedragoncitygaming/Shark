// netlib.go

// Package netlib provides high-level networking functionalities supporting TCP and UDP socket operations.

package netlib

import (
	"net"	
	"fmt"
	"errors"
)

// TCPClient represents a TCP client that connects to a given address.
type TCPClient struct {
	conn *net.TCPConn
}

// NewTCPClient creates a new TCP client.
func NewTCPClient(address string) (*TCPClient, error) {
	// Resolve TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}

	// Create TCP connection
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return &TCPClient{conn: conn}, nil
}

// Send sends data over the TCP connection.
func (c *TCPClient) Send(data []byte) (int, error) {
	return c.conn.Write(data)
}

// Close closes the TCP connection.
func (c *TCPClient) Close() error {
	return c.conn.Close()
}

// UDPClient represents a UDP client.
type UDPClient struct {
	conn *net.UDPConn
}

// NewUDPClient creates a new UDP client.
func NewUDPClient(address string) (*UDPClient, error) {
	// Resolve UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil, err
	}
	return &UDPClient{conn: conn}, nil
}

// Send sends data over the UDP connection.
func (c *UDPClient) Send(data []byte) (int, error) {
	return c.conn.Write(data)
}

// Close closes the UDP connection.
func (c *UDPClient) Close() error {
	return c.conn.Close()
}