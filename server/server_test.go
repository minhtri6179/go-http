package main

import (
	"net"
	"net/http"
	"server/handler"
	"testing"
	"time"
)

func TestSimpleHTTPServer(t *testing.T) {
	srv := &http.Server{
		Addr: "127.0.0.1:8000",
		Handler: http.TimeoutHandler(
			handler.DefaultHandler(), 2*time.Minute, ""),
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}
	listen, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		err := srv.Serve(listen)
		if err != http.ErrServerClosed {
			t.Error(err)
		}
	}()

}
