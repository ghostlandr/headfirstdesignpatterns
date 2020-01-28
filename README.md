Head First Design Patterns (practice)
=====================================

[Head First Design Patterns](http://shop.oreilly.com/product/9780596007126.do) is a book by Eric Freeman and
Elisabeth Robson which is written with Java developers in mind - the code is all Java (and the book is probably
intended to help people get better at Java). While I am not a Java developer, design patterns are a useful thing
to know, and the conversational nature of the book is great for retention, I've learned a lot so far!

This repository contains my [Go](https://golang.org/) interpretations of each chapter of examples - so far, anyway.
I'm sure eventually there will be a chapter that would require me to make so much Go boilerplate (as I lack the
ability to subclass functionality) that it may not be productive to write it all out. Even in cases like those,
however, I'm hoping that there will be value in absorbing the lesson on that particular pattern, whichever one it
ends up being.

## Patterns thus far:

### Strategy pattern

This pattern used ducks, which was enjoyable for me on a personal level - my favourite bird for sure, maybe my
favourite animal in general is the mallard duck. The idea with the strategy pattern is that rather than make a
subclass for every type and configuration of object, instead provide different strategies that a smaller set of
objects can interact with. So for ducks, there are strategies for making sound and for flying. In my go code I
ended up making a Flyer and a Quacker interface (verbs in the classic go style). A Duck then is a struct made up
of a combination of the two strategies. Like this:

```go
type Flyer interface {
	Fly()
}

type Quacker interface {
	Quack()
}

type Duck struct {
    fly Flyer
    quack Quacker
}

func NewDuck(f Flyer, q Quacker) {
    return Duck{fly: f, quack: q}
}
```

So then some example uses of strategies:

```go
mallard := NewDuck(NewItQuacks(), NewFlyWithWings())
mallard.Quack()
mallard.Fly()
```

In my case those methods just print out whatever the strategy enables, so it would print something about quacking
and flying respectively. But for a rubber duck, it would use the `NewFlyNoWings()` (that's what they had in the book,
ha). Thus when trying to make it quack and fly, it would print about quacking, but say it can't fly. Contrived
but effective for learning, I thought. I can see some applications in my own code for this one.

### Observer pattern

This is a pattern I am much more familiar with, as I have a stronger frontend skill set still. In this chapter we
implemented a weather data collection station. It starts out as a very imperative, "This thing updated, manually
update all of the classes we know are depending on us". One of the places they suggest you try and apply design
patterns is any spots in the code that are likely to change frequently. In our weather station example, it's
laid out like this: there is a weather data class that "gets" updated weather data intermittently. When it does
so, it updates three (for now) weather displays - this is our client code for this chapter. One of the principles
they talk about in chapter 3 (I think) is the Open-Closed principle: Our code should be open for extension, but
closed for modification. (This is a principle that I've longed to understand, and now it finally feels like I do
after reading a few of these chapters!) In our weather station as it stands today, every time we add a new
weather display we will need to modify the weather data station code! Not very "closed" at all. Thus we enter
the Observer pattern. Rather than adding code to our "new weather data" function in the weather data station, we
give displays the ability to subscribe to updates. Therefore if we want to add a new weather display at some point
in the future we don't have to modify the weather data station code, we just subscribe through the interface we
already have available. Here's how I did that in go:

```go
type Observer interface {
	Update(temp, humidity, pressure float64)
}

type Observable interface {
	RegisterSubscriber(o Observer)
	RemoveSubscriber(toRemove Observer)
	NotifySubscribers()
}

type WeatherData struct {
	temp, humidity, pressure float64
	observers []Observer
}
```

The register, remove, and notify functions all work like you'd expect: either append to the slice, remove from
the slice, or iterate the slice. Then it's a simple matter of calling NotifySubscribers() whenever you "get a new
update" in the WeatherData implementation.

A real life example of the observer pattern is one I used often when working in Javascript - `addEventListener`.
You add a callback that should be run on a particular event on pretty well anything in Javascript (classic), but
most often DOM nodes. Something like `myButton.addEventListener('onclick', function() { console.log('hello!') })`
where myButton is a DOM node you'd have to grab earlier.

This chapter was a good one for boosting my confidence in reading the book as I had it pretty well understood.

### Decorator pattern

Going into this chapter, I assumed we'd be using the @ sign a lot, as I'm used to that from Typescript and the like.
However, there were no @ signs at all in this chapter. This was also a relief because go doesn't have such things,
so it would have been interesting to try and figure out how to work around that.

This pattern centered around a coffee shop example. You start out with a `Beverage` class which exposes a cost
method that subclasses override to provide their cost. The coffee shop has four main beverages: dark roast,
house blend, decaf, and espresso. Then they have the usual suspects available as toppings/condiments: soy milk,
whip, mocha, and steamed milk. Initially they created a subclass for every combination of beverage, and you can
imagine the carnage this creates: `DarkRoast`, `DarkRoastMocha`, `DarkRoastMochaWhip`, `DarkRoastSoy`, ... and
on and on. Dozens of classes. And what about things like double mocha? `DarkRoastMochaMocha`? This is getting
crazy!

Enter: the decorator pattern. The other principle this chapter introduces is "programming to an interface, not
a concrete implementation". If we base our implementation off of a Beverage interface, such as this one:

```go
type Beverage interface {
    Cost() float64
    Description() string
}
```

We can create a system of decorators that take Beverages and return Beverages, and we can "decorate" (or I would
probably lean more to calling it "wrapping" but then we don't have the decorator pattern, we have the wrapper
pattern) our beverages as much as we like. Here are a few examples:

```go
e := beverage.Espresso()
e = beverage.Mocha(e)
e = beverage.Whip(e)
fmt.Printf("%s: $%.2f\n", e.Description(), e.Cost()) // output: "Espresso, Mocha, Whip: $2.29"

// You can inline them as well - a tad more gross though.
e2 := beverage.Whip(beverage.Mocha(beverage.Espresso()))
```

If you were curious, here is some of my implementation of these. It's a bit more annoying than in a true OO
language with subclassing and such, but really, I could see myself using something like this in real life.

```go
func Espresso() Beverage {
	return &espresso{}
}

type espresso struct{
	size Size
}

func (e espresso) Description() string {
	return "Espresso"
}

func (e espresso) Cost() float64 {
	return 1.99
}

func Mocha(b Beverage) Beverage {
	return mocha{Beverage: b}
}

type mocha struct{
	Beverage
}

func (m mocha) Description() string {
	return m.Beverage.Description() + ", Mocha"
}

func (m mocha) Cost() float64 {
	return m.Beverage.Cost() + .20
}
```

The Description and Cost functions make use of polymorphism (does that still exist in Go? I guess so) to pass on
their calculations. In some languages (Python is the one that comes to mind) you can create an actual "decorator"
that you add to the top of a function like this

```python
@mydecorator
def my_function():
```

and by doing so, you gain access to a few lifecycle hooks for the function and you can run your decorator before,
during(?), or after the function you're wrapping is called. In our code's case I guess we're running our code
after, since we call the wrapped function and then add our own text.

One of the things they admit to in this chapter is that this _is_ a bit cumbersome. They say this will be helped
a bit when we start learning the Factory pattern in the next chapter. In my mind the main takeaway from this
chapter is that you should program to an interface, not a concrete implementation. In go-land, this means you
should program to an interface, not a concrete struct. This enables you to do all kinds of things (dependency
injection and testing come to mind first), but I hadn't really grasped the real "why" until reading this chapter
and the previous one where they start to introduce this concept.

### Factory pattern

The biggest takeaway I took from this chapter was the Dependency Inversion Principle, which is a fancy name for
this principle: "Depend on abstractions. Do not depend on concrete classes." The inversion in the name comes from
the look of your class diagram after you've done some work to implement it. Rather than your classes flowing
down into each other, as you might normally see in inheritance, you've inverted the flow. Your high-level and
low-level modules are both depending on the same abstraction.

How does this come into play with factories? Let's use the example from the chapter: Pizza stores.

(side note: I quite enjoy the examples used in this book. I'm sure they are chosen for this very reason, but I
find them to be very relatable and reasonable. Like we're talking about "the factory pattern" but through the
lens of a pizza store and franchises, so of course it makes sense to make sure things are created the same way
every time and things like that, so providing a way for our franchisees to do that is totally necessary.)

We start out with code we've probably all written before. We have one pizza store, with a order pizza method
and it all looks like this:

```go
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type PizzaStore interface {
    OrderPizza(pizzaType string) Pizza
}

func OrderPizza(pizzaType string) Pizza {
    var p Pizza
    switch pizzaType {
    case "cheese":
        p = NewCheesePizza()
    case "pepperoni":
        p = NewPepperoniPizza()
    // ... on and on
    }

    p.Prepare()
    p.Bake()
    p.Cut()
    p.Box()
    return p
}

func main() {
    p := PizzaStore{}
    pizza := p.OrderPizza("cheese")
    // weew!
}
```

This is fine for our one store, and at least we do have some amount of abstraction here: we're using polymorphism
to treat every type of pizza the same once we get down to the prepare and ship step. However you can see that this
quickly grows out of control once we add more pizzas and more regions to our franchise. As well, you can see we're
using a method like "NewWhateverTypeOfPizza", which means we're now coding to a concrete type rather than an interface,
and it also means whenever we need to add a new kind of pizza we need to come into this code and modify it to add
that in, which violates the Open-Closed principle (open to extension, closed to modification). At the moment it is
very much open to modification - although, if Pizza functionality changed that part would be okay!

How do we improve this? I hope you said "Factories!" because if you did, you nailed it. Let's do that:

```go
type PizzaStore interface {
    CreatePizza(pizzaType string) Pizza
    OrderPizza(pizzaType string) Pizza
}

type PizzaFactory interface {
    CreatePizza(pizzaType string) Pizza
}

type pizzaFactory struct {}

func (p *pizzaFactory) CreatePizza(pizzaType string) Pizza {
    switch pizzaType {
    case "cheese":
        return NewCheesePizza()
    case "pepperoni":
        return NewPepperoniPizza()
    // ... on and on
    }
    // Error handling down here in case they gave a weird pizza type
}

func NewPizzaStore(factory PizzaFactory) PizzaStore {
    // Embed the factory in our concrete pizza store
    return &pizzaStore{PizzaFactory: factory}
}

type pizzaStore struct {
    PizzaFactory
}

func (p *pizzaStore) OrderPizza(pizzaType) Pizza {
    var p Pizza
    // Use the embedded factory function here to handle creation
    p = p.CreatePizza(pizzaType)
    p.Prepare()
    p.Bake()
    p.Cut()
    p.Box()
    return p
}
```

Why is this better? Well, now if we want to add a new pizza type we can just add it into CreatePizza. No need to
modify OrderPizza, which now follows the Open-Closed principle - it's open to extension by adding new pizzas to
CreatePizza, but closed to modification as we don't need to do that to add new pizzas.

The next part they go into in the book relies heavily on inheritance. The method above, passing a factory into
the constructor, is known as the Simple Factory. Here's what they have laid out for their Java application:

```java
public abstract class PizzaStore {
    public Pizza orderPizza(String type) {
        Pizza pizza;
        pizza = createPizza(type);
        pizza.Prepare()
        pizza.Bake()
        pizza.Cut()
        pizza.Box()
        return pizza
    }

    abstract Pizza createPizza(String type);
}
```

Now what you would do is subclass PizzaStore and override createPizza to do your pizza making. This lets us
nail down the orderPizza method so that it can't be tampered with, but allows us to put the power in the hands of
the pizza store to create pizzas however they like. In the book they talk about New York style pizza versus Chicago
style, two very different types of pizza to say the least. So you'd have a ChicagoPizzaStore creating Chicago-style,
and a NYPizzaStore creating New York-style. Life is good. In Go-land however, all we can really share is the
interface for making pizzas. We can't provide a piece of functionality for classes to embed, or at least we can't
prevent them from doing their own thing there, at least not without making some of the structs or interfaces
internal. I could see if you had a `pizza` (unexported) type and a `orderPizza` method that talked in lower-case
p pizzas, maybe you could protect that process?

The thing we've specced out in the Java code is known as a Factory Method. Responsibility for what it creates
is left up to subclasses of the main class. You've managed to decouple the creations from the creator, as any
changes we make to Pizza won't affect the creating classes, so long as the interface doesn't change.

At the end of the day I think in Go we're "stuck" with the Simple Factory idea: injecting a factory function into
our constructor method and using it either as an embedded struct or more compositionally. To use Factory Method
we'd require the ability to inherit functionality from a base class, and we can't do that with Go. We can get
kind of close though, as I'll show you in my example code.

Here's what I came up with by the end of the chapter:

```go
type PizzaStore interface {
	OrderPizza(pizzaType string) pizzas.Pizza
	CreatePizza(pizzaType string) pizzas.Pizza
}

func NewPizzaStore(f pizzas.Factory) PizzaStore {
	return &pizzaStore{
		Factory: f,
	}
}

type pizzaStore struct{
	pizzas.Factory
}

func (p pizzaStore) OrderPizza(pizzaType string) pizzas.Pizza {
	pizza := p.CreatePizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
```

Before we go on, check out how we _can_ embed the functionality in the struct itself, so it's almost like it is
inheriting functionality from the factory that we've defined elsewhere. It uses `p.CreatePizza` even though it
doesn't define that method - it gets that from the embedded struct. Let's look at what a pizzas.Factory looks like:

```go
type Factory interface {
	CreatePizza(pizzaType string) Pizza
}

func ChicagoStyleFactory() Factory {
	return chicagoFactory{}
}

type chicagoFactory struct {}

func (c chicagoFactory) CreatePizza(pizzaType string) Pizza {
	fmt.Printf("Preparing a delectable Chicago-style %s pizza\n", pizzaType)
	f := ingredients.ChicagoIngredientFactory()
	switch pizzaType {
	case "cheese":
		return Cheese("Chicago Cheese", f)
	case "pepperoni":
		return Pepperoni("Chicago Pepperoni", f)
	}
	return nil
}
```

The chapter introduces ingredient factories near the end of this chapter, so that's what the
ChicagoIngredientFactory is in the middle there. This lets us apply dependency inversion to the ingredients
list as well, not just the pizza stores. This seems like another place where inheritance would be helpful as
I have to define this CreatePizza function on every factory, rather than just inheriting it, but it's still not
too bad. And we get to create a `Cheese` in each of the factories rather than _completely_ rewriting the wheel
at least. Anyway, let's put it all together:

```go
chiPiStore = pizzastore.NewPizzaStore(pizzas.ChicagoStyleFactory())

chiPiStore.OrderPizza("cheese")
chiPiStore.OrderPizza("pepperoni")

// Imagine a similarly defined factory as the above, but with NY replacing Chicago everywhere.
// The book had much better differences between the pizzas but I got lazy.
nyPiStore = pizzastore.NewPizzaStore(pizzas.NYStyleFactory())

nyPiStore.OrderPizza("cheese")
nyPiStore.OrderPizza("pepperoni")
```

One of the lessons I keep coming back to throughout the book is "Encapsulate what varies". This seems like a
great lesson to think about in a lot of the code I write, and the patterns I've been learning so far help us to
apply it! The book has done a great job so far of building on concepts.

### Singleton pattern

Ah, finally, a simple one. This chapter is as simple as you would hope. They go over how to do get this in Java:
make your class constructor private and provide a `getInstance` method. In the `getInstance` method you
instantiate the object if you haven't already, otherwise you just provide the current instance. Simple, right?
Well, by taking that approach you will be in trouble in a multi-threaded scenario - and we should always assume
our code is multi-threaded (especially in Go, otherwise what's the point?). To get multi-threaded support, we
need to do one of the following:

1. Protect it with mutexes (essentially - there is a `synchronized` function keyword you can use in Java, but
it slows the function down by a factor of 100 :oof:)
2. Instantiate the initial object at class (or package perhaps) loading time
3. Use "double-checked locking", which basically means use a mutex but after the initial `if` check. So in cases
where you already have the object created it won't hit the mutex at all, but it still protects the case where the
object isn't instantiated yet.

That's pretty much it. The issues with number 1 are related to speed, and number 2 isn't possible in every
situation. Number 3 is a good combination of the two approaches, depending on your use case.

How can we apply this in Go? Well, let's start by implementing the three things they did. We'll discuss Thingers
in these examples. Here's one implementation:

```go
type Thinger interface {
	Thing()
}

type thing struct {
	expensiveDBConnection string
}

func (t *thing) Thing() {
	fmt.Println("Whatever exactly this is supposed to do ...")
}
```

Here's option 1, the synchronized GetInstance():

```go
var mu sync.Mutex
var t *thing

func GetInstance() Thinger {
	mu.Lock()
	defer mu.Unlock()
	if t == nil {
		t = &thing{expensiveDBConnection: "so expensive"}
	}
	return t
}
```

We always lock and unlock our mutex, even if we already have a thing created. I assume this is what the
`synchronized` keyword does or some close analogue to it. Not great, and we can clearly do better. For the Java
folks though when you're just adding a keyword to a function definition it's not bad if you can accept the
performance hit. But we can do better.

Here's package init:

```go
var t Thing = &thing{expensiveDBConnection: "So expensive it's good to do it up front"}

func GetInstance() synchronized.Thinger {
	return t
}
```

Better, but in some cases you can't instantiate things this way because you might need inputs from other packages.
If you need var definitions from other packages, you could use those in an init function in your package:

```go
var t *thing

func init() {
    t = &thing{expensiveDBConnection: otherpackage.Conn}
}

func GetInstance() Thinger {
    return t
}
```

`init()` runs after variables have been processed but before any other non-init() code runs. So it can be a good
time to do such things.

Finally, we can do double-checked locking (which is just what it makes sense to do with mutexes):

```go
var mu sync.Mutex
var t *thing

func GetInstance() Thinger {
	if t == nil {
        mu.Lock()
        defer mu.Unlock()
		t = &thing{expensiveDBConnection: "so expensive"}
	}
	return t
}
```

The only difference is that we lock the mutex inside the if check. Makes sense to me. Here's our final example,
using sync.Once:

```go
var t *thing
var initOnce sync.Once

func GetInstance() synchronized.Thinger {
	initOnce.Do(func() {
		t = &thing{expensiveDBConnection: "So expensive"}
	})
	return t
}
```

`Do` is kind of a neat function. It loads a UInt32 and if that is equal to 0 it runs your function. Otherwise
it doesn't. The running of your function is also protected by a mutex, so it's an extra secure way to run your
function _only one time_.

In summary, go has a lot of great ways to initialize singletons, and it's harder to shoot yourself in the foot,
performance-wise, compared to Java (i.e. the `synchronized` keyword on your function).
