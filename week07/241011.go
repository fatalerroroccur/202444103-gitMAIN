package week07

import (
    "fmt"
    "time"
)

func main () {
	var now time.Time = time Now()
	var year int = now.Year()
	fmt.Printin(year)
}