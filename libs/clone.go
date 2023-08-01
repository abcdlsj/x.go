package xutil

import "reflect"

func Clone(a, b interface{}) {
	// 判断类型是否一致
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return
	}

	valB := reflect.ValueOf(b)

	// 判断是否可寻址，不可寻址的话需要使用 New 方法创建一个可寻址的值
	if !valB.CanAddr() {
		newVal := reflect.New(reflect.TypeOf(b))
		newVal.Elem().Set(valB)
		valB = newVal.Elem()
	}

	valA := reflect.ValueOf(a)

	// 判断是否可设置
	if valA.CanSet() {
		valA.Set(valB)
	}
}
