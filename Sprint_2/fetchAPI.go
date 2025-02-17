package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponsee struct { 
	Data       string // тело ответа
	StatusCode int  // код ответа
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponsee, error) {
	ctxwcl, cancelctx := context.WithTimeout(ctx, timeout)
	defer cancelctx()
		select {
		case <-ctxwcl.Done():
			return nil, context.DeadlineExceeded
		default:

		

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		client := http.DefaultClient
		
		resp, err := client.Do(req)
		if err != nil {
			return nil, err // автоматически вернет DeadlineExceeded при таймауте
		}

		res, err := io.ReadAll(resp.Body)
		if err != nil{
			return nil, err
		}


		return &APIResponsee{
			Data: string(res),
			StatusCode: resp.StatusCode,
		}, nil


}
}