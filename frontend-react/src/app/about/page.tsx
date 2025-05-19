"use client";

import { motion } from "framer-motion";

export default function AboutPage() {
  return (
    <main className="p-6 text-center min-w-full min-h-[70vh] flex flex-col items-center justify-center">
      {/* About Section as Framer Motion components */}
      <motion.section
        initial={{ opacity: 0, y: 100 }}
        whileInView={{ opacity: 1, y: 0 }}
        viewport={{ once: true, amount: 0.8 }}
        transition={{ duration: 1 }}
        className="relative bg-gradient-to-br from-green-900/80 via-green-950/80 to-green-900/80 border border-green-700 shadow-[0_0_40px_rgba(0,255,100,0.4)] backdrop-blur-lg rounded-3xl px-8 py-14 mx-auto mt-32 max-w-5xl text-center"
      >
        <motion.h2
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.2, duration: 0.6 }}
          className="text-4xl font-bold mb-4 text-green-300 drop-shadow-md tracking-wider"
        >
          🌿 About the Game
        </motion.h2>

        <motion.p
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.4, duration: 0.6 }}
          className="text-lg text-green-100 mb-6 max-w-xl mx-auto leading-relaxed"
        >
          This is a Web3-powered game where players can own, trade, and monetize
          in-game assets.
        </motion.p>

        <motion.p
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.4, duration: 0.6 }}
          className="text-lg text-green-100 mb-6 max-w-xl mx-auto leading-relaxed"
        >
          We use NFTs for asset ownership and a native token economy. Build on a
          low-cost blockchain for accessibility.{" "}
        </motion.p>
      </motion.section>
    </main>
  );
}
