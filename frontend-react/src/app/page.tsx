"use client";

import { ConnectButton, WalletButton } from "@rainbow-me/rainbowkit";
import { useAccount } from "wagmi";

import ConnectWallet from "./components/ConnectWallet";
import Footer from "./components/Footer";
import { motion } from "framer-motion";
import Button from "./components/Button";

export default function HomePage() {
  const { isConnected, address } = useAccount();

  return (
    <main className="p-6 text-center min-w-full min-h-[150vh] flex flex-col items-center justify-center">
      <h1 className="text-4xl font-bold mb-4">
        Welcome to the KoalaWorld , Web3 Games{" "}
      </h1>
      <p className="text-lg mb-8">
        Own your in-game assets with NFTs and tokens
      </p>
      <a
        href="/nfts"
        className="px-6 py-3 bg-gradient-to-r from-green-500 via-lime-600 to-green-700 text-white font-bold rounded-xl shadow-lg hover:scale-105 transition duration-300"
      >
        View NFTs
      </a>

      {/* Waitlist Section - Show only when scrolled to bottom */}
      <motion.section
        initial={{ opacity: 0, y: 100 }}
        whileInView={{ opacity: 1, y: 0 }}
        viewport={{ once: true, amount: 0.8 }}
        transition={{ duration: 1 }}
        className="relative bg-gradient-to-br from-green-900/80 via-green-950/80 to-green-900/80 border border-green-700 shadow-[0_0_40px_rgba(0,255,100,0.4)] backdrop-blur-lg rounded-3xl px-8 py-14 mx-auto mt-50 max-w-5xl text-center"
      >
        <motion.h2
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.2, duration: 0.6 }}
          className="text-4xl font-bold mb-4 text-green-300 drop-shadow-md tracking-wider"
        >
          🌿 Join the Waitlist
        </motion.h2>
        <motion.p
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.4, duration: 0.6 }}
          className="text-lg text-green-100 mb-6 max-w-xl mx-auto leading-relaxed"
        >
          Be the first to explore the KoalaVerse. Get exclusive access to drops,
          in-game perks, and NFT launches.
        </motion.p>
        <motion.form
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.6, duration: 0.6 }}
          className="flex flex-col sm:flex-row justify-center items-center gap-4"
        >
          <input
            type="email"
            placeholder="Enter your email"
            className="px-5 py-3 rounded-xl text-white font-medium focus:outline-none focus:ring-2 focus:ring-green-500 w-full sm:w-72"
          />
          <button
            type="submit"
            className="px-6 py-3 bg-gradient-to-r from-green-500 via-lime-600 to-green-700 text-white font-bold rounded-xl shadow-lg hover:scale-105 transition duration-300"
          >
            Notify Me
          </button>
        </motion.form>
      </motion.section>
    </main>
  );
}
