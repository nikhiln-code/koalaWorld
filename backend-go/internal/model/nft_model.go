package model

/**
** NFT model
 */
type NFT struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Rarity string `json:"rarity"`
	Image string`json:"image"`
	Owner string`json:"owner"`
}

/**
** Mint NFT request model
*/
type MintNFTRequest struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

/**
** Transfer NFT request model
*/
type TransferNFTRequest struct {
		Name string `json:"name"`
		Address string `json:"address"`
}