package test

//
//import (
//	"fmt"
//	"math/rand"
//	"testing"
//	"time"
//)
///*
//模拟多个用户抢商品
// */
//
//type Goods struct {
//	storage int
//}
//type Good struct {
//	name string
//	price float32
//	Goods
//}
//
//type Consumer struct {
//	name string
//	req int
//}
//
//func TestGoodsConcurrency(t *testing.T) {
//	//init a good parameters
//	apple := Good{
//		name: "苹果",
//		price: 20,
//	}
//	apple.storage = 4444444
//
//	//模拟允许最多2000个用户同时操作
//	ch := make(chan []Consumer)
//	consumer := new(Consumer)
//	//模拟100个用户下单
//	var chCollection []Consumer
//	rand.Seed(time.Now().Unix())
//	for i := 0; i < 100; i++ {
//		go func() {
//		consumer.name = fmt.Sprintf("用户%d",i)
//		res := rand.Intn(50)
//		consumer.req = res
//		//go func() {
//			fmt.Println(consumer.name+"下单了")
//			chCollection = append(chCollection, *consumer)
//			ch<- chCollection
//		}()
//	}
//
//	//处理用户请求
//	err := decreaseStorage(ch, apple)
//	if err != nil {
//		return
//	}
//}
//
////下单减库存
//func decreaseStorage(ch chan []Consumer,good Good) error {
//	consumers := <-ch
//	for _, consumer := range consumers {
//		if good.storage<consumer.req {
//			fmt.Println("库存不足！")
//			break
//		}
//		fmt.Println("用户"+consumer.name+"下单成功！")
//		good.storage -= consumer.req
//	}
//	return nil
//}
