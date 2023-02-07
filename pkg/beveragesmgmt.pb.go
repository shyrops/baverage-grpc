// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: beveragesmgmt.proto

package beverage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BeverageType int32

const (
	BeverageType_BEVERAGE_TYPE_UNSPECIFIED BeverageType = 0
	BeverageType_BEVERAGE_TYPE_BEER        BeverageType = 1
	BeverageType_BEVERAGE_TYPE_VINE        BeverageType = 2
	BeverageType_BEVERAGE_TYPE_VODKA       BeverageType = 3
	BeverageType_BEVERAGE_TYPE_GIN         BeverageType = 4
	BeverageType_BEVERAGE_TYPE_RUM         BeverageType = 5
)

// Enum value maps for BeverageType.
var (
	BeverageType_name = map[int32]string{
		0: "BEVERAGE_TYPE_UNSPECIFIED",
		1: "BEVERAGE_TYPE_BEER",
		2: "BEVERAGE_TYPE_VINE",
		3: "BEVERAGE_TYPE_VODKA",
		4: "BEVERAGE_TYPE_GIN",
		5: "BEVERAGE_TYPE_RUM",
	}
	BeverageType_value = map[string]int32{
		"BEVERAGE_TYPE_UNSPECIFIED": 0,
		"BEVERAGE_TYPE_BEER":        1,
		"BEVERAGE_TYPE_VINE":        2,
		"BEVERAGE_TYPE_VODKA":       3,
		"BEVERAGE_TYPE_GIN":         4,
		"BEVERAGE_TYPE_RUM":         5,
	}
)

func (x BeverageType) Enum() *BeverageType {
	p := new(BeverageType)
	*p = x
	return p
}

func (x BeverageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BeverageType) Descriptor() protoreflect.EnumDescriptor {
	return file_beveragesmgmt_proto_enumTypes[0].Descriptor()
}

func (BeverageType) Type() protoreflect.EnumType {
	return &file_beveragesmgmt_proto_enumTypes[0]
}

func (x BeverageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BeverageType.Descriptor instead.
func (BeverageType) EnumDescriptor() ([]byte, []int) {
	return file_beveragesmgmt_proto_rawDescGZIP(), []int{0}
}

type BeverageMainAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   BeverageType `protobuf:"varint,1,opt,name=type,proto3,enum=api.BeverageType" json:"type,omitempty"`
	Name   string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Volume int32        `protobuf:"varint,3,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *BeverageMainAttributes) Reset() {
	*x = BeverageMainAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beveragesmgmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeverageMainAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeverageMainAttributes) ProtoMessage() {}

func (x *BeverageMainAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_beveragesmgmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeverageMainAttributes.ProtoReflect.Descriptor instead.
func (*BeverageMainAttributes) Descriptor() ([]byte, []int) {
	return file_beveragesmgmt_proto_rawDescGZIP(), []int{0}
}

func (x *BeverageMainAttributes) GetType() BeverageType {
	if x != nil {
		return x.Type
	}
	return BeverageType_BEVERAGE_TYPE_UNSPECIFIED
}

func (x *BeverageMainAttributes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BeverageMainAttributes) GetVolume() int32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type CreateBeverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attr *BeverageMainAttributes `protobuf:"bytes,1,opt,name=attr,proto3" json:"attr,omitempty"`
}

func (x *CreateBeverageRequest) Reset() {
	*x = CreateBeverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beveragesmgmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBeverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBeverageRequest) ProtoMessage() {}

func (x *CreateBeverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beveragesmgmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBeverageRequest.ProtoReflect.Descriptor instead.
func (*CreateBeverageRequest) Descriptor() ([]byte, []int) {
	return file_beveragesmgmt_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBeverageRequest) GetAttr() *BeverageMainAttributes {
	if x != nil {
		return x.Attr
	}
	return nil
}

type Beverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Attr  *BeverageMainAttributes `protobuf:"bytes,2,opt,name=attr,proto3" json:"attr,omitempty"`
	Price int32                   `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Beverage) Reset() {
	*x = Beverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beveragesmgmt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Beverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Beverage) ProtoMessage() {}

func (x *Beverage) ProtoReflect() protoreflect.Message {
	mi := &file_beveragesmgmt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Beverage.ProtoReflect.Descriptor instead.
func (*Beverage) Descriptor() ([]byte, []int) {
	return file_beveragesmgmt_proto_rawDescGZIP(), []int{2}
}

func (x *Beverage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Beverage) GetAttr() *BeverageMainAttributes {
	if x != nil {
		return x.Attr
	}
	return nil
}

func (x *Beverage) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_beveragesmgmt_proto protoreflect.FileDescriptor

var file_beveragesmgmt_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x65, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x6b, 0x0a, 0x16, 0x62, 0x65,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x65, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x65, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2f, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x65, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x69,
	0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x04, 0x61, 0x74, 0x74,
	0x72, 0x22, 0x61, 0x0a, 0x08, 0x42, 0x65, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a,
	0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x65, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x2a, 0xa4, 0x01, 0x0a, 0x0c, 0x42, 0x65, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x45, 0x56, 0x45, 0x52, 0x41, 0x47,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x45, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x45, 0x45, 0x52, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x42, 0x45, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49,
	0x4e, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x45, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x4f, 0x44, 0x4b, 0x41, 0x10, 0x03, 0x12, 0x15, 0x0a,
	0x11, 0x42, 0x45, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47,
	0x49, 0x4e, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x45, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x55, 0x4d, 0x10, 0x05, 0x32, 0x52, 0x0a, 0x13, 0x42,
	0x65, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x65, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x65, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x42,
	0x0b, 0x5a, 0x09, 0x2f, 0x62, 0x65, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beveragesmgmt_proto_rawDescOnce sync.Once
	file_beveragesmgmt_proto_rawDescData = file_beveragesmgmt_proto_rawDesc
)

func file_beveragesmgmt_proto_rawDescGZIP() []byte {
	file_beveragesmgmt_proto_rawDescOnce.Do(func() {
		file_beveragesmgmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_beveragesmgmt_proto_rawDescData)
	})
	return file_beveragesmgmt_proto_rawDescData
}

var file_beveragesmgmt_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_beveragesmgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_beveragesmgmt_proto_goTypes = []interface{}{
	(BeverageType)(0),              // 0: api.BeverageType
	(*BeverageMainAttributes)(nil), // 1: api.beverageMainAttributes
	(*CreateBeverageRequest)(nil),  // 2: api.CreateBeverageRequest
	(*Beverage)(nil),               // 3: api.Beverage
}
var file_beveragesmgmt_proto_depIdxs = []int32{
	0, // 0: api.beverageMainAttributes.type:type_name -> api.BeverageType
	1, // 1: api.CreateBeverageRequest.attr:type_name -> api.beverageMainAttributes
	1, // 2: api.Beverage.attr:type_name -> api.beverageMainAttributes
	2, // 3: api.BeveragesManagement.CreateBeverage:input_type -> api.CreateBeverageRequest
	3, // 4: api.BeveragesManagement.CreateBeverage:output_type -> api.Beverage
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_beveragesmgmt_proto_init() }
func file_beveragesmgmt_proto_init() {
	if File_beveragesmgmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beveragesmgmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeverageMainAttributes); i {
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
		file_beveragesmgmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBeverageRequest); i {
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
		file_beveragesmgmt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Beverage); i {
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
			RawDescriptor: file_beveragesmgmt_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_beveragesmgmt_proto_goTypes,
		DependencyIndexes: file_beveragesmgmt_proto_depIdxs,
		EnumInfos:         file_beveragesmgmt_proto_enumTypes,
		MessageInfos:      file_beveragesmgmt_proto_msgTypes,
	}.Build()
	File_beveragesmgmt_proto = out.File
	file_beveragesmgmt_proto_rawDesc = nil
	file_beveragesmgmt_proto_goTypes = nil
	file_beveragesmgmt_proto_depIdxs = nil
}
