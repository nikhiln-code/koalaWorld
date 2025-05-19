"use client";

import { useState } from "react";
import { useAccount } from "wagmi";

export default function MintPage() {
  const { address } = useAccount();
  const [recipient, setRecipient] = useState(address || "");
  const [itemName, setItemName] = useState("");
  const [action, setAction] = useState<"mint" | "transfer" | null>(null);
  const [status, setStatus] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setStatus("Processing...");

    await new Promise((res) => setTimeout(res, 1200));

    setStatus(
      `✅ ${
        action === "mint" ? "Minted" : "Transferred"
      } "${itemName}" to ${recipient}`
    );
    setItemName("");
    setAction(null);

    setTimeout(() => setStatus(""), 4000); // Clear message after 4s
  };

  return (
    <main className="min-h-screen bg-gradient-to-b from-[#0d0d0d] to-[#1a1a1a] text-white p-8 flex flex-col items-center">
      <div className="max-w-2xl w-full bg-[#141414] border border-[#333] rounded-2xl p-8 shadow-xl">
        <h1 className="text-4xl font-extrabold mb-6 text-center">
          🛠️ Mint or Transfer Item
        </h1>

        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label className="block text-sm font-medium mb-1 text-gray-300">
              Recipient Wallet
            </label>
            <input
              type="text"
              value={recipient}
              onChange={(e) => setRecipient(e.target.value)}
              className="w-full px-4 py-2 bg-[#222] border border-gray-700 rounded-lg focus:outline-none focus:border-blue-500 transition"
              placeholder="0xABC123..."
              required
            />
          </div>

          <div>
            <label className="block text-sm font-medium mb-1 text-gray-300">
              Item Name
            </label>
            <input
              type="text"
              value={itemName}
              onChange={(e) => setItemName(e.target.value)}
              className="w-full px-4 py-2 bg-[#222] border border-gray-700 rounded-lg focus:outline-none focus:border-green-500 transition"
              placeholder="Ex: Shadow Blade"
              required
            />
          </div>

          {itemName && (
            <div className="bg-[#1e1e1e] border border-gray-700 rounded-lg p-4 shadow-inner text-sm text-gray-400">
              <p className="text-white font-semibold mb-1">Preview</p>
              <p>
                🧱 <strong>Item:</strong> {itemName}
              </p>
              <p>
                📤 <strong>To:</strong> {recipient}
              </p>
              <p>
                🔐 <strong>Action:</strong>{" "}
                {action ? action.toUpperCase() : "Select below"}
              </p>
            </div>
          )}

          <div className="flex justify-between">
            <button
              type="submit"
              onClick={() => setAction("mint")}
              className="w-[48%] bg-green-600 hover:bg-green-700 text-white py-2 px-4 rounded-lg font-bold transition"
            >
              Mint
            </button>
            <button
              type="submit"
              onClick={() => setAction("transfer")}
              className="w-[48%] bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded-lg font-bold transition"
            >
              Transfer
            </button>
          </div>
        </form>

        {status && (
          <div className="mt-6 p-4 rounded-lg bg-black border border-yellow-600 text-yellow-300 font-mono text-sm text-center animate-pulse">
            {status}
          </div>
        )}
      </div>
    </main>
  );
}
