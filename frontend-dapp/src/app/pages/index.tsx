import Link from "next/link";

export default function Home() {
  return (
    <main className="flex flex-col items-center justify-center h-screen text-center bg-gradient-to-br from-black via-gray-900 to-gray-800 text-white px-6">
      <h1 className="text-5xl font-bold mb-4">Game Asset DApp</h1>
      <p className="mb-8 text-gray-400 max-w-md">
        Mint, manage, and trade your game assets on-chain. Powered by Web3 and
        Go backend.
      </p>
      <div className="flex gap-4">
        <Link
          href="/dashboard/inventory"
          className="bg-purple-600 px-6 py-2 rounded-xl hover:bg-purple-700 transition"
        >
          Enter App
        </Link>
        <Link
          href="/dashboard/mint"
          className="border border-purple-400 px-6 py-2 rounded-xl hover:bg-purple-500 hover:text-white transition"
        >
          Mint Asset
        </Link>
      </div>
    </main>
  );
}
