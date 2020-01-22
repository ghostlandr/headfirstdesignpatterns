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
