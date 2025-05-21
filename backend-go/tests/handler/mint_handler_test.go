package handler

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/handler"
	"github.com/nikhiln-code/koalaWorld/backend-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createMultipartRequest(fields map[string]string, fileField string, fileName string, 
	fileContent []byte)(*http.Request, *bytes.Buffer){
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range fields{
		_ = writer.WriteField(key, value)
	}

	if fileField != "" && fileName != ""{
		part,_ := writer.CreateFormFile(fileField, fileName)
		part.Write(fileContent)
	}

	writer.Close()
	req, _ := http.NewRequest(http.MethodPost, "/mint", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, body
}


func TestMintNFT_Success(t *testing.T){
	gin.SetMode(gin.TestMode)
	mockSvc := new(utils.MockNFTService)

	mockSvc.On("UploadToPinata", "test-jwt", mock.Anything, mock.Anything).Return("ipfs://abcd1234", nil)

	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service:   mockSvc,
	}

	fields := map[string]string{
		"name":        "Test NFT",
		"description": "An awesome NFT",
		"rarity":      "rare",
		"filename":    "image.png",
		"enviornment": "dev",
	}

	req, _ := createMultipartRequest(fields,"image", "image.png" , []byte("fake image"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	h.MintNFT(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "NFT minted successfully")
	assert.Contains(t, w.Body.String(), "ipfs_metadata_url")
	assert.Contains(t, w.Body.String(), "ipfs://abcd1234")

	mockSvc.AssertExpectations(t)
}

func TestMintNFT_MissingFields(t *testing.T){
	gin.SetMode(gin.TestMode)

	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service:   new(utils.MockNFTService),
	}

	fields := map[string]string{
		"description": "An awesome NFT",
	}

	req, _ := createMultipartRequest(fields, "image", "image.png", []byte("fake image"))

	w := httptest.NewRecorder()
	c, _:= gin.CreateTestContext(w)
	c.Request = req
	h.MintNFT(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Name, description, filename, rarity are required")
}

func TestMintNFT_missingImage(t *testing.T){
	gin.SetMode(gin.TestMode)

	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service:   new(utils.MockNFTService),
	}

	fields  := map[string] string{
		"name": "Test NFT",
		"description": "An awesome NFT",
		"rarity": "rare",
		"filename": "image.png",
		"enviornment": "dev",
	}

	req, _ := createMultipartRequest(fields, "", "", nil)

	w :=  httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	h.MintNFT(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Image is required")
	
}

func TestMintNFT_UploadError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(utils.MockNFTService)

	mockSvc.On("UploadToPinata", "test-jwt", mock.Anything, mock.Anything).
		Return("", errors.New("upload failed"))

	h := &handler.NFTHandler{
		PinataJWT: "test-jwt",
		Service:   mockSvc,
	}

	fields := map[string]string{
		"name":        "Test NFT",
		"description": "An awesome NFT",
		"rarity":      "rare",
		"filename":    "image.png",
		"enviornment": "dev",
	}

	req, _ := createMultipartRequest(fields, "image", "image.png", []byte("fake image"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	h.MintNFT(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Error uploading to Pinata")

	mockSvc.AssertExpectations(t)
}