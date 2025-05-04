package model

type Item struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Rarity string `json:"rarity"`
	Image string`json:"image"`
	Owner string`json:"owner"`
}

type MintItemRequest struct {
	Name string `json:"name"`
	Address string `json:"address"`
}


type TransferItemRequest struct {
		Name string `json:"name"`
		Address string `json:"address"`
	}