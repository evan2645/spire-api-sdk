// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: spire/api/server/debug/v1/debug.proto

package debugv1

import (
	types "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
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

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_server_debug_v1_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_server_debug_v1_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_server_debug_v1_debug_proto_rawDescGZIP(), []int{0}
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Server SVID chain
	SvidChain []*GetInfoResponse_Cert `protobuf:"bytes,1,rep,name=svid_chain,json=svidChain,proto3" json:"svid_chain,omitempty"`
	// Server uptime in seconds
	Uptime int32 `protobuf:"varint,2,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// Amount of registered agents
	AgentsCount int32 `protobuf:"varint,3,opt,name=agents_count,json=agentsCount,proto3" json:"agents_count,omitempty"`
	// Amount of federated bundles
	FederatedBundlesCount int32 `protobuf:"varint,4,opt,name=federated_bundles_count,json=federatedBundlesCount,proto3" json:"federated_bundles_count,omitempty"`
	// Amount of registration entries on database
	EntriesCount int32 `protobuf:"varint,5,opt,name=entries_count,json=entriesCount,proto3" json:"entries_count,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_server_debug_v1_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_server_debug_v1_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_spire_api_server_debug_v1_debug_proto_rawDescGZIP(), []int{1}
}

func (x *GetInfoResponse) GetSvidChain() []*GetInfoResponse_Cert {
	if x != nil {
		return x.SvidChain
	}
	return nil
}

func (x *GetInfoResponse) GetUptime() int32 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *GetInfoResponse) GetAgentsCount() int32 {
	if x != nil {
		return x.AgentsCount
	}
	return 0
}

func (x *GetInfoResponse) GetFederatedBundlesCount() int32 {
	if x != nil {
		return x.FederatedBundlesCount
	}
	return 0
}

func (x *GetInfoResponse) GetEntriesCount() int32 {
	if x != nil {
		return x.EntriesCount
	}
	return 0
}

type GetInfoResponse_Cert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Certificate SPIFFE ID
	Id *types.SPIFFEID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Expiration time
	ExpiresAt int64 `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Subject
	Subject string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *GetInfoResponse_Cert) Reset() {
	*x = GetInfoResponse_Cert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_server_debug_v1_debug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse_Cert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse_Cert) ProtoMessage() {}

func (x *GetInfoResponse_Cert) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_server_debug_v1_debug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse_Cert.ProtoReflect.Descriptor instead.
func (*GetInfoResponse_Cert) Descriptor() ([]byte, []int) {
	return file_spire_api_server_debug_v1_debug_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetInfoResponse_Cert) GetId() *types.SPIFFEID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *GetInfoResponse_Cert) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *GetInfoResponse_Cert) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

var File_spire_api_server_debug_v1_debug_proto protoreflect.FileDescriptor

var file_spire_api_server_debug_v1_debug_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xe5, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x73, 0x76, 0x69, 0x64,
	0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x52, 0x09, 0x73,
	0x76, 0x69, 0x64, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x1a, 0x6a, 0x0a, 0x04, 0x43, 0x65, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x49, 0x44, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x69, 0x0a, 0x05,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x60, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x29, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_server_debug_v1_debug_proto_rawDescOnce sync.Once
	file_spire_api_server_debug_v1_debug_proto_rawDescData = file_spire_api_server_debug_v1_debug_proto_rawDesc
)

func file_spire_api_server_debug_v1_debug_proto_rawDescGZIP() []byte {
	file_spire_api_server_debug_v1_debug_proto_rawDescOnce.Do(func() {
		file_spire_api_server_debug_v1_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_server_debug_v1_debug_proto_rawDescData)
	})
	return file_spire_api_server_debug_v1_debug_proto_rawDescData
}

var file_spire_api_server_debug_v1_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spire_api_server_debug_v1_debug_proto_goTypes = []interface{}{
	(*GetInfoRequest)(nil),       // 0: spire.api.server.debug.v1.GetInfoRequest
	(*GetInfoResponse)(nil),      // 1: spire.api.server.debug.v1.GetInfoResponse
	(*GetInfoResponse_Cert)(nil), // 2: spire.api.server.debug.v1.GetInfoResponse.Cert
	(*types.SPIFFEID)(nil),       // 3: spire.api.types.SPIFFEID
}
var file_spire_api_server_debug_v1_debug_proto_depIdxs = []int32{
	2, // 0: spire.api.server.debug.v1.GetInfoResponse.svid_chain:type_name -> spire.api.server.debug.v1.GetInfoResponse.Cert
	3, // 1: spire.api.server.debug.v1.GetInfoResponse.Cert.id:type_name -> spire.api.types.SPIFFEID
	0, // 2: spire.api.server.debug.v1.Debug.GetInfo:input_type -> spire.api.server.debug.v1.GetInfoRequest
	1, // 3: spire.api.server.debug.v1.Debug.GetInfo:output_type -> spire.api.server.debug.v1.GetInfoResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spire_api_server_debug_v1_debug_proto_init() }
func file_spire_api_server_debug_v1_debug_proto_init() {
	if File_spire_api_server_debug_v1_debug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_server_debug_v1_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoRequest); i {
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
		file_spire_api_server_debug_v1_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse); i {
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
		file_spire_api_server_debug_v1_debug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse_Cert); i {
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
			RawDescriptor: file_spire_api_server_debug_v1_debug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_api_server_debug_v1_debug_proto_goTypes,
		DependencyIndexes: file_spire_api_server_debug_v1_debug_proto_depIdxs,
		MessageInfos:      file_spire_api_server_debug_v1_debug_proto_msgTypes,
	}.Build()
	File_spire_api_server_debug_v1_debug_proto = out.File
	file_spire_api_server_debug_v1_debug_proto_rawDesc = nil
	file_spire_api_server_debug_v1_debug_proto_goTypes = nil
	file_spire_api_server_debug_v1_debug_proto_depIdxs = nil
}
