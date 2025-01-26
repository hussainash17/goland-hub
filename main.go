package main

import (
	"awesomeProject/config"
	"awesomeProject/consumers"
	"awesomeProject/routes"
	"log"
)

func main() {
	// Initialize the router
	router := routes.SetupRouter()
	config.InitDB()
	consumers.NewsConsumer()
	// Start the server
	log.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
