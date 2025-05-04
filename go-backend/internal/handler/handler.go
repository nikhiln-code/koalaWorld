package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/model"
) 


var mockInventory = []model.Item{
    {ID: "1", Name: "Shadow Blade", Rarity: "Legendary", Image: "/assets/sword.png"},
    {ID: "2", Name: "Guardian Shield", Rarity: "Epic", Image: "/assets/shield.png"},
    {ID: "3", Name: "Healing Potion", Rarity: "Common", Image: "/assets/potion.png"},
}


func GetInventory (c *gin.Context){
	c.JSON(http.StatusOK, mockInventory)
}

func MintItem (c *gin.Context){
	var req struct {
		Name string `json:"name"`
		Address string `json:"address"`
	}

	if err := c.BindJSON(&req) ; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return 
	}

	time.Sleep(time.Second)
	c.JSON(http.StatusOK, gin.H{
		"status" :"Minted",
		"item": req.Name,
		"to": req.Address,
		})
}


func TransferItem(c *gin.Context){

	var req struct {
		Name string `json:"name"`
		Address string `json:"address"`
	}

	if err := c.BindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid request"})
		return
	}

	time.Sleep(time.Second)
	c.JSON(http.StatusOK, gin.H{
		"status":"Transferred",
		"item":req.Name,
		"to": req.Address,
	})
}