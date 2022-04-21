package test

import "testing"

/*在Golang 中有一个很重要的概念那就是 逃逸分析（Escape analysis），所谓的逃逸分析指由编译器决定内存分配的位置。

分配在 栈中，则函数执行结束可自动将内存回收
分配在 堆中，则函数执行结束可交给GC（垃圾回收）处理

最终程序的执行效率和这个两种分配规则是有这重要关联的，而传值和传指针的主要区别
在于底层值是否需要拷贝,表面上看传指针不涉及值拷贝，效率肯定更高。但是实际情况
是指传针会涉及到变量逃逸到堆上，而且会增加GC的负担，所以本文我们要做的内容就是进行 逃逸分析 ,按照惯例先上结论。

栈上分配内存比在堆中分配内存有更高的效率
栈上分配的内存不需要GC处理,函数执行后自动回收
堆上分配的内存使用完毕会交给GC处理
发生逃逸时，会把栈上的申请的内存移动到堆上
指针可以减少底层值的拷贝，可以提高效率，但是会产生逃逸，但是如果拷贝的数据量小，逃逸造成的负担（堆内存分配+GC回收)会降低效率
因此选择值传递还是指针传递，变量的大小是一个很重要的分析指标

每种方式都有各自的优缺点，栈上的值，减少了 GC 的压力,但是要维护多个副本，堆上的指针，会增加 GC 的压力，但只需
维护一个值。因此选择哪种方式，依据自己的业务情况参考这个标准进行选择。
*/

type person struct {
	name string
	age  int
}

func TestAsda(t *testing.T) {
	makePerson(32, "艾玛·斯通")
	showPerson(33, "杨幂")
}

func makePerson(age int, name string) *person {
	maliya := person{name, age}
	return &maliya
}

func showPerson(age int, name string) person {
	yangmi := person{name, age}
	return yangmi
}
