// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/transformation/transformation.proto

package transformation

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Transformations struct {
	// Apply a transformation to requests.
	RequestTransformation *transformation.Transformation `protobuf:"bytes,1,opt,name=request_transformation,json=requestTransformation,proto3" json:"request_transformation,omitempty"`
	// Clear the route cache if the request transformation was applied.
	ClearRouteCache bool `protobuf:"varint,3,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	// Apply a transformation to responses.
	ResponseTransformation *transformation.Transformation `protobuf:"bytes,2,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                       `json:"-"`
	XXX_unrecognized       []byte                         `json:"-"`
	XXX_sizecache          int32                          `json:"-"`
}

func (m *Transformations) Reset()         { *m = Transformations{} }
func (m *Transformations) String() string { return proto.CompactTextString(m) }
func (*Transformations) ProtoMessage()    {}
func (*Transformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8def6c8e0694580, []int{0}
}
func (m *Transformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transformations.Unmarshal(m, b)
}
func (m *Transformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transformations.Marshal(b, m, deterministic)
}
func (m *Transformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transformations.Merge(m, src)
}
func (m *Transformations) XXX_Size() int {
	return xxx_messageInfo_Transformations.Size(m)
}
func (m *Transformations) XXX_DiscardUnknown() {
	xxx_messageInfo_Transformations.DiscardUnknown(m)
}

var xxx_messageInfo_Transformations proto.InternalMessageInfo

func (m *Transformations) GetRequestTransformation() *transformation.Transformation {
	if m != nil {
		return m.RequestTransformation
	}
	return nil
}

func (m *Transformations) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *Transformations) GetResponseTransformation() *transformation.Transformation {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

func init() {
	proto.RegisterType((*Transformations)(nil), "transformation.options.gloo.solo.io.Transformations")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/transformation/transformation.proto", fileDescriptor_d8def6c8e0694580)
}

var fileDescriptor_d8def6c8e0694580 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0xb9, 0x28, 0x22, 0x67, 0x11, 0x3c, 0x34, 0x86, 0x14, 0x21, 0x68, 0x13, 0x04, 0x77,
	0x31, 0xbe, 0x81, 0x16, 0xd6, 0x86, 0x14, 0x62, 0x73, 0x6c, 0x8e, 0xc9, 0x66, 0xf5, 0xb2, 0xb3,
	0xce, 0xce, 0xc5, 0x58, 0xfa, 0x36, 0x3e, 0x82, 0xcf, 0xe3, 0x3b, 0xd8, 0xcb, 0xee, 0xc5, 0xe2,
	0x02, 0x4a, 0xec, 0xe6, 0xff, 0x8e, 0xf9, 0xee, 0x67, 0x99, 0xf4, 0x5e, 0x1b, 0x9e, 0x57, 0x53,
	0x51, 0xe0, 0x42, 0x7a, 0x2c, 0xf1, 0xc2, 0xa0, 0xd4, 0x25, 0xa2, 0x74, 0x84, 0x8f, 0x50, 0xb0,
	0xaf, 0x93, 0x72, 0x46, 0x2e, 0x2f, 0x25, 0x3a, 0x36, 0x68, 0xbd, 0x64, 0x52, 0xd6, 0xcf, 0x90,
	0x16, 0x2a, 0xe4, 0x8d, 0x28, 0x1c, 0x21, 0x63, 0x76, 0xb6, 0x41, 0xd7, 0xbb, 0x22, 0xf8, 0x44,
	0xf8, 0x95, 0x30, 0xd8, 0xeb, 0x6b, 0x44, 0x5d, 0x82, 0x8c, 0x2b, 0xd3, 0x6a, 0x26, 0x5f, 0x48,
	0x39, 0x07, 0xe4, 0x6b, 0x49, 0x6f, 0xf2, 0x4b, 0x17, 0x58, 0x31, 0x90, 0x55, 0xa5, 0x04, 0xbb,
	0xc4, 0xd7, 0x18, 0xad, 0xdf, 0xb6, 0x5a, 0xef, 0x48, 0xa3, 0xc6, 0x38, 0xca, 0x30, 0xad, 0x69,
	0x06, 0x2b, 0xae, 0x21, 0xac, 0xb8, 0x66, 0xa7, 0x6f, 0xad, 0xb4, 0x3d, 0x69, 0x28, 0x7c, 0x96,
	0xa7, 0x1d, 0x82, 0xe7, 0x0a, 0x3c, 0xe7, 0x4d, 0x7b, 0x37, 0x19, 0x24, 0xc3, 0x83, 0xd1, 0x50,
	0xc4, 0x52, 0x42, 0x39, 0x23, 0x96, 0x23, 0x31, 0x33, 0x25, 0x03, 0x89, 0x39, 0xb3, 0x13, 0x4d,
	0xd5, 0xf8, 0x78, 0xed, 0x69, 0xe2, 0xec, 0x3c, 0x3d, 0x2c, 0x4a, 0x50, 0x94, 0x13, 0x56, 0x0c,
	0x79, 0xa1, 0x8a, 0x39, 0x74, 0x77, 0x06, 0xc9, 0x70, 0x7f, 0xdc, 0x8e, 0x1f, 0xc6, 0x81, 0xdf,
	0x04, 0x9c, 0xa9, 0xf4, 0x84, 0xc0, 0x3b, 0xb4, 0x1e, 0x36, 0xdb, 0xb4, 0xfe, 0xd9, 0xa6, 0xf3,
	0x23, 0x6a, 0xf2, 0xeb, 0xbb, 0x8f, 0xaf, 0xdd, 0xe4, 0xfd, 0xb3, 0x9f, 0x3c, 0xdc, 0x6e, 0x77,
	0x2c, 0xee, 0x49, 0xff, 0x7d, 0x30, 0xd3, 0xbd, 0xf8, 0xba, 0x57, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x03, 0x67, 0x9b, 0x59, 0x7e, 0x02, 0x00, 0x00,
}

func (this *Transformations) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transformations)
	if !ok {
		that2, ok := that.(Transformations)
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
	if !this.RequestTransformation.Equal(that1.RequestTransformation) {
		return false
	}
	if this.ClearRouteCache != that1.ClearRouteCache {
		return false
	}
	if !this.ResponseTransformation.Equal(that1.ResponseTransformation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}