// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.21.12
// source: ml.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	mi := &file_ml_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Hello struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Res           string                 `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Hello) Reset() {
	*x = Hello{}
	mi := &file_ml_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{1}
}

func (x *Hello) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

type SymptomsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SymptomsRequest) Reset() {
	*x = SymptomsRequest{}
	mi := &file_ml_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SymptomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymptomsRequest) ProtoMessage() {}

func (x *SymptomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymptomsRequest.ProtoReflect.Descriptor instead.
func (*SymptomsRequest) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{2}
}

func (x *SymptomsRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SymptomsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Symptoms      []*Symptom             `protobuf:"bytes,1,rep,name=symptoms,proto3" json:"symptoms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SymptomsResponse) Reset() {
	*x = SymptomsResponse{}
	mi := &file_ml_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SymptomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymptomsResponse) ProtoMessage() {}

func (x *SymptomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymptomsResponse.ProtoReflect.Descriptor instead.
func (*SymptomsResponse) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{3}
}

func (x *SymptomsResponse) GetSymptoms() []*Symptom {
	if x != nil {
		return x.Symptoms
	}
	return nil
}

type Symptom struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Loc           string                 `protobuf:"bytes,3,opt,name=loc,proto3" json:"loc,omitempty"`
	Confidence    float32                `protobuf:"fixed32,4,opt,name=confidence,proto3" json:"confidence,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Symptom) Reset() {
	*x = Symptom{}
	mi := &file_ml_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Symptom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Symptom) ProtoMessage() {}

func (x *Symptom) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Symptom.ProtoReflect.Descriptor instead.
func (*Symptom) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{4}
}

func (x *Symptom) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Symptom) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Symptom) GetLoc() string {
	if x != nil {
		return x.Loc
	}
	return ""
}

func (x *Symptom) GetConfidence() float32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

var File_ml_proto protoreflect.FileDescriptor

var file_ml_proto_rawDesc = string([]byte{
	0x0a, 0x08, 0x6d, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x19,
	0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x0f, 0x53, 0x79, 0x6d,
	0x70, 0x74, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x38, 0x0a, 0x10, 0x53, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x73, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d,
	0x52, 0x08, 0x73, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x22, 0x63, 0x0a, 0x07, 0x53, 0x79,
	0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x6f, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x63, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x32,
	0x66, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x33, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x12,
	0x10, 0x2e, 0x53, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x53, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x0d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x06, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x77, 0x74, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66,
	0x74, 0x68, 0x69, 0x73, 0x31, 0x32, 0x33, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_ml_proto_rawDescOnce sync.Once
	file_ml_proto_rawDescData []byte
)

func file_ml_proto_rawDescGZIP() []byte {
	file_ml_proto_rawDescOnce.Do(func() {
		file_ml_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ml_proto_rawDesc), len(file_ml_proto_rawDesc)))
	})
	return file_ml_proto_rawDescData
}

var file_ml_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ml_proto_goTypes = []any{
	(*HelloRequest)(nil),     // 0: HelloRequest
	(*Hello)(nil),            // 1: Hello
	(*SymptomsRequest)(nil),  // 2: SymptomsRequest
	(*SymptomsResponse)(nil), // 3: SymptomsResponse
	(*Symptom)(nil),          // 4: Symptom
}
var file_ml_proto_depIdxs = []int32{
	4, // 0: SymptomsResponse.symptoms:type_name -> Symptom
	2, // 1: TokensServer.FindSymptoms:input_type -> SymptomsRequest
	0, // 2: TokensServer.SayHello:input_type -> HelloRequest
	3, // 3: TokensServer.FindSymptoms:output_type -> SymptomsResponse
	1, // 4: TokensServer.SayHello:output_type -> Hello
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ml_proto_init() }
func file_ml_proto_init() {
	if File_ml_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ml_proto_rawDesc), len(file_ml_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ml_proto_goTypes,
		DependencyIndexes: file_ml_proto_depIdxs,
		MessageInfos:      file_ml_proto_msgTypes,
	}.Build()
	File_ml_proto = out.File
	file_ml_proto_goTypes = nil
	file_ml_proto_depIdxs = nil
}
