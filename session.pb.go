//*
// Standalone signaling server for the Nextcloud Spreed app.
// Copyright (C) 2024 struktur AG
//
// @author Joachim Bauch <bauch@struktur.de>
//
// @license GNU AGPL version 3 or any later version
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

package signaling

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type SessionIdData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sid           uint64                 `protobuf:"varint,1,opt,name=Sid,proto3" json:"Sid,omitempty"`
	Created       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=Created,proto3" json:"Created,omitempty"`
	BackendId     string                 `protobuf:"bytes,3,opt,name=BackendId,proto3" json:"BackendId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SessionIdData) Reset() {
	*x = SessionIdData{}
	mi := &file_session_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SessionIdData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionIdData) ProtoMessage() {}

func (x *SessionIdData) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionIdData.ProtoReflect.Descriptor instead.
func (*SessionIdData) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{0}
}

func (x *SessionIdData) GetSid() uint64 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *SessionIdData) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *SessionIdData) GetBackendId() string {
	if x != nil {
		return x.BackendId
	}
	return ""
}

var File_session_proto protoreflect.FileDescriptor

const file_session_proto_rawDesc = "" +
	"\n" +
	"\rsession.proto\x12\tsignaling\x1a\x1fgoogle/protobuf/timestamp.proto\"u\n" +
	"\rSessionIdData\x12\x10\n" +
	"\x03Sid\x18\x01 \x01(\x04R\x03Sid\x124\n" +
	"\aCreated\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\aCreated\x12\x1c\n" +
	"\tBackendId\x18\x03 \x01(\tR\tBackendIdB<Z:github.com/strukturag/nextcloud-spreed-signaling;signalingb\x06proto3"

var (
	file_session_proto_rawDescOnce sync.Once
	file_session_proto_rawDescData []byte
)

func file_session_proto_rawDescGZIP() []byte {
	file_session_proto_rawDescOnce.Do(func() {
		file_session_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_session_proto_rawDesc), len(file_session_proto_rawDesc)))
	})
	return file_session_proto_rawDescData
}

var file_session_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_session_proto_goTypes = []any{
	(*SessionIdData)(nil),         // 0: signaling.SessionIdData
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_session_proto_depIdxs = []int32{
	1, // 0: signaling.SessionIdData.Created:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_session_proto_init() }
func file_session_proto_init() {
	if File_session_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_session_proto_rawDesc), len(file_session_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_session_proto_goTypes,
		DependencyIndexes: file_session_proto_depIdxs,
		MessageInfos:      file_session_proto_msgTypes,
	}.Build()
	File_session_proto = out.File
	file_session_proto_goTypes = nil
	file_session_proto_depIdxs = nil
}
