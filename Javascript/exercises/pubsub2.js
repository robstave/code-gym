// object literal with a map of events.

const pino = require("pino");
const logger = pino();

const pubSubSimple = {
  events: new Map(),

  subscribe(event, callback) {
    // add the callback to the map of events
    if (this.events.has(event)) {
      let callbacks = this.events.get(event);
      callbacks.push(callback);
      logger.info({ count: callbacks.length, event: event }, "callback added");
    } else {
      this.events.set(event, [callback]);
      logger.info({ event: event }, "first callback added");
    }
  },

  publish(event, data) {
    cbs = this.events.get(event);

    if (cbs != null) {
      logger.info({ count: cbs.length }, "publish");
      cbs.forEach((callback) => callback(data));
    } else {
      logger.error("opps");
    }
  },
};

pubSubSimple.subscribe("foo", (data) => {
  logger.info({ data: data }, "Function1");
});
pubSubSimple.subscribe("foo", (data) => {
  logger.info({ data: data }, "Function2");
});
pubSubSimple.subscribe("bizz", (data) => {
  logger.info({ data: data }, "Function3");
});
pubSubSimple.publish("foo", 436346);
pubSubSimple.publish("bizz", 4444);
pubSubSimple.publish("buzz", 1234);
