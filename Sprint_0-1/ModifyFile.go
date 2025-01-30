package main

import (
	"os"
	"io"
	"fmt"
)

func ModifyFile(filename string, pos int, val string) {
	f, ferr := os.OpenFile(filename, os.O_WRONLY, 0600)
	if ferr != nil && ferr != io.EOF {
		fmt.Println("OpenFile error")
	}
	defer f.Close()
	f.Seek(int64(pos), 0)
	f.WriteString(val)

}