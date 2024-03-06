package day1

import "fmt"

//cap capasitas

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
}

func rangeandclose() {
	c := make(chan int, 10)

	fmt.Println("length : ", len(c))
	fmt.Println("cap : ", cap(c))

	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
