package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunAPIRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"data": map[string]string{
				"desc": "这是一款xxx小程序",
			},
			"message": "ok",
		})
	})

	// 健康检测
	router.GET("/check", HealthCheck)
}

// HealthCheck gin check健康检查，主要用于发布流程后，进行心跳检测
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"alive": true,
	})
}
