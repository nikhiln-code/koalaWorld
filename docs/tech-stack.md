## 📦 Tech Stack

### Frontend (`/frontend-react`)

- **Framework**: [Next.js (App Router)](https://nextjs.org/docs/app)
- **Language**: TypeScript
- **UI**: Tailwind CSS
- **Web3**: Wagmi, RainbowKit
- **State Management**: React Context + Local Storage
- **Testing**: Playwright (planned), Vitest (for utilities)
- **Build Tool**: pnpm, Vite (used internally via Next)

### Backend (`/backend-go`)

- **Language**: Go (Golang)
- **Framework**: Gin
- **Project Structure**: Clean Architecture (`cmd`, `internal`,`config`,`tests`)
- **Database**: PostgreSQL (via GORM planned), SQLite (for local dev)
- **Blockchain Layer**: Polygon-compatible testnet mocks (actual integration planned)
- **Logging**: Uber Zap (configurable log levels)
- **Testing**: Go test + Test utils in `/tests/utils`
- **Auth**: JWT (Mock), Wallet signature (planned)
- **Environment**: Config-based (`config.dev.yaml`,`config.prod.yaml`,`config.test.yaml`)

Go back to [README](./../README.md)
