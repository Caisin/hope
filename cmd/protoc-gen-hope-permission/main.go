package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"hope/pkg/file"
	"hope/pkg/util/str"
	"text/template"
)

var (
	tmpPath = flag.String("tmpPath", "", "模板文件路径")
)

//--hope-permission_out=./out
//protoc 通过 --foo_out 搜索插件 可执行文件 protoc-gen-hope-permission，
//也可使用参数 protoc --plugin=protoc-gen-hope-permission=D:/work/code/go/hope/cmd/protoc-gen-hope-permission/protoc-gen-hope-permission 指定插件位置
//示例参数 protoc  -I .  auth.proto --hope-permission_out=.  --proto_path=D:\work\code\go\hope\api\admin\auth\v1 --proto_path=D:\work\code\go\hope\third_party
func main() {
	flag.Parse()
	p := *tmpPath
	if str.IsBlank(p) {
		fmt.Println("template path is empty!")
		return
	}
	if !file.FileIsExisted(p) {
		fmt.Printf("template file [%s] is not exits!\n", p)
		return
	}
	parse, err := template.New("temp").Parse(file.ReadString(p))
	if err != nil {
		fmt.Printf("parse template file [%s] err: %s", p, err.Error())
		return
	}
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if len(f.Services) == 0 {
				continue
			}
			g := gen.NewGeneratedFile(f.GeneratedFilenamePrefix+"permission.txt", f.GoImportPath)
			err = parse.Execute(g, f)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
