import { NextResponse } from "next/server";

export async function GET() {
  const data = [
    {
      id: "Item-1",
      name: "Sword of Truth",
      //image: "/assets/sword.png",
      image: "https://via.placeholder.com/150",
      rarity: "Legendary",
      owner: "0x123...abc",
    },
    {
      id: "item-2",
      name: "Shield of Light",
      //image: "/assets/shield.png",
      image: "https://via.placeholder.com/150",
      rarity: "Epic",
      owner: "0x456...def",
    },
    {
      id: "item-3",
      name: "Potion of Speed",
      //image: "/assets/potion.png",
      image: "https://via.placeholder.com/150",
      rarity: "Common",
      owner: "0x789...ghi",
    },
  ];

  return NextResponse.json(data);
}
