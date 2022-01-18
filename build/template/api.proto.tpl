syntax = "proto3";

package {{.pkg}}.v1;

import "google/api/annotations.proto";
import "google/protobuf/duration.proto";
import "google/protobuf/timestamp.proto";

option go_package = "hope/api/{{.model}}/{{.pkg}}/v1;v1";

// The {{.name}} service definition.
service {{.name}} {
  // 获取{{.name}}
  rpc Get{{.name}} ({{.name}}Request) returns ({{.name}}Reply)  {
    option (google.api.http) = {
      get: "/v1/{{.pkg}}/{id}"
    };
  }
}

// 查询搜索请求
message {{.name}}Request {
    int64 id = 1;
{{- range .fields}}
    //{{.Comment}}
    {{parseType .Info.Type}} {{.Name}} = {{add .Position.Index 2}};
{{- end}}
}

// 查询搜索返回
message {{.name}}Reply {
int64 id = 1;
{{- range .fields}}
    //{{.Comment}}
    {{parseType .Info.Type}} {{.Name}} = {{add .Position.Index 2}};
{{- end}}
}