package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func hello(done chan bool) {
	fmt.Println("goroutine")
	done <- true
}

func process(num int, str string) {
	for i := 0; i <= num; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(num, str)
	}
}

func greeting(data string) {
	message := make(chan string)
	go func() {
		message <- data
	}()
	msg := <-message
	fmt.Println(msg)
}

func flow(num int, str string) {
	for i := 0; i <= num; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, str)
	}
}

func main() {
	fmt.Println("Start")
	go process(1, "A")
	go process(1, "B")
	fmt.Println("Finish")
	log.Println(runtime.NumGoroutine())
	greeting("hello")

	done := make(chan bool)

	go hello(done)

	<-done
	fmt.Println("last")

	fmt.Println("Start2!")
	go process(2, "A")
	go process(2, "B")
	fmt.Println("Finish2!")
}
