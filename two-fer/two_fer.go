// Package twofer provide one function for resolve an exercism
package twofer

// ShareWith share with the name param and me. If name not provided put 'you'.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
