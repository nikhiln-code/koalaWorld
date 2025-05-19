package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *NFTHandler) GetNFT(c *gin.Context) {
	cid := c.Query("cid")

	var (
		result string
		err    error
	)

	if cid == "" {
		result, err = h.Service.GetNFTs(h.PinataJWT)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching NFTs"})
			return
		}
	} else {
		result, err = h.Service.GetNFT(h.PinataJWT, cid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching NFT"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "NFT(s) fetched successfully",
	})
}