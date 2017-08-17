package gos

import (
	"testing"
)

func TestGos(t *testing.T) {
	g := New(`~!@#$%^&*()_+=-`, true, true, true)
	if g != nil {
		t.Log("success create gos")
	}
}

func TestSpecialCharInput(t *testing.T) {
	g := New(`~!@#$%^&*()_+=-`, true, false, false)
	e := g.Validate(`&*`)
	if e == nil {
		t.Log("test special char run as expected")
	}
}

func TestSpecialCharInvalidInput(t *testing.T) {
	g := New(`~!@#$%^&*()_+=-`, true, false, false)
	e := g.Validate(`&*:`)
	if e != nil {
		t.Log("test special char with invalid input")
	}
}
