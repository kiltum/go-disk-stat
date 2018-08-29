package main

import (
	"fmt"
)

func main() {
	a, _ := DiskFree(".")
	fmt.Printf("%d", a)
}
