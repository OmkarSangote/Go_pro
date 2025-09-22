package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}

	mux := &sync.Mutex{}

	go writeLoop(m, mux)
	go readLoop(m, mux)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mux *sync.Mutex) {
	for {
		for i := 0; i < 100; i++ {
			mux.Lock() // we Lock() the mutex before writing
			m[i] = i
			mux.Unlock() // Unlock() the mutex when we are done.
		}
	}
}

func readLoop(m map[int]int, mux *sync.Mutex) {
	for {
		mux.Lock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mux.Unlock()
	}
}
