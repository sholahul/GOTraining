package Latihan

import (
	"fmt"
	"math"
)

func Functione() {
	a := 100
	b := 300

	fmt.Println("Nilai Yang paling Kecil adalah : ", Lower(a, b))

	fmt.Println()
	var diameter float64 = 10
	var area, circumference = Calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	angka := []int{3, 9, 1, 4, 9, 1, 2, 33, 4, 3, 11, 22}

	fmt.Println("Min Max dari : ")
	fmt.Print(angka)

	min, max := GetMinMax(angka)
	fmt.Println("\nMin : ", min)
	fmt.Println("Max : ", max)

}

func Lower(nilai1 int, nilai2 int) int {
	lowest := 0
	if nilai1 < nilai2 {
		lowest = nilai1
	} else {
		lowest = nilai2
	}
	return lowest
}

func Calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

func GetMinMax(number []int) (int, int) {
	min := number[0]
	max := number[0]

	for i := 1; i < len(number); i++ {
		if number[i] < min {
			min = number[i]
		}

		if number[i] > max {
			max = number[i]
		}
	}

	return min, max
}
