syntax = "proto3";

package {{.pkg}}.v1;

import "google/api/annotations.proto";
import "google/protobuf/duration.proto";
import "google/protobuf/timestamp.proto";
import "google/protobuf/descriptor.proto";
import "pagin.proto";

option go_package = "hope/api/{{.model}}/{{.pkg}}/v1;v1";

// The {{.name}} service definition.
service {{.name}} {

  // 分页查询{{.name}}
  rpc GetPage{{.name}} ({{.name}}PageReq) returns ({{.name}}PageReply)  {
    option (google.api.http) = {
      get: "/v1/{{.apiPath}}/page"
    };
  }

  // 获取{{.name}}
  rpc Get{{.name}} ({{.name}}Req) returns ({{.name}}Reply)  {
    option (google.api.http) = {
      get: "/v1/{{.apiPath}}/{id}"
    };
  }

  // 更新{{.name}}
  rpc Update{{.name}} ({{.name}}UpdateReq) returns ({{.name}}UpdateReply)  {
    option (google.api.http) = {
      put: "/v1/{{.apiPath}}/{id}"
      body: "*"
    };
  }

  // 创建{{.name}}
  rpc Create{{.name}} ({{.name}}CreateReq) returns ({{.name}}CreateReply)  {
    option (google.api.http) = {
      post: "/v1/{{.apiPath}}"
      body: "*"
    };
  }

  // 删除{{.name}}
  rpc Delete{{.name}} ({{.name}}DeleteReq) returns ({{.name}}DeleteReply)  {
    option (google.api.http) = {
      delete: "/v1/{{.apiPath}}/{id}"
    };
  }

  // 批量删除{{.name}}
  rpc BatchDelete{{.name}} ({{.name}}BatchDeleteReq) returns ({{.name}}DeleteReply)  {
    option (google.api.http) = {
      delete: "/v1/{{.apiPath}}"
    };
  }

}

// 实体数据
message {{.name}}Data {
{{genFields .fields false 0}}
}

// 查询搜索请求
message {{.name}}PageReq {
    //分页查询参数
    pagin.Pagination pagin = 1;
    //查询条件参数
    {{.name}}Req param = 2;
}

// 查询搜索返回
message {{.name}}PageReply {
    // 返回码
    int32 code = 1;
    // 消息
    string message = 2;
    // 总条数
    int64 total = 3;
    // 查询数据
    repeated {{.name}}Data items = 4;
}


// 查询搜索请求
message {{.name}}Req {
{{genFields .fields false 0}}
}

// 查询返回
message {{.name}}Reply {
    // 返回码
    int32 code = 1;
    // 消息
    string message = 2;
    // 结果数据
    {{.name}}Data result = 3;
}

// 创建{{.name}}请求
message {{.name}}CreateReq {
{{genCreateFields .fields}}
}

// 创建{{.name}}返回
message {{.name}}CreateReply {
    // 返回码
    int32 code = 1;
    // 消息
    string message = 2;
    // 结果数据
    {{.name}}Data result = 3;
}

// 更新{{.name}}请求
message {{.name}}UpdateReq {
{{genFields .fields true 0}}
}


// 更新{{.name}}返回
message {{.name}}UpdateReply {
    // 返回码
    int32 code = 1;
    // 消息
    string message = 2;
    // 结果数据
    {{.name}}Data result = 3;
}

// 删除{{.name}}请求
message {{.name}}DeleteReq {
    int64 id = 1;
}


// 批量删除{{.name}}请求
message {{.name}}BatchDeleteReq {
    repeated int64 ids = 1;
}


// 删除{{.name}}返回
message {{.name}}DeleteReply {
    // 返回码
    int32 code = 1;
    // 消息
    string message = 2;
    // 结果
    bool result = 3;
}

