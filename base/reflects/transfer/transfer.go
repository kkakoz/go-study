package transfer

import (
	"html/template"
	"os"
	"reflect"
	"strings"
)

type data struct {
	SrcName string
	DstName string
	Fields  []string
}

const (
	No = iota
	Src
	Dst
	All
)

func Transfer(src interface{}, dst interface{}, havePkg int) {
	srcT := reflect.TypeOf(src)
	dstT := reflect.TypeOf(dst)
	fields := make([]string, 0)
	for i := 0; i < srcT.NumField(); i++ {
		f1 := srcT.Field(i)
		for i := 0; i < dstT.NumField(); i++ {
			f2 := dstT.Field(i)
			if f1.Name == f2.Name {
				fields = append(fields, f1.Name)
			}
			continue
		}
	}
	srcName := srcT.Name()
	dstName := dstT.Name()
	switch havePkg {
	case No:
	case Src:
		split := strings.Split(srcT.PkgPath(), "/")
		srcName = split[len(split) - 1] + "." + srcName
	case Dst:
		split := strings.Split(dstT.PkgPath(), "/")
		dstName = split[len(split) - 1] + "." + dstName
	case All:
		split := strings.Split(srcT.PkgPath(), "/")
		srcName = split[len(split) - 1] + "." + srcName
		split = strings.Split(dstT.PkgPath(), "/")
		dstName = split[len(split) - 1] + "." + dstName
	}
	d := data{
		SrcName: srcName,
		DstName: dstName,
		Fields:  fields,
	}
	tmpl, err := template.New("").Parse(temlContent)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, d)
	if err != nil {
		panic(err)
	}
}

var temlContent = `func (item *{{.SrcName}}) ToModel() *{{.DstName}} {
	return &{{.DstName}}{ {{range .Fields}}
		{{.}}: item.{{.}},{{end}}
	}
}
`