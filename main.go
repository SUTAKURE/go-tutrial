package main

func main() {
	for i := 1; i <= 100; i++ {
		hantei(i)
	}
}

func hantei(x int) {
	print(x)
	if x%2 == 0 {
		println("-偶数")
	} else {
		println("-奇数")
	}
}
