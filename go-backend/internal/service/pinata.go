package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type NFTService struct{	
}

func NewNFTService() *NFTService{
	return &NFTService{}
}

func (s *NFTService) UploadToPinata(jwt string, metadata NFTMetadata, imageData []byte) (string, error) {
	client := resty.New()

	//1. Upload the image 
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Create form file for the image
	part, _ := writer.CreateFormFile("file", metadata.Image)
	part.Write(imageData)
	writer.WriteField("network", "public")
	

	keyvalues := map[string]string{
		"name":        metadata.Name,
		"description": metadata.Description,
		"env":         metadata.Environment,
		"rarity":      metadata.Rarity,
		"image":       metadata.Image,
	}

	keyvalueJSON, err := json.Marshal(keyvalues)
	outer := map[string]interface{}{
		"keyvalues": string(keyvalueJSON), 
	}
	outerJSON, err := json.Marshal(outer)

	// Write keyvalues field
	writer.WriteField("keyvalues", string(outerJSON))
	// Close the multipart writer
	writer.Close()
	
	// Send the request to Pinata
	imageResp, err := client.R().
		SetHeader("Authorization", jwt).
		SetHeader("Content-Type", writer.FormDataContentType()).
		SetBody(&buf).
		Post("https://uploads.pinata.cloud/v3/files")

	// Check for errors or bad status codes
	if err != nil || imageResp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("error uploading image to Pinata: %w", err)
	}

	// Unmarshal the response
	var imagerResult map[string]interface{}
	json.Unmarshal(imageResp.Body(), &imagerResult)
	// Safely extract the CID
	data, _ := imagerResult["data"].(map[string]interface{})
	imageCid, _ := data["cid"].(string)

	// Construct the image URL
	imageUrl := "ipfs://" + imageCid
	return imageUrl, nil
}

type NFTMetadata struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"` 
	Environment string `json:"environment"`
	Rarity string `json:"rarity"`
}
