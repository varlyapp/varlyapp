import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    resolve: {
        alias: [
            { find: '@', replacement: '/src/js' },
            { find: '@wails', replacement: '/wailsjs' },
            { find: '@assets', replacement: '/src/assets' }
        ]
    },
    // build: {
    //     minify: false,
    //     assetsInlineLimit: 0,
    //     rollupOptions: {
    //         output: {
    //             entryFileNames: `assets/[name].js`,
    //             chunkFileNames: `assets/[name].js`,
    //             assetFileNames: `assets/[name].[ext]`
    //         }
    //     }
    // }
})
