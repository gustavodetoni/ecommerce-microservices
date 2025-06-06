// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/pedidos.proto

package pedidos

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

type ItemPedido struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProdutoId     int64                  `protobuf:"varint,1,opt,name=produto_id,json=produtoId,proto3" json:"produto_id,omitempty"`
	Quantidade    int32                  `protobuf:"varint,2,opt,name=quantidade,proto3" json:"quantidade,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ItemPedido) Reset() {
	*x = ItemPedido{}
	mi := &file_proto_pedidos_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemPedido) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemPedido) ProtoMessage() {}

func (x *ItemPedido) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pedidos_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemPedido.ProtoReflect.Descriptor instead.
func (*ItemPedido) Descriptor() ([]byte, []int) {
	return file_proto_pedidos_proto_rawDescGZIP(), []int{0}
}

func (x *ItemPedido) GetProdutoId() int64 {
	if x != nil {
		return x.ProdutoId
	}
	return 0
}

func (x *ItemPedido) GetQuantidade() int32 {
	if x != nil {
		return x.Quantidade
	}
	return 0
}

type CriarPedidoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClienteId     int64                  `protobuf:"varint,1,opt,name=cliente_id,json=clienteId,proto3" json:"cliente_id,omitempty"`
	Itens         []*ItemPedido          `protobuf:"bytes,2,rep,name=itens,proto3" json:"itens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CriarPedidoRequest) Reset() {
	*x = CriarPedidoRequest{}
	mi := &file_proto_pedidos_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CriarPedidoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CriarPedidoRequest) ProtoMessage() {}

func (x *CriarPedidoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pedidos_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CriarPedidoRequest.ProtoReflect.Descriptor instead.
func (*CriarPedidoRequest) Descriptor() ([]byte, []int) {
	return file_proto_pedidos_proto_rawDescGZIP(), []int{1}
}

func (x *CriarPedidoRequest) GetClienteId() int64 {
	if x != nil {
		return x.ClienteId
	}
	return 0
}

func (x *CriarPedidoRequest) GetItens() []*ItemPedido {
	if x != nil {
		return x.Itens
	}
	return nil
}

type CriarPedidoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PedidoId      int64                  `protobuf:"varint,1,opt,name=pedido_id,json=pedidoId,proto3" json:"pedido_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CriarPedidoResponse) Reset() {
	*x = CriarPedidoResponse{}
	mi := &file_proto_pedidos_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CriarPedidoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CriarPedidoResponse) ProtoMessage() {}

func (x *CriarPedidoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pedidos_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CriarPedidoResponse.ProtoReflect.Descriptor instead.
func (*CriarPedidoResponse) Descriptor() ([]byte, []int) {
	return file_proto_pedidos_proto_rawDescGZIP(), []int{2}
}

func (x *CriarPedidoResponse) GetPedidoId() int64 {
	if x != nil {
		return x.PedidoId
	}
	return 0
}

func (x *CriarPedidoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_pedidos_proto protoreflect.FileDescriptor

const file_proto_pedidos_proto_rawDesc = "" +
	"\n" +
	"\x13proto/pedidos.proto\x12\apedidos\"K\n" +
	"\n" +
	"ItemPedido\x12\x1d\n" +
	"\n" +
	"produto_id\x18\x01 \x01(\x03R\tprodutoId\x12\x1e\n" +
	"\n" +
	"quantidade\x18\x02 \x01(\x05R\n" +
	"quantidade\"^\n" +
	"\x12CriarPedidoRequest\x12\x1d\n" +
	"\n" +
	"cliente_id\x18\x01 \x01(\x03R\tclienteId\x12)\n" +
	"\x05itens\x18\x02 \x03(\v2\x13.pedidos.ItemPedidoR\x05itens\"J\n" +
	"\x13CriarPedidoResponse\x12\x1b\n" +
	"\tpedido_id\x18\x01 \x01(\x03R\bpedidoId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status2Z\n" +
	"\x0ePedidosService\x12H\n" +
	"\vCriarPedido\x12\x1b.pedidos.CriarPedidoRequest\x1a\x1c.pedidos.CriarPedidoResponseB\x1eZ\x1cecommerce-grpc/proto/pedidosb\x06proto3"

var (
	file_proto_pedidos_proto_rawDescOnce sync.Once
	file_proto_pedidos_proto_rawDescData []byte
)

func file_proto_pedidos_proto_rawDescGZIP() []byte {
	file_proto_pedidos_proto_rawDescOnce.Do(func() {
		file_proto_pedidos_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_pedidos_proto_rawDesc), len(file_proto_pedidos_proto_rawDesc)))
	})
	return file_proto_pedidos_proto_rawDescData
}

var file_proto_pedidos_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_pedidos_proto_goTypes = []any{
	(*ItemPedido)(nil),          // 0: pedidos.ItemPedido
	(*CriarPedidoRequest)(nil),  // 1: pedidos.CriarPedidoRequest
	(*CriarPedidoResponse)(nil), // 2: pedidos.CriarPedidoResponse
}
var file_proto_pedidos_proto_depIdxs = []int32{
	0, // 0: pedidos.CriarPedidoRequest.itens:type_name -> pedidos.ItemPedido
	1, // 1: pedidos.PedidosService.CriarPedido:input_type -> pedidos.CriarPedidoRequest
	2, // 2: pedidos.PedidosService.CriarPedido:output_type -> pedidos.CriarPedidoResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_pedidos_proto_init() }
func file_proto_pedidos_proto_init() {
	if File_proto_pedidos_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_pedidos_proto_rawDesc), len(file_proto_pedidos_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pedidos_proto_goTypes,
		DependencyIndexes: file_proto_pedidos_proto_depIdxs,
		MessageInfos:      file_proto_pedidos_proto_msgTypes,
	}.Build()
	File_proto_pedidos_proto = out.File
	file_proto_pedidos_proto_goTypes = nil
	file_proto_pedidos_proto_depIdxs = nil
}
