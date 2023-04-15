package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasoarruda/demo-project/golang/internal/config"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Status Health Check
// @Summary Health Check
// @Schemes
// @Description do Health Check
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {string} Health Check
// @Router /health [get]
func (m *Repository) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (m *Repository) HomePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "homepage.html.tmpl", gin.H{
		"title": "Main website",
	})
}
