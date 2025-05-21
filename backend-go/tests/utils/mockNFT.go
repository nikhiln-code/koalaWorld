package utils

import (
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/model/nft"
	"github.com/stretchr/testify/mock"
)	
type MockNFTService struct {
	mock.Mock
}

func (m *MockNFTService) GetNFTs(jwt string) (string, error) {
	args := m.Called(jwt)
	return args.String(0), args.Error(1)
}

func (m *MockNFTService) GetNFT(jwt, cid string) (string, error) {
	args := m.Called(jwt, cid)
	return args.String(0), args.Error(1)
}

func (m *MockNFTService) UploadToPinata(jwt string, metadata nft.NFTMetadata, imageData []byte) (string, error) {
	args := m.Called(jwt, metadata, imageData)
	return args.String(0), args.Error(1)
}