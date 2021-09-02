package parser

import (
	"fmt"
	"testing"
)

func TestLoadFiles(t *testing.T)  {

	innerstructs, err := LoadGoFile("./", "learn-go/base/praser_study/parser", "test_model.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range innerstructs {
		fmt.Println(*v)
	}

}


