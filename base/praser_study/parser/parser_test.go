package tools

import (
	"fmt"
	"testing"
)

func TestLoadFiles(t *testing.T)  {

	innerstructs, err := loadGoFiles("./", "learn-go/base/praser_study/tools", "model.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range innerstructs {
		fmt.Println(*v)
	}

}


