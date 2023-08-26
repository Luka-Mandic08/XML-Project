// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: reservation_service/reservation_service.proto

package reservation

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

const (
	ReservationService_Get_FullMethodName                              = "/reservation.ReservationService/Get"
	ReservationService_Create_FullMethodName                           = "/reservation.ReservationService/Create"
	ReservationService_Update_FullMethodName                           = "/reservation.ReservationService/Update"
	ReservationService_Delete_FullMethodName                           = "/reservation.ReservationService/Delete"
	ReservationService_GetAllByUserId_FullMethodName                   = "/reservation.ReservationService/GetAllByUserId"
	ReservationService_Request_FullMethodName                          = "/reservation.ReservationService/Request"
	ReservationService_CheckIfGuestHasReservations_FullMethodName      = "/reservation.ReservationService/CheckIfGuestHasReservations"
	ReservationService_CheckIfHostHasReservations_FullMethodName       = "/reservation.ReservationService/CheckIfHostHasReservations"
	ReservationService_CheckIfGuestVisitedAccommodation_FullMethodName = "/reservation.ReservationService/CheckIfGuestVisitedAccommodation"
	ReservationService_CheckIfGuestVisitedHost_FullMethodName          = "/reservation.ReservationService/CheckIfGuestVisitedHost"
	ReservationService_Approve_FullMethodName                          = "/reservation.ReservationService/Approve"
	ReservationService_Deny_FullMethodName                             = "/reservation.ReservationService/Deny"
	ReservationService_Cancel_FullMethodName                           = "/reservation.ReservationService/Cancel"
	ReservationService_UpdateOutstandingHostStatus_FullMethodName      = "/reservation.ReservationService/UpdateOutstandingHostStatus"
	ReservationService_GetOutstandingHost_FullMethodName               = "/reservation.ReservationService/GetOutstandingHost"
	ReservationService_GetAllOutstandingHosts_FullMethodName           = "/reservation.ReservationService/GetAllOutstandingHosts"
	ReservationService_GetAllForDateRange_FullMethodName               = "/reservation.ReservationService/GetAllForDateRange"
	ReservationService_GetAllByAccommodationId_FullMethodName          = "/reservation.ReservationService/GetAllByAccommodationId"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetAllByUserId(ctx context.Context, in *GetAllByUserIdRequest, opts ...grpc.CallOption) (*GetAllByAccommodationIdResponse, error)
	Request(ctx context.Context, in *RequestRequest, opts ...grpc.CallOption) (*RequestResponse, error)
	CheckIfGuestHasReservations(ctx context.Context, in *CheckReservationRequest, opts ...grpc.CallOption) (*CheckReservationResponse, error)
	CheckIfHostHasReservations(ctx context.Context, in *CheckReservationRequest, opts ...grpc.CallOption) (*CheckReservationResponse, error)
	CheckIfGuestVisitedAccommodation(ctx context.Context, in *CheckPreviousReservationRequest, opts ...grpc.CallOption) (*CheckReservationResponse, error)
	CheckIfGuestVisitedHost(ctx context.Context, in *CheckPreviousReservationRequest, opts ...grpc.CallOption) (*CheckReservationResponse, error)
	Approve(ctx context.Context, in *ApproveRequest, opts ...grpc.CallOption) (*ApproveResponse, error)
	Deny(ctx context.Context, in *DenyRequest, opts ...grpc.CallOption) (*DenyResponse, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
	UpdateOutstandingHostStatus(ctx context.Context, in *UpdateOutstandingHostStatusRequest, opts ...grpc.CallOption) (*UpdateOutstandingHostStatusResponse, error)
	GetOutstandingHost(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetRequest, error)
	GetAllOutstandingHosts(ctx context.Context, in *GetAllOutstandingHostsRequest, opts ...grpc.CallOption) (*GetAllOutstandingHostsResponse, error)
	GetAllForDateRange(ctx context.Context, in *GetAllForDateRangeRequest, opts ...grpc.CallOption) (*GetAllForDateRangeResponse, error)
	GetAllByAccommodationId(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetAllByAccommodationIdResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, ReservationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, ReservationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, ReservationService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, ReservationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllByUserId(ctx context.Context, in *GetAllByUserIdRequest, opts ...grpc.CallOption) (*GetAllByAccommodationIdResponse, error) {
	out := new(GetAllByAccommodationIdResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAllByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Request(ctx context.Context, in *RequestRequest, opts ...grpc.CallOption) (*RequestResponse, error) {
	out := new(RequestResponse)
	err := c.cc.Invoke(ctx, ReservationService_Request_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckIfGuestHasReservations(ctx context.Context, in *CheckReservationRequest, opts ...grpc.CallOption) (*CheckReservationResponse, error) {
	out := new(CheckReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_CheckIfGuestHasReservations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckIfHostHasReservations(ctx context.Context, in *CheckReservationRequest, opts ...grpc.CallOption) (*CheckReservationResponse, error) {
	out := new(CheckReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_CheckIfHostHasReservations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckIfGuestVisitedAccommodation(ctx context.Context, in *CheckPreviousReservationRequest, opts ...grpc.CallOption) (*CheckReservationResponse, error) {
	out := new(CheckReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_CheckIfGuestVisitedAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckIfGuestVisitedHost(ctx context.Context, in *CheckPreviousReservationRequest, opts ...grpc.CallOption) (*CheckReservationResponse, error) {
	out := new(CheckReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_CheckIfGuestVisitedHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Approve(ctx context.Context, in *ApproveRequest, opts ...grpc.CallOption) (*ApproveResponse, error) {
	out := new(ApproveResponse)
	err := c.cc.Invoke(ctx, ReservationService_Approve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Deny(ctx context.Context, in *DenyRequest, opts ...grpc.CallOption) (*DenyResponse, error) {
	out := new(DenyResponse)
	err := c.cc.Invoke(ctx, ReservationService_Deny_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, ReservationService_Cancel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateOutstandingHostStatus(ctx context.Context, in *UpdateOutstandingHostStatusRequest, opts ...grpc.CallOption) (*UpdateOutstandingHostStatusResponse, error) {
	out := new(UpdateOutstandingHostStatusResponse)
	err := c.cc.Invoke(ctx, ReservationService_UpdateOutstandingHostStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetOutstandingHost(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetRequest, error) {
	out := new(GetRequest)
	err := c.cc.Invoke(ctx, ReservationService_GetOutstandingHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllOutstandingHosts(ctx context.Context, in *GetAllOutstandingHostsRequest, opts ...grpc.CallOption) (*GetAllOutstandingHostsResponse, error) {
	out := new(GetAllOutstandingHostsResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAllOutstandingHosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllForDateRange(ctx context.Context, in *GetAllForDateRangeRequest, opts ...grpc.CallOption) (*GetAllForDateRangeResponse, error) {
	out := new(GetAllForDateRangeResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAllForDateRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllByAccommodationId(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetAllByAccommodationIdResponse, error) {
	out := new(GetAllByAccommodationIdResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAllByAccommodationId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetAllByUserId(context.Context, *GetAllByUserIdRequest) (*GetAllByAccommodationIdResponse, error)
	Request(context.Context, *RequestRequest) (*RequestResponse, error)
	CheckIfGuestHasReservations(context.Context, *CheckReservationRequest) (*CheckReservationResponse, error)
	CheckIfHostHasReservations(context.Context, *CheckReservationRequest) (*CheckReservationResponse, error)
	CheckIfGuestVisitedAccommodation(context.Context, *CheckPreviousReservationRequest) (*CheckReservationResponse, error)
	CheckIfGuestVisitedHost(context.Context, *CheckPreviousReservationRequest) (*CheckReservationResponse, error)
	Approve(context.Context, *ApproveRequest) (*ApproveResponse, error)
	Deny(context.Context, *DenyRequest) (*DenyResponse, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
	UpdateOutstandingHostStatus(context.Context, *UpdateOutstandingHostStatusRequest) (*UpdateOutstandingHostStatusResponse, error)
	GetOutstandingHost(context.Context, *GetRequest) (*GetRequest, error)
	GetAllOutstandingHosts(context.Context, *GetAllOutstandingHostsRequest) (*GetAllOutstandingHostsResponse, error)
	GetAllForDateRange(context.Context, *GetAllForDateRangeRequest) (*GetAllForDateRangeResponse, error)
	GetAllByAccommodationId(context.Context, *GetRequest) (*GetAllByAccommodationIdResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedReservationServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedReservationServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedReservationServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedReservationServiceServer) GetAllByUserId(context.Context, *GetAllByUserIdRequest) (*GetAllByAccommodationIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByUserId not implemented")
}
func (UnimplementedReservationServiceServer) Request(context.Context, *RequestRequest) (*RequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedReservationServiceServer) CheckIfGuestHasReservations(context.Context, *CheckReservationRequest) (*CheckReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfGuestHasReservations not implemented")
}
func (UnimplementedReservationServiceServer) CheckIfHostHasReservations(context.Context, *CheckReservationRequest) (*CheckReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfHostHasReservations not implemented")
}
func (UnimplementedReservationServiceServer) CheckIfGuestVisitedAccommodation(context.Context, *CheckPreviousReservationRequest) (*CheckReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfGuestVisitedAccommodation not implemented")
}
func (UnimplementedReservationServiceServer) CheckIfGuestVisitedHost(context.Context, *CheckPreviousReservationRequest) (*CheckReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfGuestVisitedHost not implemented")
}
func (UnimplementedReservationServiceServer) Approve(context.Context, *ApproveRequest) (*ApproveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Approve not implemented")
}
func (UnimplementedReservationServiceServer) Deny(context.Context, *DenyRequest) (*DenyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deny not implemented")
}
func (UnimplementedReservationServiceServer) Cancel(context.Context, *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedReservationServiceServer) UpdateOutstandingHostStatus(context.Context, *UpdateOutstandingHostStatusRequest) (*UpdateOutstandingHostStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOutstandingHostStatus not implemented")
}
func (UnimplementedReservationServiceServer) GetOutstandingHost(context.Context, *GetRequest) (*GetRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOutstandingHost not implemented")
}
func (UnimplementedReservationServiceServer) GetAllOutstandingHosts(context.Context, *GetAllOutstandingHostsRequest) (*GetAllOutstandingHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOutstandingHosts not implemented")
}
func (UnimplementedReservationServiceServer) GetAllForDateRange(context.Context, *GetAllForDateRangeRequest) (*GetAllForDateRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllForDateRange not implemented")
}
func (UnimplementedReservationServiceServer) GetAllByAccommodationId(context.Context, *GetRequest) (*GetAllByAccommodationIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByAccommodationId not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAllByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllByUserId(ctx, req.(*GetAllByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Request_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Request(ctx, req.(*RequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckIfGuestHasReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckIfGuestHasReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CheckIfGuestHasReservations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckIfGuestHasReservations(ctx, req.(*CheckReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckIfHostHasReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckIfHostHasReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CheckIfHostHasReservations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckIfHostHasReservations(ctx, req.(*CheckReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckIfGuestVisitedAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPreviousReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckIfGuestVisitedAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CheckIfGuestVisitedAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckIfGuestVisitedAccommodation(ctx, req.(*CheckPreviousReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckIfGuestVisitedHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPreviousReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckIfGuestVisitedHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CheckIfGuestVisitedHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckIfGuestVisitedHost(ctx, req.(*CheckPreviousReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Approve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Approve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Approve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Approve(ctx, req.(*ApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Deny_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DenyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Deny(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Deny_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Deny(ctx, req.(*DenyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Cancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateOutstandingHostStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOutstandingHostStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateOutstandingHostStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_UpdateOutstandingHostStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateOutstandingHostStatus(ctx, req.(*UpdateOutstandingHostStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetOutstandingHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetOutstandingHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetOutstandingHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetOutstandingHost(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllOutstandingHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllOutstandingHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllOutstandingHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAllOutstandingHosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllOutstandingHosts(ctx, req.(*GetAllOutstandingHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllForDateRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllForDateRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllForDateRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAllForDateRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllForDateRange(ctx, req.(*GetAllForDateRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAllByAccommodationId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllByAccommodationId(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ReservationService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ReservationService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReservationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReservationService_Delete_Handler,
		},
		{
			MethodName: "GetAllByUserId",
			Handler:    _ReservationService_GetAllByUserId_Handler,
		},
		{
			MethodName: "Request",
			Handler:    _ReservationService_Request_Handler,
		},
		{
			MethodName: "CheckIfGuestHasReservations",
			Handler:    _ReservationService_CheckIfGuestHasReservations_Handler,
		},
		{
			MethodName: "CheckIfHostHasReservations",
			Handler:    _ReservationService_CheckIfHostHasReservations_Handler,
		},
		{
			MethodName: "CheckIfGuestVisitedAccommodation",
			Handler:    _ReservationService_CheckIfGuestVisitedAccommodation_Handler,
		},
		{
			MethodName: "CheckIfGuestVisitedHost",
			Handler:    _ReservationService_CheckIfGuestVisitedHost_Handler,
		},
		{
			MethodName: "Approve",
			Handler:    _ReservationService_Approve_Handler,
		},
		{
			MethodName: "Deny",
			Handler:    _ReservationService_Deny_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _ReservationService_Cancel_Handler,
		},
		{
			MethodName: "UpdateOutstandingHostStatus",
			Handler:    _ReservationService_UpdateOutstandingHostStatus_Handler,
		},
		{
			MethodName: "GetOutstandingHost",
			Handler:    _ReservationService_GetOutstandingHost_Handler,
		},
		{
			MethodName: "GetAllOutstandingHosts",
			Handler:    _ReservationService_GetAllOutstandingHosts_Handler,
		},
		{
			MethodName: "GetAllForDateRange",
			Handler:    _ReservationService_GetAllForDateRange_Handler,
		},
		{
			MethodName: "GetAllByAccommodationId",
			Handler:    _ReservationService_GetAllByAccommodationId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation_service/reservation_service.proto",
}
