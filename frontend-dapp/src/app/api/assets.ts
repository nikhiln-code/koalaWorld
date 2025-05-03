import type { NextApiRequest, NextApiResponse } from "next";

// Fake in-memory asset store
let mockAssets = [
  {
    id: "1",
    name: "Sword of Truth",
    image: "https://via.placeholder.com/200x200?text=Sword",
    description: "Legendary sword",
    owner: "0x123",
  },
  {
    id: "2",
    name: "Shield of Valor",
    image: "https://via.placeholder.com/200x200?text=Shield",
    description: "Unbreakable shield",
    owner: "0xabc",
  },
];

export default function handler(req: NextApiRequest, res: NextApiResponse) {
  if (req.method === "GET") {
    const { owner } = req.query;

    if (owner) {
      const filtered = mockAssets.filter(
        (item) => item.owner.toLowerCase() === String(owner).toLowerCase()
      );
      return res.status(200).json(filtered);
    }

    return res.status(200).json(mockAssets);
  }

  if (req.method === "POST") {
    const { name, description, image, owner } = req.body;

    if (!name || !description || !image || !owner) {
      return res.status(400).json({ error: "Missing fields" });
    }

    const newItem = {
      id: String(Date.now()),
      name,
      description,
      image,
      owner,
    };

    mockAssets.push(newItem);
    return res.status(201).json(newItem);
  }

  return res.status(405).end();
}
