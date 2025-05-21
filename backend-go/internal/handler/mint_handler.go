package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/logger"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/model/nft"
)

/*
** MintNFT handles the minting of NFTs
** It takes the name, description, rarity, filename, enviornment from the request
** It also takes the image file from the request
** It returns the NFT metadata and the IPFS hash of the image
** It returns an error if any of the required fields are missing
 */
func (h *NFTHandler) MintNFT(c *gin.Context){
	// Initialize logger
	log := logger.SugarLog()

	log.Info("MintNFT handler: Minting NFT")

	//Get the form values
	name := c.PostForm("name")
	description := c.PostForm("description")
	rarity := c.PostForm("rarity")
	fileName := c.PostForm("filename")
	enviornment := c.PostForm("enviornment")

	//Check if any of the required fields are missing
	if name == "" || description == "" || rarity == "" || fileName == "" {
		log.Debug("Missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error":"Name, description, filename, rarity are required"})
		return
	}

	//Get the image file from the request
	file, _, err := c.Request.FormFile("image")
	if err != nil{
		log.Debug("Error getting image file")
		c.JSON(http.StatusBadRequest, gin.H{"error":"Image is required"})
		return
	}
	defer file.Close()

	//Read the image file
	imageData, err := io.ReadAll(file)
	if err != nil{
		log.Debug("Error reading image file")
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error reading image"})
		return
	}

	//Create the NFT metadata
	fileMetadata := nft.NFTMetadata{
		Name: name,
		Description: description,
		Image: fileName,
		Environment: enviornment,
		Rarity: rarity,
	}

	//Upload the image to Pinata
	metadataURL, err := h.Service.UploadToPinata(h.PinataJWT, fileMetadata, imageData)
	if err != nil{
		log.Debug("Error uploading to Pinata")
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error uploading to Pinata: %v", err)})
		return
	}
	log.Info("Minted successfully Metadata URL: ", metadataURL)
	c.JSON(http.StatusOK, gin.H{
		"ipfs_metadata_url": metadataURL,
		"message": "NFT minted successfully",
	})
}