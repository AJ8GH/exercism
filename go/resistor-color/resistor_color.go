package resistorcolor

var colors = []string{
	"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
}

// Colors returns the list of all colors.
func Colors() []string {
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for i, v := range colors {
		if v == color {
			return i
		}
	}
	return -1
}
