package test

import (
	"fmt"
	"reflect"
	"testing"
)

func GetTagValMap(v interface{}, tag string) map[string]interface{} {
	resMap := make(map[string]interface{})
	if v == nil {
		return resMap
	}
	typeOf := reflect.TypeOf(v)
	fieldNum := typeOf.NumField()
	for i := 0; i < fieldNum; i++ {
		structField := typeOf.Field(i)
		tagValue := structField.Tag.Get(tag)
		val := reflect.ValueOf(v).FieldByName(structField.Name)
		resMap[tagValue] = val.Interface()
	}
	return resMap
}

func GetStructTagList(v interface{}, tag string) []string {
	var resList []string
	if v == nil {
		return resList
	}
	var item interface{}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		values := reflect.ValueOf(v)
		if values.Len() == 0 {
			return resList
		}
		item = values.Index(0).Interface()
	case reflect.Struct:
		item = reflect.ValueOf(v).Interface()
	default:
		panic(fmt.Sprintf("type %v not support", reflect.TypeOf(v).Kind()))
	}
	typeOf := reflect.TypeOf(item)
	fieldNum := typeOf.NumField()
	for i := 0; i < fieldNum; i++ {
		resList = append(resList, typeOf.Field(i).Tag.Get(tag))
	}
	return resList
}

func Struct2MapTagList(v interface{}, tag string) []map[string]interface{} {
	var resList []map[string]interface{}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		values := reflect.ValueOf(v)
		for i := 0; i < values.Len(); i++ {
			resList = append(resList, GetTagValMap(values.Index(i).Interface(), tag))
		}
		break
	case reflect.Struct:
		val := reflect.ValueOf(v).Interface()
		resList = append(resList, GetTagValMap(val, tag))
		break
	default:
		panic(fmt.Sprintf("type %v not support", reflect.TypeOf(v).Kind()))
	}
	return resList
}

func TestGetStTagList(t *testing.T) {
	type Temp struct {
		Name   string `custom:"name"`
		Tele   int64  `custom:"tele"`
		Gender bool   `custom:"gender"`
	}
	temp := &Temp{
		Name:   "张三",
		Tele:   123059421,
		Gender: false,
	}
	fmt.Println(GetTagValMap(*temp, "custom"))
	fmt.Println(GetStructTagList(*temp, "custom"))
	fmt.Println(Struct2MapTagList(*temp, "custom"))
}
