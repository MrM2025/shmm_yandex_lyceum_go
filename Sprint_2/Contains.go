package main

import (
	"context"
	"io"
	"bytes"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	if len(seq) == 0 {
		return false, io.EOF
	   }	
	buf :=  make([]byte, 1)
	data := make([]byte, 0)

	for {
		select{
		case <- ctx.Done():
			return false, ctx.Err()

		default:
			_, err := r.Read(buf)
			if err != nil && err != io.EOF {
				return false, err
			   }
			   data = append(data, buf...)

			if bytes.Contains(data, seq) {
				return true, nil
			   }
			if err == io.EOF {
				return false, nil
			    }
			if len(data) > len(seq)-1 {
				data = data[1:]
	}
	}		
}
}