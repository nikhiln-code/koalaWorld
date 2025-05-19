package handler

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/nikhiln-code/koalaWorld/backend-go/tests/utils"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetInventory( t *testing.T){
// 	router := utils.SetupTestRouter()

// 	req, _ := http.NewRequest("GET", "/api/inventory", nil)
// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)

// 	assert.Equal(t, 200, resp.Code)
// 	assert.Contains(t, resp.Body.String(), "Shadow Blade")
// }

// func TestMintItem(t *testing.T){
// 	router := utils.SetupTestRouter()

// 	reqBody := bytes.NewBuffer(utils.GetMintPayload())
// 	req, _ := http.NewRequest("POST", "/api/mint", reqBody)
// 	req.Header.Set("Content-Type", "application/json")

// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)

// 	assert.Equal(t, 200, resp.Code)
// 	assert.Contains(t, resp.Body.String(), "Minted")
// }

// func TestMintItem_InvalidJSON(t *testing.T) {
//     router := utils.SetupTestRouter()

//     reqBody := bytes.NewBuffer([]byte(`{invalid json}`))
//     req, _ := http.NewRequest("POST", "/api/mint", reqBody)
//     req.Header.Set("Content-Type", "application/json")
//     resp := httptest.NewRecorder()

//     router.ServeHTTP(resp, req)

//     assert.Equal(t, http.StatusBadRequest, resp.Code)
// }

// func TestMintItem_MissingFields(t *testing.T) {
//     router := utils.SetupTestRouter()

//     reqBody := bytes.NewBuffer([]byte(`{"name": ""}`))
//     req, _ := http.NewRequest("POST", "/api/mint", reqBody)
//     req.Header.Set("Content-Type", "application/json")
//     resp := httptest.NewRecorder()

//     router.ServeHTTP(resp, req)

//     assert.Equal(t, http.StatusBadRequest, resp.Code)
//     assert.Contains(t, resp.Body.String(), "name and address are required")
// }

// func TestTransferItem(t *testing.T){
// 	router := utils.SetupTestRouter()

// 	reqBody := bytes.NewBuffer(utils.GetTransferPayload())
// 	req, _ := http.NewRequest("POST", "/api/transfer", reqBody)
// 	req.Header.Set("Content-Type", "application/json")

// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)

// 	assert.Equal(t, 200, resp.Code)
// 	assert.Contains(t, resp.Body.String(), "Transferred")
// }

// func TestTransferItem_EmptyPayload(t *testing.T) {
//     router := utils.SetupTestRouter()

//     reqBody := bytes.NewBuffer([]byte(``)) // empty body
//     req, _ := http.NewRequest("POST", "/api/transfer", reqBody)
//     req.Header.Set("Content-Type", "application/json")
//     resp := httptest.NewRecorder()

//     router.ServeHTTP(resp, req)

//     assert.Equal(t, http.StatusBadRequest, resp.Code)
// }


