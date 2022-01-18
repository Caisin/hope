package main

import (
	"bytes"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"fmt"
	"hope/pkg/file"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

func main() {
	target := filepath.Join(os.TempDir(), "ent")
	os.MkdirAll(target, os.ModePerm)
	defer os.RemoveAll(target)
	storage, _ := gen.NewStorage("sql")
	projectPath := "D:/work/code/go/hope"
	prods := []string{"admin", "param", "novel"}
	protoTmp := `syntax = "proto3";

package {{.pkg}}.v1;

import "google/api/annotations.proto";

option go_package = "hope/api/{{.model}}/{{.name}}/v1;v1";

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply)  {
    option (google.api.http) = {
      get: "/helloworld/v1/{name}"
    };
  }
}

// The request message containing the user's name.
message HelloRequest {
  //名字
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  //消息
  string message = 1;
}
`
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
			fmt.Errorf("%s", err.Error())
			continue
		}

		parse, err := template.New("test").Parse(protoTmp)
		if err != nil {
			fmt.Println(err.Error())
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
			m["pkg"] = name
			m["fields"] = fields
			err := parse.Execute(&bf, m)
			if err != nil {
				continue
			}
			fileName := fmt.Sprintf("%s/api/%s/v1/%s.proto", projectPath, prod, name)
			dir := path.Dir(fileName)
			file.MakeDir(dir)
			file.FileCreate(bf, fileName)
		}

		//file.RemoveAll(tmpDir)

	}

}
