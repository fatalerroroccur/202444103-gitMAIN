package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(32)
	fmt.Printf("%d점은 %s등급입니다.\n", target)

}
