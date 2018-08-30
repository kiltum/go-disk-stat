// +build !windows

package main

import (
	unx "golang.org/x/sys/unix"
)

func DiskFree(path string) (uint64, error) {
	fs := unx.Statfs_t{}
	err := unx.Statfs(path, &fs)
	if err != nil {
		return 0, err
	}
	return fs.Bfree * uint64(fs.Bsize), nil
}
