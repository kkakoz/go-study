package reflects

import (
	"fmt"
	"reflect"
)

var providers = make(map[string]provider)

var consumers = make(map[string]typeConsumer)

type provider struct {
	pType reflect.Type
}

type typeConsumer struct {
	inType []reflect.Type
}

func Provider(constructs ...interface{}) {
	//outTypes := make([]reflect.Type, 0, len(constructs))
	for _, v := range constructs {
		vType := reflect.TypeOf(v)
		oTypes := GetFuncTypeOut(vType)
		for _, oType := range oTypes {
			key := GetProviderKey(oType)
			providers[key] = provider{
				pType: oType,
			}
		}
	}
	out := providers
	fmt.Println(out)
	for _, v := range constructs {
		vType := reflect.TypeOf(v)
		consumer := typeConsumer{}
		key := GetConsumerKey(vType)
		consumers[key] = typeConsumer{
			inType: make([]reflect.Type, 0),
		}
		inTypes := GetFuncTypeIn(vType)

		for _, inType := range inTypes {
			consumer.inType = append(consumer.inType, inType)
		}
	}
	in := consumers
	fmt.Println(in)
}

func GetConsumerKey(vType reflect.Type) string {
	return ""
}

func GetProviderKey(t reflect.Type) string {
	return t.PkgPath() + ":" + t.Name()
}

