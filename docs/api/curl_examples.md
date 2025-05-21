# 🎯 API Examples

### 🌐 Fetch All NFTs

```bash
curl --request GET \
	 --url http://localhost:8085/api/v1/fetch \
```

### 🔐 Mint NFT

```bash
curl --request POST \
  --url http://localhost:8085/api/v1/mint \
  --header 'Content-Type: multipart/form-data' \
  --form name=MySecondNFT \
  --form 'description=This is my Second NFT' \
  --form image=@../frontend-react/public/koala/k5.png \
  --form rarity=2 \
  --form filename=demo.png \
  --form enviornment=dev
```
