// Implementing the weather data example from chapter 2 (Observer Pattern)
package main

import (
	"headfirstdesigntraining/observer/displays"
	"headfirstdesigntraining/observer/weatherdata"
)

func main() {
	w := weatherdata.New()
	curr := &displays.CurrentConditions{}
	fore := &displays.ForecastDisplay{}
	stat := &displays.StatisticsDisplay{}
	w.RegisterSubscriber(curr)
	w.RegisterSubscriber(fore)
	w.RegisterSubscriber(stat)
	w.NotifySubscribers(23.4, 90, 32)

	w.RemoveSubscriber(curr)
	w.NotifySubscribers(20.3, 80, 40)
	stat.Display()
	w.NotifySubscribers(10.3, 80, 40)
	stat.Display()
	w.NotifySubscribers(15.3, 80, 40)
	stat.Display()
}
