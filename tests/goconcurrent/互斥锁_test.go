package goconcurrent

import (
	"fmt"
	"sync"
	"testing"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

//资源竞争，有需要院子操作的场景，通过加锁来互斥协程访问资源
func TestMutexInCcrt(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total.value)
}
