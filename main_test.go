package main

import (
	"testing"
)

func TestBar(t *testing.T) {
	result := Bar()
	if result != "br" {
		t.Errorf("expecting bar, got %s", result)
	}
}

func TestFoo(t *testing.T) {
	result := Foo()
	if result != -10 {
		t.Errorf("expecting foo, got %d", result)
	}
}

// 関数別GoTest
func TestFooBar(t *testing.T) {
	t.Error("intentional error 1")
}

func TestQuz(t *testing.T) {
	t.Error("intentional error 1")
}
