package test

import (
	"fmt"
	"testing"
)

//defer：入栈时需要准备好执行的所有参数，否则将直接执行defer里面部分的逻辑。
type temp struct{}

func (t *temp) Add(elem int) *temp {
	fmt.Println(elem)
	return &temp{}
}

func TestDefer(t *testing.T) {
	tt := &temp{}
	defer tt.Add(1).Add(2).Add(3).Add(4)
	tt.Add(5)
}
