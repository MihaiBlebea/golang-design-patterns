package singleton_test

import (
	"testing"

	"github.com/MihaiBlebea/go-patterns/creational/singleton"
)

func TestGetSameInstance(t *testing.T) {
	s1 := singleton.GetInstance()
	s1.Increment() // value = 1

	s2 := singleton.GetInstance()
	s2.Increment() // value = 2

	if s1.Value() != 2 {
		t.Errorf("Expected %d got %d", 2, s1.Value())
	}

	if s1.Value() != s2.Value() {
		t.Errorf("Expected %d and %d to be equal", s1.Value(), s2.Value())
	}
}
