package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	fmt.Printf("Hello everyone\n")
}
