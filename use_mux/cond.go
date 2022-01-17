package main

import (
    "sync"
    "time"
    "fmt"
)

/*
cond.L.Lock() and cond.L.Unlock(): lock() and lock.Unlock() can also be used, exactly the same.

cond.Wait(): When this method is called, the outflow of the executed
operation is: Unlock() -> blocking waiting notification (i.e. waiting for
notification from Signal() or Broadcast()) -> receiving notification ->
Lock().

cond.Signal(): Notify a Wait goroutine, if there is no Wait(), no error
will be reported, Signal() notification order is based on the original join
notification list (Wait()) first in first out.

cond.Broadcast(): Notify all Wait goroutines, if there is no Wait(), no
error will be reported.

*/

func main() {
    var cond *sync.Cond = sync.NewCond(new(sync.Mutex))
    var condition int = 0

    // Consumer
    go func() {
        for {
            cond.L.Lock()
            for condition == 0 {
                cond.Wait()
            }
            condition = condition - 1
            fmt.Println(condition)
            cond.Signal()
            cond.L.Unlock()
        }
    }()
    // Producer
    for {
        time.Sleep(time.Second)
        cond.L.Lock()
        for condition == 3 {
            cond.Wait()
        }
        condition = condition + 1
        fmt.Println(condition)
        cond.Signal()
        cond.L.Unlock()
    }
}
