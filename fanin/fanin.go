package main

import (
	"fmt"
	"time"
)

func main() {
	ch := fanIn(generator("Hello"), generator("Bye"))

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}

}

// fanIn is itself a generator
func fanIn(ch1, ch2 <-chan string) <-chan string { // receives two read-only channels
	new_ch := make(chan string)

	go func() {
		for {
			select {
			case s := <-ch1:
				new_ch <- s
				fmt.Println("ch1")
			case s := <-ch2:
				new_ch <- s
				fmt.Println("ch2")
				//	new_ch <- <-ch1
			}
		}

	}() // launch two goroutine while loops to continuously pipe to new channel
	/*
		go func() {
			for {
				fmt.Println("ch2")
				new_ch <- <-ch2
			}

		}()
	*/

	return new_ch

}

func generator(msg string) <-chan string { // returns receive-only channel
	ch := make(chan string)
	fmt.Println("msg ", msg)
	go func() { // anonymous goroutine
		for i := 0; ; i++ {

			fmt.Println("generator", msg)

			ch <- fmt.Sprintf("%s %d", msg, i)

			time.Sleep(time.Second)
		}
	}()
	return ch
}
