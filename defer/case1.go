/*
 * Copyright(C) 2023 Weaxs
 */

package _defer

import (
	"fmt"
	"sync"
	"time"
)

func case1() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		defer fmt.Printf("logging job cost: %v\n", time.Since(start))
		fmt.Printf("loggine at job start: %v\n", start)
		// do sth
		time.Sleep(time.Second)
	}()
	wg.Wait()
	time.Sleep(time.Second)
}

func case1Fix() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		defer func() {
			fmt.Printf("logging job cost: %v\n", time.Since(start))
		}()
		fmt.Printf("loggine at job start: %v\n", start)
		// do sth
		time.Sleep(time.Second)
	}()
	wg.Wait()
	time.Sleep(time.Second)
}
