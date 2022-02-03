package funmap

import (
	"google.golang.org/protobuf/compiler/protogen"
	"hope/pkg/util/str"
	"strings"
)

var FuncMap = map[string]any{
	"tc":          trimComment,
	"camel2Split": str.Camel2Split,
}

func trimComment(c protogen.Comments) string {
	str := string(c)
	str = strings.ReplaceAll(str, "//", "")
	str = strings.TrimSpace(str)
	return str
}
