package main

import (
	"fmt"
	"time"
)

func process(num int, str string) {
	for i := 0; i <= num; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(num, str)
	}
}

func main() {
	fmt.Println("Start")
	go process(1, "A")
	go process(1, "B")
	fmt.Println("Finish")
}
