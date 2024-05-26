package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sathish-30/load-balancer/internal/model"
	"github.com/sathish-30/load-balancer/internal/utils"
)

var lb *model.LoadBalanacer

func main() {
	servers := []model.Server{
		// The machine name is all the service named in the docker-compose file
		utils.NewSimpleServer("http://service1:8070"),
		utils.NewSimpleServer("http://service2:8080"),
		utils.NewSimpleServer("http://service3:8090"),
	}

	lb = utils.NewLoadbalancer(os.Getenv("PORT"), servers)
	http.HandleFunc("/", handleRedirect)
	http.HandleFunc("/healh", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"Message": "Load Balancer health check",
		})
	})
	fmt.Printf("Server is running on the port %v\n", lb.Port)
	http.ListenAndServe(":"+lb.Port, nil)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	lb.ServeProxy(w, r)
}
