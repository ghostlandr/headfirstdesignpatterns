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
favourite animal in general, is the mallard duck. The idea with the strategy pattern is that rather than make a
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
ha). Thus when trying to make it quack and fly, it would print about quacking, but say it can't fly. Here is a
rubber duck (squeaks, no flying) and a wooden duck (doesn't quack, doesn't fly):

```go
rubberDuck := NewDuck(NewItQuacks(), NewFlyNoWings())
woodenDuck := NewDuck(NewSilentQuack(), NewFlyNoWings())
```

The other big thing about strategy pattern is you are able to swap strategies at runtime. So we could add some
setter methods to our Duck that allow us to update it on the fly (get it?). The example they use is of a model
duck that they later attach a rocket to. Let's see what that looks like.

```go
func (d *Duck) SetQuacker(q Quacker) {
	d.quack = q
}

func (d *Duck) SetFlyer(f Flyer) {
	d.fly = f
}

// And our new fly style
type FlyRocketPowered struct{}
func (f FlyRocketPowered) Fly() {
	fmt.Println("Flying with rocket power!")
}

func NewFlyRocketPowered() Flyer {
	return FlyRocketPowered{}
}
```

Then, in our main function:

```go
modelDuck := NewDuck(NewItQuacks(), NewFlyNoWings())
modelDuck.Quack()
modelDuck.Fly() // No flying :(
modelDuck.SetFlyer(NewFlyRocketPowered())
modelDuck.Fly() // We flyin'!
```

This book is, naturally, heavily focused on Object Oriented (OO) programming and principles. That's kind of the
point, from their perspective. One of the principles they introduce in this chapter is that you should encapsulate
the behaviours (and code) that change and move them away from the things that stay the same. Earlier on I said
that the alternative to a design like what we have now is to have a subclass for every type of duck. So say we
have classes like a MallardDuck, RubberDuck, ModelDuck, etc. that all subclass Duck and override functionality.
If we need to add or remove functionality, we now need to go into every subclass and change it. By using the
strategy pattern and allowing ducks to use different strategies based on what they need we are able to encapsulate
those changes into the strategies, rather than leaving them in the concrete ducks. Speaking of that, ConcreteDuck
would be an interesting class to implement - I'll leave it as an exercise for the reader.

The biggest takeaway from this pattern: Take what varies and encapsulate it so it won't affect the rest of your
code. A very recent real world example of this came up just the other day. In some front end code we needed to
show a different dialog based on which page of the app we were on. We could have checked the url in the dialog
itself and done a bunch of ternary logic to switch what we were showing, or we could encapsulate that change
further up and show a totally different dialog based on the page. This allows us to have two separate dialogs
that can be developed independently. If we need another dialog some day, we have an easy insertion point now too
- we don't have to add further ternary expressions or switches or anything like that. As with all OO design,
it's a matter of tradeoffs. Yes, it is slightly more code and another class and etc., but the added clarity is
worth it in this case.

### Observer pattern

This is a pattern I am much more familiar with, as I have a stronger frontend skill set. In this chapter, we
implemented a weather data collection station. Let's have a look at some code. In our weather station
example, it's laid out like this: there is a weather data class that gets updated weather data intermittently.
When it does so, it updates three (for now) weather displays - this is our client code for this chapter. Here's
the first go at it, imperative-style:

```go
type WeatherStation interface {
    UpdateWeather(temp, humidity, pressure float64)
}

type weatherStation struct {
    forecast Display
    current Display
    statistics Display
}

func NewWeatherStation(forecast, current, statistics Display) WeatherStation {
    return &weatherStation{
        forecast: forecast,
        current: current,
        statistics: statistics,
    }
}

func (w *weatherStation) UpdateWeather(temp, humidity, pressure float64) {
    // Update each of our displays!
    w.forecast.Update(temp, humidity, pressure)
    w.current.Update(temp, humidity, pressure)
    w.statistics.Update(temp, humidity, pressure)
}

// Stitch it all together in a main function (not shown)
```

This code works, but it is fragile and tightly coupled to the three specific displays we have available to us
right now. How can we use object-oriented principles and patterns to our advantage here?

One of the places they suggest you try and apply design patterns is any spots in the code that are likely
to change frequently. One of the principles they talk about in chapter 3 (I think) is the Open-Closed principle:
Our code should be open for extension but closed for modification. (This is a principle that I've longed to
understand, and now it finally feels like I do after reading a few of these chapters!) In our weather station as
it stands today, every time we add a new weather display we will need to modify the weather data station code!
Not very "closed" at all.

Enter: the Observer pattern. Rather than adding code to our "update weather"
function in the weather data station, we give displays the ability to subscribe to updates. Therefore if we want
to add a new weather display at some point in the future we don't have to modify the weather data station code,
we just subscribe through the interface we already have available. This decouples the display from the station
and vice versa. Here's what some observer and observable interfaces could look like for our weather station
scenario:

```go
type Observer interface {
	Update(temp, humidity, pressure float64)
}

type Observable interface {
	RegisterSubscriber(o Observer)
	RemoveSubscriber(toRemove Observer)
	NotifySubscribers(temp, hum, pressure float64)
}

type weatherData struct {
	temp, humidity, pressure float64
	observers []Observer
}

// Implement the Observable interface with weatherData (not shown, but described below)
```

The register, remove, and notify functions all work like you'd expect: either append to the slice, remove from
the slice, or iterate the slice and call the Update function on subscribers. Then it's a simple matter of
calling NotifySubscribers() whenever you "get a new update" in the WeatherData implementation. Here's a possible
main function implementation for the station:

```go
w := weatherdata.New()
curr := &displays.CurrentConditions{}
fore := &displays.ForecastDisplay{}
stat := &displays.StatisticsDisplay{}
w.RegisterSubscriber(curr)
w.RegisterSubscriber(fore)
w.RegisterSubscriber(stat)
w.NotifySubscribers(23.4, 90, 32)
curr.Display() // See the current temperature and such
w.NotifySubscribers(20.3, 80, 40)
curr.Display() // See the new current temperature and such
```

A real-life example of the observer pattern is one I used often when working in Javascript - `addEventListener`.
You add a callback that should be run on a particular event on pretty well anything in Javascript (classic), but
most often DOM nodes. Something like `myButton.addEventListener('onclick', function() { console.log('hello!') })`
where myButton is a DOM node you'd have to grab earlier. Any time the `onclick` event happens on `myButton`, your
anonymous function will fire, and `myButton` is none the wiser.

This chapter was a good one for boosting my confidence in reading the book as I had it pretty well understood.

### Decorator pattern

Going into this chapter, I assumed we'd be using the @ sign a lot, as I'm used to that from Typescript and the like.
However, there were no @ signs at all in this chapter. This was also a relief because go doesn't have such things,
so it would have been interesting to try and figure out how to work around that.

This pattern centred around a coffee shop example. You start out with a `Beverage` class which exposes a cost
method that subclasses override to provide their cost. The coffee shop has four main beverages: dark roast,
house blend, decaf, and espresso. Then they have the usual suspects available as toppings/condiments: soy milk,
whip, mocha, and steamed milk. Initially, they created a subclass for every combination of beverages, and you can
imagine the carnage this creates: `DarkRoast`, `DarkRoastMocha`, `DarkRoastMochaWhip`, `DarkRoastSoy`, ... and
on and on. Dozens of classes. And what about things like double mocha? `DarkRoastMochaMocha`? This is getting
crazy!

Behold the decorator pattern. The other principle this chapter introduces is "programming to an interface, not
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

If you were curious, here are some of my implementations of these. It's a bit more annoying than in a true OO
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
during(?), or after the function, you're wrapping is called. In our code's case, I guess we're running our code
after since we call the wrapped function and then add our own text. However, since we are wrapping something,
we could call the wrapped functions whenever we want.

One of the things they admit to in this chapter is that this _is_ a bit cumbersome. They say this will be helped
a bit when we start learning the Factory pattern in the next chapter. In my mind, the main takeaway from this
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

_picture of before and after dependency inversion_

How does this come into play with factories? Let's use the example from the chapter: Pizza stores.

(side note: I quite enjoy the examples used in this book. I'm sure they are chosen for this very reason, but I
find them to be very relatable and reasonable. Like we're talking about "the factory pattern" but through the
lens of a pizza store and franchises, so of course it makes sense to make sure things are created the same way
every time and things like that, so providing a way for our franchisees to do that is totally necessary.)

We start out with code we've probably all written before. We have one pizza store, with an order pizza method
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
to treat every type of pizza the same once we get down to the prepare and ship step. However, you can see that this
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

type pizzaStore struct {
    PizzaFactory
}

func NewPizzaStore(factory PizzaFactory) PizzaStore {
    // Embed the factory in our concrete pizza store
    return &pizzaStore{PizzaFactory: factory}
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

Why is this better? Well, now if we want to add a new pizza type we can just add it to CreatePizza. No need to
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
the pizza store to create pizzas however they like. In the book they talk about New York-style pizza versus Chicago
style, two very different types of pizza to say the least. So you'd have a ChicagoPizzaStore creating Chicago-style,
and an NYPizzaStore creating New York-style. Life is good. In Go-land however, all we can really share is the
interface for making pizzas. We can't provide a piece of functionality for classes to embed, or at least we can't
prevent them from doing their own thing there, at least not without making some of the structs or interfaces
internal. I could see if you had a `pizza` (unexported) type and an `orderPizza` method that talked in lower-case
p pizzas, maybe you could protect that process?

The thing we've spec'd out in the Java code is known as a Factory Method. Responsibility for what it creates
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

### Command Pattern

The command pattern lets you encapsulate a request as an object. This lets you do a couple of things: you can save
it for later invocation by some other piece of code; you can abstract object-specific details away from the code
running the commands (this makes me think of how we did the Deliverer interface in Mission Control); and it lets
you implement things like undo operations more simply since a command (can) know how to undo what it just did.

The command pattern requires three actors: a client, an invoker, and a receiver. All three of these actors work
with a Command interface. The receiver knows how to perform the actual commands in question. Let's use the example
of a light bulb (`Light`). The light exposes two methods: On, and Off. Without the command pattern, your client
would simply know exactly how to call the light's on and off methods. This is reasonable with a light, but if you
consider something more complicated - say a thermostat or something with a myriad more settings than a binary
off/on switch - being able to encapsulate what makes the thermostat go "on" in a command is where they truly shine.
All that to say, the light bulb is our receiver in this case. 

The invoker is the one who will run the actual commands. For this example, we'll use a remote control. Our remote
control has a bunch of slots that we can put commands into, and the remote invokes each one for us. The remote
control works with a command interface only. This lets us use dependency inversion, as the remote depends on
Command and the client composes commands as well. Here's an example command interface:

```go
type Command interface {
	Execute()
}
```

That's it! Now let's define a simple remote with some commands it can run:

```go
type Remote interface {
	SetCommand(slot int, on, off Command)
	OnButtonWasPressed(slot int)
	OffButtonWasPressed(slot int)
}

type remote struct {
	onCommands []Command
	offCommands []Command
}

func NewRemote() Remote {
	return &remote {
		// The 7 is arbitrary - simply a 7 slotted remote
		onCommands:  make([]Command, 7),
		offCommands: make([]Command, 7),
	}	
}

// Implementing the interface is easy
func (r *remote) SetCommand(slot int, on, off Command) {
	r.onCommands[slot] = on
	r.offCommands[slot] = off
}

func (r *remote) OnButtonWasPressed(slot int) {
	r.onCommands[slot].Execute()
}

func (r *remote) OffButtonWasPressed(slot int) {
	r.offCommands[slot].Execute()
}
```

Using the remote is pretty simple too, here's an example main.go (our client for this example):

```go
r := NewRemote()

light := NewLight("bedroom") // A light that has On and Off methods 
lOn := NewLightOnCommand(light) // Command interface satisfiers
lOff := NewLightOffCommand(light)

r.SetCommand(0, lOn, lOff) // Set the on and off buttons for the light to the first slot

r.OnButtonWasPressed(0) // Prints out "Turning bedroom light on"
r.OffButtonWasPressed(0) // Prints out "Turning bedroom light off"
```

The nice thing about this? We can add any commands we want to our remote. It doesn't matter if it's
as simple as turning on a light bulb or as complicated as turning on the stereo, setting it to CD
mode, and turning the volume up. Here's some code for a stereo:

```go
type Stereo interface {
	On()
	Off()
	SetCD()
	SetVolume(level int)
	// ... probably lots more
}

func NewStereoOnWithCD(s Stereo) Command {
	return stereoOnWithCD{s: s}
}

type stereoOnWithCD struct {
	s Stereo
}

func (s stereoOnWithCD) Execute() {
	s.s.On()
	s.s.SetCD()
	s.s.SetVolume(11) // Why 11? Cause it's one louder than 10
}

// Off is just a command that runs s.s.Off()
```

Now here's how we could use it. Imagine this code after the remote instantiation above

```go
// I didn't show you this implementation, but NewStereo returns something that satisfies the Stereo
// interface... trust me.
s := NewStereo("living room")
sOn := NewStereoOnWithCD(s)
sOff := NewStereoOff(s)

r.SetCommand(1, sOn, sOff)
r.OnButtonWasPressed(1) // Prints out for each step of the stereo execution
r.OffButtonWasPressed(1) // Prints out "Turning living room stereo off"
```

We didn't have to change our remotes code at all and we were able to tell it how to turn on and off
a brand new device. That's the power of the command pattern and dependency inversion.

For bonus points, let's add the ability to create arbitrary lists of commands to run, aka Macros.
We'll make a new command for this, the MacroCommand.

```go
type macroCommand struct {
	cmds []Command
}

// MacroCommand just executes all of its child commands!
func (m *macroCommand) Execute() {
	for _, cmd := range m.cmds {
		cmd.Execute()
	}
}

func NewMacro(cmds []Command) Command {
	return &macroCommand{cmds: cmds}
}
```

Now let's use it back in our "main function" (after we have instantiated everything):

```go
// Turn on the lights and the stereo! Party time!
macroOn := NewMacro(lOn, sOn)
macroOff := NewMacro(lOff, sOff)
r.SetCommand(2, macroOn, macroOff)

r.OnButtonWasPressed(2) // Prints all the commands for turning on lights and stereo
r.OffButtonWasPressed(2) // Prints the output from turning the lights and stereo off
```

That's pretty much it for command pattern. In the book they did some stuff with undo but that's
really as simple as you think: store the command you ran upon pressing a button and then run
its undo method if they hit the undo button.

### Adapter pattern

The adapter pattern is easy to understand, because we have lots of real life examples of it. If
you're using a laptop right now, you might even have a power adapter connected to it. A power
adapter _adapts_ the electrical current from the wall into what your laptop actually wants to use.
If you've travelled to another continent with your laptop, you may have had to use another
adapter to actually plug your power adapter into the wall. In our code, the adapter pattern is
used to make one object look like another. It makes use of object composition to do this, and
we'll explore that more, shortly.

The example they used in the book called back to the duck example from Strategy (yass, ducks!). Now
we're going to add turkeys to the mix. Let's look at their two interfaces side by side:

```go
type Duck interface {
	Quack()
	Fly()
}

type Turkey interface {
	Gobble()
	FlyShortDistance() // Turkeys are bad at flying, but they can fly short distances!
}
```

You can see that they're similar, but not quite the same. We can use the adapter pattern to make a
turkey sound and act like a duck. Here's how we might do it in go.

```go
type turkeyAdapter struct {
	t Turkey // Wrap a Turkey
}

func (t *turkeyAdapter) Quack() {
	t.t.Gobble() // Close enough, just gobble instead
}

func (t *turkeyAdapter) Fly() {
	// We can't fly well, so fly five times to make up for how much the duck would fly
	for i := 0; i < 5; i++ {
		t.t.FlyShortDistance()
	}
}
```

We now have a TurDuck. Here's some test code:

```go
func testDuck(d Duck) {
	d.Quack()
	d.Fly()
}

func main() {
	d := NewDuck()
	testDuck(d) // Works fine
	t := NewTurkey()
	turDuck := turkeyAdapter{t: t}
	testDuck(turDuck) // Also works fine! We had to fly five times but we made it.
}
```

In the case of go I guess we could just do this:

```go
// turkey implements Turkey interface above
func (t *turkey) Quack() {
	t.Gobble()
}

func (t *turkey) Fly() {
	for i := 0; i < 5; i++ {
		t.t.FlyShortDistance()
	}
}
```

Back in main:

```go
	// ... Other main function stuff above ...
	testDuck(t)
```

We implemented the Duck interface's methods on Turkey, so now our turkey struct satisfies both the
Duck and Turkey interfaces. I think this is worse than having an explicit adapter however, as now
our "turkey" struct doesn't have a great name. This could lead us to even more refactoring... or we
could define our turkeyAdapter and just wrap up a turkey whenever we need a duck. I like that more.

### Facade pattern

Adapter is used to make something look or act like something else. Facade is used to make something
simpler or easier to use. In go these are especially useful, as they not only simplify the
interfaces our code wants to use, it simplifies any tests or mock objects we would want to use in
our code as well. Imagine we had a dependency that looked something like this:

```go
type BigInterface interface {
	BigMethod1()
	BigMethod2()
	BigMethod3()
	BigMethod4()
	// ... on and on and on ...
}
```

And we used it in our code like this:

```go
type widget struct {
	builder BigInterface
}

func NewWidget(bi BigInterface) Widget {
	return &widget{ builder: bi }
}

func (w *widget) Build() {
	w.builder.BigMethod4()
}
```

You can see that we only use BigMethod4 in our actual code. If we wanted to create a mock version
of BigInterface to use in our tests, we'd have to mock out all of the methods on BigInterface to
make it compatible. Our editor could make this easier for us, but the facade pattern can help here
too. Let's make our own version of the interface that just includes the one method.

```go
type Builder interface {
	BigMethod4() // This needs to be the same signature as the one we are interested in
}

type widget struct {
	builder Builder
}

func NewWidget(b Builder) Widget {
	return &widget{ builder: b }
}

// Build method doesn't have to change!
```

We've created a facade in front of BigInterface. Now, we depend on much less of the interface (none
of it, in fact), and our testing mocks can now be simplified greatly.

### Template Method pattern

Template method is another pattern that is tough without inheritance, much like the command pattern. However,
tough does not mean impossible. We don't have inheritance based template method available to us but we do have
compositional template method available.

Template method works by allowing subclasses to override certain pieces of your algorithm. Here's an example of
what this might look like for an api handling class in Python:

```python
class ApiHandler(object):
    def handle_request(self, request):
        self.validate_args(request)
        args = self.deserialize_args(request)
        self.process(args) # Let the subclass do its work!
        self.post_process_formatting()
        return self.response()

    def process(self, args):
        pass
```

Here's how I might implement this class, as a new api handler:

```python
class TemplateHandler(ApiHandler):
    def process(self, args):
        arg = args['myThing']
        self.do_something_with_arg(arg)
        # ...
```

The power of this approach is we can put all of our validation and other things into ApiHandler and beef it up
as much as we like. Then anyone who is subclassing ApiHandler gets all of that functionality. So long as we
don't break the algorithmic flow of `handle_request`, everything will keep working as expected. Whenever I think
about "design patterns" as a general concept, the first one that comes to mind is always the template method
pattern, and I think it's because of this ApiHandler design.

The other thing they talk about under the template method umbrella is the concept of hooks. The difference
between hooks and the normal steps of the algorithm is that they are optional. The normal steps are generally
marked as abstract or they'll throw a `NotImplementedException` if you're in python. Hooks on the other hand
are just nice places to hook in and add functionality as required. If you've ever used git before you've had
the opportunity to add hooks: pre commit hook, post commit hook, pre push hook, post push hook, etc... None of
these hooks do anything unless you add one of your own.

That's all well and good, but what about our precious go? I was trying to think of how you could use struct
embedding or something to get this functionality, and with even more thinking I bet I could come up with
something (so I'll leave that as a TODO here for myself), but we do have access to one of the things they talked
about in the chapter, and that is what I'm calling compositional template method. The example they used in the
book is around sorting arrays in Java, and so we will use the example of sorting in Go!

Have you ever had to sort a slice in go? If not, you're missing out! All you need to do is implement the
interface defined in the sort package and call sort.Sort. Here's an example from the standard lib examples:

```go
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	sort.Sort(ByAge(people))
	fmt.Println(people)
```

I simplified it a bit, but this still lets you see the power of this. Here is the comment on the Sort method:

```
func Sort(data Interface)

Sort sorts data. It makes one call to data.Len to determine n, and O(n*log(n)) calls to data.Less and
data.Swap. The sort is not guaranteed to be stable.
```

It takes in a piece of data that matches its interface and returns a sorted version. This is a fantastic way
to get around the problem that lots of go programmers run into. Imagine trying to implement sorting without
knowing the type of data that you're sorting (for the sake of comparing them)? Well, just defer that as part of
a template method onto the implementer and you're done.





