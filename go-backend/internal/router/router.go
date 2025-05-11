package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/config"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/handler"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/service"
)


func SetupRouter() *gin.Engine{
	r:= gin.Default()
	conf, _ := config.LoadConfig()

	fmt.Println("config cors", conf.Cors.AllowedOrigins)
	
	// CORS configuration
    r.Use(cors.New(cors.Config{
        AllowOrigins:     conf.Cors.AllowedOrigins,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	
	mintHander := handler.NewMintHandler(conf.Pinata.JwtKEY, service.NewNFTService())

	api := r.Group("/api/v1")
	{
		api.GET("/inventory", handler.GetInventory)
		api.POST("/mint", mintHander.MintNFT)
		api.POST("/transfer", handler.TransferItem)
	}
	return r
}
