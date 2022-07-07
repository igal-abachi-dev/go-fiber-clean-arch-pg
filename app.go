package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"go-fiber-clean-arch-pg/controllers"
	"go-fiber-clean-arch-pg/logic"
	"go-fiber-clean-arch-pg/repositories"
	"log"
	"os"
	"os/signal"
	//	_ "github.com/heroku/x/hmetrics/onload"
)

//file main.go

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(favicon.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // https://github.com/valyala/fasthttp/blob/master/brotli.go
	}))
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"": "123456",
			//"admin": "123456",
		},
	}))

	/*
		cfg.Authorizer = func(user, pass string) bool {
			userPwd, exist := cfg.Users[user]
			return exist && subtle.ConstantTimeCompare(utils.UnsafeBytes(userPwd), utils.UnsafeBytes(pass)) == 1
		}*/

	//repo := repositories.NewPersonRepo()
	//personLogic := logic.NewPersonLogic(repo)

	bookRepo := repositories.NewBookRepository()
	bookLogic := logic.NewBookLogic(bookRepo)

	categoryRepo := repositories.NewCategoryRepository()
	categoryLogic := logic.NewCategoryLogic(categoryRepo)

	fileContentRepo := repositories.NewFileContentRepository()
	fileContentLogic := logic.NewFileContentLogic(fileContentRepo)

	api := app.Group("/api")
	controllers.BookController(api.Group("/book"), bookLogic)
	controllers.CategoryController(api.Group("/category"), categoryLogic)
	controllers.FileContentController(api.Group("/fileContent"), fileContentLogic)

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

	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "8080"
	}

	// Build Fiber connection URL.
	fiberConnURL := fmt.Sprintf(
		"%s:%s",
		"0.0.0.0",
		port,
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
