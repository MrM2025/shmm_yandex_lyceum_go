package main

import (
	"sync"
	//"log"
)

type Data struct {
	ID string // для упрощения содержит только ID
}

type SafeMap struct {
	m map[string]interface{}
	dr DataRetriever
	mux sync.RWMutex
}

type DataRetriever interface {
	Retrieve(ID string) (*Data, error)
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.RLock()
	data, exists := s.m[key]
	s.mux.RUnlock()

	if exists {
		return &data
	}

	data, err := s.dr.Retrieve(key)
	if err != nil {
		//log.Printf("s.dr.Retrieve(key): %s", err)
		return data
	}

	s.mux.Lock()
	defer s.mux.Unlock()

	data, exists = s.m[key]

	if exists {
		return data
	}

	s.m[key] = data

	return data

}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[key] = value	
}