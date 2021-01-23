// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_report.proto

package atlas

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ReportType int32

const (
	ReportType_REPORT_TYPE_RESERVED    ReportType = 0
	ReportType_REPORT_TYPE_UNSPECIFIED ReportType = 100
)

var ReportType_name = map[int32]string{
	0:   "REPORT_TYPE_RESERVED",
	100: "REPORT_TYPE_UNSPECIFIED",
}

var ReportType_value = map[string]int32{
	"REPORT_TYPE_RESERVED":    0,
	"REPORT_TYPE_UNSPECIFIED": 100,
}

func (x ReportType) String() string {
	return proto.EnumName(ReportType_name, int32(x))
}

func (ReportType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e80ce3fcf37ddf3, []int{0}
}

type Report struct {
	// Types that are valid to be assigned to MetadataOptional:
	//	*Report_Metadata
	MetadataOptional isReport_MetadataOptional `protobuf_oneof:"metadata_optional"`
	// Bytes is the purpose of the whole data chunk type
	// May contain any arbitrary sequence of bytes no longer than 2^32
	Bytes                []byte   `protobuf:"bytes,200,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e80ce3fcf37ddf3, []int{0}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

type isReport_MetadataOptional interface {
	isReport_MetadataOptional()
}

type Report_Metadata struct {
	Metadata *Metadata `protobuf:"bytes,100,opt,name=metadata,proto3,oneof"`
}

func (*Report_Metadata) isReport_MetadataOptional() {}

func (m *Report) GetMetadataOptional() isReport_MetadataOptional {
	if m != nil {
		return m.MetadataOptional
	}
	return nil
}

func (m *Report) GetMetadata() *Metadata {
	if x, ok := m.GetMetadataOptional().(*Report_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Report) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Report) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Report_Metadata)(nil),
	}
}

func init() {
	proto.RegisterEnum("atlas.ReportType", ReportType_name, ReportType_value)
	proto.RegisterType((*Report)(nil), "atlas.Report")
}

func init() { proto.RegisterFile("type_report.proto", fileDescriptor_1e80ce3fcf37ddf3) }

var fileDescriptor_1e80ce3fcf37ddf3 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xa9, 0x2c, 0x48,
	0x8d, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d,
	0x2c, 0xc9, 0x49, 0x2c, 0x96, 0x12, 0x06, 0xcb, 0xe4, 0xa6, 0x96, 0x24, 0xa6, 0x24, 0x96, 0x24,
	0x42, 0xe4, 0x94, 0x92, 0xb9, 0xd8, 0x82, 0xc0, 0x6a, 0x85, 0x74, 0xb9, 0x38, 0x60, 0x72, 0x12,
	0x29, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xfc, 0x7a, 0x60, 0x8d, 0x7a, 0xbe, 0x50, 0x61, 0x0f, 0x86,
	0x20, 0xb8, 0x12, 0x21, 0x51, 0x2e, 0xd6, 0xa4, 0xca, 0x92, 0xd4, 0x62, 0x89, 0x13, 0x8c, 0x0a,
	0x8c, 0x1a, 0x3c, 0x41, 0x10, 0x9e, 0x93, 0x30, 0x97, 0x20, 0x4c, 0x49, 0x7c, 0x7e, 0x41, 0x49,
	0x66, 0x7e, 0x5e, 0x62, 0x8e, 0x96, 0x33, 0x17, 0x17, 0xc4, 0x92, 0x90, 0xca, 0x82, 0x54, 0x21,
	0x09, 0x2e, 0x91, 0x20, 0xd7, 0x00, 0xff, 0xa0, 0x90, 0xf8, 0x90, 0xc8, 0x00, 0xd7, 0xf8, 0x20,
	0xd7, 0x60, 0xd7, 0xa0, 0x30, 0x57, 0x17, 0x01, 0x06, 0x21, 0x69, 0x2e, 0x71, 0x64, 0x99, 0x50,
	0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x81, 0x94, 0x24, 0x36, 0xb0, 0x83,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x06, 0x4b, 0x2e, 0xe1, 0xe1, 0x00, 0x00, 0x00,
}
