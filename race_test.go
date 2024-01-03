package main

import (
	"sync"
	"testing"
)
func TestRaceCondtion(t *testing.T){
	var state int32
	var mu sync.RWMutex

	for i:=0;i<10;i++{
		go func(i int){
			mu.Lock()
			// another way atomic.AddInt32(&state,i) same thing
			state+=int32(i)
			mu.Unlock()
		}(i)
	}
}