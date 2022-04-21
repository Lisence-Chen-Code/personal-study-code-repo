package test

import "testing"

//简要的流水模型：producers -> transmitters -> consumers

type BaseModel struct {
}

type Producer struct {
	BaseModel
}

type Transmitter struct {
	BaseModel
}

type Consumer struct {
	BaseModel
}

type BaseModelIntfs interface {
	Mission(func() chan interface{}) interface{}
}

func (p Producer) Mission(func() chan interface{}) interface{} {
	return func() {}
}

func (t Transmitter) Mission() <-chan interface{} {
	panic("implement me")
}

func (c Consumer) Mission() <-chan interface{} {
	panic("implement me")
}

func StartMission() chan interface{} {
	return nil
}

func TestName(t *testing.T) {
	var inst Producer

	doSth := func() chan interface{} {
		out := make(chan interface{})
		go func() {
			defer close(out)
			for _, n := range []int{1, 2, 3, 4} {
				out <- n
			}
		}()
		return out
	}
	inst.Mission(doSth)
}
