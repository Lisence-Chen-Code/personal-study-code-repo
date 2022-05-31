package design_modes

import (
	"sync"
	"testing"
)

//模拟场景：中介，提供租赁商品，售卖商品，收取用户一定比例佣金
//场景角色：中介，用户，商品

// -----------------------  struct --------------------
type Good struct { //商品
	Price      float64
	CosRatio   int
	Name, Desc string
}

type ProxySt struct { //中介抽象体
	Name        string
	ID          int64
	Level       float64
	CurServUser *User
	mux         *sync.Mutex
}

type User struct { // 用户
	Name    string
	Balance float64
}

// ----------------------  service ---------------------

//1.以中介能提供的服务新建一个接口
type IntfProxy interface {
	RentGood(good *Good) *ProxySt
	SellGood(good *Good) *ProxySt
}

type AfterHookDo interface {
	BalanceOpt(good *Good)
}

type BeforeHookDo interface {
	LockOpt()
}

//2.写出中介实现服务方法接口
func (p *ProxySt) BalanceOpt(good *Good) {
	p.CurServUser.Balance -= good.Price
	p.mux.Unlock()
}

func (p *ProxySt) LockOpt() {
	p.mux.Lock()
}

func (p *ProxySt) RentGood(good *Good) *ProxySt {
	p.LockOpt()
	go func() {
		//todo 租借服务
	}()
	p.BalanceOpt(good)
	return p
}

func (p *ProxySt) SellGood(good *Good) *ProxySt {
	p.LockOpt()
	go func() {
		//todo 售卖服务
	}()
	p.BalanceOpt(good)
	return p
}

type ProxyBaseSt struct {
	*ProxySt
}

func (p *ProxyBaseSt) GetOneInstance() *ProxyBaseSt {
	return new(ProxyBaseSt)
}

// -------------------------- testings ----------------------
//模拟用户通过中介租房和卖房
func TestProxyDo(t *testing.T) {
	user1 := &User{
		Name:    "张三",
		Balance: 10000.00,
	}
	user2 := &User{
		Name:    "李四",
		Balance: 20000.00,
	}
	toSellGoods := &Good{
		Price:    3000,
		CosRatio: 50,
		Name:     "别墅",
		Desc:     "张三卖别墅",
	}
	toRentGoods2 := &Good{
		Price:    1004,
		CosRatio: 30,
		Name:     "单间",
		Desc:     "张三租房",
	}
	toRentGoods := &Good{
		Price:    1004,
		CosRatio: 30,
		Name:     "单间",
		Desc:     "李四租房",
	}
	houseProxy := new(ProxyBaseSt).GetOneInstance()
	go func(proxy *ProxyBaseSt) {
		proxy.CurServUser = user1
		proxy.SellGood(toSellGoods).RentGood(toRentGoods2)
	}(houseProxy)
	go func(proxy *ProxyBaseSt) {
		proxy.CurServUser = user2
		proxy.RentGood(toRentGoods)
	}(houseProxy)
}
