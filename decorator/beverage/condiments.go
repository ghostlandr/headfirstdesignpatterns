package beverage

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

func Soy(b Beverage) Beverage {
	return soy{Beverage: b}
}

type soy struct{
	Beverage
}

func (s soy) Description() string {
	return s.Beverage.Description() + ", Soy"
}

func (s soy) Cost() float64 {
	var c float64
	size := s.Beverage.GetSize()
	switch size {
	case Tall:
		c = .1
	case Grande:
		c = .15
	case Venti:
		c = .2
	}
	return s.Beverage.Cost() + c
}

func Whip(b Beverage) Beverage {
	return whip{Beverage: b}
}

type whip struct{
	Beverage
}

func (w whip) Description() string {
	return w.Beverage.Description() + ", Whip"
}

func (w whip) Cost() float64 {
	return w.Beverage.Cost() + .10
}

func SteamedMilk(b Beverage) Beverage {
	return steamedMilk{Beverage: b}
}

type steamedMilk struct{
	Beverage
}

func (s steamedMilk) Description() string {
	return s.Beverage.Description() + ", Steamed Milk"
}

func (s steamedMilk) Cost() float64 {
	return s.Beverage.Cost() + .10
}

