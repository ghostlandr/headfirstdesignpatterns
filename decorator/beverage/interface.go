package beverage

type Beverage interface {
	Description() string
	Cost() float64
	GetSize() Size
	SetSize(s Size)
}

type Size int
const (
	Tall Size = iota
	Grande
	Venti
)
