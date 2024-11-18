package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func addContextValue(ctx context.Context) {
	fmt.Println("Result", ctx.Value("val"))
}

func ReadFile(path string) {

	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

// func worker(ch chan string) {
// 	time.Sleep(1 * time.Second)
// 	ch <- "new message"
// }

// func hello(msg string) {
// 	fmt.Println(msg)
// }

//	func channelWorker(ch chan string){
//	    ch <- ""
//	}

func longTaskWithoutChannel() {
	time.Sleep(5 * time.Second)
	fmt.Println("Finished long task")
}

// var done = make(chan bool)

// func longTask() {
// 	time.Sleep(5 * time.Second)
// 	done <- true
// }

func publish(msg chan string) {
	for i := 0; i < 5; i++ {
		msg <- fmt.Sprintf("data -> %d", i)
		time.Sleep(500 * time.Millisecond)
	}

}

func subscribe(msg chan string) {
	for {
		fmt.Println("Subscribe: ", <-msg)
	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Started", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Finised", id)

}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	// wg.Add(2)
	// wg.Add(1)
	// go worker(1, &wg) // counter = 1

	// go worker(2, &wg) // counter = 0

	// go worker(3, &wg)

	// wg.Wait()

	fmt.Println("Done")
	// msg := make(chan string)
	// defer close(msg)
	// go publish(msg)
	// go subscribe(msg)

	// <-msg
	// time.Sleep(5 * time.Second)
	// // defer close(done)

	// // go longTask()

	// // <-done

	// fmt.Println("Done")
	// go longTaskWithoutChannel()

	// time.Sleep(6 * time.Second)

	// flag := make(chan int, 1)

	// flag <- 5
	// flag <- 2

	// num := <-flag

	// num2 := <-flag
	// fmt.Println(num)
	// fmt.Println(num2)
	// msg := make(chan string)

	// for i := 0; i < 10; i++ {
	// 	go hello(fmt.Sprintf("hello %d", i))
	// }

	// time.Sleep(500 * time.Millisecond)

	// msg := make(chan string)
	// go worker(msg)

	// receive := <-msg

	// fmt.Println(receive)

	// ctx := context.Background() // until kill process
	// // ctxWithCancel, cancel := context.WithCancel(ctx)                   // until called cancel
	// // ctxWithTimeout, timeout := context.WithTimeout(ctx, 1*time.Minute) // until timeout has been called or 1 minute
	// ctxWithValue := context.WithValue(ctx, "val", 10)

	// addContextValue(ctxWithValue)
	// addContextValue(ctx)

	// ReadFile("../.env")

	// pkg.Second()

	// fmt.Println(pkg.Add())

	// var value int = 10
	// var val *int = IncreaseNestedWithPointer(&value)
	// var val1 *int = IncreaseNestedWithPointer(val)
	// var final *int = IncreaseNestedWithPointer(val1)

	// IncreaseByOneWithPointer(&value)
	// IncreaseByOneWithPointer(&value)
	// IncreaseByOneWithPointer(&value)

	// fmt.Println(value)
	// fmt.Println(&value)

	// fmt.Println("--------- val")

	// fmt.Println(val)
	// fmt.Println(&val)
	// fmt.Println(*val)

	// fmt.Println("--------- val1")

	// fmt.Println(val1)
	// fmt.Println(&val1)
	// fmt.Println(*val1)
	// fmt.Println("--------- final")

	// fmt.Println(final)
	// fmt.Println(&final)
	// fmt.Println(*final)
	// fmt.Println("---------")

	// IncreaseByOnwWithoutPointer(value)

	// fmt.Println(value)

}
