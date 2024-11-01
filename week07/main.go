package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	var month int = int(now.Month())
	var day int = int(now.Day())
	var hour int = int(now.Hour())
	var min int = int(now.Minute())
	var sec int = int(now.Second())
	fmt.Printf("오늘은 %d년 %d월 %d일입니다.\n", year, month, day)
	fmt.Printf("지금은 %d시 %d분 %d초입니다.\n", hour, min, sec)
	fmt.Println(now.Month())
}
