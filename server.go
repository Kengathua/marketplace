package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/matawis/matawis/pkg/apis"
	"github.com/matawis/matawis/pkg/config"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	config.ConnectToDb()
	api := app.Group("/api")
	apis.RegisterAPIRoutes(api, config.DB)

	args := os.Args[1:]
	port := "8000" // default port number
	if len(args) > 0 {
		port = args[0]
	}

	port = fmt.Sprintf(":%s", port)

	// Create a context that can be canceled.
	_, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure the context is canceled when the main function exits.

	// Start the server in a separate goroutine.
	go func() {
		if err := app.Listen(port); err != nil {
			log.Printf("Server error: %v\n", err)
			cancel() // Cancel the context to trigger graceful shutdown.
		}
	}()
	// Wait for an interrupt signal (e.g., SIGINT or SIGTERM).
	waitForInterrupt(cancel)

	log.Println("Shutting down gracefully...")
	// Perform any cleanup or additional shutdown tasks here.

}

func waitForInterrupt(cancelFunc context.CancelFunc) {
	// Create a channel to capture the interrupt signal.
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt) // Capture SIGINT (Ctrl+C) signals.

	// Wait for the interrupt signal.
	<-interruptChan

	// Call the cancel function to trigger graceful shutdown.
	cancelFunc()
}
