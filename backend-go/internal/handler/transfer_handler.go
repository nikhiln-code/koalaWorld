package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/model/nft"
)

/*
** This Handler is responsible for transfer the ownership of NFT from the old owner to the
** New Owner
** Internally it is expected to provide a valid wallet address
** todo: This is not implemented correctly yet
 */	
func TransferNFT(c *gin.Context){

	var req nft.TransferNFTRequest

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
