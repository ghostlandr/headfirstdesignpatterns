package pizzas

func Pepperoni() Pizza{
	return &pepperoni{}
}

type pepperoni struct{}

func (p pepperoni) Prepare() {
}

func (p pepperoni) Bake() {
}

func (p pepperoni) Cut() {
}

func (p pepperoni) Box() {
}

func Cheese() Pizza{
	return &cheese{}
}

type cheese struct{}

func (c cheese) Prepare() {
}

func (c cheese) Bake() {
}

func (c cheese) Cut() {
}

func (c cheese) Box() {
}


