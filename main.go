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
	canal := make(chan string) // Canal de comunicação entre as threads

	// Thread 2
	go func() {
		canal <- "Olá Mundo!"
		canal <- "Olá Mundo 2!"
	}() // Função Anônima

	// Thread 1
	// msg := <-canal
	fmt.Println(<-canal)
	fmt.Println(<-canal)
}
