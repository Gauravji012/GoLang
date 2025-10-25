package main

import (
	"fmt"
	"time"
)

//channels ka maiin purpose hai taking data from different end points
// taking & pputting valuese from different points
// it comes under pictures when go runs in multi-threading
// channel ki help se hm data ek goroutine se dusre goroutine mein data send kr skte h

/*
// normal functionality of channels with deadlock

// func main()  {
// 	messageChan := make(chan string)

//sending & receiving of the data behave like blocking calls

// 	messageChan <- "ping" // sending data to thee channels
// 	msg := <- messageChan // recevinng data into the channels
// 	fmt.Println(msg)
// 	 // at this place of code we are receving deadlock error

 }
*/

/*
//handling deadlock situation
// func processNum(numChan chan int) {
// 	fmt.Println("processing number ", <-numChan)
// }

// func main() {
// 	numChan := make(chan int)

// 	go processNum(numChan)

// 	numChan<-5
// 	time.Sleep(time.Second*2)

// }

*/

/*
// work as a queue // sending data to the go routine
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number ", num) // here <- is not required because range is handling that
		time.Sleep(time.Second)                // sleeep for 1 second
	}
}

func main() {
	numChan := make(chan int)

	go processNum(numChan) // make another thread
	for {                  // this is basically an infinte loop that gives random value between 0 to 100
		numChan <- rand.Intn(100)
	}
	// numChan<-5
	// time.Sleep(time.Second*2)

}
*/

/*
// receving data

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}
func main() {
	result := make(chan int)
	go sum(result, 4, 5)
	res := <- result
	fmt.Println(res)
}
*/

/*

//go routine synchronizer

// channel work as synchronization
// we can acheive the same thing as we are acheving through wait group

// lets understand this concept

// we can use anything betweeen gorroutines & channels
// but generally when we havee only 1 goroutine ->>> we use channel concept otherwise we use goroutines wait group.
func task(done chan bool) {
	defer func() { done <- true }() // inline function

	fmt.Println("processing ...")
}
func main() {
	done := make(chan bool)
	go task(done)

	<- done // blockinig as like wait of wait group
}
*/

/*
// getting data in buffer instead blocking
// previously we got the deadlock error because of buffer.
*/
/*
func main() {
	emailChan := make(chan string, 100 )// 100 is the size of buffer
	emailChan <- "1@gmail.com" // sending
	emailChan <- "2@gmail.com"

	fmt.Println(<- emailChan) // receving
	fmt.Println(<- emailChan)
}
*/

/* creating buffer queues of 100 email
 */
/*
func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true} () // at last of the execution it is calling

	for email := range emailChan {
		fmt.Println("sending email to : ", email)
		time.Sleep(time.Second)
	}
}
func main() {
	emailChan := make(chan string, 100 )// 100 is the size of buffer
	done := make(chan bool)

	go emailSender(emailChan, done)
	for i := 1; i <= 100; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("done sending !!")

	close(emailChan) // closing the channeel is a good practice // otherwise it may creatte a deadlock issue

	<- done

}
*/

/*
when we have multiple channels to send & receive the data

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()


	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <- chan1 :
			fmt.Println("received data from chan1 ", chan1Val)
		case chan2Val := <- chan2 :
			fmt.Println("received data from chan2 ", chan2Val)
		}
	}
}
*/

// make only sending & receiving channel with the help of sign
func emailSender(emailChan <-chan string, done chan<- bool) { // emailChan is only receving channel &&&& done channel is only sending channel

	defer func() { done <- true }() // at last of the execution it is calling

	for email := range emailChan {
		fmt.Println("sending email to : ", email)
		time.Sleep(time.Second)
	}
}
func main() {
	emailChan := make(chan string, 100) // 100 is the size of buffer
	done := make(chan bool)

	go emailSender(emailChan, done)
	for i := 1; i <= 100; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("done sending !!")

	close(emailChan) // closing the channeel is a good practice // otherwise it may creatte a deadlock issue

	<-done

}
