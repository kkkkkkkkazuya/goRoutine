package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"testing"
	"time"
)

func hello(done chan bool) {
	fmt.Println("goroutine")
	done <- true
}

// func process(num int, str string) {
// 	for i := 0; i <= num; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(num, str)
// 	}
// }

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
	// process(1, "A")
	// process(1, "B")
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

	result := testing.Benchmark(func(b *testing.B) { run("A", "B", "C", "D", "E") })
	fmt.Println(result.T)
}

func run(name ...string) {
	fmt.Println("Start!")
	// WaitGroupを作成する
	wg := new(sync.WaitGroup)

	// channelを処理の数分だけ作成する
	isFin := make(chan bool, len(name))

	for _, v := range name {
		// 処理1つに対して、1つ数を増やす（この例の場合は5になる）
		wg.Add(1)
		// サブスレッドに処理を任せる
		go process(v, isFin, wg)
	}

	// wg.Doneが5回通るまでブロックし続ける
	wg.Wait()
	close(isFin)
	fmt.Println("Finish!")
}

func process(name string, isFin chan bool, wg *sync.WaitGroup) {
	// wgの数を1つ減らす（この関数が終了した時）
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println(name)
	isFin <- true
}
