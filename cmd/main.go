package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	isAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

func (s *SimpleServer) Address() string {
	return s.addr
}

func (s *SimpleServer) isAlive() bool {
	// Check if the server is running up or not
	return true
}

func (s *SimpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

type SimpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

type LoadBalanacer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func newSimpleServer(addr string) *SimpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)
	return &SimpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func NewLoadbalancer(port string, servers []Server) *LoadBalanacer {
	return &LoadBalanacer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func (lb *LoadBalanacer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	// check if server is alive or not, if not switch server
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalanacer) serveProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	log.Printf("Forwarding the request to address %v\n", targetServer.Address())
	targetServer.Serve(w, r)
}

var lb *LoadBalanacer

func main() {
	servers := []Server{
		newSimpleServer("https://www.github.com"),
		newSimpleServer("https://www.youtube.com/"),
		newSimpleServer("http://www.duckduckgo.com"),
	}

	lb = NewLoadbalancer(os.Getenv("PORT"), servers)
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Server is running on the port %v\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	lb.serveProxy(w, r)
}
