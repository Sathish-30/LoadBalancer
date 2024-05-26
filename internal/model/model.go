package model

import (
	"log"
	"net/http"
	"net/http/httputil"
)

type SimpleServer struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

type LoadBalanacer struct {
	Port            string
	RoundRobinCount int
	Servers         []Server
}

func (lb *LoadBalanacer) getNextAvailableServer() Server {
	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	// check if server is alive or not, if not switch server
	// for !server.IsAlive(server.Address()) {
	// 	lb.RoundRobinCount++
	// 	server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	// }

	lb.RoundRobinCount++
	return server
}

func (lb *LoadBalanacer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	log.Printf("Forwarding the request to address %v\n", targetServer.Address())
	targetServer.Serve(w, r)
}
