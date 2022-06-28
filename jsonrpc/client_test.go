package jsonrpc

import (
	"testing"
	"time"
)

func TestSimpleCall(t *testing.T) {

	hello := Method[int, int]{Name: "hello"}

	aToB := make(chan []byte, 1)
	bToA := make(chan []byte, 1)

	a := NewClient(bToA, aToB)
	defer a.Close()
	b := NewClient(aToB, bToA)
	defer b.Close()

	hello.Listen(func(arg *int) int {
		return *arg * 2
	}, a)

	res, err := hello.Call(2, time.Second*30, b)
	if err != nil {
		t.Fatal(err)
	}
	if *res != 4 {
		t.Errorf("Expected 4 got %d", *res)
	}
}
