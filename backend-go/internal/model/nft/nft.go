package nft

/**
** NFT model
 */
type NFTMetadata struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"` 
	Environment string `json:"environment"`
	Rarity string `json:"rarity"`
}