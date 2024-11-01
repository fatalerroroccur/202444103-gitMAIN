package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isprime(n int) bool {

	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		for j := 3; j <= int(math.Sqrt(float64(n))); j += 2 {
			if n%j == 0 {
				return false
			}
			fmt.Printf("%d ", j)
		}

	}
	return true
}

func main() {
	fmt.Print("정수 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	if isprime(n) {
		fmt.Printf("\n소수 입니다.")
	} else {
		fmt.Printf("\n소수가 아닙니다.")
	}

}
