package main

import (
	"fmt"
	"time"
)

func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if _, ok := cache[n]; !ok {
			cache[n] = f(n)
		}
		return cache[n]
	}
}

func longFunc(num int) int {
	r := 0
	for i := 0; i < 10000000; i++ {
		r += num * i
	}
	return r
}

func main() {
	memoized := memoize(longFunc)
	for i := 0; i < 10; i++ {
		fmt.Println(memoized(i))
	}
	start := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Println(memoized(i))
	}
	fmt.Println(time.Since(start))
}