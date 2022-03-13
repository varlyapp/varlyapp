const path = require('path')
const mix = require('laravel-mix')

require('laravel-mix-tailwind')

// Set aliases
mix.webpackConfig({
  resolve: {
    alias: {
      '@root': path.resolve(__dirname, 'src/js'),
      '@assets': path.resolve(__dirname, 'src/assets'),
      '@layouts': path.resolve(__dirname, 'src/js/layouts'),
      '@components': path.resolve(__dirname, 'src/js/components'),
      '@pages': path.resolve(__dirname, 'src/js/pages'),
      '@utils': path.resolve(__dirname, 'src/js/utils'),
    }
  }
})

// Disable notifications
mix.disableNotifications()

// JS/SCSS compilation
mix.setPublicPath('src')
  .js('src/js/app.js', 'app.js').vue({ version: 3 })
  .sass('src/css/app.scss', 'app.css')
  .tailwind()
