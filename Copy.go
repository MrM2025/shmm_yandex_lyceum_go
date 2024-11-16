package main

import (
	"fmt"
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	p := make([]byte, n)
	readbytelength, err := r.Read(p)
	if err != nil || uint(readbytelength) > n{
		return fmt.Errorf("r_error")
	}
	_, errw := w.Write(p[:readbytelength])
	if errw != nil {
		return fmt.Errorf("w_error")
	}
	return nil
}