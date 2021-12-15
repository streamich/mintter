// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: changes.proto

package backend

import (
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

type DocumentChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TitleUpdated    string `protobuf:"bytes,1,opt,name=title_updated,json=titleUpdated,proto3" json:"title_updated,omitempty"`
	SubtitleUpdated string `protobuf:"bytes,2,opt,name=subtitle_updated,json=subtitleUpdated,proto3" json:"subtitle_updated,omitempty"`
	ContentUpdated  []byte `protobuf:"bytes,3,opt,name=content_updated,json=contentUpdated,proto3" json:"content_updated,omitempty"`
	// Only set on the first patch ever by the original author.
	Author string `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	// Only set on the first patch ever by the original author.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *DocumentChange) Reset() {
	*x = DocumentChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_changes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentChange) ProtoMessage() {}

func (x *DocumentChange) ProtoReflect() protoreflect.Message {
	mi := &file_changes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentChange.ProtoReflect.Descriptor instead.
func (*DocumentChange) Descriptor() ([]byte, []int) {
	return file_changes_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentChange) GetTitleUpdated() string {
	if x != nil {
		return x.TitleUpdated
	}
	return ""
}

func (x *DocumentChange) GetSubtitleUpdated() string {
	if x != nil {
		return x.SubtitleUpdated
	}
	return ""
}

func (x *DocumentChange) GetContentUpdated() []byte {
	if x != nil {
		return x.ContentUpdated
	}
	return nil
}

func (x *DocumentChange) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *DocumentChange) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *DocumentChange) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type DocumentFeedChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentPublished string `protobuf:"bytes,1,opt,name=document_published,json=documentPublished,proto3" json:"document_published,omitempty"`
}

func (x *DocumentFeedChange) Reset() {
	*x = DocumentFeedChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_changes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentFeedChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentFeedChange) ProtoMessage() {}

func (x *DocumentFeedChange) ProtoReflect() protoreflect.Message {
	mi := &file_changes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentFeedChange.ProtoReflect.Descriptor instead.
func (*DocumentFeedChange) Descriptor() ([]byte, []int) {
	return file_changes_proto_rawDescGZIP(), []int{1}
}

func (x *DocumentFeedChange) GetDocumentPublished() string {
	if x != nil {
		return x.DocumentPublished
	}
	return ""
}

type AccountChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewDeviceProof string `protobuf:"bytes,1,opt,name=new_device_proof,json=newDeviceProof,proto3" json:"new_device_proof,omitempty"`
	NewEmail       string `protobuf:"bytes,2,opt,name=new_email,json=newEmail,proto3" json:"new_email,omitempty"`
	NewBio         string `protobuf:"bytes,3,opt,name=new_bio,json=newBio,proto3" json:"new_bio,omitempty"`
	NewAlias       string `protobuf:"bytes,4,opt,name=new_alias,json=newAlias,proto3" json:"new_alias,omitempty"`
}

func (x *AccountChange) Reset() {
	*x = AccountChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_changes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountChange) ProtoMessage() {}

func (x *AccountChange) ProtoReflect() protoreflect.Message {
	mi := &file_changes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountChange.ProtoReflect.Descriptor instead.
func (*AccountChange) Descriptor() ([]byte, []int) {
	return file_changes_proto_rawDescGZIP(), []int{2}
}

func (x *AccountChange) GetNewDeviceProof() string {
	if x != nil {
		return x.NewDeviceProof
	}
	return ""
}

func (x *AccountChange) GetNewEmail() string {
	if x != nil {
		return x.NewEmail
	}
	return ""
}

func (x *AccountChange) GetNewBio() string {
	if x != nil {
		return x.NewBio
	}
	return ""
}

func (x *AccountChange) GetNewAlias() string {
	if x != nil {
		return x.NewAlias
	}
	return ""
}

var File_changes_proto protoreflect.FileDescriptor

var file_changes_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x6d, 0x69, 0x6e, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x0e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x43, 0x0a, 0x12, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x64, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x77, 0x5f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6e, 0x65, 0x77, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x65, 0x77, 0x5f, 0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x65, 0x77, 0x42, 0x69, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x41, 0x6c,
	0x69, 0x61, 0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_changes_proto_rawDescOnce sync.Once
	file_changes_proto_rawDescData = file_changes_proto_rawDesc
)

func file_changes_proto_rawDescGZIP() []byte {
	file_changes_proto_rawDescOnce.Do(func() {
		file_changes_proto_rawDescData = protoimpl.X.CompressGZIP(file_changes_proto_rawDescData)
	})
	return file_changes_proto_rawDescData
}

var file_changes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_changes_proto_goTypes = []interface{}{
	(*DocumentChange)(nil),        // 0: mintter.changes.DocumentChange
	(*DocumentFeedChange)(nil),    // 1: mintter.changes.DocumentFeedChange
	(*AccountChange)(nil),         // 2: mintter.changes.AccountChange
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_changes_proto_depIdxs = []int32{
	3, // 0: mintter.changes.DocumentChange.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: mintter.changes.DocumentChange.update_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_changes_proto_init() }
func file_changes_proto_init() {
	if File_changes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_changes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentChange); i {
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
		file_changes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentFeedChange); i {
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
		file_changes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountChange); i {
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
			RawDescriptor: file_changes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_changes_proto_goTypes,
		DependencyIndexes: file_changes_proto_depIdxs,
		MessageInfos:      file_changes_proto_msgTypes,
	}.Build()
	File_changes_proto = out.File
	file_changes_proto_rawDesc = nil
	file_changes_proto_goTypes = nil
	file_changes_proto_depIdxs = nil
}