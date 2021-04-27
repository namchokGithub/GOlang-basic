package main

import (
	"fmt"
	// "../pkg/school"
)

/*
 * Name: main
 * Parameter: None
 * Return: None
 * Description: Main of program
 */
func main() {
	fmt.Println("------------- Main ---------------")
	// fmt.Println("School Address:", school.GetSchoolAddressSchool())
	// //** Declare variable
	// fmt.Println("---------------- Declare variable")
	// DeclareVariable()
	// //** Declare Function
	// fmt.Println("---------------- Declare Function")
	// printSchoolAddress()
	// printSchoolAddressWithParameter("ระยอง")
	// address := getSchoolAddress()
	// fmt.Println(address)
	// resultCode, resultAddress := getSchoolAddressReturn2Param()
	// fmt.Println(resultCode)
	// fmt.Println(resultAddress)
	// //** Loop
	// fmt.Println("---------------- Loop")
	// Forloop()
	// DoWhileloop()
	// Whileloop()
	// //** IF ELSE Condition
	// fmt.Println("---------------- IF ELSE Condition")
	// IfElseStatement()
	// //** Array
	// fmt.Println("---------------- Array")
	// SampleArray()
} // End main

/*
 * Name: SampleArray
 * Parameter: None
 * Return: None
 * Description: Sample of array
 */
func SampleArray() {
	// [การประกาศ Array เปล่าๆ ว่างๆ]
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
	// [การประกาศ Array แบบมีค่าเริ่มต้น]
	var name = [3]string{"Chaiyarin", "Atikom", "Kritsana"}
	fmt.Println(name)
	// การหา Size ของ array
	fmt.Println(len(name))

	// [การประกาศตัวแปร Slice]
	nameSlice := []string{}
	nameSlice = append(nameSlice, "Chaiyarin")
	nameSlice = append(nameSlice, "Atikom")
	nameSlice = append(nameSlice, "Kritsana")
	fmt.Println(nameSlice)

	// [การประกาศตัวแปร Map]
	m := make(map[string]int)
	m["chaiyarin"] = 1
	m["atikom"] = 2
	m["kritsana"] = 3
	fmt.Println(m)
	// [เราต้องการลบ key value ของ atikom ออก เราจะใช้คำสั่ง]
	delete(m, "atikom")
	fmt.Println(m)
} // End SampleArray

/*
 * Name: IfElseStatement
 * Parameter: None
 * Return: None
 * Description: Sample of if else condition
 */
func IfElseStatement() {
	// แบบที่ 1 if else แบบปกติ
	// [if else แบบปกติ]
	x := 10
	if x < 10 {
		println("x มีค่าน้อยกว่า 10")
	} else {
		println("x มีค่ามากว่า หรือ เท่ากับ 10")
	}
	// [if else if]
	y := 10
	if y < 10 {
		println("y มีค่าน้อยกว่า 10")
	} else if y == 10 {
		println("y มีค่าเท่ากับ 10")
	} else {
		println("y มีค่ามากกว่า 10")
	}

	// แบบที่ 2 if else แบบ ทำ Process บางอย่างก่อน Check เงื่อนไข
	i := 2
	j := 3
	k := 0
	k = i + j
	if k == 5 {
		println("k มีค่าเท่ากับ 5")
	}
	// สามารถทำได้แบบนี้
	i2 := 2
	j2 := 3
	k2 := 0
	if k2 = i2 + j2; k2 == 5 {
		println("k มีค่าเท่ากับ 5")
	}
} // End IfElseStatement

/*
 * Name: Forloop
 * Parameter: None
 * Return: None
 * Description: For loop
 */
func Forloop() {
	for i := 1; i < 9; i++ {
		println(i)
	}
} // End Forloop

/*
 * Name: DoWhileloop
 * Parameter: None
 * Return: None
 * Description: Do While loop
 */
func DoWhileloop() {
	i := 1
	for {
		if i == 2 {
			println(i)
			break
		}
		i++
	}
} // End DoWhileloop

/*
 * Name: Whileloop
 * Parameter: None
 * Return: None
 * Description: While loop
 */
func Whileloop() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}
} // End Whileloop

/*
 * Name: DeclareVariable
 * Parameter: None
 * Return: None
 * Description: Sample of declare variable
 */
func DeclareVariable() {
	var a, b, c bool
	var d, e, f bool = false, false, true
	var g, h, i = false, 10, "เรียนภาษา Go"
	j, k, l := false, 10, "เรียนภาษา Go"

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", h, h)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", l, l)
} // End DeclareVariable

//** Declare Function
//* Style 1
/*
 * Name: printSchoolAddress
 * Parameter: None
 * Return: None
 * Description: Print addresss of school
 */
func printSchoolAddress() {
	fmt.Println("กรุงเทพ")
} // End printSchoolAddress

/*
 * Name: printSchoolAddressWithParameter
 * Parameter: schoolAddress string
 * Return: None
 * Description: Print addresss of school
 */
func printSchoolAddressWithParameter(schoolAddress string) {
	println(schoolAddress)
} // End printSchoolAddressWithParameter

/*
 * Name: getSchoolAddress
 * Parameter: None
 * Return: string
 * Description: Get addresss of school
 */
func getSchoolAddress() string {
	return "กรุงเทพ"
}

/*
 * Name: getSchoolAddressReturn2Param
 * Parameter: None
 * Return: int, string
 * Description: Get addresss and code of school
 */
func getSchoolAddressReturn2Param() (int, string) {
	code := 1993
	address := "กรุงเทพ"
	return code, address
}
