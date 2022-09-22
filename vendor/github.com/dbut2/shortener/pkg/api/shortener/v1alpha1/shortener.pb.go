// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: shortener/v1alpha1/shortener.proto

package shortener

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShortenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Code   *string                `protobuf:"bytes,2,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Expiry *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiry,proto3,oneof" json:"expiry,omitempty"`
	Ip     *string                `protobuf:"bytes,4,opt,name=ip,proto3,oneof" json:"ip,omitempty"`
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_v1alpha1_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_v1alpha1_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenRequest.ProtoReflect.Descriptor instead.
func (*ShortenRequest) Descriptor() ([]byte, []int) {
	return file_shortener_v1alpha1_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ShortenRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *ShortenRequest) GetExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *ShortenRequest) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

type ShortenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ShortenResponse) Reset() {
	*x = ShortenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_v1alpha1_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResponse) ProtoMessage() {}

func (x *ShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_v1alpha1_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenResponse.ProtoReflect.Descriptor instead.
func (*ShortenResponse) Descriptor() ([]byte, []int) {
	return file_shortener_v1alpha1_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type LengthenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Ip   string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *LengthenRequest) Reset() {
	*x = LengthenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_v1alpha1_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LengthenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LengthenRequest) ProtoMessage() {}

func (x *LengthenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_v1alpha1_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LengthenRequest.ProtoReflect.Descriptor instead.
func (*LengthenRequest) Descriptor() ([]byte, []int) {
	return file_shortener_v1alpha1_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *LengthenRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LengthenRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type LengthenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *LengthenResponse) Reset() {
	*x = LengthenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_v1alpha1_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LengthenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LengthenResponse) ProtoMessage() {}

func (x *LengthenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_v1alpha1_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LengthenResponse.ProtoReflect.Descriptor instead.
func (*LengthenResponse) Descriptor() ([]byte, []int) {
	return file_shortener_v1alpha1_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *LengthenResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_shortener_v1alpha1_shortener_proto protoreflect.FileDescriptor

var file_shortener_v1alpha1_shortener_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x01, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x13,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x02, 0x69, 0x70,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x70, 0x22, 0x25,
	0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x35, 0x0a, 0x0f, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x24, 0x0a, 0x10,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x32, 0xe8, 0x01, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12, 0x22,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22,
	0x08, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x08,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x12, 0x23, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x2f, 0x7b, 0x63, 0x6f, 0x64, 0x65, 0x7d, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x62, 0x75, 0x74,
	0x32, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x3b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortener_v1alpha1_shortener_proto_rawDescOnce sync.Once
	file_shortener_v1alpha1_shortener_proto_rawDescData = file_shortener_v1alpha1_shortener_proto_rawDesc
)

func file_shortener_v1alpha1_shortener_proto_rawDescGZIP() []byte {
	file_shortener_v1alpha1_shortener_proto_rawDescOnce.Do(func() {
		file_shortener_v1alpha1_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortener_v1alpha1_shortener_proto_rawDescData)
	})
	return file_shortener_v1alpha1_shortener_proto_rawDescData
}

var file_shortener_v1alpha1_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shortener_v1alpha1_shortener_proto_goTypes = []interface{}{
	(*ShortenRequest)(nil),        // 0: shortener.v1alpha1.ShortenRequest
	(*ShortenResponse)(nil),       // 1: shortener.v1alpha1.ShortenResponse
	(*LengthenRequest)(nil),       // 2: shortener.v1alpha1.LengthenRequest
	(*LengthenResponse)(nil),      // 3: shortener.v1alpha1.LengthenResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_shortener_v1alpha1_shortener_proto_depIdxs = []int32{
	4, // 0: shortener.v1alpha1.ShortenRequest.expiry:type_name -> google.protobuf.Timestamp
	0, // 1: shortener.v1alpha1.ShortService.Shorten:input_type -> shortener.v1alpha1.ShortenRequest
	2, // 2: shortener.v1alpha1.ShortService.Lengthen:input_type -> shortener.v1alpha1.LengthenRequest
	1, // 3: shortener.v1alpha1.ShortService.Shorten:output_type -> shortener.v1alpha1.ShortenResponse
	3, // 4: shortener.v1alpha1.ShortService.Lengthen:output_type -> shortener.v1alpha1.LengthenResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shortener_v1alpha1_shortener_proto_init() }
func file_shortener_v1alpha1_shortener_proto_init() {
	if File_shortener_v1alpha1_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortener_v1alpha1_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenRequest); i {
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
		file_shortener_v1alpha1_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenResponse); i {
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
		file_shortener_v1alpha1_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LengthenRequest); i {
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
		file_shortener_v1alpha1_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LengthenResponse); i {
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
	file_shortener_v1alpha1_shortener_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shortener_v1alpha1_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortener_v1alpha1_shortener_proto_goTypes,
		DependencyIndexes: file_shortener_v1alpha1_shortener_proto_depIdxs,
		MessageInfos:      file_shortener_v1alpha1_shortener_proto_msgTypes,
	}.Build()
	File_shortener_v1alpha1_shortener_proto = out.File
	file_shortener_v1alpha1_shortener_proto_rawDesc = nil
	file_shortener_v1alpha1_shortener_proto_goTypes = nil
	file_shortener_v1alpha1_shortener_proto_depIdxs = nil
}
