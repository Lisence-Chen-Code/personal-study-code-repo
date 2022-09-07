package goconcurrent

import (
	"context"
	"errors"
	"fmt"
	"sync"
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
			case <-time.After(60 * time.Second):
				fmt.Println("生产者：太久没人找我生产东西了，没什么事我就先挂了")
			}
		}
	}
	var receiveOpt = func(recch <-chan *Vehicle, receiveTimes int) int32 {
		curOptTotal := int32(0)
		for i := 0; i < receiveTimes; i++ {
			select {
			case v := <-recch:
				atomic.AddInt32(&all, int32(v.Cap))
				atomic.AddInt32(&curOptTotal, int32(v.Cap))
				recevChan <- struct{}{}
			case <-time.After(2000 * time.Millisecond): //超时关闭处理
				closeChan <- struct{}{}
				return curOptTotal
			}
		}
		closeChan <- struct{}{} //流程结束，也需要退出线程
		return curOptTotal
	}
	//搬运操作，即生产者
	go sendOpt(ch, 20000)
	//接收操作，即消费者
	//case1 :go receiveOpt(ch, 3) //无论有多少车搬运，我这边想怎么接收就怎么接口
	//case2: 模拟两个消费者
	//go func() {//模拟消费者1
	//	fmt.Println(fmt.Sprintf("消费者1本次接收了%v货物", receiveOpt(ch, 300)))
	//}()
	//
	//go func() {//模拟消费者2
	//	fmt.Println(fmt.Sprintf("消费者2本次接收了%v货物", receiveOpt(ch, 100)))
	//}()
	//case3: 模拟有超多个消费者
	consumers := make(map[string]func() int32, 0)
	for i := 0; i < 100; i++ {
		consumers[fmt.Sprintf("消费者%v", i+1)] = func() int32 {
			return receiveOpt(ch, 50)
		}
	}
	for name, doFunc := range consumers {
		go func(n string, f func() int32) {
			fmt.Println(fmt.Sprintf("%s本次接收了%v货物", n, f()))
		}(name, doFunc)
	}
	//io监听，当关闭通道有值时打印总的收到的货物量
	endSta := false
	for !endSta {
		select {
		case <-closeChan:
		case <-time.After(66 * time.Second):
			fmt.Println(fmt.Sprintf("接收完毕，消费者总共接收了%v货物", all))
			endSta = true
		}
	}
}

//正常处理需要开并发执行的任务:1.等待结果回调；2.某个任务执行失败，剩余待执行的任务不要再执行
func TestConcurrentService(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	type Mission struct {
		Index    int
		Name     string
		TodoFunc func(a int) (int, error)
	}
	todoFunc := func(a int) (int, error) {
		//假设在第1000个任务时发生报错
		if a == 1000 {
			cancelFunc()
			return 0, errors.New("unexpected error occurred")
		}
		//假设每个任务都是计算平方值
		return a * a, nil
	}
	ch := make(chan *Mission, 300)
	defer close(ch)
	var wg sync.WaitGroup
	wg.Add(30000)
	//构造一个待处理的任务列表
	missions := []*Mission{}
	for i := 0; i < 30000; i++ {
		missions = append(missions, &Mission{
			Index:    i + 1,
			Name:     fmt.Sprintf("任务%v", i+1),
			TodoFunc: todoFunc,
		})
	}
	//开启任务通道，控制并发处理任务
	go func() {
		for _, one := range missions {
			select {
			case <-ctx.Done(): //结束
				return
			default:
				ch <- one
			}
		}
	}()
	go func() {
		hasDoneCounter := 0
		for one := range ch {
			select {
			case <-ctx.Done():
				wg.Add(hasDoneCounter - 30000)
				return
			default:
				hasDoneCounter++
				wg.Done()
				res, err := one.TodoFunc(one.Index)
				if err != nil {
					fmt.Println(fmt.Sprintf("处理%v出现错误%s", one.Name, err.Error()))
				} else {
					fmt.Println(fmt.Sprintf("处理%v，任务执行结果：%v", one.Name, res))
				}
			}
		}
	}()
	wg.Wait()
	fmt.Println("任务执行完毕")
	for {
		select {
		case v := <-ctx.Done():
			fmt.Println(v) //从关闭的通道，只会拿回零值，因此可以利用这一点来控制监听多个地方的停止
			fmt.Println("上下文拿到零值，主线程退出")
			return
		default:

		}
	}
}
