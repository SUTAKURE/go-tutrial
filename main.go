package main

func main() {
	n, m := swap(10, 20)
	println(n, m)
}

func swap(x, y int) (x2, y2 int) {
	x2, y2 = y, x
	return
}
