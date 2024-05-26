package utils

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/sathish-30/load-balancer/internal/model"
)

func NewSimpleServer(addr string) *model.SimpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)
	return &model.SimpleServer{
		Addr:  addr,
		Proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func NewLoadbalancer(port string, servers []model.Server) *model.LoadBalanacer {
	return &model.LoadBalanacer{
		Port:            port,
		RoundRobinCount: 0,
		Servers:         servers,
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
