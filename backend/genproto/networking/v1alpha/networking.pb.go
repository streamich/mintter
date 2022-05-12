// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.3
// source: networking/v1alpha/networking.proto

package networking

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

// Indicates connection status of our node with a remote peer.
// Mimics libp2p connectedness.
type ConnectionStatus int32

const (
	// NotConnected means no connection to peer, and no extra information (default).
	ConnectionStatus_NOT_CONNECTED ConnectionStatus = 0
	// Connected means has an open, live connection to peer.
	ConnectionStatus_CONNECTED ConnectionStatus = 1
	// CanConnect means recently connected to peer, terminated gracefully.
	ConnectionStatus_CAN_CONNECT ConnectionStatus = 2
	// CannotConnect means recently attempted connecting but failed to connect.
	// (should signal "made effort, failed").
	ConnectionStatus_CANNOT_CONNECT ConnectionStatus = 3
)

// Enum value maps for ConnectionStatus.
var (
	ConnectionStatus_name = map[int32]string{
		0: "NOT_CONNECTED",
		1: "CONNECTED",
		2: "CAN_CONNECT",
		3: "CANNOT_CONNECT",
	}
	ConnectionStatus_value = map[string]int32{
		"NOT_CONNECTED":  0,
		"CONNECTED":      1,
		"CAN_CONNECT":    2,
		"CANNOT_CONNECT": 3,
	}
)

func (x ConnectionStatus) Enum() *ConnectionStatus {
	p := new(ConnectionStatus)
	*p = x
	return p
}

func (x ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_networking_v1alpha_networking_proto_enumTypes[0].Descriptor()
}

func (ConnectionStatus) Type() protoreflect.EnumType {
	return &file_networking_v1alpha_networking_proto_enumTypes[0]
}

func (x ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionStatus.Descriptor instead.
func (ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_networking_v1alpha_networking_proto_rawDescGZIP(), []int{0}
}

// Request to get peer's addresses.
type GetPeerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. CID-encoded Peer ID.
	PeerId string `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (x *GetPeerInfoRequest) Reset() {
	*x = GetPeerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networking_v1alpha_networking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeerInfoRequest) ProtoMessage() {}

func (x *GetPeerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_networking_v1alpha_networking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeerInfoRequest.ProtoReflect.Descriptor instead.
func (*GetPeerInfoRequest) Descriptor() ([]byte, []int) {
	return file_networking_v1alpha_networking_proto_rawDescGZIP(), []int{0}
}

func (x *GetPeerInfoRequest) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

// Request for connecting to a peer explicitly.
type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of multiaddrs for the same peer ID to attempt p2p connection.
	// For example `/ip4/10.0.0.1/tcp/55000/p2p/QmDeadBeef`.
	Addrs []string `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networking_v1alpha_networking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_networking_v1alpha_networking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_networking_v1alpha_networking_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectRequest) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

// Response for conneting to a peer.
type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networking_v1alpha_networking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_networking_v1alpha_networking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_networking_v1alpha_networking_proto_rawDescGZIP(), []int{2}
}

// Various details about a known peer.
type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of known multiaddrs of the request peer.
	Addrs []string `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty"`
	// Connection status of our node with a remote peer.
	ConnectionStatus ConnectionStatus `protobuf:"varint,2,opt,name=connection_status,json=connectionStatus,proto3,enum=com.mintter.networking.v1alpha.ConnectionStatus" json:"connection_status,omitempty"`
	// Account ID that this peer is bound to.
	AccountId string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networking_v1alpha_networking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_networking_v1alpha_networking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_networking_v1alpha_networking_proto_rawDescGZIP(), []int{3}
}

func (x *PeerInfo) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

func (x *PeerInfo) GetConnectionStatus() ConnectionStatus {
	if x != nil {
		return x.ConnectionStatus
	}
	return ConnectionStatus_NOT_CONNECTED
}

func (x *PeerInfo) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

var File_networking_v1alpha_networking_proto protoreflect.FileDescriptor

var file_networking_v1alpha_networking_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x22, 0x11, 0x0a, 0x0f,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x9e, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64,
	0x72, 0x73, 0x12, 0x5d, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x2a, 0x59, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x41, 0x4e, 0x5f, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x03, 0x32, 0xe5, 0x01, 0x0a, 0x0a,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x6b, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x69, 0x6e, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x6a, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x74, 0x65, 0x72,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x74, 0x65, 0x72,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x6d, 0x69, 0x6e, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_networking_v1alpha_networking_proto_rawDescOnce sync.Once
	file_networking_v1alpha_networking_proto_rawDescData = file_networking_v1alpha_networking_proto_rawDesc
)

func file_networking_v1alpha_networking_proto_rawDescGZIP() []byte {
	file_networking_v1alpha_networking_proto_rawDescOnce.Do(func() {
		file_networking_v1alpha_networking_proto_rawDescData = protoimpl.X.CompressGZIP(file_networking_v1alpha_networking_proto_rawDescData)
	})
	return file_networking_v1alpha_networking_proto_rawDescData
}

var file_networking_v1alpha_networking_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_networking_v1alpha_networking_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_networking_v1alpha_networking_proto_goTypes = []interface{}{
	(ConnectionStatus)(0),      // 0: com.mintter.networking.v1alpha.ConnectionStatus
	(*GetPeerInfoRequest)(nil), // 1: com.mintter.networking.v1alpha.GetPeerInfoRequest
	(*ConnectRequest)(nil),     // 2: com.mintter.networking.v1alpha.ConnectRequest
	(*ConnectResponse)(nil),    // 3: com.mintter.networking.v1alpha.ConnectResponse
	(*PeerInfo)(nil),           // 4: com.mintter.networking.v1alpha.PeerInfo
}
var file_networking_v1alpha_networking_proto_depIdxs = []int32{
	0, // 0: com.mintter.networking.v1alpha.PeerInfo.connection_status:type_name -> com.mintter.networking.v1alpha.ConnectionStatus
	1, // 1: com.mintter.networking.v1alpha.Networking.GetPeerInfo:input_type -> com.mintter.networking.v1alpha.GetPeerInfoRequest
	2, // 2: com.mintter.networking.v1alpha.Networking.Connect:input_type -> com.mintter.networking.v1alpha.ConnectRequest
	4, // 3: com.mintter.networking.v1alpha.Networking.GetPeerInfo:output_type -> com.mintter.networking.v1alpha.PeerInfo
	3, // 4: com.mintter.networking.v1alpha.Networking.Connect:output_type -> com.mintter.networking.v1alpha.ConnectResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_networking_v1alpha_networking_proto_init() }
func file_networking_v1alpha_networking_proto_init() {
	if File_networking_v1alpha_networking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_networking_v1alpha_networking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeerInfoRequest); i {
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
		file_networking_v1alpha_networking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_networking_v1alpha_networking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_networking_v1alpha_networking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
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
			RawDescriptor: file_networking_v1alpha_networking_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_networking_v1alpha_networking_proto_goTypes,
		DependencyIndexes: file_networking_v1alpha_networking_proto_depIdxs,
		EnumInfos:         file_networking_v1alpha_networking_proto_enumTypes,
		MessageInfos:      file_networking_v1alpha_networking_proto_msgTypes,
	}.Build()
	File_networking_v1alpha_networking_proto = out.File
	file_networking_v1alpha_networking_proto_rawDesc = nil
	file_networking_v1alpha_networking_proto_goTypes = nil
	file_networking_v1alpha_networking_proto_depIdxs = nil
}