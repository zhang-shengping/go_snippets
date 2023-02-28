package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("end")

	// try_queue()
	// try_queue_others()
	try_queue2()
}
