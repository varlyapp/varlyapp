const path = require('path')
const mix = require('laravel-mix')

require('laravel-mix-tailwind')

// Set aliases
// Remember to update paths in tsconfig.json
mix.webpackConfig({
  resolve: {
    alias: {
      '@root': path.resolve(__dirname, 'src/js'),
      '@assets': path.resolve(__dirname, 'src/assets'),
      '@layouts': path.resolve(__dirname, 'src/js/layouts'),
      '@components': path.resolve(__dirname, 'src/js/components'),
      '@pages': path.resolve(__dirname, 'src/js/pages'),
      '@utils': path.resolve(__dirname, 'src/js/utils'),
      '@plugins': path.resolve(__dirname, 'src/js/plugins'),
      "@wails": path.resolve(__dirname, 'wailsjs'),
    }
  }
})

// Disable notifications
mix.disableNotifications()
mix.webpackConfig({
  watchOptions: {
    aggregateTimeout: 1000,
    ignored: ['**/node_modules', '**/wailsjs']
  }
})
// TS/SCSS compilation
mix.setPublicPath('src')
  .sourceMaps()
  .ts('src/js/app.ts', 'app.js').vue({ version: 3 })
  .sass('src/css/app.scss', 'app.css')
  .tailwind()
