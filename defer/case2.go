/*
 * Copyright(C) 2023 Weaxs
 */

package _defer

import "fmt"

func case2() {
	i := 1
	j := add(&i)
	fmt.Println("i: ", i)
	fmt.Println("j: ", j)
}

func add(n *int) int {
	defer func() {
		*n = *n + 1
	}()
	return *n
}
