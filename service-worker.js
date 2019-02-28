self.addEventListener('install', function (event) {
    // Put `offline.html` page into cache
    var offlineRequest = new Request('offline.html');
    event.waitUntil(
        fetch(offlineRequest).then(function (response) {
            return caches.open('offline').then(function (cache) {
                return cache.put(offlineRequest, response);
            });
        })
    );
});

self.addEventListener('fetch', function (event) {
    if (event.request.method === 'GET') {
        console.log('Handling fetch event for', event.request.url);
        event.respondWith(
            fetch(event.request).catch(error => {
                console.log('Fetch failed; returning offline page instead.', error);
                return caches.match('offline.html');
            })
        );
    }
});

self.addEventListener('activate', function (event) {
    // Calling claim() to force a "controllerchange" event on navigator.serviceWorker
    event.waitUntil(self.clients.claim());
});