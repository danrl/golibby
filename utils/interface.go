package utils

import (
	"crypto/sha256"
	"fmt"
)

// InterfaceToKey converts an interface to a string to use as key in data
// structures such as trees. Not to be confused with a proper hash function.
func InterfaceToKey(x interface{}) string {
	return fmt.Sprintf("%#v", x)
}

// InterfaceToHash256 provides a not-so-ideal way of hashing dynamically typed
// data. It is a workaround that may be replaced with the golang-internal hashing
// function for comparables, once it is exported.
// See this thread for details: https://github.com/golang/go/issues/21195
func InterfaceToHash256(x interface{}) [32]byte {
	return sha256.Sum256([]byte(fmt.Sprintf("%#v", x)))
}
