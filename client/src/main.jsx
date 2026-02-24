import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'

if('serviceWorker' in navigator) {
  navigator.serviceWorker.register('service-worker.js')
  //  .then(reg => console.log('SW registered', reg))
  //  .catch(err => console.error('SW failed', err));
}


createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
