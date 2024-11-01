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

func main() {
	fmt.Print("정수 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	var isPrime bool = true

	if n < 2 {
		isPrime = false
	} else {
		if n%2 == 0 {
			fmt.Printf("짝수")
			isPrime = false
		} else {
			for j := 3; j <= int(math.Sqrt(float64(n))); j += 2 {
				if n%j == 0 {
					isPrime = false
					break
				}
				fmt.Printf("%d ", j)
			}

		}
	}
	if n == 2 {
		isPrime = true
	}
	if isPrime {
		fmt.Printf("\n소수 입니다.")
	} else {
		fmt.Printf("\n소수가 아닙니다.")
	}

}
