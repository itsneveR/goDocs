package main

import (
	"sync"
)

func add(w *sync.WaitGroup, num *int32) {
	defer w.Done()
	for i := 0; i < 1000; i = i + 1 {
		*num++
	} // spawn 1000 new goroutines
}
func main() {
	var n int32 = 0
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(1)
	go add(wg, &n)
	wg.Wait()
	println(n)
}
