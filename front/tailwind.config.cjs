/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    container: {
      center: true,
    },
    extend: {
      colors: {
        primary: {
          DEFAULT: "#0078FF",
          50: "#E5F2FF",
          100: "#CCE4FF",
          200: "#99C9FF",
          300: "#66AEFF",
          400: "#3393FF",
          500: "#0078FF",
          600: "#0060CC",
          700: "#004899",
          800: "#003066",
          900: "#001833",
        },
        github: "#24292E",
        linkedin: "#0073B1",
      },
    },
  },
  plugins: [],
};
