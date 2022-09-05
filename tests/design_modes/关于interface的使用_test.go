package design_modes

import (
	"fmt"
	"sync"
	"testing"
)

//原理：1.当非空interface作为方法入参时，传入的变量的结构体实现了该接口才能传入；2.当为空interface时，充当正常参数
//使用场景： 1.当我们使用到多个结构体，并想抽象出他们之中某些共同的部分，来做其他业务操作，来降低代码冗余程度。
//举例：model有三个， 模拟不同model struct的基础字段的新增，类似各个orm框架对createdAt,updatedAt,deletedAt
//字段的处理，这里用简单通用的基础crud操作代替

// ------------------------------具体各项对接口的各个实现-------------------------------
type player struct {
	// unexported fields
}

func (p *player) status() string {
	return "Player is eating"
}

func (p *player) sleep() {
	// implementation goes here
}

type monster struct {
	// unexported fields
}

func (m *monster) status() string {
	return "Monster is sleeping"
}

func (m *monster) sleep() {
	// implementation goes here
}

// ------------------------------接口组，具体项抽象出来公共的部分-------------------------
type character interface {
	status() string
	sleep()
}

//输出各个具体项的某个抽象行为
func getStatus(c character) string {
	return c.status()
}

func TestIntf(t *testing.T) {
	cyclops := new(monster)
	playerOne := new(player)
	c := character.status
	fmt.Println(getStatus(cyclops))   // Prints "Monster is sleeping"
	fmt.Println(getStatus(playerOne)) // Prints "Player is eating"
	fmt.Println(c(cyclops))
}

//-------------------------------------举例----------------------------------
type BaseCrudIntf interface {
	Create(model baseModel) bool
	Delete() bool
	Update() bool
	Select() interface{}
	GetTableName() string
}

type baseModel struct {
	BaseField string
	TableName string
}

type BaseCrudHandler struct {
}

var Handler *BaseCrudHandler
var onc sync.Once

func GetACrudHandler() *BaseCrudHandler {
	onc.Do(func() {
		Handler = new(BaseCrudHandler)
	})
	return Handler
}

type Model1 struct {
	baseModel
	Field1 string
}
type Model2 struct {
	baseModel
	Field2 string
}
type Model3 struct {
	baseModel
	Field3 string
}

func (b baseModel) GetTableName() string {
	return b.TableName
}

func (b baseModel) Create(model baseModel) bool {
	//基础model具体的insert操作
	//todo insert model
	return true
}

func (b baseModel) Delete() bool {
	panic("implement me")
}

func (b baseModel) Update() bool {
	panic("implement me")
}

func (b baseModel) Select() interface{} {
	panic("implement me")
}

func (c *BaseCrudHandler) InsertOne(intf BaseCrudIntf) bool {
	//basemodel insert option
	toCreate := new(baseModel)
	fmt.Println(fmt.Sprintf("正在对表%s做新增操作", intf.GetTableName()))
	return intf.Create(*toCreate)
}

func TestCreate(t *testing.T) {
	base := baseModel{
		BaseField: "base",
		TableName: "basemodel",
	}
	mdl1 := &Model1{
		baseModel: base,
		Field1:    "field1",
	}
	mdl2 := &Model2{
		baseModel: base,
		Field2:    "field2",
	}
	GetACrudHandler().InsertOne(mdl1)
	GetACrudHandler().InsertOne(mdl2)
}
