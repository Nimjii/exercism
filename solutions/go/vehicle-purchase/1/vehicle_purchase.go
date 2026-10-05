package purchase

func NeedsLicense(kind string) bool {
	if kind == "truck" || kind == "car" {
        return true
    }
    return false
}

func ChooseVehicle(option1, option2 string) string {
    var result string
	if option1 <= option2 {
        result = option1
    } else {
        result = option2
    }
    return result + " is clearly the better choice."
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
        return originalPrice * 0.8
    } else if age >= 10 {
        return originalPrice * 0.5
    }
    return originalPrice * 0.7
}
