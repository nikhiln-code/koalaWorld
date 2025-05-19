package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/handler"
	"github.com/nikhiln-code/koalaWorld/backend-go/tests/utils"
	"github.com/stretchr/testify/assert"
)




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
	mockSvc.On("GetNFT", "test-jwt", "xyz1234").Return(`{"id":"xyz1234","name":"Koola Sydney"}`, nil)

	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service: mockSvc,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/nfts?cid=xyz1234", nil)


	h.GetNFT(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Koola Sydney" )
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

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Error fetching NFTs")
	mockSvc.AssertExpectations(t)
}


func TestGetNFT_Error_Returns500(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := new(utils.MockNFTService)
	mockSvc.On("GetNFT", "test-jwt", "fail123").Return("", assert.AnError)

	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service:   mockSvc,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/nfts?cid=fail123", nil)

	h.GetNFT(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Error fetching NFT")
	mockSvc.AssertExpectations(t)
}
