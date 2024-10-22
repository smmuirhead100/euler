package main

import (
	"all_euler"
	"fmt"
)

func main() {
	res := all_euler.CountingSundaysOnFirstSince("January", 1901, "December", 2000)
	fmt.Println(res)
}
