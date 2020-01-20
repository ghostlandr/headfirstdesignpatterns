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
	w.SetMeasurements(23.4, 90, 32)

	w.RemoveSubscriber(curr)
	w.SetMeasurements(20.3, 80, 40)
	stat.Display()
	w.SetMeasurements(10.3, 80, 40)
	stat.Display()
	w.SetMeasurements(15.3, 80, 40)
	stat.Display()
}
