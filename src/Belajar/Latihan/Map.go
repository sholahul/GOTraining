package Latihan

import "fmt"

func Map() {
	//define map versi 1
	var chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("======MAP VERSI 1")
	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	//define map fersi 2
	data := map[string]string{}
	data["nama"] = "Sholahul Fajri"
	data["password"] = "Fajri"

	fmt.Println("\n\n======MAP VERSI 2")
	fmt.Println(data)

	// define map versi 3
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	fmt.Println("\n\n======MAP VERSI 3")
	fmt.Print(chicken2)

	//define map versi 4
	fmt.Println("\n\n======MAP VERSI 4")
	var chicken4 = make(map[string]int)
	chicken4["Nomor1"] = 1
	chicken4["Nomor2"] = 2
	// fmt.Print("\n", chicken4)
	for key, val := range chicken4 {
		fmt.Println(key, ":", val)
	}

	//delete map
	delete(chicken4, "Nomor1")
	fmt.Println("Sisa : ", len(chicken4))

	//kombinasi map dan slice
	GetAllUser := []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	// fmt.Print(GetAllUser)
	for key, val := range GetAllUser {
		fmt.Println(key, ":", val)
	}
}
