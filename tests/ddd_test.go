package test

import (
	"sync"
)

type ModelOpts interface {
	InsertOrUpdate(s string)
	GetToOptModelData() []byte
}

type BaseModelFields struct {
	Field1     string
	Field2     string
	Field3     string
	Field4     string
	ModelsData []byte
	mute       sync.RWMutex
}

type Model1 struct {
	BaseModelFields
	Field5 string
}

type Model2 struct {
	BaseModelFields
	Field6 string
}

type Model3 struct {
	BaseModelFields
	Field7 string
}

func (m *Model1) InsertOrUpdate(s string) {
	m.mute.Lock()
	//fmt.Println(fmt.Sprintf("插入数据%s", m.)
	m.mute.Unlock()
	//panic("implement me")
}

func (m *Model1) GetToOptModelData() []byte {

	//panic("implement me")
	return nil
}
