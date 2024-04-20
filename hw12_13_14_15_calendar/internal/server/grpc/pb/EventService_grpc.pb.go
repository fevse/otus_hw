// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: EventService.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Calendar_Create_FullMethodName         = "/event.Calendar/Create"
	Calendar_Update_FullMethodName         = "/event.Calendar/Update"
	Calendar_Delete_FullMethodName         = "/event.Calendar/Delete"
	Calendar_Show_FullMethodName           = "/event.Calendar/Show"
	Calendar_ShowEventDay_FullMethodName   = "/event.Calendar/ShowEventDay"
	Calendar_ShowEventWeek_FullMethodName  = "/event.Calendar/ShowEventWeek"
	Calendar_ShowEventMonth_FullMethodName = "/event.Calendar/ShowEventMonth"
)

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarClient interface {
	Create(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *IDEvent, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Show(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Events, error)
	ShowEventDay(ctx context.Context, in *Date, opts ...grpc.CallOption) (*Events, error)
	ShowEventWeek(ctx context.Context, in *Date, opts ...grpc.CallOption) (*Events, error)
	ShowEventMonth(ctx context.Context, in *Date, opts ...grpc.CallOption) (*Events, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) Create(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Calendar_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) Update(ctx context.Context, in *IDEvent, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Calendar_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Calendar_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) Show(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, Calendar_Show_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ShowEventDay(ctx context.Context, in *Date, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, Calendar_ShowEventDay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ShowEventWeek(ctx context.Context, in *Date, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, Calendar_ShowEventWeek_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ShowEventMonth(ctx context.Context, in *Date, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, Calendar_ShowEventMonth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
// All implementations must embed UnimplementedCalendarServer
// for forward compatibility
type CalendarServer interface {
	Create(context.Context, *Event) (*emptypb.Empty, error)
	Update(context.Context, *IDEvent) (*emptypb.Empty, error)
	Delete(context.Context, *ID) (*emptypb.Empty, error)
	Show(context.Context, *emptypb.Empty) (*Events, error)
	ShowEventDay(context.Context, *Date) (*Events, error)
	ShowEventWeek(context.Context, *Date) (*Events, error)
	ShowEventMonth(context.Context, *Date) (*Events, error)
	mustEmbedUnimplementedCalendarServer()
}

// UnimplementedCalendarServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (UnimplementedCalendarServer) Create(context.Context, *Event) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCalendarServer) Update(context.Context, *IDEvent) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCalendarServer) Delete(context.Context, *ID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCalendarServer) Show(context.Context, *emptypb.Empty) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedCalendarServer) ShowEventDay(context.Context, *Date) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowEventDay not implemented")
}
func (UnimplementedCalendarServer) ShowEventWeek(context.Context, *Date) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowEventWeek not implemented")
}
func (UnimplementedCalendarServer) ShowEventMonth(context.Context, *Date) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowEventMonth not implemented")
}
func (UnimplementedCalendarServer) mustEmbedUnimplementedCalendarServer() {}

// UnsafeCalendarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarServer will
// result in compilation errors.
type UnsafeCalendarServer interface {
	mustEmbedUnimplementedCalendarServer()
}

func RegisterCalendarServer(s grpc.ServiceRegistrar, srv CalendarServer) {
	s.RegisterService(&Calendar_ServiceDesc, srv)
}

func _Calendar_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Create(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Update(ctx, req.(*IDEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Delete(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_Show_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Show(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ShowEventDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Date)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ShowEventDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_ShowEventDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ShowEventDay(ctx, req.(*Date))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ShowEventWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Date)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ShowEventWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_ShowEventWeek_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ShowEventWeek(ctx, req.(*Date))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ShowEventMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Date)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ShowEventMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_ShowEventMonth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ShowEventMonth(ctx, req.(*Date))
	}
	return interceptor(ctx, in, info, handler)
}

// Calendar_ServiceDesc is the grpc.ServiceDesc for Calendar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calendar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Calendar_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Calendar_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Calendar_Delete_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _Calendar_Show_Handler,
		},
		{
			MethodName: "ShowEventDay",
			Handler:    _Calendar_ShowEventDay_Handler,
		},
		{
			MethodName: "ShowEventWeek",
			Handler:    _Calendar_ShowEventWeek_Handler,
		},
		{
			MethodName: "ShowEventMonth",
			Handler:    _Calendar_ShowEventMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "EventService.proto",
}
