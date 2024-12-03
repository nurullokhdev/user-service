package main

import (
	"testing"
)

func TestSucks(t *testing.T) {
	str := "Asshole of the year"
	t.Fatal([]byte(str))
}
