package main

import (
	"fmt"
	"time"
)

func main() {
	//var dates [3]time.Time{time.unix(1,0),time.unix(1447920000,0),time.unix()}
	//fmt.Println{dates[1]}
	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1947920001, 0),
	}
	fmt.Println(dates[0], dates[1], dates[2])
}
