package main

import "fmt"

func main() {
	for i := 1; i <= 100; i = i + 1 {
		if i%2 == 0 {
			fmt.Println(i)
			fmt.Println("偶数")
		} else {
			fmt.Println(i)
			fmt.Println("奇数")
		}
	}
}
