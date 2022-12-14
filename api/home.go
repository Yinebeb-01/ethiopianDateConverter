package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Show Homepage godoc
// @Summary      Home page
// @Description  Home_page
// @Tags         homepage
// @Accept       json
// @Produce      json
// @Success      200  {object}  json
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       / [get]
func HomePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "EthioGrego Server",
	})
}
