package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/model"
)
	
func TransferItem(c *gin.Context){

	var req model.TransferNFTRequest

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid JSON"})
		return
	}

	time.Sleep(time.Second)
	c.JSON(http.StatusOK, gin.H{
		"status":"Transferred",
		"item":req.Name,
		"to": req.Address,
	})
}
