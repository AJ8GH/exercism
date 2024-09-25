package elon

import "fmt"

func (c *Car) Drive() {
	if c.battery-c.batteryDrain <= 0 {
		return
	}
	c.distance += c.speed
	c.battery -= c.batteryDrain
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
	distanceOverSpeed := float64(trackDistance) / float64(c.speed)
	return (distanceOverSpeed)*float64(c.batteryDrain) <= float64(c.battery)
}
