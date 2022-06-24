package utils

import "github.com/davecgh/go-spew/spew"

// Debug -
func Debug(a ...interface{}) {
	spew.Dump(a...)
}
