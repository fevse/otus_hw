package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type tClient struct {
	address string
	timeout time.Duration
	conn    net.Conn
	in      io.ReadCloser
	out     io.Writer
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	var conn net.Conn
	return &tClient{address, timeout, conn, in, out}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.

func (tC *tClient) Connect() error {
	conn, err := net.DialTimeout("tcp", tC.address, tC.timeout)
	if err != nil {
		return err
	}
	tC.conn = conn
	return nil
}

func (tC *tClient) Close() error {
	if err := tC.conn.Close(); err != nil {
		return err
	}
	return nil
}

func (tC *tClient) Send() error {
	if _, err := io.Copy(tC.conn, tC.in); err != nil {
		return err
	}
	return nil
}

func (tC *tClient) Receive() error {
	if _, err := io.Copy(tC.out, tC.conn); err != nil {
		return err
	}
	return nil
}
