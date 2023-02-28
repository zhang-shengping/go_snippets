package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"k8s.io/client-go/util/workqueue"
)

type task struct {
	id int
}

var wq *workqueue.Type

func try_queue2() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// struct 空值可用, struct 可以用来调用绑定在struct 指针上的method？
	var wg sync.WaitGroup

	tasks := 5
	wq = workqueue.New()
	wg.Add(1)
	go func() {
		for i := 0; i < tasks; i++ {
			// mimic a work load
			time.Sleep(time.Duration(r1.Intn(10)) * time.Second)

			wq.Add(task{id: i})
			fmt.Println("add work:", i)
		}
		wq.ShutDown()
	}()

	go func() {

		for {
			task, fin := wq.Get()
			fmt.Println("get task", task, "finish is", fin, "Len:", wq.Len())
			wq.Done(task)
			if fin {
				break
			}
		}

		if wq.ShuttingDown() {
			fmt.Println("finish", wq.ShuttingDown())
			wg.Done()
			return
		}
	}()

	wg.Wait()
}
