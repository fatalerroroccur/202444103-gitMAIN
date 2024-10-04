package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 55
	f := 12.9
	var b bool
	fmt.Printf("valuse i : %d, value f : %f\n", i, f)
	//fmt.Printf("%d + = %f", i, i)
	fmt.Printf("%f\n", float64(i)*f)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(float64(i)))
	fmt.Printf("%t", b)
}
