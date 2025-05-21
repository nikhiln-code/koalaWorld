package handler

import (
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/service"
) 

type NFTHandler struct{
	PinataJWT string
	Service service.NFTServiceInterface
}

func NewNFTHandler(jwt string, svc service.NFTServiceInterface) *NFTHandler{
	return &NFTHandler{
		PinataJWT: jwt,
		Service: svc,
	}
}