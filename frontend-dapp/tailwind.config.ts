import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/app/**/*.{js,ts,jsx,tsx}",
    "./src/app/components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      keyframes: {
        shimmer: {
          "0%": { backgroundPosition: "-200% 0" },
          "100%": { backgroundPosition: "200% 0" },
        },
      },
      animation: {
        shimmer: "shimmer 2s linear infinite",
      },
      colors: {
        bronze: {
          light: "#cd7f32",
          DEFAULT: "#b87333",
          dark: "#8c5a26",
        },
      },
    },
  },
  plugins: [],
};

export default config;
