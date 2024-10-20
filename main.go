package main

import (
	"all_euler"
	"fmt"
)

func main() {
	res := all_euler.LatticePaths(20, 20, map[string]int{})
	fmt.Println(res)
}
