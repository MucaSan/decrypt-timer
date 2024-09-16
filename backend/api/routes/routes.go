package routes

import (
	"api/api/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	config := cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"}, // Replace with the frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	r.Use(cors.New(config))
	r.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"WELCOME TO THE API!": "TEST"}) })
	r.POST("/caesar", controllers.EncryptDescription)
	r.Run()
}
