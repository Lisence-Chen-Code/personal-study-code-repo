package goconcurrent

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

type Vehicle struct {
	Type string
	Cap  int //承载量
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

//----------------------------------------------------------------------------------------------------------------

//多协程对通道操作的理解：
//上述将通道比作桥梁，那么协程就相当于小道，多个协程即多条小道

//模拟开多个协程对通道进行io操作
//运货：多辆车运货，将东西搬过桥
func TestConcurrentOptChannel(t *testing.T) {
	status := false
	all := 0
	ch := make(chan *Vehicle, 4) //桥长度4，宽度1，一次最多4辆车进来
	defer close(ch)
	for i := 0; i < 10; i++ {
		go func(idx int) { //协程开启，模拟多条道路向通道驶入车辆
			ch <- &Vehicle{
				Type: "car",
				Cap:  idx,
			}
		}(i)
	}
	for !status {
		select {
		case v := <-ch: //对搬到的货物进行处理
			all += v.Cap
		case <-time.After(200 * time.Millisecond): //这里控制阻塞200毫秒，即通道里面的东西已经被读完了时弹出
			status = true
			fmt.Println(fmt.Sprintf("总共运过来%v吨货物", all))
			break
		}
	}
}

//运货：多辆车运货，将东西搬过桥，用单向通道来解耦io操作,这样可以取货的人可以精准地控制需要接口的运货的车辆，将搬运和接收解耦开
func TestDecouplingIO(t *testing.T) {
	all := int32(0)
	ch := make(chan *Vehicle, 5) //开启一个通道，专门用来车辆搬运货物
	closeChan := make(chan struct{}, 1)
	defer close(closeChan)
	recevChan := make(chan struct{}, 1) //控制消费者不需要再生产东西的回滚信号通道
	defer close(recevChan)
	var sendOpt = func(sendch chan<- *Vehicle, genTimes int) {
		sendch <- &Vehicle{} //添加初始信号，开启发送操作
		for i := 0; i < genTimes; i++ {
			select {
			case <-recevChan: //当成功完成一次接收操作之后，向接收通道发送信号，此处监听，有接收到信号才需要继续发送
				sendch <- &Vehicle{
					Type: "car",
					Cap:  i + 1,
				}
			case <-time.After(2 * time.Second):
				fmt.Println("太久没人找我生产东西了，没什么事我就先挂了")
			}
		}
	}
	var receiveOpt = func(recch <-chan *Vehicle, receiveTimes int) {
		for i := 0; i < receiveTimes; i++ {
			select {
			case v := <-recch:
				atomic.AddInt32(&all, int32(v.Cap))
				recevChan <- struct{}{}
			case <-time.After(2000 * time.Millisecond): //超时关闭处理
				closeChan <- struct{}{}
				return
			}
		}
		closeChan <- struct{}{} //流程结束，也需要退出线程
	}
	//搬运操作，即生产者
	go sendOpt(ch, 200)
	//接收操作，即消费者
	go receiveOpt(ch, 3) //无论有多少车搬运，我这边想怎么接收就怎么接口
	//io监听，当关闭通道有值时打印总的收到的货物量
	select {
	case <-closeChan:
		fmt.Println(all)
		fmt.Println("接收完毕")
	}
}

//
