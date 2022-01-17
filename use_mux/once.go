package main

import (
    "fmt"
    "sync"
)

var once *sync.Once = new(sync.Once)

func doSomething(){
    fmt.Println("DO SOMETHING")
}

func main(){
    ch := make(chan byte, 1)
    for i:= 0 ; i < 10; i++ {
        go func(){
            fmt.Println("test", i)
            once.Do(doSomething)
            ch <- 1
        }()
    }
    fmt.Println("Finished")
    <- ch
}
