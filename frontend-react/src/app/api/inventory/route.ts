import { NextResponse } from "next/server";

export async function fetchInventory() {
  const res = await fetch("http://localhost:8080/api/inventory");
  if (!res.ok) throw new Error("Failed to fetch inventory");
  return res.json();
}

export async function mintAsset(data: { name: string; tyoe: string }) {
  const res = await fetch("http://localhost:8080/api/mint", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });

  if (!res.ok) throw new Error("Failed to mint asset");
  return res.json();
}

export async function transferAsset(data: { to: string; assetId: string }) {
  const res = await fetch("http://localhost:8080/api/transfer", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Failed to transfer asset");
  return res.json();
}
