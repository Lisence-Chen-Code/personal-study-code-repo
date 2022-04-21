package design_modes

import (
	"fmt"
	"testing"
)

type Customer interface { // 接待者
	Accept(Visitor)
}
type Visitor interface { // 访问者
	Visit(Customer)
}

type EnterpriseCustomer struct {
	name string
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{name: name}
}
func (e *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(e)
}

type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{name: name}
}
func (e *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(e)
}

type CustomerCol struct {
	customers []Customer // 接待者集合
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer) // 叠加
}
func (c *CustomerCol) Accept(visitor Visitor) { // 每个服务者都接手访问者
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

type AnalysisVisitor struct {
}

func (*AnalysisVisitor) Visit(customer Customer) {
	//实现访问者的接口，重写其访问方法
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Println("analysis enterprise customer", c.name)
	}
}

type ServiceRequestVisitor struct {
}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Println("serving enterprise customer", c.name)
	case *IndividualCustomer:
		fmt.Println("serving individual customer", c.name)
	}
}

/**
	场景描述：
	接待者接待访问者，接待者结构体增加多个接待者，结束后开始执行接
·	待业务accept，里面是不同接待者身份，现在有个访问者接口，因为访问者身份的不同，理应存
	在多个访问者接口题的具体实现，并具体实现接口的访问visit方法，在接待者的accept方法里面传入具体的访问者实例，
*/
func TestSth2(t *testing.T) {
	c := CustomerCol{}                        //接待者
	c.Add(NewEnterpriseCustomer("Microsoft")) //添加多个接待者
	//c.Add(visitor.NewEnterpriseCustomer("Google"))
	c.Add(NewIndividualCustomer("Billgates"))
	c.Add(NewEnterpriseCustomer("Google"))
	c.Accept(&ServiceRequestVisitor{})

	c.Accept(&AnalysisVisitor{}) //接待者接收访问者，根据传入的结构体具体实现的方向来实现接口
}

/**
实现场景：公司外来访客系统的功能
访客身份（目的）：面试者，外包员工，正式员工，其他
*/
type BaseIntf interface {
}

type BaseModel struct {
}

type Receiver interface {
	BaseIntf
	ResponseHandle(Visitors)
}

type Visitors interface {
	BaseIntf
	DoSth(Receiver)
}

type ReceiverProcessor struct {
	Receivers []Receiver
}

type Interviewer struct {
	BaseModel
}

type OutSource struct {
	BaseModel
}

type Official struct {
	BaseModel
}

type Others struct {
	BaseModel
}

type HR struct {
}

type Boss struct {
}

func (r *ReceiverProcessor) responseHandle(visitors Visitors) {
	//allot service for different visitors
	for _, e := range r.Receivers {
		e.ResponseHandle(visitors)
	}
}
