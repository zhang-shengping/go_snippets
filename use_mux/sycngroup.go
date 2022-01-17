package main

import "sync"

func add(w *sync.WaitGroup, num *int) {
    // wg.Done() and wg.Add(-1) are exactly equivalent
    defer w.Done()
    *num= *num + 1
}
func main() {
    var n int = 0
    // is used to wait for a group
    // of goroutines to finish executing, and control is blocked until the group of
    // goroutines finishes executing.
    var wg *sync.WaitGroup = new(sync.WaitGroup)
    wg.Add(1000)
    for i := 0; i < 1000; i = i + 1 {
	go add(wg, &n)
    } // spawn 1000 new goroutines
    wg.Wait()
    println(n)
}
