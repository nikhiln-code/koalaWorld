### 🔧 Setup local development or for production refer [Docker & Deployment](./docs/docker.md)

We are using pinata to mint the NFTs and store them on IPFS.
Please refer to [Pinata](./docs/setup-pinata.md) to create and add the pinata api keys.

```bash
# Clone repo
git clone https://github.com/your-org/koala-world.git
cd koala-world
# -------------------------------------
# Backend
cd backend-go
# Set up env vars
cp  config/config.yaml_example config/config.dev.yaml

# Install deps
go mod tidy

# Run the app
go run cmd/main.go

#--------------------------------------
# Frontend
cd frontend-react
# Install deps
pnpm install
# Run in development mode
pnpm dev
```

Go back to [README](./../README.md)
