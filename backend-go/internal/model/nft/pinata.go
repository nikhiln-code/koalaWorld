package nft

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

/**
** Pinata NFT Response model 
*/
var PinataNFTResponse struct {
	Data struct {
		Files []struct {
			ID            string                 `json:"id"`
			Name          string                 `json:"name"`
			CID           string                 `json:"cid"`
			Size          int                    `json:"size"`
			NumberOfFiles int                    `json:"number_of_files"`
			MimeType      string                 `json:"mime_type"`
			GroupID       *string                `json:"group_id"`
			Keyvalues     map[string]interface{} `json:"keyvalues"`
			CreatedAt     string                 `json:"created_at"`
		} `json:"files"`
		NextPageToken string `json:"next_page_token"`
	} `json:"data"`
}