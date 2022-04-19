package main

import (
	"fiber-rest-api/controllers"
	"fiber-rest-api/logic"
	"fiber-rest-api/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"os/signal"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	repo := repositories.NewRepo()
	personLogic := logic.NewService(repo)

	api := app.Group("/api")
	controllers.PersonController(api.Group("/person"), personLogic)
	StartServer(app)
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServer(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Fatal("Oops... Server is not shutting down! Reason: ", err)
		}

		close(idleConnsClosed)
	}()

	// Build Fiber connection URL.
	fiberConnURL := fmt.Sprintf(
		"%s:%s",
		"0.0.0.0",
		"8080",
		//os.Getenv("SERVER_HOST"),
		//os.Getenv("SERVER_PORT"),
	)

	//	127.0.0.1 || null == bind to the loopback device, i.e. gofiber is only accessible by processes on the same machine.
	//	0.0.0.0 == bind to all network cards on that machine, i.e. the gofiber process is likely accessible to remote machines.

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
