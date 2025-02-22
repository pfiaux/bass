// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: runtime.proto

package proto

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

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Inner:
	//	*RunResponse_Progress
	Inner isRunResponse_Inner `protobuf_oneof:"inner"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runtime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{0}
}

func (m *RunResponse) GetInner() isRunResponse_Inner {
	if m != nil {
		return m.Inner
	}
	return nil
}

func (x *RunResponse) GetProgress() *Progress {
	if x, ok := x.GetInner().(*RunResponse_Progress); ok {
		return x.Progress
	}
	return nil
}

type isRunResponse_Inner interface {
	isRunResponse_Inner()
}

type RunResponse_Progress struct {
	Progress *Progress `protobuf:"bytes,1,opt,name=progress,proto3,oneof"`
}

func (*RunResponse_Progress) isRunResponse_Inner() {}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Inner:
	//	*ReadResponse_Progress
	//	*ReadResponse_Output
	Inner isReadResponse_Inner `protobuf_oneof:"inner"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runtime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{1}
}

func (m *ReadResponse) GetInner() isReadResponse_Inner {
	if m != nil {
		return m.Inner
	}
	return nil
}

func (x *ReadResponse) GetProgress() *Progress {
	if x, ok := x.GetInner().(*ReadResponse_Progress); ok {
		return x.Progress
	}
	return nil
}

func (x *ReadResponse) GetOutput() []byte {
	if x, ok := x.GetInner().(*ReadResponse_Output); ok {
		return x.Output
	}
	return nil
}

type isReadResponse_Inner interface {
	isReadResponse_Inner()
}

type ReadResponse_Progress struct {
	Progress *Progress `protobuf:"bytes,1,opt,name=progress,proto3,oneof"`
}

type ReadResponse_Output struct {
	Output []byte `protobuf:"bytes,2,opt,name=output,proto3,oneof"`
}

func (*ReadResponse_Progress) isReadResponse_Inner() {}

func (*ReadResponse_Output) isReadResponse_Inner() {}

type Bytes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Bytes) Reset() {
	*x = Bytes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runtime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bytes) ProtoMessage() {}

func (x *Bytes) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bytes.ProtoReflect.Descriptor instead.
func (*Bytes) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{2}
}

func (x *Bytes) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_runtime_proto protoreflect.FileDescriptor

var file_runtime_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x62, 0x61, 0x73, 0x73, 0x1a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x44, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x07,
	0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x5f, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42,
	0x07, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x1b, 0x0a, 0x05, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xf0, 0x01, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x62,
	0x61, 0x73, 0x73, 0x2e, 0x54, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x66, 0x1a, 0x13, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x54, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12,
	0x0b, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x54, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x11, 0x2e, 0x62,
	0x61, 0x73, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x2b, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0b, 0x2e, 0x62, 0x61,
	0x73, 0x73, 0x2e, 0x54, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x26, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0b, 0x2e, 0x62, 0x61, 0x73,
	0x73, 0x2e, 0x54, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x54, 0x68,
	0x75, 0x6e, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runtime_proto_rawDescOnce sync.Once
	file_runtime_proto_rawDescData = file_runtime_proto_rawDesc
)

func file_runtime_proto_rawDescGZIP() []byte {
	file_runtime_proto_rawDescOnce.Do(func() {
		file_runtime_proto_rawDescData = protoimpl.X.CompressGZIP(file_runtime_proto_rawDescData)
	})
	return file_runtime_proto_rawDescData
}

var file_runtime_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_runtime_proto_goTypes = []interface{}{
	(*RunResponse)(nil),  // 0: bass.RunResponse
	(*ReadResponse)(nil), // 1: bass.ReadResponse
	(*Bytes)(nil),        // 2: bass.Bytes
	(*Progress)(nil),     // 3: bass.Progress
	(*ImageRef)(nil),     // 4: bass.ThunkImageRef
	(*Thunk)(nil),        // 5: bass.Thunk
	(*ThunkPath)(nil),    // 6: bass.ThunkPath
}
var file_runtime_proto_depIdxs = []int32{
	3, // 0: bass.RunResponse.progress:type_name -> bass.Progress
	3, // 1: bass.ReadResponse.progress:type_name -> bass.Progress
	4, // 2: bass.Runtime.Resolve:input_type -> bass.ThunkImageRef
	5, // 3: bass.Runtime.Run:input_type -> bass.Thunk
	5, // 4: bass.Runtime.Read:input_type -> bass.Thunk
	5, // 5: bass.Runtime.Export:input_type -> bass.Thunk
	6, // 6: bass.Runtime.ExportPath:input_type -> bass.ThunkPath
	4, // 7: bass.Runtime.Resolve:output_type -> bass.ThunkImageRef
	0, // 8: bass.Runtime.Run:output_type -> bass.RunResponse
	1, // 9: bass.Runtime.Read:output_type -> bass.ReadResponse
	2, // 10: bass.Runtime.Export:output_type -> bass.Bytes
	2, // 11: bass.Runtime.ExportPath:output_type -> bass.Bytes
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_runtime_proto_init() }
func file_runtime_proto_init() {
	if File_runtime_proto != nil {
		return
	}
	file_progress_proto_init()
	file_bass_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_runtime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
		file_runtime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_runtime_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bytes); i {
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
	file_runtime_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RunResponse_Progress)(nil),
	}
	file_runtime_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ReadResponse_Progress)(nil),
		(*ReadResponse_Output)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_runtime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runtime_proto_goTypes,
		DependencyIndexes: file_runtime_proto_depIdxs,
		MessageInfos:      file_runtime_proto_msgTypes,
	}.Build()
	File_runtime_proto = out.File
	file_runtime_proto_rawDesc = nil
	file_runtime_proto_goTypes = nil
	file_runtime_proto_depIdxs = nil
}
