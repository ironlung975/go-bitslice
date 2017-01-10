package main

import (
	"fmt"
	"time"
)

func main() {
	// Test to see if bitwise operations on the largest native bit values are equivalent to operations on slices of bytes
	length := 8
	bs1 := make([]byte, length)
	bs2 := make([]byte, length)
	bs1[0] = byte(255)
	bs1[1] = byte(1)
	{
		start := time.Now()
		for i, _ := range bs1 {
			bs2[i] |= bs1[i]
		}
		elapsed := time.Since(start)
		fmt.Println("Reading bytes and operating on bytes took ", elapsed)
	}

	ui1 := uint64(256)
	ui2 := uint64(0)
	{
		start := time.Now()
		ui2 |= ui1
		elapsed := time.Since(start)
		fmt.Println("Operating on uint64 took ", elapsed)
	}

	passed := true
	for i, _ := range bs1 {
		if bs2[i] != bs1[i] {
			passed = false
		}
	}
	fmt.Println("slice of bytes operation passed: ", passed)
	fmt.Println("uint operation passed: ", ui1 == ui2)
}
