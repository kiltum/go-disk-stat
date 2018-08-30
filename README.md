# go-disk-stat

This is very simple module for getting available and total 
disk space from Go programs.

Tested on Linux, Windows and OS X

## Using
```go

import (
	ds "github.com/kiltum/go-disk-stat"
	"fmt"
)

func main() {
	a, t, _ := ds.DiskStat(".")
	fmt.Printf("Available %.2f Total %.2f", ds.ToGb(a), ds.ToMb(t))
}

```