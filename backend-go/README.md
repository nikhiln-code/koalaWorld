## Backend Go

This is the backend project for the Koala world Game.
It hostst the services which are used by the Game UI.

# Game Assets:

1. NFTs -> Partially done
2. Game tokens -> to be done.

## It have Rest API's for the following services:

1. Getting the list of NFTs
2. Hosting the NFTs to the IPFS storage.
3. Minting the NFTs to the blockchain.

## Building the project .

### Prerequisites

1. Go 1.18+
2. Docker
3. Docker Compose

### Building the project

1. Clone the repository
2. Run `docker-compose up --build`
3. The APIs will be available at `http://localhost:8085`
