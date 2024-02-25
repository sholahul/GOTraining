package Latihan

import "fmt"

func Variadic() {
	fmt.Println("Fungsi Variadic")
	num := Hitungs(3, 4, 5, 6, 7, 8, 9)

	fmt.Println("Nilai : ", num)
}

func Hitungs(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
