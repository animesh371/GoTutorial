package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func bufferOverflow(ch chan int) {
	ch <- 3
	ch <- 5
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	//s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 2)
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c // receive from c

	// fmt.Println(x, y, x+y)
	// bufferOverflow(c)
	// x, y, z := <-c, <-c, 4
	// fmt.Println(x, y, z)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
