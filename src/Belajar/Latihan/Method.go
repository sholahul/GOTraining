package Latihan

import "fmt"

type MyString string

type datas struct {
	username string
	password string
}

func Method() {
	fmt.Println("Method")
	var s MyString = "Kirito"
	s.Hallo()

	var d = datas{"sholahul", "ganteng"}
	d.Login_Awal()
}

func (s MyString) Hallo() {
	fmt.Println("Hi, Welcome ", s)
}

func (d datas) Login_Awal() {
	fmt.Println("Halo bos : ", d.username)

}
