package main

import "sync"
// shared data
var num int = 0

func add(lc *sync.Mutex, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 100000; i = i + 1 {
        lc.Lock()
        num = num + 1
        lc.Unlock()
    }
}
func minus(lc *sync.Mutex, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 100000; i = i + 1 {
        lc.Lock()
        num = num - 1
        lc.Unlock()
    }
}
func main() {
    var mutex *sync.Mutex = new(sync.Mutex)
    var wg *sync.WaitGroup = new(sync.WaitGroup)
    wg.Add(2)
    go add(mutex, wg)
    go minus(mutex, wg)
    wg.Wait()

    println(num) // 0
}

// sync.RWMutex
// sync.RWMutex type has two other methods: RLock() and RUnlock().

/*

If a goroutine wants to get RLock and other goroutines have only RLock, it will
not block the goroutine that wants to get RLock, if there is a goroutine with
Lock, it will block the goroutine that wants to get RLock until all other
goroutines with Lock have released their Lock.

If a goroutine wants to get Lock, other goroutines with RLock or Lock will block
the current goroutine that wants to get Lock until all other goroutines have
RLock and Lock released.

*/
