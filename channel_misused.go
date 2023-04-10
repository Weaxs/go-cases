/*
 * Copyright(C) 2023 Weaxs
 */

package go_wrong_cases

import (
	"fmt"
	"sync"
	"time"
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
