package cars

const oneHundredPercent = 100.0
const secondsInMinute = 60.0
const costPerCar = 10_000
const costPerTenCars = 95_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / oneHundredPercent)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / secondsInMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	remainder := carsCount % 10
	return uint((((carsCount - remainder) / 10) * costPerTenCars) +
		(remainder * costPerCar))
}
