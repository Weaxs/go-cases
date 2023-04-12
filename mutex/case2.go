/*
 * Copyright(C) 2023 Weaxs
 */

package mutex

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex

func case2() {
	go a1()
	time.Sleep(2 * time.Second)
	fmt.Println("w call lock")
	mu.Lock()
	defer mu.Unlock()
}

func a1() {
	fmt.Println("a1 call RLock")
	mu.RLock()
	fmt.Println("a1 RLocked")

	defer mu.RUnlock()
	b2()
}

func b2() {
	time.Sleep(5 * time.Second)
	c3()
}

func c3() {
	fmt.Println("c3 call RLock")
	mu.RLock()
	fmt.Println("c3 RLocked")
	defer mu.RUnlock()
}
