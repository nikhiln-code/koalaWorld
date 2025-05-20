package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/model"
)

const(
	PINATA_UPLOAD_URL              = "https://uploads.pinata.cloud/v3/files"
	PINATA_FETCH_ALL_NFT_URL       = "https://api.pinata.cloud/v3/files/public"
	PINATA_FETCH_SPECIFIC_NFT_URL  = "https://api.pinata.cloud/v3/files/public?cid="
)

type NFTService struct{	
}

func NewNFTService() *NFTService{
	return &NFTService{}
}

type NFTMetadata struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"` 
	Environment string `json:"environment"`
	Rarity string `json:"rarity"`
}

type NFTServiceInterface interface{
	GetNFTs(jwt string) (string, error)
	GetNFT(jwt, cid string) (string, error)
	UploadToPinata(jwt string, metadata NFTMetadata, imageData []byte) (string, error)
}

/*
** Service method to upload the NFT to Pinata , should have an valid jwt token provided in the config
** It expects valid NFTMetadata containing Name, Description, Image, enviornment, rarity
** It talks to pinata via API calls on https://uploads.pinata.cloud/v3/files
** We are using Resty for api calls 
** Once successful this method will return the CID which can be used to fetch the NFT details.
*/
func (s *NFTService) UploadToPinata(jwt string, metadata NFTMetadata, imageData []byte) (string, error) {
	client := resty.New()

	//1. Upload the image 
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Create form file for the image
	part, _ := writer.CreateFormFile("file", metadata.Image)
	part.Write(imageData)
	writer.WriteField("network", "public") // Uploading to a public network in pinata
	
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
		Post(PINATA_UPLOAD_URL)

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

/*
** GetNFT Service method will return the details of specific NFT with the given cid
** a valid JWT token is required which will be fetched from the config file for this project
** It gives a call to Pinata api on https://api.pinata.cloud/v3/files/public?cid=
** 
*/
func (s *NFTService) GetNFT(jwt ,cid string) (string, error) {
	client := resty.New()

	resp, err := client.R().
	SetHeader("Authorization", jwt).
	SetHeader("Content-Type", "application/json").
	Get(PINATA_FETCH_SPECIFIC_NFT_URL + cid)

	// Check for errors or bad status codes
	if err != nil || resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("error fetching NFT from Pinata: %w", err)
	}

	var response = model.PinataNFTResponse

	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return string(resp.Body()), nil
}

/*
** GetNFT Service method will return all available public NFTs 
** a valid JWT token is required which will be fetched from the config file for this project
** It gives a call to Pinata api on https://api.pinata.cloud/v3/files/public
** 
*/
func (s *NFTService) GetNFTs(jwt string) (string, error) {
	client := resty.New()

	resp, err := client.R().
	SetHeader("Authorization", jwt).
	SetHeader("Content-Type", "application/json").
	Get(PINATA_FETCH_ALL_NFT_URL)

	// Check for errors or bad status codes
	if err != nil || resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("error fetching NFTs from Pinata: %w", err)
	}

	var response = model.PinataNFTResponse
	
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	result, err := json.Marshal(&response)
	if err != nil {
		return "", fmt.Errorf("failed to marshal final JSON: %w", err)
	}

	return string(result), nil
}