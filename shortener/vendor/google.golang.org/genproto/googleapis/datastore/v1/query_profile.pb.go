// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: google/datastore/v1/query_profile.proto

package datastore

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The mode in which the query request must be processed.
type QueryMode int32

const (
	// The default mode. Only the query results are returned.
	QueryMode_NORMAL QueryMode = 0
	// This mode returns only the query plan, without any results or execution
	// statistics information.
	QueryMode_PLAN QueryMode = 1
	// This mode returns both the query plan and the execution statistics along
	// with the results.
	QueryMode_PROFILE QueryMode = 2
)

// Enum value maps for QueryMode.
var (
	QueryMode_name = map[int32]string{
		0: "NORMAL",
		1: "PLAN",
		2: "PROFILE",
	}
	QueryMode_value = map[string]int32{
		"NORMAL":  0,
		"PLAN":    1,
		"PROFILE": 2,
	}
)

func (x QueryMode) Enum() *QueryMode {
	p := new(QueryMode)
	*p = x
	return p
}

func (x QueryMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_datastore_v1_query_profile_proto_enumTypes[0].Descriptor()
}

func (QueryMode) Type() protoreflect.EnumType {
	return &file_google_datastore_v1_query_profile_proto_enumTypes[0]
}

func (x QueryMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryMode.Descriptor instead.
func (QueryMode) EnumDescriptor() ([]byte, []int) {
	return file_google_datastore_v1_query_profile_proto_rawDescGZIP(), []int{0}
}

// Plan for the query.
type QueryPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Planning phase information for the query. It will include:
	//
	//	{
	//	  "indexes_used": [
	//	    {"query_scope": "Collection", "properties": "(foo ASC, __name__ ASC)"},
	//	    {"query_scope": "Collection", "properties": "(bar ASC, __name__ ASC)"}
	//	  ]
	//	}
	PlanInfo *structpb.Struct `protobuf:"bytes,1,opt,name=plan_info,json=planInfo,proto3" json:"plan_info,omitempty"`
}

func (x *QueryPlan) Reset() {
	*x = QueryPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_v1_query_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPlan) ProtoMessage() {}

func (x *QueryPlan) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_v1_query_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPlan.ProtoReflect.Descriptor instead.
func (*QueryPlan) Descriptor() ([]byte, []int) {
	return file_google_datastore_v1_query_profile_proto_rawDescGZIP(), []int{0}
}

func (x *QueryPlan) GetPlanInfo() *structpb.Struct {
	if x != nil {
		return x.PlanInfo
	}
	return nil
}

// Planning and execution statistics for the query.
type ResultSetStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Plan for the query.
	QueryPlan *QueryPlan `protobuf:"bytes,1,opt,name=query_plan,json=queryPlan,proto3" json:"query_plan,omitempty"`
	// Aggregated statistics from the execution of the query.
	//
	// This will only be present when the request specifies `PROFILE` mode.
	// For example, a query will return the statistics including:
	//
	//	{
	//	  "results_returned": "20",
	//	  "documents_scanned": "20",
	//	  "indexes_entries_scanned": "10050",
	//	  "total_execution_time": "100.7 msecs"
	//	}
	QueryStats *structpb.Struct `protobuf:"bytes,2,opt,name=query_stats,json=queryStats,proto3" json:"query_stats,omitempty"`
}

func (x *ResultSetStats) Reset() {
	*x = ResultSetStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_v1_query_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultSetStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultSetStats) ProtoMessage() {}

func (x *ResultSetStats) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_v1_query_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultSetStats.ProtoReflect.Descriptor instead.
func (*ResultSetStats) Descriptor() ([]byte, []int) {
	return file_google_datastore_v1_query_profile_proto_rawDescGZIP(), []int{1}
}

func (x *ResultSetStats) GetQueryPlan() *QueryPlan {
	if x != nil {
		return x.QueryPlan
	}
	return nil
}

func (x *ResultSetStats) GetQueryStats() *structpb.Struct {
	if x != nil {
		return x.QueryStats
	}
	return nil
}

var File_google_datastore_v1_query_profile_proto protoreflect.FileDescriptor

var file_google_datastore_v1_query_profile_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x09,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x89, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61,
	0x6e, 0x12, 0x38, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2a, 0x2e, 0x0a, 0x09, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d,
	0x41, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4c, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x02, 0x42, 0xc3, 0x01, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c,
	0x56, 0x31, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_datastore_v1_query_profile_proto_rawDescOnce sync.Once
	file_google_datastore_v1_query_profile_proto_rawDescData = file_google_datastore_v1_query_profile_proto_rawDesc
)

func file_google_datastore_v1_query_profile_proto_rawDescGZIP() []byte {
	file_google_datastore_v1_query_profile_proto_rawDescOnce.Do(func() {
		file_google_datastore_v1_query_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_datastore_v1_query_profile_proto_rawDescData)
	})
	return file_google_datastore_v1_query_profile_proto_rawDescData
}

var file_google_datastore_v1_query_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_datastore_v1_query_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_datastore_v1_query_profile_proto_goTypes = []interface{}{
	(QueryMode)(0),          // 0: google.datastore.v1.QueryMode
	(*QueryPlan)(nil),       // 1: google.datastore.v1.QueryPlan
	(*ResultSetStats)(nil),  // 2: google.datastore.v1.ResultSetStats
	(*structpb.Struct)(nil), // 3: google.protobuf.Struct
}
var file_google_datastore_v1_query_profile_proto_depIdxs = []int32{
	3, // 0: google.datastore.v1.QueryPlan.plan_info:type_name -> google.protobuf.Struct
	1, // 1: google.datastore.v1.ResultSetStats.query_plan:type_name -> google.datastore.v1.QueryPlan
	3, // 2: google.datastore.v1.ResultSetStats.query_stats:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_datastore_v1_query_profile_proto_init() }
func file_google_datastore_v1_query_profile_proto_init() {
	if File_google_datastore_v1_query_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_datastore_v1_query_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPlan); i {
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
		file_google_datastore_v1_query_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultSetStats); i {
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
			RawDescriptor: file_google_datastore_v1_query_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_datastore_v1_query_profile_proto_goTypes,
		DependencyIndexes: file_google_datastore_v1_query_profile_proto_depIdxs,
		EnumInfos:         file_google_datastore_v1_query_profile_proto_enumTypes,
		MessageInfos:      file_google_datastore_v1_query_profile_proto_msgTypes,
	}.Build()
	File_google_datastore_v1_query_profile_proto = out.File
	file_google_datastore_v1_query_profile_proto_rawDesc = nil
	file_google_datastore_v1_query_profile_proto_goTypes = nil
	file_google_datastore_v1_query_profile_proto_depIdxs = nil
}
