package main

import "fmt"

func main() {
	var names = [4]string{"apple", "pie", "banana", "grepe"}

	fmt.Println("Panjang array name : ", len(names))

	for i := 0; i < len(names); i++ {
		fmt.Println("Array Ke-", i, ": ", names[i])
	}
}
