package purchase

const choiceMessage = " is clearly the better choice."
const priceUnder3Years = 0.8
const priceUnder10Years = 0.7
const priceAfter10Years = 0.5

var licenseNeeders = map[string]bool{"car": true, "truck": true}

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return licenseNeeders[kind]
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + choiceMessage
	}
	return option2 + choiceMessage
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3:
		return originalPrice * priceUnder3Years
	case age < 10:
		return originalPrice * priceUnder10Years
	default:
		return originalPrice * priceAfter10Years
	}
}
