package parser

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/loader"
	"log"
	"strings"
)

type parseTagFunc func(s string) string

var (
	conf     loader.Config
	parseTag parseTagFunc = defaultParseTagFunc
)

// UsageCfg ... config tools some feature
type UsageCfg struct {
	// Dir 设置需要对那一路径下的文件进行解析
	Dir string
	// ExportDir 设置生成的文件的存放地址
	ExportDir string
	// ExportFilename 指定生成的文件名字
	ExportFilename string
	// ExportPkgName 指定生成的文件的包名
	ExportPkgName string
	// ExportStructSuffix 指定生成的新的机构体后缀
	ExportStructSuffix string
	// ModelImportPath 指定源文件所在包的导入路径
	ModelImportPath string
	// Filenames 指定需要解析的.go源文件 文件名字
	Filenames []string
}

// ParseAndGenerate parse all input go files and
// get wanted struct info save with innerStruct, and then generate file
func ParseAndGenerate(cfg *UsageCfg) error {
	// parse
	ises, err := LoadGoFile(cfg.Dir, cfg.ModelImportPath, cfg.Filenames...)
	if err != nil {
		return err
	}
	fmt.Println(ises)

	//// generate
	//generateFile(ises, &outfileCfg{
	//	exportFilename:  path.Join(cfg.ExportDir, cfg.ExportFilename),
	//	exportPkgName:   cfg.ExportPkgName,
	//	modelImportPath: cfg.ModelImportPath,
	//})

	return nil
}

// SetCustomParseTagFunc use user's custom parseTag func
func SetCustomParseTagFunc(f parseTagFunc) {
	parseTag = f
}

// Exported, and specified type
func LoadGoFile(dir string, importPath string, filenames ...string) ([]*InnerStruct, error) {
	conf.Cwd = dir
	conf.CreateFromFilenames(importPath, filenames...)
	prog, err := conf.Load()
	if err != nil {
		log.Println("load program err:", err)
		return nil, err
	}

	return loopProgramCreated(prog.Created), nil
}

// loopProgramCreated to loo and filter:
// 1. unexported type
// 2. bultin types
// 3. only specified style struct name
func loopProgramCreated(
	created []*loader.PackageInfo,
) (innerStructs []*InnerStruct) {
	for _, pkgInfo := range created {
		pkgName := pkgInfo.Pkg.Name()
		defs := pkgInfo.Defs

		// imports := pkgInfo.Pkg.Imports()
		// for _, imp := range imports {
		// 	log.Println(imp.Path(), imp.Name())
		// }

		for indent, obj := range defs {
			if !indent.IsExported() ||
				obj == nil {
				continue
			}

			st, ok := obj.Type().Underlying().(*types.Struct)
			if !ok {
				log.Println("not a struct, skip this")
				continue
			}
			is := new(InnerStruct)

			is.Content = st.String()
			is.PkgName = pkgName
			is.Name = obj.Name()
			is.Fields = parseStructFields(st)

			innerStructs = append(innerStructs, is)
		}
	}
	return
}

type Field struct {
	Name string
	Typ  string
	Tag  string
}

type InnerStruct struct {
	Fields  []*Field
	Content string
	Name    string
	PkgName string
}

// parseStructFields parse fields
func parseStructFields(st *types.Struct) []*Field {
	flds := make([]*Field, 0, st.NumFields())

	for i := 0; i < st.NumFields(); i++ {
		fld := st.Field(i)
		isField := new(Field)

		isField.Name = fld.Name()
		isField.Tag = parseTag(st.Tag(i))
		isField.Typ = fld.Type().String()

		flds = append(flds, isField)
	}
	return flds
}

func defaultParseTagFunc(s string) string {
	s = strings.Replace(s, `"`, "", -1)
	splited := strings.Split(s, ":")
	return splited[len(splited)-1]
}

type GeneratorFunc func(ins []*InnerStruct, importPath string)
