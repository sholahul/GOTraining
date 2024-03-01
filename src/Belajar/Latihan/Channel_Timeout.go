package Latihan

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}
}

func retreiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 10):
			fmt.Println("timeout. no activities under 10 seconds")
			break loop
		}
	}
}

func Chanell_Timeout() {
	runtime.GOMAXPROCS(5)

	var messages = make(chan int)

	go sendData(messages)
	retreiveData(messages)
}
