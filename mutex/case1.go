/*
 * Copyright(C) 2023 Weaxs
 */

package mutex

import (
	"fmt"
	"sync"
	"time"
)

func case1() {
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	wg.Add(5)

	// read lock
	go func() {
		defer wg.Done()
		mu.RLock()
		fmt.Println("reader0 RLock")
		// do str
		time.Sleep(5 * time.Second)
		fmt.Println("reader0 Call RUnLock")
		mu.RUnlock()
		fmt.Println("reader0 RUnlocked")
	}()

	time.Sleep(time.Second)
	// write lock
	go func() {
		defer wg.Done()
		fmt.Println("writer1 Call Lock")
		mu.Lock()
		fmt.Println("writer1 Locked")
		time.Sleep(5 * time.Second)
		mu.Unlock()
		fmt.Println("writer1 Unlocked")
	}()

	time.Sleep(time.Second)
	for i := 1; i <= 3; i++ {
		i := i
		go func() {
			wg.Done()
			mu.RLock()
			fmt.Printf("new reader%d RLock\n", i)
			time.Sleep(3 * time.Second)
			mu.RUnlock()
			fmt.Printf("new reader%d RUnlock\n", i)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("all new readers are ready")
	wg.Wait()
}
