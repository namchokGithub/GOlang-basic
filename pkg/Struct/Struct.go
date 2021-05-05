package Struct

import "fmt"

type customer struct { // การประกาศโครงสร้าง struct
	firstname string
	lastname  string
	code      int
	phone     string
}

func Kstruct() {
	cus := customer{
		firstname: "Chaiyarin",
		lastname:  "Niamsuwan",
		code:      111990,
		phone:     "085661234",
	} // การกำหนดค่าเริ่มต้น ให้ customer struct
	fmt.Println(cus)
	fmt.Println(cus.firstname)
	fmt.Println(cus.lastname)
	fmt.Println(cus.code)
	fmt.Println(cus.phone)
	cus.firstname = "Atikom"
	cus.lastname = "Sombutjalearn"
	fmt.Println(cus)

	// [การใช้ Struct กับ Array]
	customerLists := [3]customer{}
	customerLists[0] = customer{
		firstname: "Chaiyarin",
		lastname:  "Niamsuwan",
		code:      111990,
		phone:     "085661234",
	}
	customerLists[1] = customer{
		firstname: "Atikom",
		lastname:  "Sombutjalearn",
		code:      111991,
		phone:     "085664321",
	}
	customerLists[2] = customer{
		firstname: "Kritsana",
		lastname:  "Punyaphon",
		code:      111992,
		phone:     "085662344",
	}
	fmt.Println(customerLists)

	//[การใช้ Struct กับ Slice]
	customerLists2 := []customer{}
	customer12 := customer{
		firstname: "Chaiyarin",
		lastname:  "Niamsuwan",
		code:      111990,
		phone:     "085661234",
	}
	customer22 := customer{
		firstname: "Atikom",
		lastname:  "Sombutjalearn",
		code:      111991,
		phone:     "085664321",
	}
	customer32 := customer{
		firstname: "Kritsana",
		lastname:  "Punyaphon",
		code:      111992,
		phone:     "085662344",
	}
	customerLists2 = append(customerLists2, customer12)
	customerLists2 = append(customerLists2, customer22)
	customerLists2 = append(customerLists2, customer32)
	fmt.Println(customerLists)
}
