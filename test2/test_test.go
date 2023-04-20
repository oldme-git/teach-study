package test2

import (
	"fmt"
	"testing"
)

const (
	Like = 1 << iota
	Collect
	Comment
)

func TestA(t *testing.T) {
	fmt.Printf("%b", 0b1110&0b0001)
}
