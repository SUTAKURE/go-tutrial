package main

import "fmt"

func main() {
	fmt.Println("price")

	num := 1
	switch num {
	case 1:
		fmt.Printf("I \n")
		fallthrough
	case 2:
		fmt.Printf("am \n")
		fallthrough
	case 3:
		fmt.Printf("yyh-gl \n")
	}
}
