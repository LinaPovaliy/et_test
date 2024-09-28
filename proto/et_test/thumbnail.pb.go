// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: thumbnail.proto

package et_test

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

type GetThumbnailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetThumbnailRequest) Reset() {
	*x = GetThumbnailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thumbnail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThumbnailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThumbnailRequest) ProtoMessage() {}

func (x *GetThumbnailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thumbnail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThumbnailRequest.ProtoReflect.Descriptor instead.
func (*GetThumbnailRequest) Descriptor() ([]byte, []int) {
	return file_thumbnail_proto_rawDescGZIP(), []int{0}
}

func (x *GetThumbnailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetThumbnailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetThumbnailResponse) Reset() {
	*x = GetThumbnailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thumbnail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThumbnailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThumbnailResponse) ProtoMessage() {}

func (x *GetThumbnailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thumbnail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThumbnailResponse.ProtoReflect.Descriptor instead.
func (*GetThumbnailResponse) Descriptor() ([]byte, []int) {
	return file_thumbnail_proto_rawDescGZIP(), []int{1}
}

func (x *GetThumbnailResponse) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_thumbnail_proto protoreflect.FileDescriptor

var file_thumbnail_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x32, 0x5b, 0x0a,
	0x10, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f,
	0x65, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thumbnail_proto_rawDescOnce sync.Once
	file_thumbnail_proto_rawDescData = file_thumbnail_proto_rawDesc
)

func file_thumbnail_proto_rawDescGZIP() []byte {
	file_thumbnail_proto_rawDescOnce.Do(func() {
		file_thumbnail_proto_rawDescData = protoimpl.X.CompressGZIP(file_thumbnail_proto_rawDescData)
	})
	return file_thumbnail_proto_rawDescData
}

var file_thumbnail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_thumbnail_proto_goTypes = []any{
	(*GetThumbnailRequest)(nil),  // 0: proto.GetThumbnailRequest
	(*GetThumbnailResponse)(nil), // 1: proto.GetThumbnailResponse
}
var file_thumbnail_proto_depIdxs = []int32{
	0, // 0: proto.ThumbnailService.GetThumbnail:input_type -> proto.GetThumbnailRequest
	1, // 1: proto.ThumbnailService.GetThumbnail:output_type -> proto.GetThumbnailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_thumbnail_proto_init() }
func file_thumbnail_proto_init() {
	if File_thumbnail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thumbnail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetThumbnailRequest); i {
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
		file_thumbnail_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetThumbnailResponse); i {
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
			RawDescriptor: file_thumbnail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thumbnail_proto_goTypes,
		DependencyIndexes: file_thumbnail_proto_depIdxs,
		MessageInfos:      file_thumbnail_proto_msgTypes,
	}.Build()
	File_thumbnail_proto = out.File
	file_thumbnail_proto_rawDesc = nil
	file_thumbnail_proto_goTypes = nil
	file_thumbnail_proto_depIdxs = nil
}
