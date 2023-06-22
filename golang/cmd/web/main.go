package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lucasoarruda/demo-project/golang/internal/config"
)

var app config.AppConfig

var (
	// Version is the current version of the application
	Version = "1.0.0"
	// BuildTime is the time the binary was built
	BuildTime = "1970-01-01T00:00:00Z"
	// GitCommit is the git commit that was compiled. This will be filled in by the compiler.
	GitCommit = "Testing"
)

func main() {
	app.BuildInfo = fmt.Sprintf("\n Version: %s || \n BuildTime: %s || \n GitCommit: %s", Version, BuildTime, GitCommit)
	if os.Getenv("GO_PORT") == "" {
		app.Port = ":8000"
	} else {
		app.Port = os.Getenv("GO_PORT")
	}
	log.Println(app.BuildInfo)
	fmt.Println("Listening on port", app.Port)

	router := routes(&app)
	s := &http.Server{
		Addr:         app.Port,
		Handler:      router,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}
	// Start the HTTP server in a goroutine
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s\n", err)
		}
	}()

	// Listen for the SIGINT signal and shut down the server
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigint

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %s\n", err)
	}

	fmt.Println("Server shut down.")
}
