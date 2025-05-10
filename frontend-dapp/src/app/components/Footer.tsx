import Link from "next/link";

export default function Footer() {
  return (
    <footer className="bg-gradient-to-r from-green-900 via-black to-green-900 text-green-300 py-6 px-4 mt-12 shadow-inner border-t border-green-700">
      <div className="max-w-7xl mx-auto flex flex-col md:flex-row justify-between items-center space-y-4 md:space-y-0 text-sm">
        <div className="flex flex-col items-center md:items-start">
          <p className="font-semibold tracking-wide">
            © {new Date().getFullYear()} KoalaWorld
          </p>
          <p className="text-lime-400">A Web3 Gaming Experience</p>
        </div>
        <div className="flex space-x-6">
          <Link
            href="/about"
            className="hover:text-white transition-transform transform hover:scale-105"
          >
            About
          </Link>
          <Link
            href="/nfts"
            className="hover:text-white transition-transform transform hover:scale-105"
          >
            NFTs
          </Link>
          <Link
            href="/marketplace"
            className="hover:text-white transition-transform transform hover:scale-105"
          >
            Marketplace
          </Link>
        </div>
      </div>
      <div className="text-center text-l font-sans text-lime-500 mt-6">
        Made with ❤️ by Nikhil
      </div>
    </footer>
  );
}
