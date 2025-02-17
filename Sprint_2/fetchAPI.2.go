package main

import (
	"context"
	"net/http"
	"io"
	"time"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

type APIResponse struct {
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int // код ответа
	Err        error // ошибка, если возникла
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	var slice []*APIResponse
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for _, i := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				mu.Lock()
				slice = append(slice, &APIResponse{
					URL: u,
					Err: nil,
				})
				mu.Unlock()
			}

			resp, err := http.DefaultClient.Do(req)
			if err!= nil {
				mu.Lock()
				slice = append(slice, &APIResponse{
					URL: u,
					Err: err,
				})
				mu.Unlock()
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err!= nil {
				mu.Lock()
				slice = append(slice, &APIResponse{
					URL: u,
					StatusCode: resp.StatusCode,
					Err: err,
				})
				mu.Unlock()
				return
			}

			mu.Lock()
			slice = append(slice, &APIResponse{
				URL: u,
				Data: string(body),
				StatusCode: resp.StatusCode,
				Err: nil,
			})
			mu.Unlock()
	}(i)}
	wg.Wait()
	return slice
}