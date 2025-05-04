package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/handler"
)

func SetupTestRouter() *gin.Engine {
    r := gin.Default()
    api := r.Group("/api")
    {
        api.GET("/inventory", handler.GetInventory)
        api.POST("/mint", handler.MintItem)
        api.POST("/transfer", handler.TransferItem)
    }
    return r
}
