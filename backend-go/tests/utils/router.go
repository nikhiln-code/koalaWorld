package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/handler"
)

func SetupTestRouter() *gin.Engine {
    r := gin.Default()
    api := r.Group("/api")
    {
        api.GET("/fetch", handler.GetInventory)
      //  api.POST("/mint", handler.MintItem)
        api.POST("/transfer", handler.TransferNFT)
    }
    return r
}
