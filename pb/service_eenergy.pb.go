// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.6
// source: service_eenergy.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_eenergy_proto protoreflect.FileDescriptor

var file_service_eenergy_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x65, 0x6e, 0x65, 0x72, 0x67,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x72, 0x70, 0x63, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x70, 0x63, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x72, 0x70, 0x63, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x65,
	0x72, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf9, 0x09, 0x0a, 0x0e, 0x45, 0x65, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x92, 0x41, 0x34, 0x12, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a,
	0x21, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f,
	0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x7d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x40, 0x92, 0x41, 0x2a, 0x12, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x32, 0x08, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x12, 0xa3, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x69, 0x92, 0x41, 0x4d, 0x12, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x1a, 0x3f, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74,
	0x6f, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x20, 0x26, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x90, 0x01, 0x0a, 0x0b, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x92, 0x41, 0x35,
	0x12, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x25,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0xb4, 0x01,
	0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x45, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x92, 0x41, 0x4a, 0x12, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x1a, 0x37, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x20, 0x66,
	0x72, 0x6f, 0x6d, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x61,
	0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22,
	0x13, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x12, 0x98, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x92, 0x41, 0x32, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x55, 0x73,
	0x65, 0x72, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0xbc, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x92, 0x41, 0x3b, 0x12, 0x0f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x20, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x28, 0x55, 0x73, 0x65,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x93,
	0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x92, 0x41, 0x33, 0x12, 0x0e, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x42, 0x74, 0x92, 0x41, 0x51, 0x12, 0x4f, 0x0a, 0x07, 0x45, 0x65, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x22, 0x3f, 0x0a, 0x0c, 0x41, 0x68, 0x6d, 0x65, 0x64, 0x20, 0x52, 0x61,
	0x64, 0x77, 0x61, 0x6e, 0x12, 0x1b, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x64, 0x77, 0x61, 0x6e,
	0x6e, 0x1a, 0x12, 0x61, 0x72, 0x61, 0x64, 0x77, 0x61, 0x6e, 0x6e, 0x40, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x64, 0x77, 0x61, 0x6e, 0x6e, 0x2f,
	0x65, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_service_eenergy_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),          // 0: pb.CreateUserRequest
	(*UpdateUserRequest)(nil),          // 1: pb.UpdateUserRequest
	(*LoginUserRequest)(nil),           // 2: pb.LoginUserRequest
	(*VerifyEmailRequest)(nil),         // 3: pb.VerifyEmailRequest
	(*TransferEnergyRequest)(nil),      // 4: pb.TransferEnergyRequest
	(*ListUserAccountsRequest)(nil),    // 5: pb.ListUserAccountsRequest
	(*ListAccountEntriesRequest)(nil),  // 6: pb.ListAccountEntriesRequest
	(*CreateAccountRequest)(nil),       // 7: pb.CreateAccountRequest
	(*CreateUserResponse)(nil),         // 8: pb.CreateUserResponse
	(*UpdateUserResponse)(nil),         // 9: pb.UpdateUserResponse
	(*LoginUserResponse)(nil),          // 10: pb.LoginUserResponse
	(*VerifyEmailResponse)(nil),        // 11: pb.VerifyEmailResponse
	(*TransferEnergyResponse)(nil),     // 12: pb.TransferEnergyResponse
	(*ListUserAccountsResponse)(nil),   // 13: pb.ListUserAccountsResponse
	(*ListAccountEntriesResponse)(nil), // 14: pb.ListAccountEntriesResponse
	(*CreateAccountResponse)(nil),      // 15: pb.CreateAccountResponse
}
var file_service_eenergy_proto_depIdxs = []int32{
	0,  // 0: pb.EenergyService.CreateUser:input_type -> pb.CreateUserRequest
	1,  // 1: pb.EenergyService.UpdateUser:input_type -> pb.UpdateUserRequest
	2,  // 2: pb.EenergyService.LoginUser:input_type -> pb.LoginUserRequest
	3,  // 3: pb.EenergyService.VerifyEmail:input_type -> pb.VerifyEmailRequest
	4,  // 4: pb.EenergyService.TransferEnergy:input_type -> pb.TransferEnergyRequest
	5,  // 5: pb.EenergyService.ListUserAccounts:input_type -> pb.ListUserAccountsRequest
	6,  // 6: pb.EenergyService.ListAccountEntries:input_type -> pb.ListAccountEntriesRequest
	7,  // 7: pb.EenergyService.CreateAccount:input_type -> pb.CreateAccountRequest
	8,  // 8: pb.EenergyService.CreateUser:output_type -> pb.CreateUserResponse
	9,  // 9: pb.EenergyService.UpdateUser:output_type -> pb.UpdateUserResponse
	10, // 10: pb.EenergyService.LoginUser:output_type -> pb.LoginUserResponse
	11, // 11: pb.EenergyService.VerifyEmail:output_type -> pb.VerifyEmailResponse
	12, // 12: pb.EenergyService.TransferEnergy:output_type -> pb.TransferEnergyResponse
	13, // 13: pb.EenergyService.ListUserAccounts:output_type -> pb.ListUserAccountsResponse
	14, // 14: pb.EenergyService.ListAccountEntries:output_type -> pb.ListAccountEntriesResponse
	15, // 15: pb.EenergyService.CreateAccount:output_type -> pb.CreateAccountResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_eenergy_proto_init() }
func file_service_eenergy_proto_init() {
	if File_service_eenergy_proto != nil {
		return
	}
	file_rpc_create_account_proto_init()
	file_rpc_create_user_proto_init()
	file_rpc_list_account_entries_proto_init()
	file_rpc_list_user_accounts_proto_init()
	file_rpc_login_user_proto_init()
	file_rpc_transfer_energy_proto_init()
	file_rpc_update_user_proto_init()
	file_rpc_verify_email_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_eenergy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_eenergy_proto_goTypes,
		DependencyIndexes: file_service_eenergy_proto_depIdxs,
	}.Build()
	File_service_eenergy_proto = out.File
	file_service_eenergy_proto_rawDesc = nil
	file_service_eenergy_proto_goTypes = nil
	file_service_eenergy_proto_depIdxs = nil
}
