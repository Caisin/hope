package main

import (
	"encoding/json"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
	"fmt"
	"hope/pkg/file"
	"hope/pkg/util/str"
	"os"
	"path/filepath"
	"testing"
	"text/template"
	"time"
)

func TestJson(t *testing.T) {
	dir := "D:/work/code/go/hope/api"
	fiels, err := file.WalkDir(dir, "swagger.json")
	if err != nil {
		panic(err)
	}
	swg := map[string]interface{}{
		"swagger": "3.0.3",
		"info": map[string]interface{}{
			"title":   "hope文档",
			"version": 1.0,
		},
		"produces": []string{"application/json"},
		"consumes": []string{"application/json"},
	}

	for _, f := range fiels {
		str := file.ReadString(f)
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(str), &m)
		if err != nil {
			continue
		}
		tags, ok := swg["tags"]
		mt := m["tags"]
		if ok {
			switch mt.(type) {
			case []interface{}:
				swg["tags"] = append(tags.([]interface{}), mt.([]interface{})...)
			}
		} else {
			swg["tags"] = mt
		}

		paths, ok := swg["paths"]
		if ok {
			switch m["paths"].(type) {
			case map[string]interface{}:
				mp := m["paths"].(map[string]interface{})
				p := paths.(map[string]interface{})
				for k, v := range mp {
					p[k] = v
				}
				swg["paths"] = p
			}
		} else {
			swg["paths"] = m["paths"]
		}

		definitions, ok := swg["definitions"]
		if ok {
			switch m["definitions"].(type) {
			case map[string]interface{}:
				mp := m["definitions"].(map[string]interface{})
				p := definitions.(map[string]interface{})
				for k, v := range mp {
					p[k] = v
				}
				swg["definitions"] = p
			}
		} else {
			swg["definitions"] = m["definitions"]
		}
	}
	bytes, _ := json.Marshal(swg)
	bf := str.NewBuffer()
	bf.Write(bytes)
	file.FileCreate(bf, dir+"/swagger.json")
}

func TestAdmin(t *testing.T) {
	target := filepath.Join(os.TempDir(), "ent")
	os.MkdirAll(target, os.ModePerm)
	defer os.RemoveAll(target)
	storage, _ := gen.NewStorage("sql")
	projectPath := "D:/work/code/go/hope"
	admPath := "D:/work/code/vue/vue-vben-admin"
	prods := []string{"admin", "param", "novel"}
	funcMap := template.FuncMap{
		"tsType":             tsType,
		"formTime":           formTime,
		"fcp":                fcp,
		"dealTime":           dealTime,
		"hasStr":             hasStr,
		"genCondition":       genCondition,
		"genCreateSetFields": genCreateSetFields,
		"genCreateFields":    genCreateFields,
		"genFields":          genFields,
		"add":                add,
		"parseType":          parseType,
	}
	//dataTs模板
	dataTsTemplate := createTemplate(projectPath, "data.ts.template", funcMap)

	//dataTs模板
	indexTemplate := createTemplate(projectPath, "index.vue.template", funcMap)

	//dataTs模板
	modalTemplate := createTemplate(projectPath, "Modal.vue.template", funcMap)

	//dataTs模板
	modelTemplate := createTemplate(projectPath, "model.template", funcMap)

	//dataTs模板
	apiTemplate := createTemplate(projectPath, "api.ts.template", funcMap)
	id := 0
	bf := str.NewBuffer()
	for _, prod := range prods {
		tmpDir := fmt.Sprintf("%s/build/%s/schema", projectPath, prod)
		err := file.MakeDir(tmpDir)
		if err != nil {
			return
		}
		sprintf := fmt.Sprintf("%s/apps/%s/internal/data/ent/schema", projectPath, prod)
		err = file.CopyDir(sprintf, tmpDir)
		if err != nil {
			_ = fmt.Errorf("%s", err.Error())
			continue
		}

		graph, err := entc.LoadGraph(tmpDir, &gen.Config{
			Storage: storage,
			IDType:  &field.TypeInfo{Type: field.TypeInt64},
			Target:  target,
			Package: "entgo.io/ent/entc/integration/ent",
		})
		if err != nil {
			fmt.Printf("%s", err.Error())
			continue
		}
		schemas := graph.Schemas
		for _, sc := range schemas {
			fields := sc.Fields
			name := sc.Name
			//sc.Name
			m := make(map[string]interface{})
			m["model"] = prod
			m["name"] = name
			lower := str.LeftLower(name)
			m["pkg"] = lower
			m["apiPath"] = str.Camel2Split(name, "/")
			m["fields"] = fields

			dataTsFileName := fmt.Sprintf("%s/src/views/%s/%s/%s.data.ts", admPath, prod, lower, lower)
			genFile(dataTsTemplate, m, dataTsFileName)
			indexFileName := fmt.Sprintf("%s/src/views/%s/%s/index.vue", admPath, prod, lower)
			genFile(indexTemplate, m, indexFileName)
			modalFileName := fmt.Sprintf("%s/src/views/%s/%s/%sModal.vue", admPath, prod, lower, name)
			genFile(modalTemplate, m, modalFileName)
			modelFileName := fmt.Sprintf("%s/src/api/%s/model/%s.ts", admPath, prod, lower)
			genFile(modelTemplate, m, modelFileName)
			apiFileName := fmt.Sprintf("%s/src/api/%s/%s.ts", admPath, prod, lower)
			genFile(apiTemplate, m, apiFileName)
		}

		id = genMenu(id, prod, schemas, bf)
		file.RemoveAll(tmpDir)
	}
	file.FileCreate(bf, projectPath+"/menu.sql")
}

func createTemplate(projectPath, tempName string, funcMap template.FuncMap) *template.Template {
	tmpPath := fmt.Sprintf("%s/build/template/adm/%s", projectPath, tempName)
	tempStr, _ := file.ReadStr(tmpPath)
	tp, err := template.New(tempName).Funcs(funcMap).Parse(tempStr)
	if err != nil {
		panic(err)
	}
	return tp
}

func tsType(tp field.Type) string {
	switch tp {
	case
		field.TypeBool:
		return "boolean"
	case
		field.TypeInt,
		field.TypeUint,
		field.TypeInt8,
		field.TypeUint8,
		field.TypeInt16,
		field.TypeUint16,
		field.TypeInt32,
		field.TypeFloat32,
		field.TypeUint32:
		return "number"
	case
		field.TypeString,
		field.TypeInt64,
		field.TypeTime:
		return "string"
	default:
		return "string"
	}
}

//处理时间
func dealTime(f load.Field) string {
	if f.Info.Type == field.TypeTime {
		return "\n    format: (text) => formatToDateTime(text),"
	}
	return ""
}

//表单时间处理
func formTime(f load.Field) string {
	if f.Info.Type == field.TypeTime {
		return "\n    componentProps: {\n      showTime: true,\n    },"
	}
	return ""
}

//通过field取表单组件
func fcp(f load.Field) string {
	tp := f.Info.Type
	switch {
	case
		tp == field.TypeTime:
		return "DatePicker"
	case
		tp == field.TypeBool:
		return "Switch"
	case
		f.Name == "desc",
		f.Name == "summary",
		f.Name == "content",
		f.Name == "remark":
		return "InputTextArea"
	case
		tp == field.TypeInt,
		tp == field.TypeUint,
		tp == field.TypeInt8,
		tp == field.TypeUint8,
		tp == field.TypeInt16,
		tp == field.TypeUint16,
		tp == field.TypeInt32,
		tp == field.TypeFloat32,
		tp == field.TypeUint32:
		return "InputNumber"
		/*	case
			field.TypeString,
			field.TypeInt64,
			field.TypeTime:
			return "InputNumber"*/
	default:
		return "Input"
	}
}
func TestName(t *testing.T) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now)
}
func genMenu(id int, mod string, scs []*load.Schema, bf *str.Buffer) int {
	pId := 0
	now := time.Now().Format("2006-01-02 15:04:05")
	id++
	menu := fmt.Sprintf(`
INSERT INTO sys_menus (id, title, icon, path, paths,
                       menu_type, action, permission, component, sort,
                       created_at, updated_at, create_by, update_by, tenant_id,
                       parent_id, name, redirect, ignore_keep_alive, hide_breadcrumb,
                       hide_children_in_menu, hide_menu, frame_src)
VALUES (%d, '%s', null, '/%s', '/',
        '%s', null, '%s', '%s', null,
        '%s', '%s', 1, 1, 1,
        %d, '%s', null, null, null,
        null, null, null);
`,
		id, mod+"管理", mod,
		"D", "", "Layout",
		now, now,
		pId, str.LeftUpper(mod),
	)
	bf.WriteString(menu)
	pId = id
	for _, sc := range scs {
		name := sc.Name
		title := name
		s, ok := tableNameMapping[sc.Name]
		if ok {
			title = s
		}
		lwName := str.LeftLower(name)
		path := str.LeftLower(sc.Name)
		basePerm := str.Camel2Split(name, ":")
		component := fmt.Sprintf("/%s/%s/index", mod, lwName)
		id++
		fpId := id
		query := fmt.Sprintf(`
INSERT INTO sys_menus (id, title, icon, path, paths,
                       menu_type, action, permission, component, sort,
                       created_at, updated_at, create_by, update_by, tenant_id,
                       parent_id, name, redirect, ignore_keep_alive, hide_breadcrumb,
                       hide_children_in_menu, hide_menu, frame_src)
VALUES (%d, '%s', null, '/%s/%s', '/',
        '%s', null, '%s', '%s', null,
        '%s', '%s', 1, 1, 1,
        %d, '%s', '', null, null,
        null, null, null);
`,
			id, title+"管理", mod, path,
			"M", basePerm+":query", component,
			now, now,
			pId, name,
		)
		bf.WriteString(query)
		id++
		create := fmt.Sprintf(`
INSERT INTO sys_menus (id, title, icon, path, paths,
                       menu_type, action, permission, component, sort,
                       created_at, updated_at, create_by, update_by, tenant_id,
                       parent_id, name, redirect, ignore_keep_alive, hide_breadcrumb,
                       hide_children_in_menu, hide_menu, frame_src)
VALUES (%d, '%s', null, '', '/',
        '%s', null, '%s', '%s', null,
        '%s', '%s', 1, 1, 1,
        %d, '%s', '', null, null,
        null, null, null);
`,
			id, title+"新增",
			"F", basePerm+":create", "",
			now, now,
			fpId, "",
		)
		bf.WriteString(create)
		id++
		del := fmt.Sprintf(`
INSERT INTO sys_menus (id, title, icon, path, paths,
                       menu_type, action, permission, component, sort,
                       created_at, updated_at, create_by, update_by, tenant_id,
                       parent_id, name, redirect, ignore_keep_alive, hide_breadcrumb,
                       hide_children_in_menu, hide_menu, frame_src)
VALUES (%d, '%s', null, '', '/',
        '%s', null, '%s', '%s', null,
        '%s', '%s', 1, 1, 1,
        %d, '', '', null, null,
        null, null, null);
`,
			id, title+"删除",
			"F", basePerm+":delete", "",
			now, now,
			fpId,
		)
		bf.WriteString(del)
		id++
		batchDelete := fmt.Sprintf(`
INSERT INTO sys_menus (id, title, icon, path, paths,
                       menu_type, action, permission, component, sort,
                       created_at, updated_at, create_by, update_by, tenant_id,
                       parent_id, name, redirect, ignore_keep_alive, hide_breadcrumb,
                       hide_children_in_menu, hide_menu, frame_src)
VALUES (%d, '%s', null, '', '/',
        '%s', null, '%s', '%s', null,
        '%s', '%s', 1, 1, 1,
        %d, '', '', null, null,
        null, null, null);
`,
			id, title+"批量删除",
			"F", basePerm+":batchDelete", "",
			now, now,
			fpId,
		)
		bf.WriteString(batchDelete)
		id++
		update := fmt.Sprintf(`
INSERT INTO sys_menus (id, title, icon, path, paths,
                       menu_type, action, permission, component, sort,
                       created_at, updated_at, create_by, update_by, tenant_id,
                       parent_id, name, redirect, ignore_keep_alive, hide_breadcrumb,
                       hide_children_in_menu, hide_menu, frame_src)
VALUES (%d, '%s', null, '', '/',
        '%s', null, '%s', '%s', null,
        '%s', '%s', 1, 1, 1,
        %d, '', '', null, null,
        null, null, null);
`,
			id, title+"修改",
			"F", basePerm+":update", "",
			now, now,
			fpId,
		)
		bf.WriteString(update)

	}
	return id
}
