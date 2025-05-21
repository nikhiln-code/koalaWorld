package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/logger"
)

/*
** This handler is responsible for fetching the NFTs from the Pinata in IPFS format
*/
func (h *NFTHandler) GetNFT(c *gin.Context) {
	// Initialize logger
	log := logger.SugarLog()
	
	log.Info("GetNFT handler: Fetching NFTs")
	
	cid := c.Query("cid")
	log.Debug("Fetching NFTs for cid: %s", cid)

	var (
		result string
		err    error
	)

	if cid == "" { // When cid is not given it will fetch all the avaible public nfts
		log.Debug("cid not provided : Fetching all NFTs")
		result, err = h.Service.GetNFTs(h.PinataJWT)
		if err != nil {

			c.JSON(http.StatusBadGateway, gin.H{
				"error": "Unable to fetch NFTs",
				"details":err.Error(),
			})
			return
		}
	} else { // try to fetch the nft with the given cid
		log.Debug("Fetching NFTs for cid: %s", cid)
		// Validate CID format (basic check for now)
		if len(cid) < 59 {
			log.Debug("Invalid CID format length expected as 59 it has : %s", len(cid), "chars")
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid CID",
				"details": "CID must be a valid IPFS hash (minimum 59 characters)",
			})
			return
		}


		result, err = h.Service.GetNFT(h.PinataJWT, cid)
		if err != nil {
			log.Error("Error fetching NFT:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error fetching NFT",
				"details": err.Error(),
			})
			return
		}
	}

	var parsedData interface{}
	if err := json.Unmarshal([]byte(result), &parsedData); err != nil {
		log.Error("Failed to parse pinata response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse pinata response",
			"details": err.Error(),
		})
		return
	}

	log.Debug("NFTs fetched successfully")
	c.JSON(http.StatusOK, gin.H{
		"data":   parsedData,
		"message": "NFT(s) fetched successfully",
	})
}