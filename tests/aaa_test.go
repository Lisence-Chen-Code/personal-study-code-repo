package test

import (
	"net/http"
	"testing"
)

//调度测试
func TestSched(t *testing.T) {
	done := false
	ch := make(chan bool, 1)

	go func() {
		done = true
		ch <- done
	}()
	for !done {
		select {
		case <-ch:
			println("not done !") // 并不内联执行
		}
	}
	println("done !")
	// 启动一个 pprof http server
	if err := http.ListenAndServe(":7899", nil); err != nil {
		panic(err.Error())
	}
	select {}
}
