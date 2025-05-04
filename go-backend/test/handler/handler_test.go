package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/go-backend/internal/handler"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine{
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/inventory", handler.GetInventory)
		api.POST("/mint", handler.MintItem)
		api.POST("/transfer", handler.TransferItem)
	}
	return r
}

func TestGetInventory( t *testing.T){
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/inventory", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Shadow Blade")
}

func TestMintItem(t *testing.T){
	router:= setupTestRouter()

	body := map[string]string {"name": "Excalibur", "address":"0x123"}
	jsonValue, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/api/mint", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Minted")
}


func TestTransferItem(t *testing.T){
	router:= setupTestRouter()

	body := map[string]string {"name": "Guardian Shield", "address":"0x789"}
	jsonValue, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/api/transfer", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Transferred")
}


