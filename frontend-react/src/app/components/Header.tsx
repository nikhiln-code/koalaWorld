"use client";

import Link from "next/link";
import ConnectWallet from "./ConnectWallet";
import { useAccount } from "wagmi";
import { useState } from "react";
import { Bars3Icon, XMarkIcon } from "@heroicons/react/24/outline";

export default function Header() {
  const { isConnected, address } = useAccount();
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  return (
    <header className="p-4 border-b  border-lime-200 shadow-lg relative bg-gradient-to-r from-green-900 via-black to-green-900 text-white font-sans tracking-wide">
      <div className="flex flex-col md:flex-row items-center justify-between max-w-7xl mx-auto gap-4">
        <div className="text-3xl md:text-5xl font-bold text-lime-500 animate-pulse drop-shadow-glow-green capitalize font-bruce">
          Welcome to 🎋 KoalaWorld<sup className="text-3xl">®</sup>
        </div>

        <div className="flex items-center gap-4 md:ml-auto">
          {/* Wallet Connect */}
          <div className="relative">
            <div className="bg-gradient-to-tl from-green-600 via-lime-500 to-green-700 p-1 rounded-full shadow-lg hover:scale-110 transition-transform">
              <div className="bg-black rounded-full px-4 py-2 text-white font-semibold text-sm">
                <ConnectWallet />
              </div>
            </div>
            {isConnected && (
              <p className="mt-2 text-sm text-lime-400 animate-fade-in text-center">
                Connected: {address}
              </p>
            )}
          </div>

          {/* Menu Toggle */}
          <button
            className="z-50 hover:scale-110 transition-transform hover:text-lime-400"
            onClick={() => setIsMenuOpen(!isMenuOpen)}
          >
            {isMenuOpen ? (
              <XMarkIcon className="h-8 w-8 text-lime-400" />
            ) : (
              <Bars3Icon className="h-8 w-8 text-green-300" />
            )}
          </button>
        </div>
      </div>

      {/* Toggleable nav (for all screens) */}
      {isMenuOpen && (
        <nav className="absolute top-full left-0 w-full bg-gradient-to-br from-black via-green-950 to-black shadow-2xl py-8 px-4 flex flex-col items-center space-y-6 z-40 backdrop-blur-md border-t border-green-700 animate-slide-down rounded-b-xl">
          <Link
            href="/"
            className="text-xl font-bold text-green-300 hover:text-white transition duration-300 transform hover:scale-105 hover:underline"
            onClick={() => setIsMenuOpen(false)}
          >
            🏠 Home
          </Link>
          <Link
            href="/about"
            className="text-xl font-bold text-green-300 hover:text-white transition duration-300 transform hover:scale-105 hover:underline"
            onClick={() => setIsMenuOpen(false)}
          >
            📜 About
          </Link>
          <Link
            href="/nfts"
            className="text-xl font-bold text-green-300 hover:text-white transition duration-300 transform hover:scale-105 hover:underline"
            onClick={() => setIsMenuOpen(false)}
          >
            🐨 NFTs
          </Link>
          <Link
            href="/marketplace"
            className="text-xl font-bold text-green-300 hover:text-white transition duration-300 transform hover:scale-105 hover:underline"
            onClick={() => setIsMenuOpen(false)}
          >
            🛒 Marketplace
          </Link>
          <div className="mt-4">
            <div className="bg-gradient-to-br from-green-600 via-lime-500 to-green-700 p-1 rounded-full shadow-lg hover:scale-105 transition-transform">
              <div className="bg-black rounded-full px-4 py-2 text-white font-semibold text-sm">
                <ConnectWallet />
              </div>
            </div>
          </div>
          {isConnected && (
            <p className="mt-2 text-sm text-lime-400 animate-fade-in">
              Connected: {address}
            </p>
          )}
        </nav>
      )}
    </header>
  );
}
