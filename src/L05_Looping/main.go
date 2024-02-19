package main

import "fmt"

func main() {
	// using for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Printf("\nBilangan Genap : ")
	var j = 1
	for j < 10 {
		if j%2 == 0 {
			fmt.Print(j, " ")
		}
		j++
	}

	// for tanpa argumen
	var k = 0
	fmt.Print("\nBilangan Ganjil : ")
	for {
		if k%2 == 1 {
			fmt.Print(k, " ")
		}
		k++
		if k == 10 {
			break
		}
	}

	fmt.Printf("\nBintang Naik:")

	for i := 0; i < 10; i++ {
		for j = 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println("\nBintang Turun:")

	for i := 10; i > 0; i-- {
		for j = 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println("\nBintang Segitiga")
	rows := 5
	for A := 1; A <= rows; A++ {
		// inner loop spaces before stars
		for B := 1; B <= rows-A; B++ {
			fmt.Print("  ")
		}

		//inner loop for printing starts
		for C := 1; C <= 2*A-1; C++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	// using do

	// using do while

}
