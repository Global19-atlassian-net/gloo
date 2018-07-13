// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: upstream_spec.proto

/*
Package aws is a generated protocol buffer package.

It is generated from these files:
	upstream_spec.proto
	function_spec.proto

It has these top-level messages:
	UpstreamSpec
	FunctionSpec
*/
package aws

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Upstream Spec for AWS Lambda Upstreams
type UpstreamSpec struct {
	// The AWS Region in which to run Lambda Functions
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  ```
	SecretRef string `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
}

func (m *UpstreamSpec) Reset()                    { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string            { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()               {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) { return fileDescriptorUpstreamSpec, []int{0} }

func (m *UpstreamSpec) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() string {
	if m != nil {
		return m.SecretRef
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "gloo.api.v1.UpstreamSpec")
}
func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if this.SecretRef != that1.SecretRef {
		return false
	}
	return true
}

func init() { proto.RegisterFile("upstream_spec.proto", fileDescriptorUpstreamSpec) }

var fileDescriptorUpstreamSpec = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2d, 0x28, 0x2e,
	0x29, 0x4a, 0x4d, 0xcc, 0x8d, 0x2f, 0x2e, 0x48, 0x4d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4e, 0xcf, 0xc9, 0xcf, 0xd7, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0x8b, 0xeb, 0x83, 0x58, 0x10, 0x25, 0x4a, 0xae, 0x5c, 0x3c, 0xa1, 0x50, 0x9d,
	0xc1, 0x05, 0xa9, 0xc9, 0x42, 0x62, 0x5c, 0x6c, 0x45, 0xa9, 0xe9, 0x99, 0xf9, 0x79, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x90, 0x2c, 0x17, 0x57, 0x71, 0x6a, 0x72, 0x51, 0x6a,
	0x49, 0x7c, 0x51, 0x6a, 0x9a, 0x04, 0x13, 0x58, 0x8e, 0x13, 0x22, 0x12, 0x94, 0x9a, 0xe6, 0xa4,
	0xbb, 0xe2, 0x91, 0x1c, 0x63, 0x94, 0x7a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e,
	0xae, 0x7e, 0x71, 0x7e, 0x4e, 0xbe, 0x6e, 0x66, 0xbe, 0x3e, 0xc8, 0x09, 0xfa, 0x05, 0xd9, 0xe9,
	0xfa, 0x05, 0x39, 0xa5, 0xe9, 0x99, 0x79, 0xc5, 0xfa, 0x89, 0xe5, 0xc5, 0x49, 0x6c, 0x60, 0xcb,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x41, 0x7e, 0x31, 0xb6, 0x00, 0x00, 0x00,
}
