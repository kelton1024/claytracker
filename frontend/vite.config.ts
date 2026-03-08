import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
    plugins: [
        tailwindcss(),
    ],
    server: {
        proxy: {
            '/insert' : {
                target: 'http://127.0.0.1:8080',
                changeOrigin: true,
            }
        }
    }
});