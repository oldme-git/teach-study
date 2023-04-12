package test2

import (
	"fmt"
	"testing"
)

type Nei struct {
	s string
}

func (n Nei) NeiA() {
	fmt.Println(n.s)
}

type Wai struct {
	*Nei
}

func (w *Wai) WaiA() {
	fmt.Println("waiA")
}

func TestA(t *testing.T) {
	w := &Wai{
		&Nei{s: "saaa"},
	}
	w.NeiA()
}
