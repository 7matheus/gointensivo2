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
	go task("Tarefa 1")
	go task("Tarefa 2")
	task("Tarefa 3")
}
