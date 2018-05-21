package polyglot

import (
	"bufio"
	"net"
)

type BufferedConn struct {
	reader *bufio.Reader
	net.Conn
}

func NewBufferedConn(conn net.Conn) BufferedConn {
	return BufferedConn{bufio.NewReader(conn), conn}
}

func NewBufferedConnSize(conn net.Conn, size int) BufferedConn {
	return BufferedConn{bufio.NewReaderSize(conn, size), conn}
}

func (conn BufferedConn) Peek(n int) ([]byte, error) {
	return conn.reader.Peek(n)
}

func (conn BufferedConn) Read(p []byte) (int, error) {
	return conn.reader.Read(p)
}
