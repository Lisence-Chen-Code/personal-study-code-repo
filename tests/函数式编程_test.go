package test

//
//import (
//	"fmt"
//	"io"
//	"testing"
//)
//
///*函数式编程，是指忽略（通常是不允许）可变数据（以避免它处可改变的数据引发的边际效应），
//忽略程序执行状态（不允许隐式的、隐藏的、不可见的状态），通过函数作为入参，函数作为返回
//值的方式进行计算，通过不断的推进（迭代、递归）这种计算，从而从输入得到输出的编程范式。在
//函数式编程范式中，没有过程式编程所常见的概念：语句，过程控制（条件，循环等等）。此外，在函
//数式编程范式中，具有引用透明（Referential Transparency）的特性，此概念的含义是函数的运
//行仅仅和入参有关，入参相同则出参必然总是相同，函数本身（被视作f(x)）所完成的变换是确定的。*/
//
////what's more,go https://juejin.cn/post/6877505132620333064#heading-0
//
//func add(a, b int) int { return a+b }
//func sub(a, b int) int { return a-b }
//
//var operators map[string]func(a, b int) int
//
//func init(){
//	operators = map[string]func(a, b int) int {
//		"+": add,
//		"-": sub,
//
//}
//
////入参参数加执行的业务操作
////总控+不同运算子=业务逻辑替换，实现接口的拓展通用性
//func calculator(a, b int, op string) int {
//	if fn, _ := operators[op];len(op)!=0 && fn!=nil{
//		return fn(a, b)
//	}
//	return 0
//}
//
//func aaa(args ...interface{})  {
//
//}
//
//func bbb(args ...interface{})  {
//
//}
//
//func TestCaller(t *testing.T) {
//	caller :=
//}
//
//
////func TestFunctor(t *testing.T) {
////	_ = calculator(1, 2, "+")
////}
//
////函数式编程实现递归-----------------------------
//func factorialTailRecursive(num int) int {
//	return factorial(1, num)
//}
//
//func factorial(accumulator, val int) int {
//	if val == 1 {
//		return accumulator
//	}
//	return factorial(accumulator*val, val-1)
//}
//
//func TestRecursion(t *testing.T) {
//	fmt.Println(factorialTailRecursive(10)) // 3628800
//}
//
////斐波那契数列来实现高阶函数的递归
//func fibonacci() func() int {
//	a, b := 0, 1
//
//	return func() int {
//		a, b = b, a+b
//		return a
//	}
//}
//
//func TestFibonacci(t *testing.T) {
//	f := fibonacci()
//
//	for i := 0; i < 10; i++ {
//		fmt.Println(f())
//	}
//}
////使用高阶/匿名函数的一个重要用途是捕俘变量和延迟计算，
////也即所谓的惰性计算（Lazy evaluations。-------------------------------
//func doSth(){
//	var err error
//	defer func(){ //defer俘虏计算外部作用域的变量，延迟计算
//		if err != nil {
//			println(err.Error())
//		}
//		println(err.Error())
//	}()
//
//	//defer err.Error()   //计算的是俘虏时刻变量的值，为nil
//
//	// ...
//	err = io.EOF
//	return
//}
//func TestLazyEval(t *testing.T) {
//	doSth() // printed: EOF
//}
////额外的特例：循环变量，循环变量并不被延迟计算
//func wrongLoop(){
//	for i:=0; i<10; i++ {
//		go func(){
//			println(i)
//		}()
//	}
//}
//func correctLoop(){
//	for i:=0; i<10; i++ {
//		go func(ix int){
//			println(ix)
//			//time.Sleep(time.Second)
//		}(i)
//	}
//}
//func TestLoop(t *testing.T) {
//	// 1. 结果会是 全部的 10
//	// 2. 在新版本的 Golang 中，将无法通过编译，报错为：
//	// loop variable i captured by func literal
//	//wrongLoop()
//	//正确
//	correctLoop()
//}
//func BenchmarkSth1(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		correctLoop()
//	}
//}
//
//
//
