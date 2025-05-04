package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nikhiln-code/koalaWorld/go-backend/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetInventory( t *testing.T){
	router := utils.SetupTestRouter()

	req, _ := http.NewRequest("GET", "/api/inventory", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Shadow Blade")
}

func TestMintItem(t *testing.T){
	router := utils.SetupTestRouter()

	reqBody := bytes.NewBuffer(utils.GetMintPayload())
	req, _ := http.NewRequest("POST", "/api/mint", reqBody)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Minted")
}


func TestTransferItem(t *testing.T){
	router := utils.SetupTestRouter()

	reqBody := bytes.NewBuffer(utils.GetTransferPayload())
	req, _ := http.NewRequest("POST", "/api/transfer", reqBody)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Transferred")
}


