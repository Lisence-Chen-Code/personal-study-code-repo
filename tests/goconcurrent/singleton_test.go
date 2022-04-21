package goconcurrent

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

type Model struct {
	Field1 string
	Field2 string
	Field3 string
}

type SingletonSt struct {
	SpecificSt *Model
	once       sync.Once
}

func (s *SingletonSt) GetSingletonInstance(st Model) *Model {
	s.once.Do(func() {
		s.SpecificSt = &st
	})
	return s.SpecificSt
}

func TestSingleton(t *testing.T) {
	st := &SingletonSt{SpecificSt: &Model{Field1: "111"}} //初始化单例对象值
	for i := 1; i < 5; i++ {
		fmt.Println(fmt.Sprintf("%s", st.GetSingletonInstance(Model{Field1: strconv.Itoa(i)}).Field1))
	}
}
