package pizzas

func ChicagoStyleFactory() Factory {
	return chicagoFactory{}
}

func NewYorkStyleFactory() Factory {
	return nyFactory{}
}
