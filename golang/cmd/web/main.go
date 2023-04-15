package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/lucasoarruda/demo-project/golang/internal/config"
)

var app config.AppConfig

func main() {
	if os.Getenv("GO_PORT") == "" {
		app.Port = ":8000"
	} else {
		app.Port = os.Getenv("GO_PORT")
	}
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
