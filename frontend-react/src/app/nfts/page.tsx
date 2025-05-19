"use client";

import { useEffect, useState } from "react";
import { fetchInventory } from "../api/inventory/route";

type Item = {
  id: string;
  name: string;
  image: string;
  rarity: string;
  owner: string;
};

export default function InventoryPage() {
  const [items, setItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchInventory().then((data) => {
      setItems(data);
      setLoading(false);
    });
  }, []);

  return (
    <main className="p-6 max-w-4xl min-h-screen text-white">
      <h1 className="text-3xl font-bold mb-6">🧰 Game Inventory</h1>

      {loading ? (
        <p>Loading...</p>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
          {items.map((item) => (
            <div
              key={item.id}
              className="bg-[#1a1a1a] border border-gray-700 rounded-xl p-4 shadow-lg hover:shadow-xl transition"
            >
              <img
                src={item.image}
                alt={item.name}
                className="w-full h-40 object-cover rounded mb-3"
              />
              <h2 className="text-xl font-semibold">{item.name}</h2>
              <p className="text-sm text-gray-400">Rarity: {item.rarity}</p>
              <p className="text-xs text-gray-500">Owner: {item.owner}</p>
            </div>
          ))}
        </div>
      )}
    </main>
  );
}
