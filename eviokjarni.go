package evio

import "net"

func (c *stdconn) GetConn() net.Conn { return c.conn }
