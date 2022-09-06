package goconcurrent

import (
	"fmt"
	"testing"
)

type Vehicle struct {
	Type string
}

//通道的使用：
//将通道比作一座桥梁，通道的读写相当于车辆进入驶出桥的操作，关闭操作相当于为修好的桥设卡检修，正在用最普通的车型，无法继续进入
func TestChannel(t *testing.T) {
	var ch chan int        //ch这时为nil，可理解为桥正在建设，因此对于正在建设的桥梁：驶入车辆和驶出车辆操作都需要等待桥梁修好，所以阻塞；对于正在修理的桥梁，因为还没修好，如果做设置关卡的操作，将会发生恐慌
	ch = make(chan int, 2) //ch这时为一座空的桥，对于驶入车辆可正常操作，驶出车辆需要桥上有车，那么会阻塞住一只等待车汇入桥；设置关卡，因为桥建好了，可正常设卡
	ch <- 5                //ch这时就是一座有车辆在行驶的桥梁，可以正常设卡，驶入车辆，驶出车辆
	ch <- 4                //ch这时已经满了，桥梁达到荷载，接下来：无法进行驶入操作，除非有车驶出，所以阻塞；可以正常驶出；可以正常设卡
	close(ch)              //正常关闭设卡
	fmt.Println(<-ch)      //ch设卡了，仍然可以正常驶出车辆
	fmt.Println(<-ch)      //ch已经空了
	fmt.Println(<-ch)      //此时只能获取到通道类型0值
	close(ch)              //ch已经关闭，再设卡会引起恐慌
	ch <- 2                //ch已经关闭设卡，此时再向里面驶入车辆必定引起恐慌
}

//模拟对已经关闭的通道进行output操作
func TestOutputClosedChannel(t *testing.T) {
	ch := make(chan *Vehicle, 2)
	ch <- &Vehicle{Type: "car"}
	ch <- &Vehicle{Type: "bike"}
	close(ch)
	for one := range ch {
		fmt.Println(fmt.Sprintf("%v", one))
	}
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	temp := <-ch
	fmt.Println(temp)
}

//模拟对已经关闭的通道input操作
func TestInputClosedChannel(t *testing.T) {
	ch := make(chan *Vehicle, 2)
	ch <- &Vehicle{Type: "car"}
	ch <- &Vehicle{Type: "bike"}
	close(ch)
	ch <- &Vehicle{Type: "motorcycle"}
	fmt.Println(<-ch)
}
