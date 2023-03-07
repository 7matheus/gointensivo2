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

	go publish(canal)

	for x := range canal {
		fmt.Println(x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
