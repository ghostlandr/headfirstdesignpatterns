package beverage

func Espresso() Beverage {
	return &espresso{}
}

type espresso struct{
	size Size
}

func (e espresso) GetSize() Size {
	return e.size
}

func (e *espresso) SetSize(s Size) {
	e.size = s
}

func (e espresso) Description() string {
	return "Espresso"
}

func (e espresso) Cost() float64 {
	return 1.99
}

func HouseBlend() Beverage {
	return &houseBlend{}
}

type houseBlend struct{
	size Size
}

func (h houseBlend) GetSize() Size {
	return h.size
}

func (h *houseBlend) SetSize(s Size) {
	h.size = s
}

func (h houseBlend) Description() string {
	return "House Blend"
}

func (h houseBlend) Cost() float64 {
	return .89
}

func DarkRoast() Beverage {
	return darkRoast{}
}

type darkRoast struct{
	size Size
}

func (d darkRoast) GetSize() Size {
	return d.size
}

func (d darkRoast) SetSize(s Size) {
	d.size = s
}

func (d darkRoast) Description() string {
	return "Dark Roast"
}

func (d darkRoast) Cost() float64 {
	return .99
}

func Decaf() Beverage {
	return &decaf{}
}

type decaf struct{
	size Size
}

func (d decaf) GetSize() Size {
	return d.size
}

func (d *decaf) SetSize(s Size) {
	d.size = s
}

func (d decaf) Description() string {
	return "Decaf"
}

func (d decaf) Cost() float64 {
	return 1.05
}

