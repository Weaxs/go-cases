/*
 * Copyright(C) 2023 Weaxs
 */

package once

import (
	"fmt"
	"sync"
	"time"
)

func case1() {
	count := 0
	once := sync.Once{}
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("recovered from once")
			}
		}()

		// panic
		once.Do(func() {
			fmt.Println("once in goroutine")
			count = 1 / count
		})
	}()
	time.Sleep(time.Second)

	// not run
	once.Do(func() {
		fmt.Println("once in main")
		count = 1 / count
	})
	fmt.Println("end")
}
