FROM node:23.9.0-slim AS builder

# Install pnpm
RUN npm install -g pnpm@10.7.0

WORKDIR /app

# Only copy necessary files for installing dependencies
COPY package.json pnpm-lock.yaml ./

# Install dependencies (only prod if NODE_ENV=production)
RUN pnpm install --frozen-lockfile

# Copy the rest of the code
COPY . .

# Build the Next.js app
RUN pnpm build

# -----------------------------------

# Stage 2: Runtime
FROM node:18-slim AS runner

ENV NODE_ENV=production

WORKDIR /app

# Install pnpm
RUN npm install -g pnpm@10.7.0

# Only copy what's needed to run the app
COPY --from=builder /app/.next .next
COPY --from=builder /app/public public
COPY --from=builder /app/package.json .
COPY --from=builder /app/pnpm-lock.yaml .

# Install only prod dependencies
RUN pnpm install --prod --frozen-lockfile

EXPOSE 3000
CMD ["pnpm", "dev"]