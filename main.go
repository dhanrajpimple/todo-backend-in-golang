package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"todo-app/config"
	"todo-app/routes"
	"github.com/rs/cors"
	"github.com/joho/godotenv"
)

// isPortAvailable checks if a port is available by trying to bind to it.
func isPortAvailable(port string) bool {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return false // port is not available
	}
	ln.Close() // port is available, so close it immediately
	return true
}

// getAvailablePort returns the first available port from the environment variable or defaults to 8000.
func getAvailablePort(defaultPort string) string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Check if the desired port is available, and try a fallback if not
	if !isPortAvailable(port) {
		log.Printf("Port %s is already in use. Trying a fallback port...\n", port)
		for i := 1; i <= 5; i++ {
			fallbackPort := fmt.Sprintf("%d", 8000+i)
			if isPortAvailable(fallbackPort) {
				log.Printf("Using fallback port: %s\n", fallbackPort)
				return fallbackPort
			}
		}
		log.Fatal("No available port found.")
	}

	return port
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up the CORS handler
	corsHandler := cors.Default()

	// Get the available port, defaulting to 8000
	port := getAvailablePort("8000")

	// Connect to MongoDB
	config.ConnectToMongoDB()

	// Set up routes
	router := routes.SetupRouter()

	// Start the server with the available port and CORS handler
	log.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler.Handler(router)))
}
