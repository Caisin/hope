package main

import (
	"errors"
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"hope/cmd/protoc-gen-hope-permission/funmap"
	"hope/pkg/file"
	"hope/pkg/util/str"
	"text/template"
)

var (
	tmpPath     string
	outFileName string
)

//--hope-permission_out=./out
//protoc 通过 --foo_out 搜索插件 可执行文件 protoc-gen-hope-permission，
//也可使用参数 protoc --plugin=protoc-gen-hope-permission=D:/work/code/go/hope/cmd/protoc-gen-hope-permission/protoc-gen-hope-permission 指定插件位置
//示例参数 protoc  -I .  auth.proto --hope-permission_out=.  --proto_path=D:\work\code\go\hope\api\admin\auth\v1 --proto_path=D:\work\code\go\hope\third_party
//示例参数 protoc --plugin=protoc-gen-hope-permission -I .  auth.proto --hope-permission_out=tmpPath=persql.gohtml:.  --proto_path=/Users/caisin/study/code/go/hope/api/admin/auth/v1 --proto_path=/Users/caisin/study/code/go/hope/third_party
func main() {

	var flags flag.FlagSet
	flags.StringVar(&tmpPath, "tmpPath", "tp.gohtml", "模板文件路径")
	flags.StringVar(&outFileName, "outFileName", "out.txt", "模板文件路径")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		p := tmpPath
		println(tmpPath)
		if str.IsBlank(p) {
			return errors.New("template path is empty")
		}
		if !file.FileIsExisted(p) {
			return errors.New(fmt.Sprintf("template file [%s] is not exits!\n", p))
		}
		parse, err := template.New("temp").Funcs(funmap.FuncMap).Parse(file.ReadString(p))
		if err != nil {
			return errors.New(fmt.Sprintf("parse template file [%s] err: %s", p, err.Error()))
		}
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if len(f.Services) == 0 {
				continue
			}
			g := gen.NewGeneratedFile(f.GeneratedFilenamePrefix+"/"+outFileName, f.GoImportPath)
			err = parse.Execute(g, f)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
