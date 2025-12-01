

// Pub/Sub (Publish-Subscribe) Pattern
// A simple event system that decouples publishers and subscribers.
// Publishers emit events without knowing who's listening; subscribers
// register callbacks and get notified when events are published.

// npm install pino (optional logger; not used in this minimal example)

// The pubsub object is a simple event manager.
const pubsub = {
  // events: a registry of event names mapped to arrays of callbacks.
  // Example structure: { "message1": [callback1, callback2], "message2": [callback3] }
  events: {},

  // subscribe(event, callback): register a callback to be invoked when 'event' is published.
  // If this is the first subscription for 'event', initialize an empty array.
  subscribe(event, callback) {
    this.events[event] = this.events[event] || [];
    this.events[event].push(callback);
  },

  // publish(event, data): trigger all callbacks subscribed to 'event', passing 'data' to each.
  // If no subscriptions exist for 'event', the array is empty and no callbacks run.
  publish(event, data) {
    (this.events[event] || []).forEach((callback) => callback(data));
  },
};

// Register three separate callbacks for "message1".
// Each callback will be invoked when "message1" is published.
pubsub.subscribe("message1", (data) => console.log("Received mess1 call1:", data));
pubsub.subscribe("message1", (data) => console.log("Received mess1 call2:", data));
pubsub.subscribe("message1", (data) => console.log("Received mess1 call3:", data));

// Register two callbacks for "message2".
pubsub.subscribe("message2", (data) => console.log("Received mess2 call1:", data));
pubsub.subscribe("message2", (data) => console.log("Received mess2 call2:", data));

// Publish "message1" — all three "message1" callbacks will be invoked with "Hello, aaa".
// Expected output:
//   Received mess1 call1: Hello, aaa
//   Received mess1 call2: Hello, aaa
//   Received mess1 call3: Hello, aaa
pubsub.publish("message1", "Hello, aaa");

// Publish "message2" — both "message2" callbacks will be invoked with "Hello, bbb".
// Expected output:
//   Received mess2 call1: Hello, bbb
//   Received mess2 call2: Hello, bbb
pubsub.publish("message2", "Hello, bbb");

