// index.tsx
import { ConnectButton } from "@rainbow-me/rainbowkit";
import { useAccount } from "wagmi";

export default function Home() {
  const { address, isConnected } = useAccount();

  return (
    <main className="min-h-screen bg-black text-white p-10">
      <h1 className="text-2xl mb-4">RainbowKit + Wagmi Test</h1>
      <ConnectButton />
      {isConnected && <p className="mt-4">Connected: {address}</p>}
    </main>
  );
}
