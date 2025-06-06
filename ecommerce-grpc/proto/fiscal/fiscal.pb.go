// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/fiscal.proto

package fiscal

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type EmitirNotaFiscalRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PedidoId      int64                  `protobuf:"varint,1,opt,name=pedido_id,json=pedidoId,proto3" json:"pedido_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmitirNotaFiscalRequest) Reset() {
	*x = EmitirNotaFiscalRequest{}
	mi := &file_proto_fiscal_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmitirNotaFiscalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmitirNotaFiscalRequest) ProtoMessage() {}

func (x *EmitirNotaFiscalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fiscal_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmitirNotaFiscalRequest.ProtoReflect.Descriptor instead.
func (*EmitirNotaFiscalRequest) Descriptor() ([]byte, []int) {
	return file_proto_fiscal_proto_rawDescGZIP(), []int{0}
}

func (x *EmitirNotaFiscalRequest) GetPedidoId() int64 {
	if x != nil {
		return x.PedidoId
	}
	return 0
}

type EmitirNotaFiscalResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	NotaFiscalId  int64                  `protobuf:"varint,2,opt,name=nota_fiscal_id,json=notaFiscalId,proto3" json:"nota_fiscal_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmitirNotaFiscalResponse) Reset() {
	*x = EmitirNotaFiscalResponse{}
	mi := &file_proto_fiscal_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmitirNotaFiscalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmitirNotaFiscalResponse) ProtoMessage() {}

func (x *EmitirNotaFiscalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fiscal_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmitirNotaFiscalResponse.ProtoReflect.Descriptor instead.
func (*EmitirNotaFiscalResponse) Descriptor() ([]byte, []int) {
	return file_proto_fiscal_proto_rawDescGZIP(), []int{1}
}

func (x *EmitirNotaFiscalResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EmitirNotaFiscalResponse) GetNotaFiscalId() int64 {
	if x != nil {
		return x.NotaFiscalId
	}
	return 0
}

var File_proto_fiscal_proto protoreflect.FileDescriptor

const file_proto_fiscal_proto_rawDesc = "" +
	"\n" +
	"\x12proto/fiscal.proto\x12\x06fiscal\"6\n" +
	"\x17EmitirNotaFiscalRequest\x12\x1b\n" +
	"\tpedido_id\x18\x01 \x01(\x03R\bpedidoId\"X\n" +
	"\x18EmitirNotaFiscalResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12$\n" +
	"\x0enota_fiscal_id\x18\x02 \x01(\x03R\fnotaFiscalId2f\n" +
	"\rFiscalService\x12U\n" +
	"\x10EmitirNotaFiscal\x12\x1f.fiscal.EmitirNotaFiscalRequest\x1a .fiscal.EmitirNotaFiscalResponseB\x1dZ\x1becommerce-grpc/proto/fiscalb\x06proto3"

var (
	file_proto_fiscal_proto_rawDescOnce sync.Once
	file_proto_fiscal_proto_rawDescData []byte
)

func file_proto_fiscal_proto_rawDescGZIP() []byte {
	file_proto_fiscal_proto_rawDescOnce.Do(func() {
		file_proto_fiscal_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_fiscal_proto_rawDesc), len(file_proto_fiscal_proto_rawDesc)))
	})
	return file_proto_fiscal_proto_rawDescData
}

var file_proto_fiscal_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_fiscal_proto_goTypes = []any{
	(*EmitirNotaFiscalRequest)(nil),  // 0: fiscal.EmitirNotaFiscalRequest
	(*EmitirNotaFiscalResponse)(nil), // 1: fiscal.EmitirNotaFiscalResponse
}
var file_proto_fiscal_proto_depIdxs = []int32{
	0, // 0: fiscal.FiscalService.EmitirNotaFiscal:input_type -> fiscal.EmitirNotaFiscalRequest
	1, // 1: fiscal.FiscalService.EmitirNotaFiscal:output_type -> fiscal.EmitirNotaFiscalResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_fiscal_proto_init() }
func file_proto_fiscal_proto_init() {
	if File_proto_fiscal_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_fiscal_proto_rawDesc), len(file_proto_fiscal_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fiscal_proto_goTypes,
		DependencyIndexes: file_proto_fiscal_proto_depIdxs,
		MessageInfos:      file_proto_fiscal_proto_msgTypes,
	}.Build()
	File_proto_fiscal_proto = out.File
	file_proto_fiscal_proto_goTypes = nil
	file_proto_fiscal_proto_depIdxs = nil
}
