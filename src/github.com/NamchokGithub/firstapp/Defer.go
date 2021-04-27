package main

import (
	"fmt"
)

func main() {
	// printFirst()
	// printFinish()
	// printSecond()
	// fmt.Println("--------------------------")
	// printFirst()
	// defer printFinish()
	// printSecond()
	fmt.Println("--------------------------")
	defer printFirst()
	defer printFinish()
	defer printSecond()
}
func printFirst() {
	fmt.Println("First")
}
func printSecond() {
	fmt.Println("Second")
}
func printFinish() {
	fmt.Println("Close")
}
