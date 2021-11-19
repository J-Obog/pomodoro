module.exports = {
  purge: ["./src/**/*.{js,jsx,ts,tsx}", "./public/index.html"],
  darkMode: false,
  theme: {
    extend: {
      fontFamily: {
        "space-mono": ['"Space Mono"', "monospace"],
        "roboto-cond": ['"Roboto Condensed"', "sans-serif"],
        quicksand: ['"Quicksand"', "sans-serif"],
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
