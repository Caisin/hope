package entgen

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

//ParseGraph 解析Graph
func ParseGraph(schemaPath string, idType field.Type) (*gen.Graph, error) {
	return entc.LoadGraph(schemaPath, &gen.Config{
		IDType: &field.TypeInfo{Type: idType},
	})
}
