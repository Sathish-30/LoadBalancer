package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sathish-30/load-balancer/internal/model"
	"github.com/sathish-30/load-balancer/internal/utils"
)

var lb *model.LoadBalanacer

func main() {
	servers := []model.Server{
		utils.NewSimpleServer("https://www.github.com"),
		utils.NewSimpleServer("https://www.youtube.com/"),
		utils.NewSimpleServer("http://www.duckduckgo.com"),
	}

	lb = utils.NewLoadbalancer(os.Getenv("PORT"), servers)
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Server is running on the port %v\n", lb.Port)
	http.ListenAndServe(":"+lb.Port, nil)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	lb.ServeProxy(w, r)
}
