package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	i := 55
	var f float32 = 4.30
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Printf("%s\n", strings.Title("inhatc"))
	fmt.Println(math.Ceil(3.99))
	fmt.Printf("value i : %d\n", i)
	fmt.Print("value i : ", i, "\n")
	fmt.Println("value i :", i)
}
