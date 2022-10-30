package goconvertstruct

import "reflect"

func ConvertStruct(instance interface{}) {
	objType := reflect.TypeOf(instance)
	println(objType)
}
