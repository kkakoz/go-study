package reflects

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Username string
	Password string
}

func Test(t *testing.T)  {
	a := func(user User) {
		fmt.Println("user = ", user)
	}
	ftypea := reflect.TypeOf(a)
	fValue := reflect.ValueOf(a)
	types := GetFuncTypeIn(ftypea)
	values := make([]reflect.Value, 0, len(types))
	for _, ftype := range types {
		if ftype.Kind() == reflect.Struct {
			value := reflect.New(ftype)
			for _, field := range GetStructField(ftype) {
				fmt.Println(field.Tag)
				form, ok := field.Tag.Lookup("form")
				if !ok {
					form = field.Name
				}
				fmt.Println("form = ", form)
				switch field.Type.Kind() {
				case reflect.Int:
					value.Elem().FieldByName(field.Name).SetInt(10)
				case reflect.String:
					value.Elem().FieldByName(field.Name).SetString("zhangsan")
				}
			}
			values = append(values, value.Elem())
			//fmt.Println("reflct value = ", value)
		} else {
			fmt.Println(ftype)
		}
	}
	fValue.Call(values)

}
