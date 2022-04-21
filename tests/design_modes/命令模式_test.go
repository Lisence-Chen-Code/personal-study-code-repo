package design_modes

/**
命令模式：
*/

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type Command interface {
	Execute() // 执行
}

type MotherBoard struct {
}

func (*MotherBoard) WashClothes() {
	fmt.Println("MM去洗衣服")
}
func (*MotherBoard) WarmBed() {
	fmt.Println("MM去暖床")
}

type MMCommand1 struct {
	mb *MotherBoard
}

func NewMMCommand1(mb *MotherBoard) *MMCommand1 {
	return &MMCommand1{mb: mb}
}
func (mmc1 *MMCommand1) Execute() {
	mmc1.mb.WashClothes() // 洗衣服
}

type MMCommand2 struct {
	mb *MotherBoard
}

func NewMMCommand2(mb *MotherBoard) *MMCommand2 {
	return &MMCommand2{mb: mb}
}
func (mmc1 *MMCommand2) Execute() {
	mmc1.mb.WarmBed() //暖床
}

type Box struct {
	WashClothes Command
	WarmBed     Command
}

func NewBox(washClothes, warmBed Command) *Box {
	return &Box{
		WashClothes: washClothes,
		WarmBed:     warmBed,
	}
}
func (b *Box) GoWashClothes() {
	fmt.Println("给妹子买了一束花")
	b.WashClothes.Execute()
}
func (b *Box) GoWarmBed() {
	fmt.Println("给妹子买了iPhone")
	b.WarmBed.Execute()
}

func TestSth(t *testing.T) {
	//toBeCharge := "2015-01-01"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	//timeLayout := "2006-01-02"                             //转化所需模板
	//loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	//theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	//sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(sr)                                                 //打印输出时间戳 1420041600
	//
	////时间戳转日期
	//dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	//fmt.Println(dataTimeStr)
	TimeStrIncr("2015-01-01", 5)
}
func TimeStrIncr(timePar string, days int64) (timeRes string) {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02", timePar, loc)
	timeRes = time.Unix(theTime.Unix()+86400*days, 0).Format("2006-01-02")
	fmt.Println(timeRes)
	return
}

func Testasdasasd(t *testing.T) {
	sd := time.Now().Unix()
	fmt.Println(reflect.TypeOf(sd))
	fmt.Println(time.Unix(sd, 0).Format("2006-01-02 15:04:05"))
	sd += 86400 //增加一天
	fmt.Println(time.Unix(sd, 0).Format("2006-01-02 15:04:05"))
}
