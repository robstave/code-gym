

// npm install pino

const pubsub = {
  events: {},
  subscribe(event, callback) {
    this.events[event] = this.events[event] || [];
    this.events[event].push(callback);
  },
  publish(event, data) {
    (this.events[event] || []).forEach((callback) => callback(data));
  },
};

pubsub.subscribe("message1", (data) => console.log("Received mess1 call1:", data));
pubsub.subscribe("message1", (data) => console.log("Received mess1 call2:", data));
pubsub.subscribe("message1", (data) => console.log("Received mess1 call3:", data));

pubsub.subscribe("message2", (data) => console.log("Received mess2 call1:", data));
pubsub.subscribe("message2", (data) => console.log("Received mess2 call2:", data));


pubsub.publish("message1", "Hello, aaa");
pubsub.publish("message2", "Hello, bbb");

