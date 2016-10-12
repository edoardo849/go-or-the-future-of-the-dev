package main

import "fmt"

// START OMIT
func main() {
	c := make(chan int64)
	defer close(c)

	go fib(c, 1, 1)

	for i := 0; i < 10; i++ {
		x := <-c
		fmt.Println(x)
	}
}

func fib(c chan int64, a int64, b int64) {
	c <- a
	go fib(c, b, a+b)
}

// END OMIT
