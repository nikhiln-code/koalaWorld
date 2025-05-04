package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/handler"
)


func SetupRouter() *gin.Engine{
	r:= gin.Default()
	
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/inventory", handler.GetInventory)
		api.POST("/mint", handler.MintItem)
		api.POST("/transfer", handler.TransferItem)
	}
	return r
}
