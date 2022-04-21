package test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestOptPtr(t *testing.T) {
	data := []int{1, 2, 3}
	//for i := 0; i < len(data); i++ {
	//	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&data[0])) + uintptr(i)*unsafe.Sizeof(data[0]))
	//	fmt.Printf("%d ", *(*int)(unsafe.Pointer(ptr)))
	//}
	fmt.Printf("\n")
	str := []string{"a", "b", "c"}
	for i := 0; i < len(data); i++ {
		ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&str[0])) + uintptr(i)*unsafe.Sizeof(str[0]))
		fmt.Printf("%s ", *(*string)(unsafe.Pointer(ptr)))
	}
	fmt.Printf("\n")
	// 利用指针修改下标为1的值
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&str[0])) + uintptr(1)*unsafe.Sizeof(str[0]))
	fmt.Printf("-------------------start personal testing -------------------")
	//s := reflect.TypeOf(ptr)
	*(*string)(unsafe.Pointer(ptr)) = "d"
	fmt.Println(str)
	// 利用指针修改下标为0的值
	ptr = unsafe.Pointer(uintptr(unsafe.Pointer(&data[0])) + uintptr(0)*unsafe.Sizeof(data[0]))
	*(*int)(unsafe.Pointer(ptr)) = 5
	fmt.Println(data)
}
