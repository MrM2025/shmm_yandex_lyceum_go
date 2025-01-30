package main

import (
	"fmt"
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	got, err1 := GetUTFLength([]byte("hello"))
	expected := 5

	if got != expected && err1 != nil {
		t.Fatalf(`GetUTFLength("hello", 5) = %d, want %d`, got, expected)
	}
}

func TestGetUTFLength1(t *testing.T) {
	got, err := GetUTFLength([]byte{255})
	expected := fmt.Errorf("invalid utf8")

	if err != expected && got != 0 {
		t.Fatalf(`GetUTFLength([]byte{255}, 0) = %d, want %s`, err, expected)
	}
}
