package Latihan

import "fmt"

type MyString string

func Method() {
	fmt.Println("Method")
	var s MyString = "Kirito"
	s.Hallo()

}

func (s MyString) Hallo() {
	fmt.Println("Hi, Welcome ", s)
}
