// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoServiceClient interface {
	// Creates a new account by submitting the transaction
	CreateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Updates an account by submitting the transaction
	UpdateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Initiates a transfer by submitting the transaction
	CryptoTransfer(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Deletes and account by submitting the transaction
	CryptoDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Adds a livehash
	AddLiveHash(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Deletes a livehash
	DeleteLiveHash(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Retrieves a livehash for an account
	GetLiveHash(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Returns all transactions in the last 180s of consensus time for which the given account was the effective payer <b>and</b> network property <tt>ledger.keepRecordsInState</tt> was <tt>true</tt>.
	GetAccountRecords(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves the balance of an account
	CryptoGetBalance(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves the metadata of an account
	GetAccountInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves the latest receipt for a transaction that is either awaiting consensus, or reached consensus in the last 180 seconds
	GetTransactionReceipts(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// (NOT CURRENTLY SUPPORTED) Returns the records of transactions recently funded by an account
	GetFastTransactionRecord(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Retrieves the record of a transaction that is either awaiting consensus, or reached consensus in the last 180 seconds
	GetTxRecordByTxID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// (NOT CURRENTLY SUPPORTED) Retrieves the stakers for a node by account id
	GetStakersByAccountID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) CreateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/createAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) UpdateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/updateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CryptoTransfer(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/cryptoTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CryptoDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/cryptoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) AddLiveHash(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/addLiveHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DeleteLiveHash(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/deleteLiveHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetLiveHash(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getLiveHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetAccountRecords(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getAccountRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CryptoGetBalance(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/cryptoGetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetAccountInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetTransactionReceipts(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getTransactionReceipts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetFastTransactionRecord(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getFastTransactionRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetTxRecordByTxID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getTxRecordByTxID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetStakersByAccountID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getStakersByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
// All implementations must embed UnimplementedCryptoServiceServer
// for forward compatibility
type CryptoServiceServer interface {
	// Creates a new account by submitting the transaction
	CreateAccount(context.Context, *Transaction) (*TransactionResponse, error)
	// Updates an account by submitting the transaction
	UpdateAccount(context.Context, *Transaction) (*TransactionResponse, error)
	// Initiates a transfer by submitting the transaction
	CryptoTransfer(context.Context, *Transaction) (*TransactionResponse, error)
	// Deletes and account by submitting the transaction
	CryptoDelete(context.Context, *Transaction) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Adds a livehash
	AddLiveHash(context.Context, *Transaction) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Deletes a livehash
	DeleteLiveHash(context.Context, *Transaction) (*TransactionResponse, error)
	// (NOT CURRENTLY SUPPORTED) Retrieves a livehash for an account
	GetLiveHash(context.Context, *Query) (*Response, error)
	// Returns all transactions in the last 180s of consensus time for which the given account was the effective payer <b>and</b> network property <tt>ledger.keepRecordsInState</tt> was <tt>true</tt>.
	GetAccountRecords(context.Context, *Query) (*Response, error)
	// Retrieves the balance of an account
	CryptoGetBalance(context.Context, *Query) (*Response, error)
	// Retrieves the metadata of an account
	GetAccountInfo(context.Context, *Query) (*Response, error)
	// Retrieves the latest receipt for a transaction that is either awaiting consensus, or reached consensus in the last 180 seconds
	GetTransactionReceipts(context.Context, *Query) (*Response, error)
	// (NOT CURRENTLY SUPPORTED) Returns the records of transactions recently funded by an account
	GetFastTransactionRecord(context.Context, *Query) (*Response, error)
	// Retrieves the record of a transaction that is either awaiting consensus, or reached consensus in the last 180 seconds
	GetTxRecordByTxID(context.Context, *Query) (*Response, error)
	// (NOT CURRENTLY SUPPORTED) Retrieves the stakers for a node by account id
	GetStakersByAccountID(context.Context, *Query) (*Response, error)
	mustEmbedUnimplementedCryptoServiceServer()
}

// UnimplementedCryptoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (UnimplementedCryptoServiceServer) CreateAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedCryptoServiceServer) UpdateAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedCryptoServiceServer) CryptoTransfer(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoTransfer not implemented")
}
func (UnimplementedCryptoServiceServer) CryptoDelete(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoDelete not implemented")
}
func (UnimplementedCryptoServiceServer) AddLiveHash(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLiveHash not implemented")
}
func (UnimplementedCryptoServiceServer) DeleteLiveHash(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLiveHash not implemented")
}
func (UnimplementedCryptoServiceServer) GetLiveHash(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiveHash not implemented")
}
func (UnimplementedCryptoServiceServer) GetAccountRecords(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountRecords not implemented")
}
func (UnimplementedCryptoServiceServer) CryptoGetBalance(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoGetBalance not implemented")
}
func (UnimplementedCryptoServiceServer) GetAccountInfo(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
func (UnimplementedCryptoServiceServer) GetTransactionReceipts(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionReceipts not implemented")
}
func (UnimplementedCryptoServiceServer) GetFastTransactionRecord(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFastTransactionRecord not implemented")
}
func (UnimplementedCryptoServiceServer) GetTxRecordByTxID(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxRecordByTxID not implemented")
}
func (UnimplementedCryptoServiceServer) GetStakersByAccountID(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStakersByAccountID not implemented")
}
func (UnimplementedCryptoServiceServer) mustEmbedUnimplementedCryptoServiceServer() {}

// UnsafeCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoServiceServer will
// result in compilation errors.
type UnsafeCryptoServiceServer interface {
	mustEmbedUnimplementedCryptoServiceServer()
}

func RegisterCryptoServiceServer(s grpc.ServiceRegistrar, srv CryptoServiceServer) {
	s.RegisterService(&CryptoService_ServiceDesc, srv)
}

func _CryptoService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/createAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CreateAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/updateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).UpdateAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CryptoTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CryptoTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/cryptoTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CryptoTransfer(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CryptoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CryptoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/cryptoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CryptoDelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_AddLiveHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).AddLiveHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/addLiveHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).AddLiveHash(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DeleteLiveHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).DeleteLiveHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/deleteLiveHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).DeleteLiveHash(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetLiveHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetLiveHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/getLiveHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetLiveHash(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetAccountRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetAccountRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/getAccountRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetAccountRecords(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CryptoGetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CryptoGetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/cryptoGetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CryptoGetBalance(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/getAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetAccountInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetTransactionReceipts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetTransactionReceipts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/getTransactionReceipts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetTransactionReceipts(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetFastTransactionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetFastTransactionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/getFastTransactionRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetFastTransactionRecord(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetTxRecordByTxID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetTxRecordByTxID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/getTxRecordByTxID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetTxRecordByTxID(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetStakersByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetStakersByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/getStakersByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetStakersByAccountID(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptoService_ServiceDesc is the grpc.ServiceDesc for CryptoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createAccount",
			Handler:    _CryptoService_CreateAccount_Handler,
		},
		{
			MethodName: "updateAccount",
			Handler:    _CryptoService_UpdateAccount_Handler,
		},
		{
			MethodName: "cryptoTransfer",
			Handler:    _CryptoService_CryptoTransfer_Handler,
		},
		{
			MethodName: "cryptoDelete",
			Handler:    _CryptoService_CryptoDelete_Handler,
		},
		{
			MethodName: "addLiveHash",
			Handler:    _CryptoService_AddLiveHash_Handler,
		},
		{
			MethodName: "deleteLiveHash",
			Handler:    _CryptoService_DeleteLiveHash_Handler,
		},
		{
			MethodName: "getLiveHash",
			Handler:    _CryptoService_GetLiveHash_Handler,
		},
		{
			MethodName: "getAccountRecords",
			Handler:    _CryptoService_GetAccountRecords_Handler,
		},
		{
			MethodName: "cryptoGetBalance",
			Handler:    _CryptoService_CryptoGetBalance_Handler,
		},
		{
			MethodName: "getAccountInfo",
			Handler:    _CryptoService_GetAccountInfo_Handler,
		},
		{
			MethodName: "getTransactionReceipts",
			Handler:    _CryptoService_GetTransactionReceipts_Handler,
		},
		{
			MethodName: "getFastTransactionRecord",
			Handler:    _CryptoService_GetFastTransactionRecord_Handler,
		},
		{
			MethodName: "getTxRecordByTxID",
			Handler:    _CryptoService_GetTxRecordByTxID_Handler,
		},
		{
			MethodName: "getStakersByAccountID",
			Handler:    _CryptoService_GetStakersByAccountID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/crypto_service.proto",
}
