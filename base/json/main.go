package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
)

func main() {
	m := map[string]interface{}{
		"a": 1,
		"b": "t",
		"c": 1.1,
	}

	data, _ := json.Marshal(m)
	fmt.Println(data)

	var jiter = jsoniter.Config{
		EscapeHTML:                    false,
		MarshalFloatWith6Digits:       true, // will lose precession
		ObjectFieldMustBeSimpleString: true, // do not unescape object field
		UseNumber:                     true,
	}.Froze()

	// json unmarshal
	tar1 := map[string]interface{}{}
	err := json.Unmarshal(data, &tar1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tar1)

	tar2 := map[string]interface{}{}
	err = jiter.Unmarshal(data, &tar2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tar2)
	a := tar2["c"].(json.Number)
	fmt.Println(a.Float64())
}
