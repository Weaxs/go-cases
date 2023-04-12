/*
 * Copyright(C) 2023 Weaxs
 */

package channel

import (
	"fmt"
	"sync"
)

func case1() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)

	wg.Add(4)
	for j := 0; j < 4; j++ {
		go func() {
			for {
				task := <-ch
				fmt.Println(task)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
