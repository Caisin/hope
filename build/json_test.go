package main

import (
	"encoding/json"
	"hope/pkg/file"
	"hope/pkg/util/str"
	"testing"
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
