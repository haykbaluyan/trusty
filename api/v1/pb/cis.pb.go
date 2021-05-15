// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: cis.proto

package pb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "github.com/gogo/googleapis/google/api"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Trust int32

const (
	Trust_Any     Trust = 0 // default
	Trust_Public  Trust = 1
	Trust_Private Trust = 2
)

// Enum value maps for Trust.
var (
	Trust_name = map[int32]string{
		0: "Any",
		1: "Public",
		2: "Private",
	}
	Trust_value = map[string]int32{
		"Any":     0,
		"Public":  1,
		"Private": 2,
	}
)

func (x Trust) Enum() *Trust {
	p := new(Trust)
	*p = x
	return p
}

func (x Trust) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Trust) Descriptor() protoreflect.EnumDescriptor {
	return file_cis_proto_enumTypes[0].Descriptor()
}

func (Trust) Type() protoreflect.EnumType {
	return &file_cis_proto_enumTypes[0]
}

func (x Trust) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Trust.Descriptor instead.
func (Trust) EnumDescriptor() ([]byte, []int) {
	return file_cis_proto_rawDescGZIP(), []int{0}
}

// Root provides X509 Root Cert information
type RootCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the certificate
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Skid provides Subject Key Identifier
	Skid string `protobuf:"bytes,2,opt,name=skid,proto3" json:"skid,omitempty"`
	// NotBefore is the time when the validity period starts
	NotBefore *timestamp.Timestamp `protobuf:"bytes,3,opt,name=not_before,proto3" json:"not_before,omitempty"`
	// NotAfter is the time when the validity period starts
	NotAfter *timestamp.Timestamp `protobuf:"bytes,4,opt,name=not_after,proto3" json:"not_after,omitempty"`
	// Subject name
	Subject string `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	// SHA256 thnumbprint of the cert
	Sha256 string `protobuf:"bytes,6,opt,name=sha256,proto3" json:"sha256,omitempty"`
	// Trust scope
	Trust Trust `protobuf:"varint,7,opt,name=trust,proto3,enum=pb.Trust" json:"trust,omitempty"`
	// PEM encoded certificate
	Pem string `protobuf:"bytes,8,opt,name=pem,proto3" json:"pem,omitempty"`
}

func (x *RootCertificate) Reset() {
	*x = RootCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootCertificate) ProtoMessage() {}

func (x *RootCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_cis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootCertificate.ProtoReflect.Descriptor instead.
func (*RootCertificate) Descriptor() ([]byte, []int) {
	return file_cis_proto_rawDescGZIP(), []int{0}
}

func (x *RootCertificate) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RootCertificate) GetSkid() string {
	if x != nil {
		return x.Skid
	}
	return ""
}

func (x *RootCertificate) GetNotBefore() *timestamp.Timestamp {
	if x != nil {
		return x.NotBefore
	}
	return nil
}

func (x *RootCertificate) GetNotAfter() *timestamp.Timestamp {
	if x != nil {
		return x.NotAfter
	}
	return nil
}

func (x *RootCertificate) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *RootCertificate) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (x *RootCertificate) GetTrust() Trust {
	if x != nil {
		return x.Trust
	}
	return Trust_Any
}

func (x *RootCertificate) GetPem() string {
	if x != nil {
		return x.Pem
	}
	return ""
}

// RootsResponse provides response for GetRootsRequest
type RootsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roots []*RootCertificate `protobuf:"bytes,1,rep,name=roots,proto3" json:"roots,omitempty"`
}

func (x *RootsResponse) Reset() {
	*x = RootsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootsResponse) ProtoMessage() {}

func (x *RootsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootsResponse.ProtoReflect.Descriptor instead.
func (*RootsResponse) Descriptor() ([]byte, []int) {
	return file_cis_proto_rawDescGZIP(), []int{1}
}

func (x *RootsResponse) GetRoots() []*RootCertificate {
	if x != nil {
		return x.Roots
	}
	return nil
}

// Certificate provides X509 Certificate information
type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the certificate
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// OrgId of the certificate, only used with Org scope
	OrgId int64 `protobuf:"varint,2,opt,name=org_id,proto3" json:"org_id,omitempty"`
	// Skid provides Subject Key Identifier
	Skid string `protobuf:"bytes,3,opt,name=skid,proto3" json:"skid,omitempty"`
	// Ikid provides Issuer Key Identifier
	Ikid string `protobuf:"bytes,4,opt,name=ikid,proto3" json:"ikid,omitempty"`
	// SerialNumber provides Serial Number
	SerialNumber string `protobuf:"bytes,5,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	// NotBefore is the time when the validity period starts
	NotBefore *timestamp.Timestamp `protobuf:"bytes,6,opt,name=not_before,proto3" json:"not_before,omitempty"`
	// NotAfter is the time when the validity period starts
	NotAfter *timestamp.Timestamp `protobuf:"bytes,7,opt,name=not_after,proto3" json:"not_after,omitempty"`
	// Subject name
	Subject string `protobuf:"bytes,8,opt,name=subject,proto3" json:"subject,omitempty"`
	// Issuer name
	Issuer string `protobuf:"bytes,9,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// SHA256 thnumbprint of the cert
	Sha256 string `protobuf:"bytes,10,opt,name=sha256,proto3" json:"sha256,omitempty"`
	// Profile of the certificate
	Profile string `protobuf:"bytes,11,opt,name=profile,proto3" json:"profile,omitempty"`
	// Pem encoded certificate
	Pem string `protobuf:"bytes,12,opt,name=pem,proto3" json:"pem,omitempty"`
	// IssuersPem provides PEM encoded issuers
	IssuersPem string `protobuf:"bytes,13,opt,name=issuers_pem,proto3" json:"issuers_pem,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_cis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_cis_proto_rawDescGZIP(), []int{2}
}

func (x *Certificate) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Certificate) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *Certificate) GetSkid() string {
	if x != nil {
		return x.Skid
	}
	return ""
}

func (x *Certificate) GetIkid() string {
	if x != nil {
		return x.Ikid
	}
	return ""
}

func (x *Certificate) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *Certificate) GetNotBefore() *timestamp.Timestamp {
	if x != nil {
		return x.NotBefore
	}
	return nil
}

func (x *Certificate) GetNotAfter() *timestamp.Timestamp {
	if x != nil {
		return x.NotAfter
	}
	return nil
}

func (x *Certificate) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Certificate) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Certificate) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (x *Certificate) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *Certificate) GetPem() string {
	if x != nil {
		return x.Pem
	}
	return ""
}

func (x *Certificate) GetIssuersPem() string {
	if x != nil {
		return x.IssuersPem
	}
	return ""
}

var File_cis_proto protoreflect.FileDescriptor

var file_cis_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x02, 0x0a,
	0x0f, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6b, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x1f, 0x0a, 0x05,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x05, 0x74, 0x72, 0x75, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x65, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x65, 0x6d, 0x22,
	0x3a, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x74, 0x73, 0x22, 0x90, 0x03, 0x0a, 0x0b,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x72, 0x67,
	0x5f, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6b, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6b, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x3a, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6e, 0x6f, 0x74,
	0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32,
	0x35, 0x36, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x65,
	0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0b,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x65, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x65, 0x6d, 0x2a, 0x29,
	0x0a, 0x05, 0x54, 0x72, 0x75, 0x73, 0x74, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x6e, 0x79, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x10, 0x02, 0x32, 0x56, 0x0a, 0x0f, 0x43, 0x65, 0x72,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x05,
	0x52, 0x6f, 0x6f, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6f,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x6f, 0x74,
	0x73, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6b, 0x73, 0x70, 0x61, 0x6e, 0x64, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cis_proto_rawDescOnce sync.Once
	file_cis_proto_rawDescData = file_cis_proto_rawDesc
)

func file_cis_proto_rawDescGZIP() []byte {
	file_cis_proto_rawDescOnce.Do(func() {
		file_cis_proto_rawDescData = protoimpl.X.CompressGZIP(file_cis_proto_rawDescData)
	})
	return file_cis_proto_rawDescData
}

var file_cis_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cis_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cis_proto_goTypes = []interface{}{
	(Trust)(0),                  // 0: pb.Trust
	(*RootCertificate)(nil),     // 1: pb.RootCertificate
	(*RootsResponse)(nil),       // 2: pb.RootsResponse
	(*Certificate)(nil),         // 3: pb.Certificate
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*EmptyRequest)(nil),        // 5: pb.EmptyRequest
}
var file_cis_proto_depIdxs = []int32{
	4, // 0: pb.RootCertificate.not_before:type_name -> google.protobuf.Timestamp
	4, // 1: pb.RootCertificate.not_after:type_name -> google.protobuf.Timestamp
	0, // 2: pb.RootCertificate.trust:type_name -> pb.Trust
	1, // 3: pb.RootsResponse.roots:type_name -> pb.RootCertificate
	4, // 4: pb.Certificate.not_before:type_name -> google.protobuf.Timestamp
	4, // 5: pb.Certificate.not_after:type_name -> google.protobuf.Timestamp
	5, // 6: pb.CertInfoService.Roots:input_type -> pb.EmptyRequest
	2, // 7: pb.CertInfoService.Roots:output_type -> pb.RootsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cis_proto_init() }
func file_cis_proto_init() {
	if File_cis_proto != nil {
		return
	}
	file_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RootCertificate); i {
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
		file_cis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RootsResponse); i {
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
		file_cis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
			RawDescriptor: file_cis_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cis_proto_goTypes,
		DependencyIndexes: file_cis_proto_depIdxs,
		EnumInfos:         file_cis_proto_enumTypes,
		MessageInfos:      file_cis_proto_msgTypes,
	}.Build()
	File_cis_proto = out.File
	file_cis_proto_rawDesc = nil
	file_cis_proto_goTypes = nil
	file_cis_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CertInfoServiceClient is the client API for CertInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertInfoServiceClient interface {
	// Roots returns the root CAs
	Roots(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RootsResponse, error)
}

type certInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertInfoServiceClient(cc grpc.ClientConnInterface) CertInfoServiceClient {
	return &certInfoServiceClient{cc}
}

func (c *certInfoServiceClient) Roots(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RootsResponse, error) {
	out := new(RootsResponse)
	err := c.cc.Invoke(ctx, "/pb.CertInfoService/Roots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertInfoServiceServer is the server API for CertInfoService service.
type CertInfoServiceServer interface {
	// Roots returns the root CAs
	Roots(context.Context, *EmptyRequest) (*RootsResponse, error)
}

// UnimplementedCertInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCertInfoServiceServer struct {
}

func (*UnimplementedCertInfoServiceServer) Roots(context.Context, *EmptyRequest) (*RootsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Roots not implemented")
}

func RegisterCertInfoServiceServer(s *grpc.Server, srv CertInfoServiceServer) {
	s.RegisterService(&_CertInfoService_serviceDesc, srv)
}

func _CertInfoService_Roots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertInfoServiceServer).Roots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CertInfoService/Roots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertInfoServiceServer).Roots(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CertInfoService",
	HandlerType: (*CertInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Roots",
			Handler:    _CertInfoService_Roots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cis.proto",
}