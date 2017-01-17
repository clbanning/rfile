package rfile

import (
	"fmt"
	"testing"
)

func TestTail(t *testing.T) {
	fmt.Println("\n---------- TestTail")
	tail, err := Tail("LICENSE", 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(tail) != 3 {
		t.Fatal("tail has len:", len(tail))
	}
	for i, s := range tail {
		fmt.Println(i, ":", s)
	}
}

