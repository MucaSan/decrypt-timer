package routes

import (
	"api/api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"WELCOME TO THE API!": "TEST"}) })
	r.POST("/caesar", controllers.EncryptDescription)
	r.Run()
}
