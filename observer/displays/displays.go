package displays

import (
	"log"
)

type CurrentConditions struct {}

func (c *CurrentConditions) Update(temp, humidity, pressure float64) {
	log.Printf("Getting updated! %f %f %f", temp, humidity, pressure)
}

type StatisticsDisplay struct {
	temps, humidities, pressures []float64
}

func (c *StatisticsDisplay) Update(temp, humidity, pressure float64) {
	c.temps = append(c.temps, temp)
}

func (c StatisticsDisplay) Display() {
	avgTmp := getAvgFloat(c.temps)
	log.Printf("Over %d reports, the average temperature is %f", len(c.temps), avgTmp)
}

func getAvgFloat(floats []float64) float64 {
	var total float64
	for _, f := range floats {
		total += f
	}
	return total / float64(len(floats))
}

type ForecastDisplay struct {}

func (c *ForecastDisplay) Update(temp, humidity, pressure float64) {

}

