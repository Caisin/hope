package pagin

import (
	"entgo.io/ent/dialect/sql"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"reflect"
	"sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//分页对象
type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//第page页
	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	//每页大小
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	//总条数
	Total int64 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	//排序字段
	Field string `protobuf:"bytes,4,opt,name=field,proto3" json:"field,omitempty"`
	//排序类型
	Order string `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_pagin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_pagin_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Pagination) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Pagination) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

var File_pagin_proto protoreflect.FileDescriptor

var file_pagin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x22, 0x7e, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x16, 0x5a, 0x14, 0x68, 0x6f, 0x70, 0x65, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x3b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pagin_proto_rawDescOnce sync.Once
	file_pagin_proto_rawDescData = file_pagin_proto_rawDesc
)

func file_pagin_proto_rawDescGZIP() []byte {
	file_pagin_proto_rawDescOnce.Do(func() {
		file_pagin_proto_rawDescData = protoimpl.X.CompressGZIP(file_pagin_proto_rawDescData)
	})
	return file_pagin_proto_rawDescData
}

var file_pagin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pagin_proto_goTypes = []interface{}{
	(*Pagination)(nil), // 0: pagin.Pagination
}
var file_pagin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pagin_proto_init() }
func file_pagin_proto_init() {
	if File_pagin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pagin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pagin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pagin_proto_goTypes,
		DependencyIndexes: file_pagin_proto_depIdxs,
		MessageInfos:      file_pagin_proto_msgTypes,
	}.Build()
	File_pagin_proto = out.File
	file_pagin_proto_rawDesc = nil
	file_pagin_proto_goTypes = nil
	file_pagin_proto_depIdxs = nil
}

//下面是改过的方法

func (x *Pagination) IsDesc() bool {
	return "descend" == x.GetOrder()
}

func (x *Pagination) NeedOrder() bool {
	return x.GetField() != ""
}

func (x *Pagination) OrderFun(desc, asc func(fields ...string) func(*sql.Selector)) func(*sql.Selector) {
	if !x.NeedOrder() {
		return nil
	}
	if x.IsDesc() {
		return desc(x.GetField())
	}
	return asc(x.GetField())
}

func (x *Pagination) GetPage() int64 {
	if x != nil {
		if x.Page <= 0 {
			x.Page = 1
		}
		return x.Page
	}
	return 1
}

func (x *Pagination) GetPageSize() int64 {
	if x != nil {
		if x.PageSize <= 0 {
			x.PageSize = 10
		}
		return x.PageSize
	}
	return 10
}

func (x *Pagination) GetOffSet() int64 {
	if x != nil {
		return (x.GetPage() - 1) * x.GetPageSize()
	}
	return 0
}
