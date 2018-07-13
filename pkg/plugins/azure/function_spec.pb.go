// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: function_spec.proto

package azure

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Function Spec for Functions on Azure Functions Upstreams
type FunctionSpec struct {
	// The Name of the Azure Function as it appears in the Azure Functions Portal
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// Auth Level can bve either "anonymous" "function" or "admin"
	// See https://vincentlauzon.com/2017/12/04/azure-functions-http-authorization-levels/ for more details
	AuthLevel string `protobuf:"bytes,2,opt,name=auth_level,json=authLevel,proto3" json:"auth_level,omitempty"`
}

func (m *FunctionSpec) Reset()                    { *m = FunctionSpec{} }
func (m *FunctionSpec) String() string            { return proto.CompactTextString(m) }
func (*FunctionSpec) ProtoMessage()               {}
func (*FunctionSpec) Descriptor() ([]byte, []int) { return fileDescriptorFunctionSpec, []int{0} }

func (m *FunctionSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *FunctionSpec) GetAuthLevel() string {
	if m != nil {
		return m.AuthLevel
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionSpec)(nil), "gloo.api.v1.FunctionSpec")
}
func (this *FunctionSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionSpec)
	if !ok {
		that2, ok := that.(FunctionSpec)
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
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if this.AuthLevel != that1.AuthLevel {
		return false
	}
	return true
}

func init() { proto.RegisterFile("function_spec.proto", fileDescriptorFunctionSpec) }

var fileDescriptorFunctionSpec = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2b, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0x8b, 0x2f, 0x2e, 0x48, 0x4d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4e, 0xcf, 0xc9, 0xcf, 0xd7, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0x8b, 0xeb, 0x83, 0x58, 0x10, 0x25, 0x4a, 0x41, 0x5c, 0x3c, 0x6e, 0x50, 0x9d,
	0xc1, 0x05, 0xa9, 0xc9, 0x42, 0xca, 0x5c, 0xbc, 0x70, 0x93, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x78, 0x60, 0x82, 0x7e, 0x89, 0xb9, 0xa9, 0x42, 0xb2, 0x5c, 0x5c,
	0x89, 0xa5, 0x25, 0x19, 0xf1, 0x39, 0xa9, 0x65, 0xa9, 0x39, 0x12, 0x4c, 0x60, 0x15, 0x9c, 0x20,
	0x11, 0x1f, 0x90, 0x80, 0x93, 0xfe, 0x8a, 0x47, 0x72, 0x8c, 0x51, 0x9a, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc5, 0xf9, 0x39, 0xf9, 0xba, 0x99, 0xf9, 0xfa, 0x20,
	0xf7, 0xe8, 0x17, 0x64, 0xa7, 0xeb, 0x17, 0xe4, 0x94, 0xa6, 0x67, 0xe6, 0x15, 0xeb, 0x27, 0x56,
	0x95, 0x16, 0xa5, 0x26, 0xb1, 0x81, 0xdd, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x1d,
	0xda, 0x65, 0xc5, 0x00, 0x00, 0x00,
}
