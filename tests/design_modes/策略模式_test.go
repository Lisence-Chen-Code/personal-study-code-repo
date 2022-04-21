package design_modes

import (
	"fmt"
	"testing"
)

type MMContext struct {
	Name        string
	Age         int
	paoStrategy MMStrategy
}

func NewMMContext(name string, age int, strategy MMStrategy) *MMContext {
	return &MMContext{name, age, strategy}
}
func (mm *MMContext) Pao() {
	mm.paoStrategy.Pao(mm)
}

type MMStrategy interface {
	Pao(*MMContext)
}
type Girl struct {
}

func (*Girl) Pao(ctx *MMContext) {
	fmt.Println("girl for world", ctx.Name)
}

type Women struct {
}

func (*Women) Pao(ctx *MMContext) {
	fmt.Println("women for function", ctx.Name)
}

func TestSth3(t *testing.T) {
	//ctx := strategy.NewMMContext("marry", 18, &strategy.Girl{})
	ctx := NewMMContext("alis", 28, &Women{})
	ctx.Pao()
}
