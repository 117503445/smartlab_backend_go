package api

import (
	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary 状态检查
// @Description 返回 pong
// @ID ping
// @Accept  json
// @Produce  json
// @Success 200 {string} string "pong"
// @Router /ping [post]
func Ping(c *gin.Context) {
	c.JSON(200, "pong")
}
