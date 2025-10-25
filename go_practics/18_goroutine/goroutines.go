package main

import (
	"fmt"
	"sync"
	// "time"
)

//goroutines ==> light weight threadds / run with multiple threads

// func task(id int) {
// 	fmt.Println("doing task for ::", id)
// }

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		// task(i) blocking system call i.e. it is exeecuting 0 to 10 by sequence
// 		// if we want to run it in asynchronously then we just add go. main() is also run in async mode so when we just add go before this line then it schedule it schelduler but before exeecuting our main() completes its execution so we couldnot see any result. for this we add sleep method of time package and see the result

// 		go task(i) // just add go to apply multi-threadding concept

// 		go func(i int) {
// 			fmt.Println(i)
// 		}(i)

// 	}

// **********************
// 	time.Sleep(time.Second * 1) // generally we do not use slleep method like this because we do not know about the time it will take to finish the execution. So we have the concept oof waitgroup in go-routine.

// }

//wait-group concept for blocking sync untill the current program execution completes its exeecution

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task for : ", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)

	}
	// time.Sleep(time.Second * 1)
	wg.Wait()
}
