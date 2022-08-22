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
	process(1, "A")
	process(1, "B")
	fmt.Println("Finish")
	greeting("hello")

	done := make(chan bool)

	go hello(done)

	<-done
	fmt.Println("last")

	ch1 := make(chan bool)
	ch2 := make(chan bool)

	fmt.Println("Start2!")

	go func() {
		flow(2, "A")
		ch1 <- true
	}()

	go func() {
		flow(2, "B")
		ch2 <- true
	}()

	<-ch1
	<-ch2

	fmt.Println("Finish2!")
	log.Println(runtime.NumGoroutine())

}
