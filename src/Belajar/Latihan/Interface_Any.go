package Latihan

import (
	"fmt"
	"strings"
)

type persons struct {
	name string
	age  int
}

func Any_interface() {
	//tipe data any
	fmt.Println("============Interface===========")
	tipe_any()

	fmt.Println("\n\n============Any===========")
	tipe_interface()

	fmt.Println("\n\n============Casting Interface===========")
	casting_interface()

	fmt.Println("\n\n============Casting Pointer===========")
	casting_pointer()

	fmt.Println("\n\n============Combine interface with slice and nmap ===========")
	combine_nmap_slice_interface()
}

func tipe_any() {
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	var data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println()
	fmt.Println(data)
}

func tipe_interface() {
	var data = map[string]any{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}
	fmt.Println(data)
	for key, val := range data {
		fmt.Println(key, ":", val)
	}

	// Print key and val specifically for "breakfast"
	if breakfast, ok := data["breakfast"].([]string); ok {
		fmt.Println("Breakfast:")
		for i, item := range breakfast {
			fmt.Printf("%d: %s\n", i+1, item)
		}
	} else {
		fmt.Println("Breakfast not found or not of type []string")
	}
}

func casting_interface() {
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")
}

func casting_pointer() {
	var secret interface{} = &persons{name: "wick", age: 27}
	var name = secret.(*persons).name
	var age = secret.(*persons).age
	fmt.Println("Name :", name)
	fmt.Println("Age :", age)
}

func combine_nmap_slice_interface() {
	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10}, //map
		[]string{"manggo", "pineapple", "papaya"},                 //array
		"orange", //string
	}

	for _, each := range fruits {
		fmt.Println(each)
	}

}
