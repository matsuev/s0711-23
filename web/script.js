// Use WebSocket transport endpoint.
const centrifuge = new Centrifuge('ws://127.0.0.1:8000/connection/websocket', {
   data: {
      username: "root",
      password: "secret"
   }
});

// Allocate Subscription to a channel.
const sub = centrifuge.newSubscription('news');

// React on `news` channel real-time publications.
sub.on('publication', function (ctx)
{
   console.log(ctx.data);
});

// Trigger subscribe process.
sub.subscribe();

// Trigger actual connection establishement.
centrifuge.connect();