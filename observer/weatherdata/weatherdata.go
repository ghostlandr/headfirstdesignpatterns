package weatherdata

type Observer interface {
	Update(temp, humidity, pressure float64)
}

type Observable interface {
	RegisterSubscriber(o Observer)
	RemoveSubscriber(toRemove Observer)
	NotifySubscribers(temp, hum, pres float64)
}

func New() Observable {
	return &WeatherData{
	}
}

type WeatherData struct {
	temp, humidity, pressure float64
	observers []Observer
}

// RegisterSubscriber adds an observer to the update queue
func (w *WeatherData) RegisterSubscriber(o Observer) {
	w.observers = append(w.observers, o)
}

// RemoveSubscriber removes an observer from the update queue
func (w *WeatherData) RemoveSubscriber(toRemove Observer) {
	for i, o := range w.observers {
		if o == toRemove {
			// Slice out the observer to be removed
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherData) NotifySubscribers(temp, hum, pres float64) {
	w.temp = temp
	w.humidity = hum
	w.pressure = pres
	for _, o := range w.observers {
		o.Update(w.temp, w.humidity, w.pressure)
	}
}
