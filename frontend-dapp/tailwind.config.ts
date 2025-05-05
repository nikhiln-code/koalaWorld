import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        "bronze-light": "#f8c88d",
        "bronze-dark": "#b87333",
      },
      backgroundImage: {
        shine:
          "linear-gradient(90deg, rgba(255,255,255,0) 0%, rgba(255,255,255,0.5) 50%, rgba(255,255,255,0) 100%)",
      },
      animation: {
        shine: "shine 2s linear infinite",
      },
      keyframes: {
        shine: {
          "0%": { backgroundPosition: "-200%" },
          "100%": { backgroundPosition: "200%" },
        },
      },
    },
  },
  plugins: [
    function ({ addComponents }) {
      addComponents({
        ".text-shine": {
          position: "relative",
          overflow: "hidden",
          "@apply text-transparent bg-clip-text bg-gradient-to-r from-bronze-light to-bronze-dark":
            {},
        },
        ".text-shine::after": {
          content: '""',
          position: "absolute",
          inset: "0",
          backgroundImage:
            "linear-gradient(90deg, transparent 0%, white 50%, transparent 100%)",
          backgroundRepeat: "no-repeat",
          backgroundSize: "200% 100%",
          animation: "shine 2s linear infinite",
          zIndex: 0,
        },
        ".text-shine > *": {
          position: "relative",
          zIndex: 10,
        },
      });
    },
  ],
};

export default config;
