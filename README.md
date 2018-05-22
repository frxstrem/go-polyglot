# go-polyglot

`go-polyglot` is inspired by the [httpolyglot](https://github.com/mscdex/httpolyglot) Node.js module.

The polyglot listener uses the first byte sent by the client to detect a TLS connection: if the first byte is `22`, then the connection is identified as TLS, otherwise it's identified as an unencrypted connection.

This works well for HTTP/HTTPS, and may work for other protocols as well, but may not work well if the client sends a first byte that cannot be distinguished for a TLS handshake, or if the client does not initiate the handshake at all.

**Note!** This package is experimental and indented for development use only. In production settings, you should use dedicated ports for TLS and non-TLS traffic.

## Usage

Use `polyglot.NewTLSPolyglot(net.Listener, *tls.Config)` to create a new polyglot listener.

If an incoming connection is identified as TLS, it will return a TLS server connection (as if wrapped with `tls.Server(net.Conn, *tls.Config)`; otherwise it will return the connection itself.

## HTTP/HTTPS example

This example shows how to use `go-polyglot` to host HTTP and HTTPS on the same port.

```go
package main

import (
  "crypto/tls"
  "fmt"
  "net"
  "net/http"

  polyglot "github.com/frxstrem/go-polyglot"
)

func main() {
  bindAddress := ":8888"
  certFile := "ssl/server.pem"
  keyFile := "ssl/server.key"

  var err error

  // create TLS config
  tlsConfig := &tls.Config{}
  tlsConfig.Certificates = make([]tls.Certificate, 1)
  tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)
  must(err)

  // listen to port
  listener, err := net.Listen("tcp", bindAddress)
  must(err)

  // create polyglot
  pg := polyglot.NewTLSPolyglot(listener, tlsConfig)

  // serve HTTP
  err = http.Serve(pg, http.HandlerFunc(myHandler))
  must(err)
}

func myHandler(res http.ResponseWriter, req *http.Request) {
  // use req.TLS to detect HTTPS
  isHTTPS := (req.TLS != nil)

  res.Header().Set("Content-Type", "text/plain")
  res.WriteHeader(http.StatusOK)
  fmt.Fprintf(res, "Is HTTPS: %v\n", isHTTPS)
}

// if any error occurs, simply panic
// in production, you should do proper error handling instead
func must(err error) {
  if err != nil {
    panic(err)
  }
}
```
