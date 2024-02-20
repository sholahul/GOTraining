package main

import (
	"fmt"
)

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

	//for - range
	var xs = "123"
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	fmt.Println()
	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	fmt.Println()
	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}

	fmt.Println()
	for range kvs {
		fmt.Println("Done")
	}

	//BREAK AND CONTINUE
	fmt.Println()
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}
	fmt.Println()
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

}
