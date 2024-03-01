package Latihan

import "fmt"

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silakan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}

func deferr() {
	defer fmt.Println("halo")
	fmt.Println("selamat datang")

	orderSomeFood("pizza")
	orderSomeFood("burger")
}

func Defer_Exit() {
	fmt.Println("===Fungsi Defer====")
	deferr()

	fmt.Println("\n\n===Fungsi Exit====")
}
