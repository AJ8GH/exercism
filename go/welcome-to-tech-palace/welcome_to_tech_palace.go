package techpalace

import (
	"strings"
)

const welcomeMessage = "Welcome to the Tech Palace, "
const star = "*"
const newLine = "\n"
const empty = ""

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return welcomeMessage + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat(star, numStarsPerLine)
	return strings.Join([]string{border, welcomeMsg, border}, newLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, star, empty))
}
