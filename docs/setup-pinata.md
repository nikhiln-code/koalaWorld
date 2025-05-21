## 🔐 IPFS Integration via Pinata

Koala World uses [Pinata](https://pinata.cloud/) to upload images and metadata for NFTs to IPFS.

### 🧾 How to Get API Keys

1. Sign up at [https://app.pinata.cloud/](https://app.pinata.cloud/)
2. Navigate to **API Keys**
3. Click **New Key** and give it the required permissions:
   - Pin Files
   - Read Pinned Files
   - Remove Files (optional)
4. Copy the **API Key**
5. Add them to your `config.{env}.yaml` file:

```config.{env}.yaml
pinata:
  jwt_key: "Bearer <YOUR_PINATA_JWT_KEY>"
```

Go back to [README](./../README.md)
