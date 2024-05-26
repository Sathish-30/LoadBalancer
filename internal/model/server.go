package model

import (
	"net/http"
)

type Server interface {
	Address() string
	IsAlive(addr string) bool
	Serve(w http.ResponseWriter, r *http.Request)
}

func (s *SimpleServer) Address() string {
	return s.Addr
}

func (s *SimpleServer) IsAlive(addr string) bool {
	// Check if the server is running up or not
	// timeout := 5 * time.Second
	// conn, err := net.DialTimeout("tcp", addr, timeout)
	// if err != nil {
	// 	return false
	// }
	// defer conn.Close()
	return true
}

func (s *SimpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.Proxy.ServeHTTP(w, r)
}
