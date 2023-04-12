/*
 * Copyright(C) 2023 Weaxs
 */

package shadow

import "fmt"

// https://github.com/kubernetes/kubernetes/pull/86886/files
func case1() {
	v1 := 1
	if v1 == 1 {
		v1, v2 := 2, 3
		fmt.Printf("Inner: %d, %d\n", v1, v2)
	}
	fmt.Printf("Outer: %d\n", v1)
}
