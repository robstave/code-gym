// Base component
class Coffee {
  cost() {
    return 5;
  }

  description() {
    return "Plain coffee";
  }
}

// Decorator base class
class CoffeeDecorator {
  constructor(coffee) {
    this.coffee = coffee;
  }

  cost() {
    return this.coffee.cost();
  }

  description() {
    return this.coffee.description();
  }
}

// Concrete decorators
class MilkDecorator extends CoffeeDecorator {
  cost() {
    return super.cost() + 1.5;
  }

  description() {
    return super.description() + ", milk";
  }
}

class SugarDecorator extends CoffeeDecorator {
  cost() {
    return super.cost() + 0.5;
  }

  description() {
    return super.description() + ", sugar";
  }
}

// Usage
let myCoffee = new Coffee();
console.log(myCoffee.description(), "$" + myCoffee.cost());

myCoffee = new MilkDecorator(myCoffee);
console.log(myCoffee.description(), "$" + myCoffee.cost());

myCoffee = new SugarDecorator(myCoffee);
console.log(myCoffee.description(), "$" + myCoffee.cost());
