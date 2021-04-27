package main

import (
	"fmt"
)

/*
 * Name: main
 * Parameter: None
 * Return: None
 * Description: Main of program
 */
func main() {
	//** Declare variable
	fmt.Println("---------------- Declare variable")
	DeclareVariable()
	//** Declare Function
	fmt.Println("---------------- Declare Function")
	printSchoolAddress()
	printSchoolAddressWithParameter("ระยอง")
	address := getSchoolAddress()
	fmt.Println(address)

} // End main

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
