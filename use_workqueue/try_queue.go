package main

import (
	"fmt"
	"math/rand"
	"time"
)

// type task struct {
// 	id int
// }

// var wq *workqueue.Type

// func try_queue() {
// 	s1 := rand.NewSource(time.Now().UnixNano())
// 	r1 := rand.New(s1)

// 	// struct 空值可用, struct 可以用来调用绑定在struct 指针上的method？
// 	var wg sync.WaitGroup

// 	tasks := 5
// 	wq = workqueue.New()

// 	for i := 0; i < tasks; i++ {
// 		wg.Add(1)
// 		work := i

// 		go func() {
// 			// mimic a work load
// 			time.Sleep(time.Duration(r1.Intn(10)) * time.Second)

// 			wq.Add(task{id: work})
// 			fmt.Println("add work:", work)
// 			wg.Done()
// 		}()
// 	}

// 	i := 1

// 	for {

// 		task, fin := wq.Get()
// 		fmt.Println("get task", task, "finish is", fin, "Len:", wq.Len())
// 		wq.Done(task)

// 		if i == tasks {
// 			wq.ShutDown()
// 		}
// 		i++

// 		if wq.ShuttingDown() {
// 			fmt.Println("finish", wq.ShuttingDown())
// 			break
// 		}
// 	}

// 	wg.Wait()
// }

func try_queue() {
	stop := make(chan int, 1)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	tasks := 5

	for i := 0; i < tasks; i++ {
		work := i

		go func() {
			// mimic a work load
			time.Sleep(time.Duration(r1.Intn(10)) * time.Second)

			wq.Add(task{id: work})
			fmt.Println("add work:", work)
		}()
	}

	i := 1

	for {
		task, fin := wq.Get()
		fmt.Println("get task", task, "finish is", fin, "Len:", wq.Len())
		wq.Done(task)

		if i == tasks {
			wq.ShutDown()
		}
		i++

		if wq.ShuttingDown() {
			fmt.Println("finish", wq.ShuttingDown())
			break
		}
	}
	<-stop
}
