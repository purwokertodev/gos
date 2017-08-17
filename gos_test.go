package gos

import (
	"testing"
)

func TestGos(t *testing.T) {
	g := New(`~!@#$%^&*()_+=-`, true, true, true)
	if g.allowEmpty{
		t.Log("success create gos")
	}
}

func TestSpecialCharInput(t *testing.T) {
	g := New(`~!@#$%^&*()_+=-`, true, false, false)
	e := g.Validate(`&*`)
	if e != nil {
		t.Error("test special char run as expected")
	}
}
