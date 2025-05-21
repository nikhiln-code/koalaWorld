package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/handler"
	"github.com/nikhiln-code/koalaWorld/backend-go/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	utils.SetupTestLogger()
	code := m.Run()
	os.Exit(code)
}

func TestGetNFTs_NoCID_ReturnsAllNFTs(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := new(utils.MockNFTService)
	mockSvc.On("GetNFTs", "test-jwt").Return(`{"nfts": ["NFT1", "NFT2"]}`, nil)

	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service:   mockSvc,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/nfts", nil)

	h.GetNFT(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "NFT1")
	mockSvc.AssertExpectations(t)
}

func TestGetNFT_WithCID_ReturnsNFT(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := new(utils.MockNFTService)
	mockSvc.On("GetNFT", "test-jwt", "bafkreichrsckrl2uxbqcddfsmvussq2jedw5po7sb23uab52tfsr7gycufa").Return(`{"id":"bafkreichrsckrl2uxbqcddfsmvussq2jedw5po7sb23uab52tfsr7gycfa","name":"Koola Sydney"}`, nil)

	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service: mockSvc,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/nfts?cid=bafkreichrsckrl2uxbqcddfsmvussq2jedw5po7sb23uab52tfsr7gycufa", nil)


	h.GetNFT(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Koola Sydney" )
	mockSvc.AssertExpectations(t)
}

func TestGetNFT_WithCID_NOT_HAVING_59_CHARS(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := new(utils.MockNFTService)
	
	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service: mockSvc,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/nfts?cid=baf", nil)


	h.GetNFT(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid CID" )
	assert.Contains(t, w.Body.String(), "CID must be a valid IPFS hash (minimum 59 characters)")
	mockSvc.AssertExpectations(t)
}

func TestGetNFTs_Error_Returns500(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := new(utils.MockNFTService)
	mockSvc.On("GetNFTs", "test-jwt").Return("", assert.AnError)

	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service:   mockSvc,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/nfts", nil)

	h.GetNFT(c)

	assert.Equal(t, http.StatusBadGateway, w.Code)
	assert.Contains(t, w.Body.String(), "Unable to fetch NFTs")
	mockSvc.AssertExpectations(t)
}


func TestGetNFT_Error_Returns500(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := new(utils.MockNFTService)
	mockSvc.On("GetNFT", "test-jwt", "bafkreichrsckrl2uxbqcddfsmvussq2jedw5po7sb23uab52tfsr7gycufa").Return("", assert.AnError)

	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service:   mockSvc,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/nfts?cid=bafkreichrsckrl2uxbqcddfsmvussq2jedw5po7sb23uab52tfsr7gycufa", nil)

	h.GetNFT(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Error fetching NFT")
	mockSvc.AssertExpectations(t)
}
