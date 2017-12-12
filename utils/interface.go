package utils

import "fmt"

// InterfaceToKey converts an interface to a string to use as key in data
// structures such as trees. Not to be confused with a proper hash function.
func InterfaceToKey(x interface{}) string {
	return fmt.Sprintf("%#v", x)
}
