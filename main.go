package main

import (
	"fmt"
)

func main() {
	a, t, _ := DiskStat(".")
	fmt.Printf("%.2f %.2f", float64(a/(1014*1024)), float64(t/(1024*1024)))
}
