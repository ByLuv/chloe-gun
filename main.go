package main

import (
	"fmt"
	"time"

	"github.com/is5/r"
)

func main() {
	fmt.Println("Yes? Go!")

	Seed(time.Now().Second())
	for _, i := range r.R(1, 10) {
		fmt.Printf("#%03d - %d\n", i, Random())
	}
}
