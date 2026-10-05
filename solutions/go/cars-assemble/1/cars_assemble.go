package cars

const MinutesPerHour uint = 60
const PricePerCar uint = 10000
const BulkMultiplier float64 = 0.95
const CarsInBulk uint = 10

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / float64(100)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / float64(MinutesPerHour))
}

func CalculateCost(carsCount int) uint {
    var bulkPrice = uint(float64(PricePerCar * CarsInBulk) * BulkMultiplier)
    var bulks = uint(carsCount / int(CarsInBulk))
    var singles = uint(carsCount % int(CarsInBulk))
	return singles * PricePerCar + bulks * bulkPrice
}
