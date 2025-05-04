package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/model"
) 


var mockInventory = []model.Item{
    {ID: "1", Name: "Shadow Blade", Rarity: "Legendary", Image: "/assets/sword.png", Owner: "0x123.."},
    {ID: "2", Name: "Guardian Shield", Rarity: "Epic", Image: "/assets/shield.png", Owner: "0xC56.."},
    {ID: "3", Name: "Healing Potion", Rarity: "Common", Image: "/assets/potion.png", Owner: "0x769.."},
}


func GetInventory (c *gin.Context){
	c.JSON(http.StatusOK, mockInventory)
}

func MintItem (c *gin.Context){
	var req model.MintItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if req.Name == "" || req.Address == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "name and address are required"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "Minted", "item": req})
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

