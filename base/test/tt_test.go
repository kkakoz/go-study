package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestJson(t *testing.T) {
	a := []int{}
	data := []byte("[1,2,3]")
	err := json.Unmarshal(data, &a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)

	curT, err := time.Parse("2006-01-02", "2001-01-01")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(curT)
	fmt.Println(curT.Unix())
}

type IA interface {
	AA()
}

type A struct {

}

func (receiver A) AA() {

}

func TestT(t *testing.T) {
	ias := make([]IA, 0)
	ias = append(ias, A{})
	handleA(ias)
}

func handleA(as []IA) {
	for _, v := range as {
		v.AA()
	}
}

