package funmap

import (
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

var FuncMap = map[string]any{
	"tc": trimComment,
}

func trimComment(c protogen.Comments) string {
	str := string(c)
	str = strings.ReplaceAll(str, "//", "")
	str = strings.TrimSpace(str)
	return str
}
