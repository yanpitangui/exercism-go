// Package twofer give you the ability to know what to say as you give away an extra cookie.
package twofer

import "fmt"

// ShareWith is a function that accepts the name of the person you are giving the cookie to, and returns the formatted string with the phrase.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
