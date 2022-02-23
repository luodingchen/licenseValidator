package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"verifyLinux/controllers"
)

func InitRouters() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("template/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/verify.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "verify.html", nil)
	})

	router.GET("/hardwareMsg/download", controllers.GenerateHardwareMsg)
	router.GET("/license/verify", controllers.VerifyLicense)
	router.POST("/license/upload", controllers.UploadLicense)

	return router
}
