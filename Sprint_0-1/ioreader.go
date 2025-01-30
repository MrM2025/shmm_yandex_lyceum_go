package main

import (
	"fmt"
	"io"
	"strings"
)

type customReader struct {
}

func (cr *customReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("read error")
}

func NewCustomReader() *customReader {
	return &customReader{}
}

func ReadIntoBuffer(r io.Reader) (string, error) {

	var buffer []byte

	anotherbyte := make([]byte, 1)
	for {
		_, err := r.Read(anotherbyte)
		if err != nil {
			if err.Error() != io.EOF.Error() {
				return string(buffer), err
			} else {
				return string(buffer), nil
			}
		}
		buffer = append(buffer, anotherbyte...)
	}
}

func ReadString(r io.Reader) (string, error) {

	return ReadIntoBuffer(r)
}

func main() {

	tests := []struct {
		name     string
		input    io.Reader
		expected string
		wantErr  bool
	}{
		{
			name:     "Valid Empty input",
			input:    strings.NewReader(""),
			expected: "",
			wantErr:  false,
		},
		{
			name:     "Valid input",
			input:    strings.NewReader("Hello, World!"),
			expected: "Hello, World!",
			wantErr:  false,
		},
		{
			name:    "Invalid reader",
			input:   NewCustomReader(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		got, err := ReadString(tt.input)
		if (err != nil) != tt.wantErr {
			fmt.Println(fmt.Errorf("ReadString() error = %v, wantErr %v", err, tt.wantErr))
			fmt.Println(got, err)
		} else if got != tt.expected {
			fmt.Println(fmt.Errorf("ReadString() = %v, expected...", tt.expected))
			fmt.Println(got, err)
		} else {
			fmt.Println(got, err)
		}
	}

}
