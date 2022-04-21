package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/***
	sync.Map的性能高体现在读操作远多于写操作的时候。 极端情况下，只有读操作时，是普通map的性能的44.3倍。

反过来，如果是全写，没有读，那么sync.Map还不如加普通map+mutex锁呢。只有普通map性能的一半。

建议使用sync.Map时一定要考虑读定比例。当写操作只占总操作的<=1/10的时候，使用sync.Map性能会明显高很多。

原文链接：https://blog.csdn.net/wyg_031113/article/details/106282340
*/

/**
对map资源进行高并发的写操作，map并不支持并被panic。
*/
func TestConcurrentMapWrites1(t *testing.T) {
	m := make(map[int]int, 1000)
	go do(m)
	go do(m)

	time.Sleep(1 * time.Second)
	fmt.Println(m)
}

func do(m map[int]int) {
	i := 0
	for i < 1000 {
		m[i] = 1
		i++
	}
}

/**
使用同步组读写锁，对资源操作时上锁处理,相当于关键部分代码变成同步的了
*/
func TestConcurrentMapWrites2(t *testing.T) {
	m := map[int]int{1: 1}
	var s sync.RWMutex
	var do1 = func(m map[int]int) {
		i := 0
		for i < 10000 {
			// 加锁
			s.Lock()
			m[1] = 1
			// 解锁
			s.Unlock()
			i++
		}
	}
	go do1(m)
	go do1(m)
	time.Sleep(1 * time.Second)
	fmt.Println(m)
}

/**
使用同步map：
*/
func TestConcurrentMapWrites3(t *testing.T) {
	m := sync.Map{}
	m.Store(1, 1)
	var do = func(m sync.Map) {
		i := 0
		for i < 10000 {
			m.Store(1, 1)
			i++
		}
	}
	go do(m)
	go do(m)
	time.Sleep(1 * time.Second)
	fmt.Println(m.Load(1))
}
