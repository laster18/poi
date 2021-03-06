module.exports = {
  purge: [
    './src/pages/**/*.{js,ts,jsx,tsx}',
    './src/components/**/*.{js,ts,jsx,tsx}',
  ],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
  },
  variants: {
    extend: {
      transitionDuration: ['hover', 'focus'],
      opacity: ['disabled'],
    },
    outline: ['responsive', 'focus', 'hover', 'active'],
    visibility: ['responsive', 'hover', 'focus', 'group-hover'],
  },
  plugins: [],
}
