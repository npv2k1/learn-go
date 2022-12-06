package main

import "testing"

func TestAdd(t *testing.T) {

	got := Add(4, 5)
	want := 9
	// fmt.Println("got: ", got)
	// fmt.Println("want: ", want)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
