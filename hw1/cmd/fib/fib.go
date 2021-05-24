package main

import (
	"fmt"
	"github.com/igvasilyev/go-core-2-homeworks/hw1/pkg/fib"
)

func main() {
	nums := []int{1, 2, 3, 10, 20, 21, 25}
	for _, n := range nums {
		if n > 20 {
			fmt.Printf("максимум 20, передано %d\n", n)
			continue
		}
		fmt.Printf("%d: %d\n", n, fib.Num(n))
	}
}
