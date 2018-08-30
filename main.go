package main

import (
	"fmt"
)

func main() {
	a, t, _ := DiskStat(".")
	fmt.Printf("%d %d", a, t)
}
