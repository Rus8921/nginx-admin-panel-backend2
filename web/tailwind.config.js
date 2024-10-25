/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    screens: {
      sm: '480px',
      md: '768px',
      lg: '976px',
      xl: '1440px',
    },
    colors: {
      'bg-clr': '#f9f9f9',
      'main-clr': '#32ade6',
      'scndry-clr': '#1a6d94',
      'txt-clr': '#000000',
      'scndry-txt-clr': '#b1b1b1',
      'highlight': '#eef9fe',
      'scndry-highlight': '#f3f3f3',
      'white': '#ffffff',
      'black': '#000000',
      'green': '#5ce06c',
      'red': '#cc332e',
    },
    fontFamily: {
      roboto: ["-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif"],
      robotoLogo: ["-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif",
        {
          fontFeatureSettings: '"ss01"',
        }],
      mono: ["source-code-pro, Menlo, Monaco, Consolas, 'Courier New', monospace"],
    },
    fontSize: {
      'xl': '3.75rem',
      'lg': '2.25rem',
      'md': '1.5rem',
      'sm': '1.25rem',
    },
    letterSpacing: {
      tighter: '-.08em',
      tight: '-.04em',
      normal: '.07em',
    },
    extend: {
      boxShadow: {
        'r-xs': '30px 0 30px -30px #b1b1b133',
        'b-xs': '0 30px 30px -30px #b1b1b133',
        '2xs': '0 0 30px 0 #b1b1b11A',
      },
      gridTemplateColumns: {
        'adaptive-cards': 'repeat(auto-fill, minmax(360px, 1fr))',
      }
    },
  },
  plugins: [],
};
