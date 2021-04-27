// [การทำ List จาก Array หรือ Slice เป็น JSON]
package main

import (
	"encoding/json" // มีการ Import Encoding JSON เข้ามา "fmt"
	"fmt"
)

type customer struct { // มีการปรับ Field ให้ขึ้นต้นตัวใหญ่
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func main() {
	customerLists := []customer{}
	customer1 := customer{
		Firstname: "Chaiyarin",
		Lastname:  "Niamsuwan",
		Code:      111990,
		Phone:     "085661234",
	}
	customer2 := customer{
		Firstname: "Atikom",
		Lastname:  "Sombutjalearn",
		Code:      111991,
		Phone:     "085664321",
	}
	customer3 := customer{
		Firstname: "Kritsana",
		Lastname:  "Punyaphon",
		Code:      111992,
		Phone:     "085662344",
	}
	customerLists = append(customerLists, customer1)
	customerLists = append(customerLists, customer2)
	customerLists = append(customerLists, customer3)
	customerListsJson, err := json.Marshal(customerLists)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(customerListsJson))
}
