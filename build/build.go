package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
	"fmt"
	"hope/pkg/file"
	"hope/pkg/util/str"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	target := filepath.Join(os.TempDir(), "ent")
	os.MkdirAll(target, os.ModePerm)
	defer os.RemoveAll(target)
	storage, _ := gen.NewStorage("sql")
	projectPath := "D:/work/code/go/hope"
	prods := []string{"admin", "param", "novel"}
	funcMap := template.FuncMap{
		"genCondition":       genCondition,
		"genCreateSetFields": genCreateSetFields,
		"genCreateFields":    genCreateFields,
		"genFields":          genFields,
		"add":                add,
		"parseType":          parseType,
	}
	//api模板
	apiTmpPath := fmt.Sprintf("%s/build/template/api.proto.tpl", projectPath)
	apiReadStr, _ := file.ReadStr(apiTmpPath)
	apiTemplate, err := template.New("apiTemplate").Funcs(funcMap).Parse(apiReadStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//biz 模板
	bizTmpPath := fmt.Sprintf("%s/build/template/biz.tpl", projectPath)
	bizReadStr, _ := file.ReadStr(bizTmpPath)
	bizTemplate, err := template.New("bizTemplate").Funcs(funcMap).Parse(bizReadStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//service 模板
	serviceTmpPath := fmt.Sprintf("%s/build/template/service.tpl", projectPath)
	serviceReadStr, _ := file.ReadStr(serviceTmpPath)
	serviceTemplate, err := template.New("serviceTemplate").Funcs(funcMap).Parse(serviceReadStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//data 模板
	dataTmpPath := fmt.Sprintf("%s/build/template/data.tpl", projectPath)
	dataReadStr, _ := file.ReadStr(dataTmpPath)
	dataTemplate, err := template.New("dataTemplate").Funcs(funcMap).Parse(dataReadStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, prod := range prods {
		tmpDir := fmt.Sprintf("%s/build/%s/schema", projectPath, prod)
		file.MakeDir(tmpDir)
		sprintf := fmt.Sprintf("%s/apps/%s/internal/data/ent/schema", projectPath, prod)
		err := file.CopyDir(sprintf, tmpDir)
		if err != nil {
			fmt.Errorf("%s", err.Error())
			continue
		}

		graph, err := entc.LoadGraph(tmpDir, &gen.Config{
			Storage: storage,
			IDType:  &field.TypeInfo{Type: field.TypeInt64},
			Target:  target,
			Package: "entgo.io/ent/entc/integration/ent",
			/*Templates: []*gen.Template{
				gen.MustParse(gen.NewTemplate("template").
					Funcs(gen.Funcs).
					ParseGlob("../integration/ent/template/*.tmpl")),
			},*/

		})
		if err != nil {
			fmt.Printf("%s", err.Error())
			continue
		}
		schemas := graph.Schemas
		for _, sc := range schemas {
			genConvert(projectPath, prod, sc)
			fields := sc.Fields
			name := sc.Name
			//sc.Name
			log.Printf("%v", fields)
			m := make(map[string]interface{})
			m["model"] = prod
			m["name"] = name
			m["llName"] = str.LeftLower(name)
			lower := strings.ToLower(name)
			m["pkg"] = lower
			m["fields"] = fields

			//生成api
			apiFileName := fmt.Sprintf("%s/api/%s/%s/v1/%s.proto", projectPath, prod, lower, str.Camel2Case(name))
			genFile(apiTemplate, m, apiFileName)

			//生成biz
			bizFileName := fmt.Sprintf("%s/apps/%s/internal/biz/%s.go", projectPath, prod, str.Camel2Case(name))
			genFile(bizTemplate, m, bizFileName)

			//生成data
			dataFileName := fmt.Sprintf("%s/apps/%s/internal/data/%s.go", projectPath, prod, str.Camel2Case(name))
			genFile(dataTemplate, m, dataFileName)

			//生成data
			serviceFileName := fmt.Sprintf("%s/apps/%s/internal/service/%s.go", projectPath, prod, str.Camel2Case(name))
			genFile(serviceTemplate, m, serviceFileName)
		}
		genProvider(projectPath, prod, schemas)
		genRegServer(projectPath, prod, schemas)
		file.RemoveAll(tmpDir)
	}
}

//根据模板
func genFile(tmp *template.Template, m interface{}, fileName string) {
	bf := str.NewBuffer()
	err := tmp.Execute(bf, m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dir := path.Dir(fileName)
	file.MakeDir(dir)
	file.FileCreate(bf, fileName)
}

func parseType(f *load.Field) string {
	t := f.Info.Type
	//r := f.Info.RType
	switch {
	case t == field.TypeInt:
		return "int32"
	case t == field.TypeTime:
		return "google.protobuf.Timestamp"
	case f.Info.Ident == "time.Duration":
		return "google.protobuf.Duration"
	default:
		return t.String()
	}
}

func genFields(fields []*load.Field, isModify bool) string {
	bf := str.NewBuffer()
	bf.Append("    int64 id = 1;\n")
	tab := "    "
	space := " "
	//len := len(fields)
	for i, f := range fields {
		name := f.Name
		if name == "createdAt" && isModify {
			break
		}
		//备注
		bf.Append(tab).Append("//").Append(strings.ReplaceAll(f.Comment, "\n", " ")).Append("\n")
		//字段
		bf.Append(tab).Append(parseType(f)).Append(space).Append(name).
			Append(space).Append("=").Append(space).Append(i + 2).Append(";\n")
	}
	return bf.String()
}

func genCreateFields(fields []*load.Field) string {
	bf := str.NewBuffer()
	//bf.Append("    int64 id = 1;\n")
	tab := "    "
	space := " "
	//len := len(fields)
	for i, f := range fields {
		name := f.Name
		if name == "createdAt" {
			break
		}
		//备注
		bf.Append(tab).Append("//").Append(f.Comment).Append("\n")
		//字段
		bf.Append(tab).Append(parseType(f)).Append(space).Append(name).
			Append(space).Append("=").Append(space).Append(i + 1).Append(";\n")
	}
	return bf.String()
}

//生成创建SetXxx
func genCreateSetFields(fields []*load.Field) string {
	bf := str.NewBuffer()
	//bf.Append("    int64 id = 1;\n")
	tab := "    "
	//len := len(fields)
	for _, f := range fields {
		name := f.Name
		if name == "createdAt" {
			break
		}
		entName, protoName := GetEntNameAndProtoName(f)
		inner := "req." + protoName
		t := f.Info.Type
		switch {
		case t == field.TypeTime:
			inner += ".AsTime()"
		case f.Info.Ident == "time.Duration":
			inner += ".AsDuration()"
		case t == field.TypeEnum:
			inner = fmt.Sprintf("%s(%s)", f.Info.Ident, inner)
		case str.IsNotBlank(f.Info.Ident):
			inner = fmt.Sprintf("%s(%s)", f.Info.Ident, inner)
		}
		//todo duration 字段处理
		//备注
		bf.Append(tab).Append(fmt.Sprintf("Set%s(%s).\n", entName, inner))
	}
	return bf.String()
}

func GetEntNameAndProtoName(f *load.Field) (entName, protoName string) {
	name := f.Name
	name = str.Case2Camel(name)
	leftUpper := str.LeftUpper(name)
	entName, protoName = leftUpper, leftUpper
	//字段
	switch {
	case f.Info.Type == field.TypeEnum:
	case strings.EqualFold(name, "url"):
		entName = "URL"
	case strings.EqualFold(name, "Md5code"):
		protoName = "Md5Code"
	case strings.Contains(name, "State"):
		fmt.Println(name)
	}
	return
}

//生成查询条件
func genCondition(pkg string, fields []*load.Field) string {
	bf := str.NewBuffer()
	strTmp := `if str.IsBlank(req.%s) {
		list = append(list, %s.%sContains(req.%s))
	}
	`
	numTmp := `if req.%s > 0 {
		list = append(list, %s.%s(req.%s))
	}
	`
	timeTmp := `if req.%s.IsValid() && !req.%s.AsTime().IsZero() {
		list = append(list, %s.%sGTE(req.%s.AsTime()))
	}
	`

	enumTmp := `%s := %s(req.%s)
	if %sValidator(%s)==nil {
		list = append(list, %sEQ(%s))
	}
`
	for _, f := range fields {
		entName, protoName := GetEntNameAndProtoName(f)
		t := f.Info.Type
		switch {
		case
			t == field.TypeInt,
			t == field.TypeInt32,
			t == field.TypeInt8,
			t == field.TypeInt16,
			t == field.TypeUint8,
			t == field.TypeUint16,
			t == field.TypeUint32,
			t == field.TypeUint,
			t == field.TypeUint64,
			t == field.TypeFloat32,
			t == field.TypeFloat64,
			t == field.TypeInt64:
			if f.Info.Ident == "time.Duration" {
				protoName += ".AsDuration()"
			}
			bf.Append(fmt.Sprintf(numTmp, protoName, pkg, entName, protoName))
		case t == field.TypeTime:
			bf.Append(fmt.Sprintf(timeTmp, protoName, protoName, pkg, entName, protoName))
		case t == field.TypeString:
			bf.Append(fmt.Sprintf(strTmp, protoName, pkg, entName, protoName))
		case t == field.TypeEnum:
			name := f.Name
			ident := f.Info.Ident
			bf.Append(fmt.Sprintf(enumTmp, name, ident, protoName, ident, name, ident, name))
		case str.IsNotBlank(f.Info.Ident):
			name := f.Name
			ident := f.Info.Ident
			bf.Append(fmt.Sprintf(enumTmp, name, ident, protoName, ident, name, ident, name))
		}
		//todo duration 字段处理
		//备注

	}
	return bf.String()
}

func add(a, b int) int {
	return a + b
}

//生成转化代码
func genConvert(projectPath, model string, sc *load.Schema) bool {
	bf := str.NewBuffer()
	bf.Append(fmt.Sprintf(`// Code generated by Caisin. DO NOT EDIT.
// source: %s

package convert

`, fmt.Sprintf("apps/%s/internal/data/ent/schema/%s.go", model, str.Camel2Case(sc.Name))))
	importTemp := `import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/%s/%s/v1"
	"hope/apps/%s/internal/data/ent"
)`
	name := sc.Name
	funTmp := `
func %s(v *%s) *%s {
	if v == nil {
		return nil
	}
	return &%s{
%s
	}
}

`
	update2Data := str.NewBuffer()
	data2Update := str.NewBuffer()
	update2Data.Append("\t\tID:       v.Id,\n")
	data2Update.Append("\t\tId:       v.ID,\n")
	create2Data := str.NewBuffer()
	data2Create := str.NewBuffer()
	req2Data := str.NewBuffer()
	data2Req := str.NewBuffer()
	reply2Data := str.NewBuffer()
	data2Reply := str.NewBuffer()
	reply2Data.Append("\t\tID:       v.Id,\n")
	data2Reply.Append("\t\tId:       v.ID,\n")
	fields := sc.Fields
	l := len(fields)
	duflag := false
	emflag := false
	if sc.Name == "PayOrder" {
		fmt.Println()
	}
	for i, f := range fields {
		entName, protoName := GetEntNameAndProtoName(f)
		toEnt := ""
		toProto := ""
		t := f.Info.Type
		switch {
		case t == field.TypeTime:
			toEnt = fmt.Sprintf("\t\t%s:       v.%s.AsTime(),\n", entName, protoName)
			toProto = fmt.Sprintf("\t\t%s:       timestamppb.New(v.%s),\n", protoName, entName)
		case f.Info.Ident == "time.Duration":
			toEnt = fmt.Sprintf("\t\t%s:       v.%s.AsDuration(),\n", entName, protoName)
			toProto = fmt.Sprintf("\t\t%s:       durationpb.New(v.%s),\n", protoName, entName)
			duflag = true
		case t == field.TypeEnum:
			toEnt = fmt.Sprintf("\t\t%s:       %s(v.%s),\n", entName, f.Info.Ident, protoName)
			toProto = fmt.Sprintf("\t\t%s:       string(v.%s),\n", protoName, entName)
			emflag = true
		case str.IsNotBlank(f.Info.Ident):
			toEnt = fmt.Sprintf("\t\t%s:       %s(v.%s),\n", entName, f.Info.Ident, protoName)
			toProto = fmt.Sprintf("\t\t%s:       %s(v.%s),\n", protoName, f.Info.Type, entName)
			emflag = true
		default:
			toEnt = fmt.Sprintf("\t\t%s:       v.%s,\n", entName, protoName)
			toProto = fmt.Sprintf("\t\t%s:       v.%s,\n", protoName, entName)

		}
		reply2Data.Append(toEnt)
		data2Reply.Append(toProto)
		if l-i > 6 {
			create2Data.Append(toEnt)
			data2Create.Append(toProto)
			update2Data.Append(toEnt)
			data2Update.Append(toProto)
			req2Data.Append(toEnt)
			data2Req.Append(toProto)
		}
	}
	lower := strings.ToLower(name)
	if duflag {
		importTemp = strings.ReplaceAll(importTemp, "\"google.golang.org/protobuf/types/known/timestamppb\"",
			"\"google.golang.org/protobuf/types/known/timestamppb\"\n\t\"google.golang.org/protobuf/types/known/durationpb\"")
	}
	if emflag {
		importTemp = strings.ReplaceAll(importTemp, "\"google.golang.org/protobuf/types/known/timestamppb\"",
			fmt.Sprintf("\"google.golang.org/protobuf/types/known/timestamppb\"\n\t\"hope/apps/%s/internal/data/ent/%s\"", model, lower))
	}
	bf.Append(fmt.Sprintf(importTemp, model, lower, model))
	appendFun(bf, update2Data, data2Update, funTmp, name, "Update", "Req")
	appendFun(bf, create2Data, data2Create, funTmp, name, "Create", "Req")
	appendFun(bf, req2Data, data2Req, funTmp, name, "", "Req")
	appendFun(bf, reply2Data, data2Reply, funTmp, name, "", "Reply")
	appendFun(bf, update2Data, data2Update, funTmp, name, "Update", "Reply")
	appendFun(bf, reply2Data, data2Reply, funTmp, name, "Create", "Reply")
	fileName := fmt.Sprintf("%s/apps/%s/internal/convert/%s.go", projectPath, model, str.Camel2Case(name))
	dir := path.Dir(fileName)
	file.MakeDir(dir)
	file.FileCreate(bf, fileName)
	return true
}

func appendFun(bf, toData, toReq *str.Buffer, funTmp, name, mode, typ string) {
	modeName := name + mode
	entName := "ent." + name
	v1Name := "v1." + modeName + typ
	bf.Append(fmt.Sprintf(funTmp, modeName+typ+"2Data", v1Name, entName, entName, toData.String()))
	bf.Append(fmt.Sprintf(funTmp, name+"Data2"+mode+typ, entName, v1Name, v1Name, toReq.String()))
}

func genProvider(projectPath, prod string, scs []*load.Schema) {
	tmp := `package %s

import "github.com/google/wire"

// ProviderSet is %s providers.
var ProviderSet = wire.NewSet(
%s
)
`
	biz := str.NewBuffer()
	data := str.NewBuffer()
	//NewEntClient, NewRedisClient, NewData
	data.Append("\tNewEntClient,\n")
	data.Append("\tNewRedisClient,\n")
	data.Append("\tNewData,\n")
	service := str.NewBuffer()
	for _, sc := range scs {
		name := sc.Name
		biz.Append("\tNew").Append(name).Append("UseCase,\n")
		data.Append("\tNew").Append(name).Append("Repo,\n")
		service.Append("\tNew").Append(name).Append("Service,\n")
	}
	inPath := fmt.Sprintf("%s/apps/%s/internal/", projectPath, prod)

	bizStr := "biz"
	file.FileCreate(str.NewBuffer().Append(fmt.Sprintf(tmp, bizStr, bizStr, biz)), inPath+bizStr+"/biz.go")

	dataStr := "data"
	file.FileCreate(str.NewBuffer().Append(fmt.Sprintf(tmp, dataStr, dataStr, data)), inPath+dataStr+"/provider.go")

	serviceStr := "service"
	file.FileCreate(str.NewBuffer().Append(fmt.Sprintf(tmp, serviceStr, serviceStr, service)), inPath+serviceStr+"/service.go")
}

func genRegServer(projectPath, prod string, scs []*load.Schema) {
	tmp := `package server

import (
	"github.com/go-kratos/kratos/v2/transport/{{.svcType}}"
{{ range .scs }}
	{{.pkg}} "hope/api/{{$.mode}}/{{.pkg}}/v1"
{{- end }}
	"hope/apps/{{.mode}}/internal/service"
)

func Register{{.upSvcType}}Server(
{{ range .scs }}
	{{.llName}}Service *service.{{.name}}Service,
{{- end }}
) []func(*{{.svcType}}.Server) {
	list := make([]func(*{{.svcType}}.Server), 0)
{{ range .scs }}
	list = append(list, func(srv *{{$.svcType}}.Server) {
		{{.pkg}}.Register{{.name}}{{$.regType}}Server(srv, {{.llName}}Service)
	})
{{- end }}
	return list
}
`
	parse, err := template.New("tpl").Parse(tmp)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	list := make([]map[string]interface{}, 0)
	for _, sc := range scs {
		item := make(map[string]interface{})
		name := sc.Name
		item["name"] = name
		item["llName"] = str.LeftLower(name)
		item["pkg"] = strings.ToLower(name)
		list = append(list, item)
	}
	m["scs"] = list
	m["mode"] = prod
	m["upSvcType"] = "HTTP"
	m["svcType"] = "http"
	m["regType"] = "HTTP"
	httpFileName := fmt.Sprintf("%s/apps/%s/internal/server/reg_http.go", projectPath, prod)
	genFile(parse, m, httpFileName)
	m["upSvcType"] = "GRPC"
	m["svcType"] = "grpc"
	m["regType"] = ""
	grpcFileName := fmt.Sprintf("%s/apps/%s/internal/server/reg_grpc.go", projectPath, prod)
	genFile(parse, m, grpcFileName)
}
