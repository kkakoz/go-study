package reflects

import "reflect"

func New(p reflect.Type) reflect.Value {
	if p.Kind() == reflect.Ptr {
		p = p.Elem()
	}
	return reflect.New(p)
}
