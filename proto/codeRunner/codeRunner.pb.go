// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/codeRunner/codeRunner.proto

package CodeRunner

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CodeLanguage int32

const (
	CodeLanguage_Cpp  CodeLanguage = 0
	CodeLanguage_Java CodeLanguage = 1
)

// Enum value maps for CodeLanguage.
var (
	CodeLanguage_name = map[int32]string{
		0: "Cpp",
		1: "Java",
	}
	CodeLanguage_value = map[string]int32{
		"Cpp":  0,
		"Java": 1,
	}
)

func (x CodeLanguage) Enum() *CodeLanguage {
	p := new(CodeLanguage)
	*p = x
	return p
}

func (x CodeLanguage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeLanguage) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_codeRunner_codeRunner_proto_enumTypes[0].Descriptor()
}

func (CodeLanguage) Type() protoreflect.EnumType {
	return &file_proto_codeRunner_codeRunner_proto_enumTypes[0]
}

func (x CodeLanguage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodeLanguage.Descriptor instead.
func (CodeLanguage) EnumDescriptor() ([]byte, []int) {
	return file_proto_codeRunner_codeRunner_proto_rawDescGZIP(), []int{0}
}

type ResultType int32

const (
	ResultType_Sucess                ResultType = 0
	ResultType_Compile_Error         ResultType = 1
	ResultType_Runtime_Error         ResultType = 2
	ResultType_Time_Limit_Exceeded   ResultType = 3
	ResultType_Memory_Limit_Exceeded ResultType = 4
	ResultType_System_Error          ResultType = 5
)

// Enum value maps for ResultType.
var (
	ResultType_name = map[int32]string{
		0: "Sucess",
		1: "Compile_Error",
		2: "Runtime_Error",
		3: "Time_Limit_Exceeded",
		4: "Memory_Limit_Exceeded",
		5: "System_Error",
	}
	ResultType_value = map[string]int32{
		"Sucess":                0,
		"Compile_Error":         1,
		"Runtime_Error":         2,
		"Time_Limit_Exceeded":   3,
		"Memory_Limit_Exceeded": 4,
		"System_Error":          5,
	}
)

func (x ResultType) Enum() *ResultType {
	p := new(ResultType)
	*p = x
	return p
}

func (x ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_codeRunner_codeRunner_proto_enumTypes[1].Descriptor()
}

func (ResultType) Type() protoreflect.EnumType {
	return &file_proto_codeRunner_codeRunner_proto_enumTypes[1]
}

func (x ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultType.Descriptor instead.
func (ResultType) EnumDescriptor() ([]byte, []int) {
	return file_proto_codeRunner_codeRunner_proto_rawDescGZIP(), []int{1}
}

type CodeRunnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language_ CodeLanguage `protobuf:"varint,1,opt,name=language_,json=language,proto3,enum=CodeRunner.CodeLanguage" json:"language_,omitempty"`
	Code_     string       `protobuf:"bytes,2,opt,name=code_,json=code,proto3" json:"code_,omitempty"`
	Input_    string       `protobuf:"bytes,3,opt,name=input_,json=input,proto3" json:"input_,omitempty"`
}

func (x *CodeRunnerRequest) Reset() {
	*x = CodeRunnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_codeRunner_codeRunner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeRunnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeRunnerRequest) ProtoMessage() {}

func (x *CodeRunnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_codeRunner_codeRunner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeRunnerRequest.ProtoReflect.Descriptor instead.
func (*CodeRunnerRequest) Descriptor() ([]byte, []int) {
	return file_proto_codeRunner_codeRunner_proto_rawDescGZIP(), []int{0}
}

func (x *CodeRunnerRequest) GetLanguage_() CodeLanguage {
	if x != nil {
		return x.Language_
	}
	return CodeLanguage_Cpp
}

func (x *CodeRunnerRequest) GetCode_() string {
	if x != nil {
		return x.Code_
	}
	return ""
}

func (x *CodeRunnerRequest) GetInput_() string {
	if x != nil {
		return x.Input_
	}
	return ""
}

type CodeRunnerRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type_   ResultType `protobuf:"varint,1,opt,name=type_,json=type,proto3,enum=CodeRunner.ResultType" json:"type_,omitempty"`
	Result_ string     `protobuf:"bytes,2,opt,name=result_,json=result,proto3" json:"result_,omitempty"`
}

func (x *CodeRunnerRespone) Reset() {
	*x = CodeRunnerRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_codeRunner_codeRunner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeRunnerRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeRunnerRespone) ProtoMessage() {}

func (x *CodeRunnerRespone) ProtoReflect() protoreflect.Message {
	mi := &file_proto_codeRunner_codeRunner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeRunnerRespone.ProtoReflect.Descriptor instead.
func (*CodeRunnerRespone) Descriptor() ([]byte, []int) {
	return file_proto_codeRunner_codeRunner_proto_rawDescGZIP(), []int{1}
}

func (x *CodeRunnerRespone) GetType_() ResultType {
	if x != nil {
		return x.Type_
	}
	return ResultType_Sucess
}

func (x *CodeRunnerRespone) GetResult_() string {
	if x != nil {
		return x.Result_
	}
	return ""
}

var File_proto_codeRunner_codeRunner_proto protoreflect.FileDescriptor

var file_proto_codeRunner_codeRunner_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x22,
	0x76, 0x0a, 0x11, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x59, 0x0a, 0x11, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x2b, 0x0a, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x5f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2a, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x70, 0x70, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a,
	0x61, 0x76, 0x61, 0x10, 0x01, 0x2a, 0x84, 0x01, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x75, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x5f, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x69, 0x6d, 0x65, 0x5f, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x5f, 0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x03, 0x12,
	0x19, 0x0a, 0x15, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x5f,
	0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x5f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x05, 0x32, 0x55, 0x0a, 0x0a,
	0x63, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x05, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_codeRunner_codeRunner_proto_rawDescOnce sync.Once
	file_proto_codeRunner_codeRunner_proto_rawDescData = file_proto_codeRunner_codeRunner_proto_rawDesc
)

func file_proto_codeRunner_codeRunner_proto_rawDescGZIP() []byte {
	file_proto_codeRunner_codeRunner_proto_rawDescOnce.Do(func() {
		file_proto_codeRunner_codeRunner_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_codeRunner_codeRunner_proto_rawDescData)
	})
	return file_proto_codeRunner_codeRunner_proto_rawDescData
}

var file_proto_codeRunner_codeRunner_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_codeRunner_codeRunner_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_codeRunner_codeRunner_proto_goTypes = []interface{}{
	(CodeLanguage)(0),         // 0: CodeRunner.codeLanguage
	(ResultType)(0),           // 1: CodeRunner.resultType
	(*CodeRunnerRequest)(nil), // 2: CodeRunner.codeRunnerRequest
	(*CodeRunnerRespone)(nil), // 3: CodeRunner.codeRunnerRespone
}
var file_proto_codeRunner_codeRunner_proto_depIdxs = []int32{
	0, // 0: CodeRunner.codeRunnerRequest.language_:type_name -> CodeRunner.codeLanguage
	1, // 1: CodeRunner.codeRunnerRespone.type_:type_name -> CodeRunner.resultType
	2, // 2: CodeRunner.codeRunner.judge:input_type -> CodeRunner.codeRunnerRequest
	3, // 3: CodeRunner.codeRunner.judge:output_type -> CodeRunner.codeRunnerRespone
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_codeRunner_codeRunner_proto_init() }
func file_proto_codeRunner_codeRunner_proto_init() {
	if File_proto_codeRunner_codeRunner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_codeRunner_codeRunner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeRunnerRequest); i {
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
		file_proto_codeRunner_codeRunner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeRunnerRespone); i {
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
			RawDescriptor: file_proto_codeRunner_codeRunner_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_codeRunner_codeRunner_proto_goTypes,
		DependencyIndexes: file_proto_codeRunner_codeRunner_proto_depIdxs,
		EnumInfos:         file_proto_codeRunner_codeRunner_proto_enumTypes,
		MessageInfos:      file_proto_codeRunner_codeRunner_proto_msgTypes,
	}.Build()
	File_proto_codeRunner_codeRunner_proto = out.File
	file_proto_codeRunner_codeRunner_proto_rawDesc = nil
	file_proto_codeRunner_codeRunner_proto_goTypes = nil
	file_proto_codeRunner_codeRunner_proto_depIdxs = nil
}
