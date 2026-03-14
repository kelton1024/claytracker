import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
    plugins: [
        tailwindcss(),
    ],
    server: {
        host: '0.0.0.0',
        proxy: {
            '/insert' : {
                target: 'http://claytracker_backend_1:8080',
                changeOrigin: true,
            }
        }
    }
});
