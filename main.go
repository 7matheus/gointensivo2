package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	canal := make(chan int) // Canal de comunicação entre as threads

	// T2
	go publish(canal)

	// T3
	go reader(canal)

	time.Sleep(time.Second * 5)
}

func reader(canal chan int) {
	for x := range canal {
		// value := <-canal
		fmt.Println(x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
