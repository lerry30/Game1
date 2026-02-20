import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  plugins: [react(), tailwindcss()],
   server: {
    proxy: {
      '/assets': {
        target: 'http://localhost:8080/game/game1',
        changeOrigin: true,
      }
    }
  }
})