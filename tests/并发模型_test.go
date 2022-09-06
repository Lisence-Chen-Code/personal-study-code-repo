package test

import (
	"fmt"
	"testing"
	"time"
)

//根据通道的双向性进行解耦。
//例：利用通道进行消息推送与消费

func SendMsg(ch chan<- string, msg string) {
	ch <- msg
}

func ConsumeMsg(ch <-chan string) {
	for one := range ch {
		fmt.Println(one)
	}
}

func TestGoRoutines(t *testing.T) {
	ch := make(chan string, 50)
	go func() {
		for _, one := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "1", "1", "1", "1", "1", "1", "1", "1", "1"} {
			go func(e string) {
				SendMsg(ch, e)
			}(one)
		}
	}()
	go func() {
		ConsumeMsg(ch)
	}()
	time.Sleep(5 * time.Second)
}
