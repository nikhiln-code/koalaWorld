"use client";

import "./globals.css";
import { ReactNode } from "react";
import { Inter, Geist } from "next/font/google";
import { Toaster } from "sonner";
import { WagmiProvider } from "wagmi"; // Import WagmiProvider
import { RainbowKitProvider, getDefaultConfig } from "@rainbow-me/rainbowkit"; // RainbowKit
import { polygonMumbai } from "wagmi/chains"; // Your chain (e.g., Polygon Mumbai)
import { http } from "wagmi";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"; // React Query
import Footer from "./components/Footer";
import Header from "./components/Header";
import BackgroundEffect from "./components/BackgroundEffect";

const geist = Inter({
  subsets: ["latin"],
});
const queryClient = new QueryClient();
const inter = Inter({ subsets: ["latin"] });
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
      <body className={geist.className}>
        <WagmiProvider config={config}>
          <QueryClientProvider client={queryClient}>
            <RainbowKitProvider>
              <div className="relative min-h-screen overflow-hidden text-white">
                {/* <div className="absolute inset-0 z-0 bg-[url('/jungle-bg.jpeg')] bg-center opacity-20" /> */}

                {/* Subtle animated gradient jungle theme */}
                {/* <div className="absolute inset-0 z-0 bg-gradient-to-br from-green-900 via-black to-green-800 animate-gradient bg-[length:400%_400%]" /> */}
                <BackgroundEffect />

                <div className="relative z-10 ">
                  <Header />
                  {children}
                  <Footer />
                </div>
              </div>
            </RainbowKitProvider>
          </QueryClientProvider>
        </WagmiProvider>
        <Toaster richColors position="top-center" />
      </body>
    </html>
  );
}
