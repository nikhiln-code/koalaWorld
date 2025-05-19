package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/model"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/service"
) 

type NFTHandler struct{
	PinataJWT string
	Service service.NFTServiceInterface
}

func NewNFTHandler(jwt string, svc service.NFTServiceInterface) *NFTHandler{
	return &NFTHandler{
		PinataJWT: jwt,
		Service: svc,
	}
}


var mockInventory = []model.NFT{
    {ID: "1", Name: "Shadow Blade", Rarity: "Legendary", Image: "/assets/sword.png", Owner: "0x123.."},
    {ID: "2", Name: "Guardian Shield", Rarity: "Epic", Image: "/assets/shield.png", Owner: "0xC56.."},
    {ID: "3", Name: "Healing Potion", Rarity: "Common", Image: "/assets/potion.png", Owner: "0x769.."},
}


func GetInventory (c *gin.Context){
	c.JSON(http.StatusOK, mockInventory)
}