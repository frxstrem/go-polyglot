package polyglot

import (
	"crypto/tls"
	"net"
)

type TLSPolyglot struct {
	net.Listener
	tlsConfig *tls.Config
}

func NewTLSPolyglot(listener net.Listener, tlsConfig *tls.Config) net.Listener {
  l := new(TLSPolyglot)
  l.Listener = listener
  l.tlsConfig = tlsConfig
  return l
}

func (pg *TLSPolyglot) Accept() (net.Conn, error) {
	conn, err := pg.Listener.Accept()
	if err != nil {
		return nil, err
	}

	buffConn := NewBufferedConn(conn)

	p, err := buffConn.Peek(1)
	if err != nil {
		return nil, err
	}

	if p[0] == 22 {
		return tls.Server(buffConn, pg.tlsConfig), nil
	} else {
		return buffConn, nil
	}
}
