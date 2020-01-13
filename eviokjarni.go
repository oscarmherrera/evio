package evio

import (
	"errors"
	"syscall"
)

func (c *stdconn) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 0 {
		nn, err := c.conn.Write(p)
		if err != nil {
			return n, err
		}
		p = p[nn:]
	}
	return n, nil
}
func (c *stdudpconn) Write(p []byte) (n int, err error) {
	return 0, errors.New("sending through udp not supported")
}

func (c *conn) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 0 {
		nn, err := syscall.Write(c.fd, p)
		if err != nil {
			return n, err
		}
		p = p[nn:]
	}
	return n, nil
}
