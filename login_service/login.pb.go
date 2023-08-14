package github.com/franklinwen0516/RPC/login_service

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ERR_CODE int32

const (
	ERR_CODE_CODE_NO_ERROR          ERR_CODE = 0
	ERR_CODE_CODE_ERR_MISSING_PARAM ERR_CODE = -1
	ERR_CODE_CODE_ERR_INVALID_PARAM ERR_CODE = -2
	ERR_CODE_CODE_ERR_UNKNOWN       ERR_CODE = -1001
)

// Enum value maps for ERR_CODE.
var (
	ERR_CODE_name = map[int32]string{
		0:     "CODE_NO_ERROR",
		-1:    "CODE_ERR_MISSING_PARAM",
		-2:    "CODE_ERR_INVALID_PARAM",
		-1001: "CODE_ERR_UNKNOWN",
	}
	ERR_CODE_value = map[string]int32{
		"CODE_NO_ERROR":          0,
		"CODE_ERR_MISSING_PARAM": -1,
		"CODE_ERR_INVALID_PARAM": -2,
		"CODE_ERR_UNKNOWN":       -1001,
	}
)

func (x ERR_CODE) Enum() *ERR_CODE {
	p := new(ERR_CODE)
	*p = x
	return p
}

func (x ERR_CODE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERR_CODE) Descriptor() protoreflect.EnumDescriptor {
	return file_login_proto_enumTypes[0].Descriptor()
}

func (ERR_CODE) Type() protoreflect.EnumType {
	return &file_login_proto_enumTypes[0]
}

func (x ERR_CODE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERR_CODE.Descriptor instead.
func (ERR_CODE) EnumDescriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

type BioKeyLoginRequset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountUuid      string `protobuf:"bytes,1,opt,name=account_uuid,json=accountUuid,proto3" json:"account_uuid,omitempty"`
	EncryptedBioinfo string `protobuf:"bytes,2,opt,name=encrypted_bioinfo,json=encryptedBioinfo,proto3" json:"encrypted_bioinfo,omitempty"`
	Timestamp        int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *BioKeyLoginRequset) Reset() {
	*x = BioKeyLoginRequset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BioKeyLoginRequset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BioKeyLoginRequset) ProtoMessage() {}

func (x *BioKeyLoginRequset) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BioKeyLoginRequset.ProtoReflect.Descriptor instead.
func (*BioKeyLoginRequset) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *BioKeyLoginRequset) GetAccountUuid() string {
	if x != nil {
		return x.AccountUuid
	}
	return ""
}

func (x *BioKeyLoginRequset) GetEncryptedBioinfo() string {
	if x != nil {
		return x.EncryptedBioinfo
	}
	return ""
}

func (x *BioKeyLoginRequset) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type BioKeyLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *CommonRspHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Amount int32            `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BioKeyLoginResponse) Reset() {
	*x = BioKeyLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BioKeyLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BioKeyLoginResponse) ProtoMessage() {}

func (x *BioKeyLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BioKeyLoginResponse.ProtoReflect.Descriptor instead.
func (*BioKeyLoginResponse) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *BioKeyLoginResponse) GetHeader() *CommonRspHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BioKeyLoginResponse) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CommonRspHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret    int32  `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *CommonRspHeader) Reset() {
	*x = CommonRspHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRspHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRspHeader) ProtoMessage() {}

func (x *CommonRspHeader) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRspHeader.ProtoReflect.Descriptor instead.
func (*CommonRspHeader) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

func (x *CommonRspHeader) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *CommonRspHeader) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01,
	0x0a, 0x12, 0x42, 0x69, 0x6f, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x73, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x69, 0x6f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x42, 0x69, 0x6f,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x57, 0x0a, 0x13, 0x42, 0x69, 0x6f, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x0f, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a, 0x86, 0x01, 0x0a, 0x08, 0x45, 0x52, 0x52,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x16, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52,
	0x41, 0x4d, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x23, 0x0a,
	0x16, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0x01, 0x12, 0x1d, 0x0a, 0x10, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x97, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x32, 0x4c, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x42, 0x69,
	0x6f, 0x4b, 0x65, 0x79, 0x12, 0x13, 0x2e, 0x42, 0x69, 0x6f, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x73, 0x65, 0x74, 0x1a, 0x14, 0x2e, 0x42, 0x69, 0x6f, 0x4b,
	0x65, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x16, 0x5a, 0x14, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_login_proto_goTypes = []interface{}{
	(ERR_CODE)(0),               // 0: ERR_CODE
	(*BioKeyLoginRequset)(nil),  // 1: BioKeyLoginRequset
	(*BioKeyLoginResponse)(nil), // 2: BioKeyLoginResponse
	(*CommonRspHeader)(nil),     // 3: CommonRspHeader
}
var file_login_proto_depIdxs = []int32{
	3, // 0: BioKeyLoginResponse.header:type_name -> CommonRspHeader
	1, // 1: LoginService.LoginWithBioKey:input_type -> BioKeyLoginRequset
	2, // 2: LoginService.LoginWithBioKey:output_type -> BioKeyLoginResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BioKeyLoginRequset); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BioKeyLoginResponse); i {
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
		file_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRspHeader); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		EnumInfos:         file_login_proto_enumTypes,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
