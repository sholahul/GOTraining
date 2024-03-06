package day1

import "fmt"

func ScaleFunc(v *vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Indirection() {
	v := vertex{3, 4}
	v.Scale(2)
	fmt.Println(v)

	ScaleFunc(&v, 10)
	fmt.Println(v)

	fmt.Println("--------------")

	q := &vertex{2, 3}
	q.Scale(3)
	fmt.Println(q)
}
