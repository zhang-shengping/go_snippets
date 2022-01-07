package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    sigs :=make(chan os.Signal, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    done := make(chan bool, 1)

    go func(){
        sig := <- sigs
        fmt.Println()
        fmt.Println(sig)
        // 可以使用 true/false, 只有一处可以收到消息
        // done <- true
        // done <- false
        // 可以使用 close, close 的好处是所有订阅 <-done 的地方都可以收到消息
        close(done)
        // 如果以上两者都不使用，那么 main 中的 done block
    }()

    fmt.Println("awaiting singal")
    <-done
    <-done
    // test:=<-done
    // fmt.Println(test)
    fmt.Println("exiting")
}
