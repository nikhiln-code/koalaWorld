export async function fetchAssets(owner?: string) {
  const res = await fetch(`/api/assets${owner ? `?owner=${owner}` : ""}`);
  if (!res.ok) throw new Error("Failed to fetch assets");
  return res.json();
}

export async function mintAsset(data: {
  name: string;
  description: string;
  image: string;
  owner: string;
}) {
  const res = await fetch("/api/assets", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });

  if (!res.ok) throw new Error("Minting failed");
  return res.json();
}
