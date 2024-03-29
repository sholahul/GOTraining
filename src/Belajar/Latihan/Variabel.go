package Latihan

import "fmt"

func Variabell() {
	/*format deklrasi
	var <nama-variabel> <tipe-data>
	var <nama-variabel> <tipe-data> = <nilai>
	*/

	//deklrasi string
	var firstName string = "john"
	var lastName string = "wick"

	// deklrasi tanpa var
	middlename := "pantau"

	//deklrasi multi var
	var awal, tengah, akhir string
	awal = "sholahul"
	tengah = "fajri"
	akhir = "tsi"

	// multiple var
	awal1, tengah1, akhir1 := "Senang", "Berpetualang", "Jungler"

	//multiple var string, integer, float
	integ, float, str := 1, 1.1, "stringv"

	//var boolean
	var bool1 bool = false
	bool2 := true

	//define string
	var message = `Nama saya adalah fajri.
	Bagian Pengembangan TSI.
	Pengelolaan core banking`

	//undefined variabel string blackhole bangke wkwk
	_ = "hayo"

	//pointer
	name := new(string)
	name2 := "jono"

	name = &name2

	//int dan uint
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	//float32 dan float64
	var decimalNumber = 2.62

	//konstanta (tidak dapat diubah)
	const phi = 3.14

	//multiple konstanta
	const (
		square         = "kotak"
		istoday  bool  = true
		numeric  uint8 = 1
		floatnum       = 2.2
	)

	//multiple konstanta dengan nilai sama
	const (
		nilai1 = 0
		nilai2
	)

	//multiple konstanta gabungan
	const (
		today string = "senin"
		sekarang
		istoday2 = true
	)

	//multiple konstanta in row
	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"

	fmt.Printf("halo %s %s %s!\n", firstName, middlename, lastName)
	fmt.Printf("halo %s %s %s!\n", awal, tengah, akhir)
	fmt.Printf("halo %s %s %s!\n", awal1, tengah1, akhir1)
	fmt.Printf("halo %d %f %s!\n", integ, float, str)

	// using %t untuk boolean dan printf, tidak dapat menggunakan Println
	fmt.Printf("terkadang %t dan %t \n", bool1, bool2)
	fmt.Println(message)

	//concat string menggunakan printf
	// fmt.Printf("%s", awal, +"ganteng")
	fmt.Println(name2, name)
	fmt.Printf("nama : %s \n", name2)

	// print pointer using *
	fmt.Printf("nama : %s", *name)

	fmt.Printf("\nbilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	fmt.Printf("bilangan phi : %.2f\n", phi)
	fmt.Printf("Square : %s, Istoday : %t; Numeric : %d, floatnum : %.2f", square, istoday, numeric, floatnum)

	fmt.Printf("\nnilai %d dan %d\n", nilai1, nilai2)

	fmt.Printf("Sekarang %s dan %s pasti %t", today, sekarang, istoday2)

	fmt.Printf("\nnilai : %d %d %s %s", satu, dua, three, four)
}
