package main

import "io"

func Contains(r io.Reader, seq []byte) (bool, error) {
	var tf bool = false
	onebyteslice := make([]byte, 1)
	var l int
	
	for {
		_, err :=  r.Read(onebyteslice) 
	
		if err != nil {
			return tf, err
		}
		if l == len(seq) - 1 {
			if onebyteslice[0] == seq[l]{
				tf = true
				return tf, nil
			}
		}
		if onebyteslice[0] == seq[l] {
			l++
			continue
		}
	}

}
func main() {
	var r io.Reader
	Contains(r, []byte{123, 123, 123})
}

