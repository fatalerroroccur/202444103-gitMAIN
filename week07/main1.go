package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#ck"
	alpha := "l#l"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Printf("%s\n", fixed)
	fmt.Printf(replacer.Replace(alpha))
}
