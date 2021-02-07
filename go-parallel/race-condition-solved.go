package main

import (
	"sync"
)

func main() {
	a := []int{1, 2, 3}

	mu := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			a = append(a, i)
			mu.Unlock()
		}(i)
	}
}
