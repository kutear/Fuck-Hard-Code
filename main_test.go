package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	a := &A{str: "name"}
	hello(a)
	fmt.Println(a)
}

func hello(a *A) {
	a.str = "hi"
}

type A struct {
	name []B
	str  string
}
type B struct {
	name string
}
