package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index godoc
// @Summary index
// @Schemes
// @Description do ping
// @Tags Index
// @Accept json
// @Produce json
// @Success 200 {string} Index
// @Router / [get]
func Index(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"message": "hello"})
}
