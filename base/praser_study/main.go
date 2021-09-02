package main

import (
	"fmt"
	"learn-go/base/praser_study/parser"
	"os"
	"strings"
	"text/template"
)

type tempData struct {
	SrcName    string
	TargetName string
	Fields     []*parser.Field
}

func main() {
	f, err := os.Open("/Users/pengxc/code/github/go-study/base/praser_study/parser/test_model.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Name())
	innerstructs, err := parser.LoadGoFile("./parser/", "learn-go/base/praser_study/parser", "/Users/pengxc/code/github/go-study/base/praser_study/parser/test_model.go")
	if err != nil {
		panic(err)
	}
	for _, v := range innerstructs {
		fmt.Println(*v)
	}

	var suffixName = "DTO"

	for _, target := range innerstructs {
		if strings.HasSuffix(target.Name, suffixName) {
			v2Name := strings.TrimRight(target.Name, suffixName)
			for _, src := range innerstructs {
				if src.Name == v2Name {
					curFunc := func(src, target *parser.InnerStruct) *tempData {
						data := &tempData{}
						sameField := make([]*parser.Field, 0)
						for _, field := range src.Fields {
							for _, targetF := range target.Fields {
								if field.Name == targetF.Name {
									sameField = append(sameField, field)
									continue
								}
							}
						}
						data.SrcName = src.Name
						data.TargetName = target.Name
						data.Fields= sameField
						return data

					}
					data := curFunc(src, target)
					tmpl, err := template.New("").Parse(temlContent)
					if err != nil {
						panic(err)
					}
					file, err := os.Open("./tmpl")
					if err != nil {
						panic(err)
					}
					defer file.Close()
					err = tmpl.Execute(file, data)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
	//tmpl, err := template.New("").Parse(temlContent)
	//if err != nil {
	//	panic(err)
	//}
	//
	//file, err := os.OpenFile("./temp")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//tmpl.Execute(file, )

}

var temlContent = `import {{.Path}}


func (item *{{.Model}}) To{{.DTO}}() *{{.DTO}} {
	return &{{.DTO}}{
		ID:   item.ID,
		Name: item.Name,
		Age:  item.Age,
	}
}

`
