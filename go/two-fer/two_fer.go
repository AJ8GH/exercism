// Package twofer implements a function to share with someone by name
package twofer

import "fmt"

// ShareWith addresses the subject by name or as "you"
func ShareWith(name string) string {
	return fmt.Sprint("One for ", nameOrYou(name), ", one for me.")
}

func nameOrYou(name string) string {
	if name == "" {
		return "you"
	}
	return name
}
