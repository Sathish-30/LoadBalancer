# Load Balancer in Go

This repository contains a simple load balancer written in Go. The load balancer forwards incoming requests to a list of backend servers in a round-robin manner.

## Code Explanation

### Main Components

1. **Server Interface**: Defines the methods that any server should implement.
2. **SimpleServer Struct**: Implements the Server interface and represents a backend server.
3. **LoadBalancer Struct**: Manages a list of servers and distributes the load among them.

### Functions

- `newSimpleServer(addr string)`: Creates a new instance of SimpleServer.
- `NewLoadbalancer(port string, servers []Server)`: Creates a new load balancer with the specified port and servers.
- `handleErr(err error)`: Handles errors by printing them.
- `(*LoadBalanacer) getNextAvailableServer()`: Returns the next available server in a round-robin manner.
- `(*LoadBalanacer) serveProxy(w http.ResponseWriter, r *http.Request)`: Forwards the incoming request to the selected server.
- `main()`: Initializes the servers and starts the load balancer.

### Running the Load Balancer

To run the load balancer, you need to have Docker and docker-compose installed. You can update the port in the docker-compose.

```docker-compose
docker compose up -d