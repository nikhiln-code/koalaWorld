package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/handler"
)


func SetupRouter() *gin.Engine{
	r:= gin.Default()
	
    // CORS configuration
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	api := r.Group("/api")
	{
		api.GET("/inventory", handler.GetInventory)
		api.POST("/mint", handler.MintItem)
		api.POST("/transfer", handler.TransferItem)
	}
	return r
}
