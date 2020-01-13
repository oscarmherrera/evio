package evio

import "net"

func (c *stdconn) GetConn() net.Conn    { return c.conn }
func (c *stdudpconn) GetConn() net.Conn { return nil }
func (c *conn) GetConn() net.Conn       { return nil }
