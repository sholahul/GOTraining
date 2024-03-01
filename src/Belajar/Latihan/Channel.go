package Latihan

import (
	"fmt"
	"runtime"
	"sync"
)

func Channel() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	var wg sync.WaitGroup

	var sayHelloTo = func(who string) {
		defer wg.Done()
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	wg.Add(4)
	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	go func() {
		wg.Wait()
		close(messages)
	}()
	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}
