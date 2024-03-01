package Latihan

import "fmt"

//define struct using func init()
var Pelajar = struct {
	Name  string
	Grade string
}{}

//func init pertama kali akan dicall pada saat program berjalan
func init() {
	Pelajar.Name = "Kirigaya Kazuto"
	Pelajar.Grade = "SSS"
}

//Huruf kapital diawal menandakan public
func Public_Private() {
	fmt.Println("hello")
	introduce("fajri")

	fmt.Println("\n\nFungsi init : ")
	fmt.Println("Name : ", Pelajar.Name)
	fmt.Println("Grade : ", Pelajar.Grade)
}

//Huruf kecil diawal menandakan private
func introduce(name string) {
	fmt.Println("nama saya", name)
}
