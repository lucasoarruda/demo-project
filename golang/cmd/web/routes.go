package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"html/template"
	"log"

	docs "github.com/lucasoarruda/demo-project/golang/docs"
	"github.com/lucasoarruda/demo-project/golang/internal/config"
	"github.com/lucasoarruda/demo-project/golang/internal/handlers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       Swagger Autodeluge API
// @version     1.0
// @description This is external acess to the autodeluge-service.

//go:embed templates/*.html.tmpl

var f embed.FS

// @BasePath /api/v1
func routes(app *config.AppConfig) *gin.Engine {
	router := gin.Default()
	tmpl, err := template.ParseFS(f, "templates/*.html.tmpl")
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}

	// Set the HTML template for the router
	router.SetHTMLTemplate(tmpl)
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)
	handlers.Repo = &handlers.Repository{App: app}
	//v1 := router.Group("/api/v1")
	v1 := router.Group("")
	{

		// timezones
		v1.GET("/", handlers.Repo.HomePage)

		// health
		v1.GET("/health", handlers.Repo.Status)
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
