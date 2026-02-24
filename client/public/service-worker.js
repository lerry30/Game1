self.addEventListener('install', (event) => {
  self.skipWaiting();
});

self.addEventListener('activate', (event) => {
  event.waitUntil(clients.claim()); // take control of existing pages NOW
});

self.addEventListener('fetch', (event) => {
  try {
    const url = new URL(event.request.url);
    //console.log('SW caught:', url.pathname); // add this first to confirm it's intercepting
    
    if (url.pathname.includes('/assets')) {
      const newUrl = `http://localhost:8080/game/game1${url.pathname}`;
      //console.log('Redirecting to ', newUrl);

      event.respondWith(
        fetch(newUrl).then(res => {
          if (!res.ok) throw new Error(`Failed: ${res.status}`);
          return res;
        }).catch(err => {
          console.error('Fetch failed:', err);
          return new Response('Asset not found', { status: 404 });
        })
      );
    }
  } catch(error) {
    console.log('SW Error: ', error);
  }
});