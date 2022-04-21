package test

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
}

func (p Person) test1() {
	p.Name = "www.gopher.cc"
	fmt.Println(p.Name)
}
func (p *Person) test2() {
	p.Name = "www.gopher.cc"
	fmt.Println(p.Name)
}
func TestSth(t *testing.T) {
	p := new(Person) // 此处p为指针类型
	p.Name = "gopher.cc"
	p.test1() // 形式上传入的是指针，实际上却是值拷贝
	fmt.Println(p.Name)
	p2 := Person{} // 此处为值类型
	p2.Name = "gopher.cc"
	p2.test2() // 形式上传入的是值类型，实际上是却是地址拷贝
	fmt.Println(p2.Name)
}
