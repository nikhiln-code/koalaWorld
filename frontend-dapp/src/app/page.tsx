"use client";

import { ConnectButton } from "@rainbow-me/rainbowkit";
import { useAccount } from "wagmi";
import Footer from "./components/Footer";

export default function HomePage() {
  const { isConnected, address } = useAccount();

  return (
    <>
      <div className="flex flex-col min-h-screen">
        <main className=" bg-black text-white p-10 text-center">
          <h1 className="text-5xl md:text-6xl font-bruce text-gradient-bronze shine-effect drop-shadow-glow mb-6 capitalize">
            Welcome to KoalaWorld<sup>&reg;</sup>
          </h1>
          <h1 className="text-5xl text-shine">Welcome to KoalaWorld</h1>
          {isConnected && <p className="mt-4">Connected: {address}</p>}
        </main>
        <Footer />
      </div>
    </>
  );
}
