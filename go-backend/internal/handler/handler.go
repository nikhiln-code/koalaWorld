package handler

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/model"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/service"
) 


var mockInventory = []model.Item{
    {ID: "1", Name: "Shadow Blade", Rarity: "Legendary", Image: "/assets/sword.png", Owner: "0x123.."},
    {ID: "2", Name: "Guardian Shield", Rarity: "Epic", Image: "/assets/shield.png", Owner: "0xC56.."},
    {ID: "3", Name: "Healing Potion", Rarity: "Common", Image: "/assets/potion.png", Owner: "0x769.."},
}


func GetInventory (c *gin.Context){
	c.JSON(http.StatusOK, mockInventory)
}

type MintHandler struct{
	PinataJWT string
	Service *service.NFTService
}

func NewMintHandler(jwt string, svc *service.NFTService) *MintHandler{
	return &MintHandler{
		PinataJWT: jwt,
		Service: svc,
	}
}

func (h *MintHandler) MintNFT(c *gin.Context){
	name := c.PostForm("name")
	description := c.PostForm("description")

	if name == "" || description == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Name and description are required"})
		return
	}
	file, _, err := c.Request.FormFile("image")
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Image is required"})
		return
	}
	defer file.Close()

	imageData, err := io.ReadAll(file)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error reading image"})
		return
	}

	metadataURL, err := h.Service.UploadToPinata(h.PinataJWT, name, description, imageData)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error uploading to Pinata: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{


		"ipfs_metadata_url": metadataURL,
		"message": "NFT minted successfully",
	})

}

func TransferItem(c *gin.Context){

	var req model.TransferItemRequest

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

