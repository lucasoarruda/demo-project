package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	app.BuildInfo = fmt.Sprintf("Version: %s || \n BuildTime: %s || \n GitCommit: %s", Version, BuildTime, GitCommit)
	if os.Getenv("GO_PORT") == "" {
		app.Port = ":8000"
	} else {
		app.Port = os.Getenv("GO_PORT")
	}
	log.Println(app.BuildInfo)
	router := routes(&app)
	s := &http.Server{
		Addr:         app.Port,
		Handler:      router,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Println("Listening on port", app.Port)
	_, err = fmt.Scanln()
	if err != nil {
		return
	}
}
