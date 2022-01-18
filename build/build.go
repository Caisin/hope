package main

import (
	"bytes"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
	"fmt"
	"hope/pkg/file"
	"hope/pkg/str"
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
	s := fmt.Sprintf("%s/build/template/api.proto.tpl", projectPath)
	readStr, _ := file.ReadStr(s)
	parse, err := template.New("api").Funcs(
		template.FuncMap{
			"genPageFields":   genPageFields,
			"genCreateFields": genCreateFields,
			"genFields":       genFields,
			"add":             add,
			"parseType":       parseType,
		}).Parse(readStr)
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
			IDType:  &field.TypeInfo{Type: field.TypeInt},
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

		for _, sc := range graph.Schemas {
			fields := sc.Fields
			name := sc.Name
			//sc.Name
			log.Printf("%v", fields)
			bf := bytes.Buffer{}
			m := make(map[string]interface{})
			m["model"] = prod
			m["name"] = name
			lower := strings.ToLower(name)
			m["pkg"] = lower
			m["fields"] = fields
			err := parse.Execute(&bf, m)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fileName := fmt.Sprintf("%s/api/%s/%s/v1/%s.proto", projectPath, prod, lower, str.Camel2Case(name))
			dir := path.Dir(fileName)
			file.MakeDir(dir)
			file.FileCreate(bf, fileName)
		}

		//file.RemoveAll(tmpDir)

	}

}

func parseType(t field.Type) string {
	switch t.String() {
	case "int":
		return "int32"
	case "time.Time":
		return "google.protobuf.Timestamp"
	case "time.Duration":
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
		bf.Append(tab).Append("//").Append(f.Comment).Append("\n")
		//字段
		bf.Append(tab).Append(parseType(f.Info.Type)).Append(space).Append(name).
			Append(space).Append("=").Append(space).Append(i + 2).Append(";\n")
	}
	return bf.String()
}

func genPageFields(fields []*load.Field) string {
	bf := str.NewBuffer()
	bf.Append("    int64 page = 1;\n")
	bf.Append("    int64 pageSize = 2;\n")
	bf.Append("    int64 id = 3;\n")
	tab := "    "
	space := " "
	//len := len(fields)
	for i, f := range fields {
		name := f.Name
		//备注
		bf.Append(tab).Append("//").Append(f.Comment).Append("\n")
		//字段
		bf.Append(tab).Append(parseType(f.Info.Type)).Append(space).Append(name).
			Append(space).Append("=").Append(space).Append(i + 4).Append(";\n")
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
		bf.Append(tab).Append(parseType(f.Info.Type)).Append(space).Append(name).
			Append(space).Append("=").Append(space).Append(i + 1).Append(";\n")
	}
	return bf.String()
}

func add(a, b int) int {
	return a + b
}
