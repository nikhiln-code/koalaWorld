"use client";

import "./globals.css";
import { ReactNode } from "react";
import { Toaster } from "sonner";
import { WagmiProvider } from "wagmi"; // Import WagmiProvider
import { RainbowKitProvider, getDefaultConfig } from "@rainbow-me/rainbowkit"; // RainbowKit
import { polygonMumbai } from "wagmi/chains"; // Your chain (e.g., Polygon Mumbai)
import { http } from "wagmi";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"; // React Query

const queryClient = new QueryClient();

// Create Wagmi config
const config = getDefaultConfig({
  appName: "Game DApp",
  projectId: "YOUR_PROJECT_ID", // Replace with your WalletConnect project ID
  chains: [polygonMumbai],
  transports: {
    [polygonMumbai.id]: http(), // Connect to Polygon Mumbai RPC
  },
  ssr: true,
});

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <head />
      <body>
        <WagmiProvider config={config}>
          <QueryClientProvider client={queryClient}>
            <RainbowKitProvider>{children}</RainbowKitProvider>
          </QueryClientProvider>
        </WagmiProvider>
        <Toaster richColors position="top-center" />
      </body>
    </html>
  );
}
