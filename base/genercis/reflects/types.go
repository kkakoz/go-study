package reflects

import (
	"reflect"
)

func GetType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

func GetFuncTypeIn(ftype reflect.Type) []reflect.Type {
	if ftype.Kind() != reflect.Func {
		return nil
	}
	numIn := ftype.NumIn()

	types := make([]reflect.Type, 0, numIn)
	for i := 0; i < numIn; i++ {
		inType := ftype.In(i)
		types = append(types, inType)
	}
	return types
}

func GetFuncTypeOut(ftype reflect.Type) []reflect.Type {
	if ftype.Kind() != reflect.Func {
		return nil
	}
	numIn := ftype.NumOut()

	types := make([]reflect.Type, 0, numIn)
	for i := 0; i < numIn; i++ {
		inType := ftype.In(i)
		types = append(types, inType)
	}
	return types
}

func GetStructField(t reflect.Type) []reflect.StructField {
	if t.Kind() != reflect.Struct {
		return nil
	}
	numF := t.NumField()
	fields := make([]reflect.StructField, 0, numF)
	for i := 0; i < numF; i++ {
		inType := t.Field(i)
		fields = append(fields, inType)
	}
	return fields
}