package main

import (
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	random := rand.Intn(6)
	println(random)
}
