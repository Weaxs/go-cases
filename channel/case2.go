/*
 * Copyright(C) 2023 Weaxs
 */

package channel

import (
	"fmt"
	"time"
)

// case: https://github.com/kubernetes/kubernetes/pull/83925
func case2() {
	ch := make(chan string)
	go func() {
		// do stn then put result
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): //timeout
		return
	}
}
