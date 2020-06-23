// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/waf/waf.proto

package waf

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	waf "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/waf"
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

type Settings struct {
	// Disable waf on this resource (if omitted defaults to false).
	// If a route/virtual host is configured with WAF, you must explicitly disable its WAF,
	// i.e., it will not inherit the disabled status of its parent
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Custom massage to display if an intervention occurs.
	CustomInterventionMessage string `protobuf:"bytes,2,opt,name=custom_intervention_message,json=customInterventionMessage,proto3" json:"custom_intervention_message,omitempty"`
	// Add OWASP core rule set
	// if nil will not be added
	CoreRuleSet *CoreRuleSet `protobuf:"bytes,3,opt,name=core_rule_set,json=coreRuleSet,proto3" json:"core_rule_set,omitempty"`
	// Custom rule sets rules to add
	RuleSets []*waf.RuleSet `protobuf:"bytes,4,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	// Audit Log settings
	AuditLogging         *waf.AuditLogging `protobuf:"bytes,5,opt,name=audit_logging,json=auditLogging,proto3" json:"audit_logging,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_0151c80aefddd633, []int{0}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Settings) GetCustomInterventionMessage() string {
	if m != nil {
		return m.CustomInterventionMessage
	}
	return ""
}

func (m *Settings) GetCoreRuleSet() *CoreRuleSet {
	if m != nil {
		return m.CoreRuleSet
	}
	return nil
}

func (m *Settings) GetRuleSets() []*waf.RuleSet {
	if m != nil {
		return m.RuleSets
	}
	return nil
}

func (m *Settings) GetAuditLogging() *waf.AuditLogging {
	if m != nil {
		return m.AuditLogging
	}
	return nil
}

type CoreRuleSet struct {
	// Optional custom settings for the OWASP core rule set.
	// For an example on the configuration options see: https://github.com/SpiderLabs/owasp-modsecurity-crs/blob/v3.2/dev/crs-setup.conf.example
	// The same rules apply to these options as do to the `RuleSet`s. The file option is better if possible.
	//
	// Types that are valid to be assigned to CustomSettingsType:
	//	*CoreRuleSet_CustomSettingsString
	//	*CoreRuleSet_CustomSettingsFile
	CustomSettingsType   isCoreRuleSet_CustomSettingsType `protobuf_oneof:"CustomSettingsType"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CoreRuleSet) Reset()         { *m = CoreRuleSet{} }
func (m *CoreRuleSet) String() string { return proto.CompactTextString(m) }
func (*CoreRuleSet) ProtoMessage()    {}
func (*CoreRuleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0151c80aefddd633, []int{1}
}
func (m *CoreRuleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreRuleSet.Unmarshal(m, b)
}
func (m *CoreRuleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreRuleSet.Marshal(b, m, deterministic)
}
func (m *CoreRuleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreRuleSet.Merge(m, src)
}
func (m *CoreRuleSet) XXX_Size() int {
	return xxx_messageInfo_CoreRuleSet.Size(m)
}
func (m *CoreRuleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreRuleSet.DiscardUnknown(m)
}

var xxx_messageInfo_CoreRuleSet proto.InternalMessageInfo

type isCoreRuleSet_CustomSettingsType interface {
	isCoreRuleSet_CustomSettingsType()
	Equal(interface{}) bool
}

type CoreRuleSet_CustomSettingsString struct {
	CustomSettingsString string `protobuf:"bytes,2,opt,name=custom_settings_string,json=customSettingsString,proto3,oneof" json:"custom_settings_string,omitempty"`
}
type CoreRuleSet_CustomSettingsFile struct {
	CustomSettingsFile string `protobuf:"bytes,3,opt,name=custom_settings_file,json=customSettingsFile,proto3,oneof" json:"custom_settings_file,omitempty"`
}

func (*CoreRuleSet_CustomSettingsString) isCoreRuleSet_CustomSettingsType() {}
func (*CoreRuleSet_CustomSettingsFile) isCoreRuleSet_CustomSettingsType()   {}

func (m *CoreRuleSet) GetCustomSettingsType() isCoreRuleSet_CustomSettingsType {
	if m != nil {
		return m.CustomSettingsType
	}
	return nil
}

func (m *CoreRuleSet) GetCustomSettingsString() string {
	if x, ok := m.GetCustomSettingsType().(*CoreRuleSet_CustomSettingsString); ok {
		return x.CustomSettingsString
	}
	return ""
}

func (m *CoreRuleSet) GetCustomSettingsFile() string {
	if x, ok := m.GetCustomSettingsType().(*CoreRuleSet_CustomSettingsFile); ok {
		return x.CustomSettingsFile
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CoreRuleSet) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CoreRuleSet_CustomSettingsString)(nil),
		(*CoreRuleSet_CustomSettingsFile)(nil),
	}
}

func init() {
	proto.RegisterType((*Settings)(nil), "waf.options.gloo.solo.io.Settings")
	proto.RegisterType((*CoreRuleSet)(nil), "waf.options.gloo.solo.io.CoreRuleSet")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/waf/waf.proto", fileDescriptor_0151c80aefddd633)
}

var fileDescriptor_0151c80aefddd633 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x65, 0xdb, 0x82, 0x52, 0x87, 0x5e, 0xac, 0x08, 0x2d, 0x41, 0x42, 0x51, 0x25, 0xa4, 0x5c,
	0xb0, 0x21, 0x08, 0x8e, 0x95, 0x68, 0x25, 0xa0, 0x12, 0x39, 0xb0, 0xe1, 0xd4, 0xcb, 0x6a, 0xb3,
	0x99, 0xb8, 0x06, 0xc7, 0xb3, 0xf2, 0xcc, 0x86, 0xe4, 0x2b, 0xf8, 0x0d, 0x3e, 0x81, 0x8f, 0xe1,
	0xc4, 0x3f, 0x70, 0x47, 0xde, 0x4d, 0x20, 0x05, 0xb5, 0xea, 0x61, 0xa5, 0x79, 0x6f, 0xe6, 0xbd,
	0xb5, 0x9f, 0x47, 0x8c, 0x8d, 0xe5, 0xcb, 0x7a, 0xaa, 0x4a, 0x5c, 0x68, 0x42, 0x87, 0x4f, 0x2d,
	0x6a, 0xe3, 0x10, 0x75, 0x15, 0xf0, 0x13, 0x94, 0x4c, 0x2d, 0x2a, 0x2a, 0xab, 0x97, 0xcf, 0x35,
	0x78, 0x86, 0x50, 0x05, 0x4b, 0xa0, 0xb1, 0x62, 0x8b, 0x9e, 0xf4, 0x97, 0x62, 0x1e, 0x3f, 0x55,
	0x05, 0x64, 0x94, 0x69, 0x2c, 0x37, 0x2d, 0x15, 0x95, 0x2a, 0x9a, 0x2a, 0x8b, 0xfd, 0x93, 0x6b,
	0x5c, 0x61, 0xc5, 0x10, 0x7c, 0xe1, 0x34, 0xf8, 0x25, 0xae, 0x1b, 0xe8, 0xe9, 0x7f, 0xe7, 0x7e,
	0xcf, 0xa0, 0xc1, 0xa6, 0xd4, 0xb1, 0xda, 0xb0, 0x12, 0x56, 0xdc, 0x92, 0xb0, 0xe2, 0x96, 0x3b,
	0xfe, 0xb1, 0x27, 0x3a, 0x13, 0x60, 0xb6, 0xde, 0x90, 0xec, 0x8b, 0xce, 0xcc, 0x52, 0x31, 0x75,
	0x30, 0x4b, 0x93, 0x41, 0x32, 0xec, 0x64, 0x7f, 0xb0, 0x3c, 0x11, 0x8f, 0xca, 0x9a, 0x18, 0x17,
	0xb9, 0x8d, 0xf7, 0x5a, 0x82, 0x8f, 0xe7, 0xce, 0x17, 0x40, 0x54, 0x18, 0x48, 0xf7, 0x06, 0xc9,
	0xf0, 0x30, 0x7b, 0xd8, 0x8e, 0x9c, 0xef, 0x4c, 0x8c, 0xdb, 0x01, 0x79, 0x2e, 0x8e, 0x4a, 0x0c,
	0x90, 0x87, 0xda, 0x41, 0x4e, 0xc0, 0xe9, 0xfe, 0x20, 0x19, 0x76, 0x47, 0x4f, 0xd4, 0x75, 0x21,
	0xa8, 0x33, 0x0c, 0x90, 0xd5, 0x0e, 0x26, 0xc0, 0x59, 0xb7, 0xfc, 0x0b, 0xe4, 0x58, 0x1c, 0x6e,
	0x5d, 0x28, 0x3d, 0x18, 0xec, 0x0f, 0xbb, 0xa3, 0x67, 0xaa, 0x49, 0x44, 0x95, 0xe8, 0xe7, 0xd6,
	0xa8, 0xb9, 0x75, 0x0c, 0x41, 0x5d, 0x32, 0x57, 0x6a, 0x81, 0x33, 0x82, 0xb2, 0x0e, 0x96, 0xd7,
	0x6a, 0x39, 0x52, 0x5b, 0xc7, 0x4e, 0x68, 0x0b, 0x92, 0x17, 0xe2, 0xa8, 0xa8, 0x67, 0x96, 0x73,
	0x87, 0xc6, 0x58, 0x6f, 0xd2, 0xbb, 0xcd, 0xc9, 0x5e, 0xde, 0xda, 0xf2, 0x75, 0x54, 0xbf, 0x6f,
	0xc5, 0xd9, 0xfd, 0x62, 0x07, 0x1d, 0x7f, 0x4d, 0x44, 0x77, 0xe7, 0x1e, 0xf2, 0x95, 0x78, 0xb0,
	0x49, 0x91, 0x36, 0xa1, 0xe7, 0xc4, 0x21, 0xfe, 0xb4, 0x09, 0xf0, 0xdd, 0x9d, 0xac, 0xd7, 0xf6,
	0xb7, 0x6f, 0x32, 0x69, 0xba, 0x72, 0x24, 0x7a, 0xff, 0xea, 0xe6, 0xd6, 0x41, 0x13, 0x62, 0x54,
	0xc9, 0xab, 0xaa, 0x37, 0xd6, 0xc1, 0x69, 0x4f, 0xc8, 0xb3, 0x2b, 0xec, 0xc7, 0x75, 0x05, 0xa7,
	0x1f, 0xbe, 0xff, 0x3a, 0x48, 0xbe, 0xfd, 0x7c, 0x9c, 0x5c, 0xbc, 0xbd, 0xdd, 0x36, 0x57, 0x9f,
	0xcd, 0xcd, 0x1b, 0x3d, 0xbd, 0xd7, 0xac, 0xd2, 0x8b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39,
	0xa4, 0x2e, 0x7c, 0x1f, 0x03, 0x00, 0x00,
}

func (this *Settings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
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
	if this.Disabled != that1.Disabled {
		return false
	}
	if this.CustomInterventionMessage != that1.CustomInterventionMessage {
		return false
	}
	if !this.CoreRuleSet.Equal(that1.CoreRuleSet) {
		return false
	}
	if len(this.RuleSets) != len(that1.RuleSets) {
		return false
	}
	for i := range this.RuleSets {
		if !this.RuleSets[i].Equal(that1.RuleSets[i]) {
			return false
		}
	}
	if !this.AuditLogging.Equal(that1.AuditLogging) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CoreRuleSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet)
	if !ok {
		that2, ok := that.(CoreRuleSet)
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
	if that1.CustomSettingsType == nil {
		if this.CustomSettingsType != nil {
			return false
		}
	} else if this.CustomSettingsType == nil {
		return false
	} else if !this.CustomSettingsType.Equal(that1.CustomSettingsType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CoreRuleSet_CustomSettingsString) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet_CustomSettingsString)
	if !ok {
		that2, ok := that.(CoreRuleSet_CustomSettingsString)
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
	if this.CustomSettingsString != that1.CustomSettingsString {
		return false
	}
	return true
}
func (this *CoreRuleSet_CustomSettingsFile) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet_CustomSettingsFile)
	if !ok {
		that2, ok := that.(CoreRuleSet_CustomSettingsFile)
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
	if this.CustomSettingsFile != that1.CustomSettingsFile {
		return false
	}
	return true
}
