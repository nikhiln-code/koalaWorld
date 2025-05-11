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

func (s *NFTService) UploadToPinata(jwt, name, description string, imageData []byte) (string, error){
	client := resty.New()

	//1. Upload the image 
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, _ := writer.CreateFormFile("file", "nft.png")
	part.Write(imageData)
	writer.Close()	

	imageResp , err := client.R().
	SetHeader("Authorization", jwt).
	SetHeader("Content-Type", writer.FormDataContentType()).
	SetBody(&buf).
	Post("https://api.pinata.cloud/pinning/pinFileToIPFS")


	if err != nil || imageResp.StatusCode() != http.StatusOK{
		return "", fmt.Errorf("Error uploading image to Pinata: %w", err)
	}


	var imagerResult map[string]interface{}
	json.Unmarshal(imageResp.Body(), &imagerResult)
	imageCid := imagerResult["IpfsHash"].(string)
	imageUrl :="ipfs://" +imageCid


	//2. Upload metadata JSON
	metadata := Metadata{
		Name: name,
		Description: description,
		Image: imageUrl,
	}
	metaBytes, _ := json.Marshal(metadata)

	metaResp, err := client.R().
		SetHeader("Authorization", jwt).
		SetHeader("Content-Type", "application/json").
		SetBody(metaBytes).
		Post("https://api.pinata.cloud/pinning/pinJSONToIPFS")

	if err != nil || metaResp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("metadata upload failed: %v", err)
	}

	var metaResult map[string]interface{}
	json.Unmarshal(metaResp.Body(), &metaResult)
	metaCid := metaResult["IpfsHash"].(string)

	return "ipfs://" + metaCid, nil
}

type Metadata struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"` 
}
