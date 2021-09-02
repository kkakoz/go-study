package main

import (
	"fmt"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"learn-go/base/reflects"
	"reflect"
)

func filter1[T any](slice []T, f func(T) bool) []T {
	newSlice := make([]T, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func main() {
	user := newStruct[User]()
	fmt.Println(user)
	noNull(user, func(v User) {
		fmt.Println("name = ", v.Name)
	})
	users := newSlice[User]()
	fmt.Println(users)
}

type User struct {
	Name     string
	Password string
}

func newStruct[T any]() *T {
	temp := new(T)
	fValue := reflect.ValueOf(temp)
	tempT := reflect.TypeOf(*temp)
	fields := reflects.GetStructField(tempT)
	for _, field := range fields {
		fmt.Println(generator.CamelCase(field.Name))
		switch field.Type.Kind() {
		case reflect.Int:
			fValue.Elem().FieldByName(field.Name).SetInt(10)
		case reflect.String:
			fValue.Elem().FieldByName(field.Name).SetString("test")
		}
	}
	return temp
}

func newSlice[T any]() []T {
	temp := make([]T, 0)
	return temp
}

func noNull[T any](v *T, f func(v T)) {
	if v != nil {
		f(*v)
	}
}