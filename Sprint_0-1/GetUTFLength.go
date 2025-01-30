package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}

func main() {
	len, err := GetUTFLength([]byte("hello"))
	fmt.Println(len, err)
	len, err = GetUTFLength([]byte{255})
	fmt.Println(len, err)
}
