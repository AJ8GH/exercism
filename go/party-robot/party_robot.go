package partyrobot

import "fmt"

const welcomeMessage = "Welcome to my party, %s!"
const birthdayMessage = "Happy birthday %s! You are now %d years old!"
const tableMessage = `You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.
You will be sitting next to %s.`

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf(welcomeMessage, name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf(birthdayMessage, name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return Welcome(name) + "\n" + fmt.Sprintf(tableMessage, table, direction, distance, neighbor)
}
