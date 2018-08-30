package main

import "testing"

func TestDiskStat(t *testing.T) {
	a, b, err := DiskStat(".")

	if err != nil {
		t.Errorf("%s \n", err)
	}

	if a == 0 {
		t.Errorf("Available Space 0. Looks like error")
	}

	if b == 0 {
		t.Errorf("Total Space 0. Definitely error")
	}
}
