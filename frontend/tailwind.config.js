module.exports = {
  mode: 'jit',
  content: [
    './src/index.html',
    './src/css/app.scss',
    './src/js/**/*.{vue,js}'
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms')
  ],
}
