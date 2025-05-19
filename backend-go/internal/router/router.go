package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/config"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/handler"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/service"
)

/*
** Setup the Gin router
** It sets up the routes and handlers for the application
** It also have the CORS configuration
**/


func SetupRouter() *gin.Engine{
	r:= gin.Default()
	
	//Load the configuration to get the Cors Allowed origins
	conf, _ := config.LoadConfig()

	//Check if the Cors Allowed origins are set in the configuration file
	if conf.Cors.AllowedOrigins == nil || len(conf.Cors.AllowedOrigins) == 0{
		panic("CORS Allowed Origins are not set in the configuration")
	}

	// CORS configuration
    r.Use(cors.New(cors.Config{
        AllowOrigins:     conf.Cors.AllowedOrigins,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	nftHander := handler.NewNFTHandler(conf.Pinata.JwtKEY, service.NewNFTService())

	api := r.Group("/api/v1")
	{
		api.GET("/fetch", nftHander.GetNFT)
		api.POST("/transfer", handler.TransferItem)
		api.POST("/mint", nftHander.MintNFT)
	}

	return r
}