// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateGatewayRequest struct {
	// Hex encoded mac address.
	Mac string `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
	// A name for the gateway
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// A description for the gateway
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Latitude of the gateway -90 to 90
	Latitude float64 `protobuf:"fixed64,4,opt,name=latitude" json:"latitude,omitempty"`
	// Longitude of the gateway -180 to 180
	Longitude float64 `protobuf:"fixed64,5,opt,name=longitude" json:"longitude,omitempty"`
	// Altitude of the gateway in meters
	Altitude float64 `protobuf:"fixed64,6,opt,name=altitude" json:"altitude,omitempty"`
	// ID of the organization to which this gateway belongs.
	OrganizationID int64 `protobuf:"varint,7,opt,name=organizationID" json:"organizationID,omitempty"`
}

func (m *CreateGatewayRequest) Reset()                    { *m = CreateGatewayRequest{} }
func (m *CreateGatewayRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateGatewayRequest) ProtoMessage()               {}
func (*CreateGatewayRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *CreateGatewayRequest) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *CreateGatewayRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateGatewayRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateGatewayRequest) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *CreateGatewayRequest) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *CreateGatewayRequest) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *CreateGatewayRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

type CreateGatewayResponse struct {
}

func (m *CreateGatewayResponse) Reset()                    { *m = CreateGatewayResponse{} }
func (m *CreateGatewayResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateGatewayResponse) ProtoMessage()               {}
func (*CreateGatewayResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type GetGatewayRequest struct {
	// Hex encoded mac address of the node.
	Mac string `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
}

func (m *GetGatewayRequest) Reset()                    { *m = GetGatewayRequest{} }
func (m *GetGatewayRequest) String() string            { return proto.CompactTextString(m) }
func (*GetGatewayRequest) ProtoMessage()               {}
func (*GetGatewayRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *GetGatewayRequest) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

type GetGatewayResponse struct {
	// Hex encoded mac address.
	Mac string `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
	// A name for the gateway
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// A description for the gateway
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Latitude of the gateway -90.0 to 90.0
	Latitude float64 `protobuf:"fixed64,4,opt,name=latitude" json:"latitude,omitempty"`
	// Longitude of the gateway -180.0 to 180.0
	Longitude float64 `protobuf:"fixed64,5,opt,name=longitude" json:"longitude,omitempty"`
	// Altitude of the gateway in meters
	Altitude float64 `protobuf:"fixed64,6,opt,name=altitude" json:"altitude,omitempty"`
	// Creation timestamp of the record
	CreatedAt string `protobuf:"bytes,7,opt,name=createdAt" json:"createdAt,omitempty"`
	// Last update timestamp of the record
	UpdatedAt string `protobuf:"bytes,8,opt,name=updatedAt" json:"updatedAt,omitempty"`
	// The timestamp of the first data from the gateway.
	FirstSeenAt string `protobuf:"bytes,9,opt,name=firstSeenAt" json:"firstSeenAt,omitempty"`
	// The timestamp of the most recent data from the gateway.
	LastSeenAt string `protobuf:"bytes,10,opt,name=lastSeenAt" json:"lastSeenAt,omitempty"`
	// ID of the organization to which this gateway belongs.
	OrganizationID int64 `protobuf:"varint,11,opt,name=organizationID" json:"organizationID,omitempty"`
}

func (m *GetGatewayResponse) Reset()                    { *m = GetGatewayResponse{} }
func (m *GetGatewayResponse) String() string            { return proto.CompactTextString(m) }
func (*GetGatewayResponse) ProtoMessage()               {}
func (*GetGatewayResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *GetGatewayResponse) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *GetGatewayResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetGatewayResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetGatewayResponse) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *GetGatewayResponse) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *GetGatewayResponse) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *GetGatewayResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetGatewayResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *GetGatewayResponse) GetFirstSeenAt() string {
	if m != nil {
		return m.FirstSeenAt
	}
	return ""
}

func (m *GetGatewayResponse) GetLastSeenAt() string {
	if m != nil {
		return m.LastSeenAt
	}
	return ""
}

func (m *GetGatewayResponse) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

type DeleteGatewayRequest struct {
	// Hex encoded mac address.
	Mac string `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
}

func (m *DeleteGatewayRequest) Reset()                    { *m = DeleteGatewayRequest{} }
func (m *DeleteGatewayRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteGatewayRequest) ProtoMessage()               {}
func (*DeleteGatewayRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *DeleteGatewayRequest) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

type DeleteGatewayResponse struct {
}

func (m *DeleteGatewayResponse) Reset()                    { *m = DeleteGatewayResponse{} }
func (m *DeleteGatewayResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteGatewayResponse) ProtoMessage()               {}
func (*DeleteGatewayResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

type ListGatewayRequest struct {
	// Max number of nodes to return in the result-set.
	Limit int32 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset of the result-set (for pagination).
	Offset int32 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	// ID of the organization for which to filter on, when left blank the
	// response will return all gateways to which the user has access to.
	OrganizationID int64 `protobuf:"varint,3,opt,name=organizationID" json:"organizationID,omitempty"`
}

func (m *ListGatewayRequest) Reset()                    { *m = ListGatewayRequest{} }
func (m *ListGatewayRequest) String() string            { return proto.CompactTextString(m) }
func (*ListGatewayRequest) ProtoMessage()               {}
func (*ListGatewayRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *ListGatewayRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListGatewayRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListGatewayRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

type ListGatewayItem struct {
	// Hex encoded mac address.
	Mac string `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
	// A name for the gateway
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// A description for the gateway
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Creation timestamp of the record
	CreatedAt string `protobuf:"bytes,4,opt,name=createdAt" json:"createdAt,omitempty"`
	// Last update timestamp of the record
	UpdatedAt string `protobuf:"bytes,5,opt,name=updatedAt" json:"updatedAt,omitempty"`
	// ID of the organization to which this gateway belongs.
	OrganizationID int64 `protobuf:"varint,6,opt,name=organizationID" json:"organizationID,omitempty"`
}

func (m *ListGatewayItem) Reset()                    { *m = ListGatewayItem{} }
func (m *ListGatewayItem) String() string            { return proto.CompactTextString(m) }
func (*ListGatewayItem) ProtoMessage()               {}
func (*ListGatewayItem) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *ListGatewayItem) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *ListGatewayItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListGatewayItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ListGatewayItem) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ListGatewayItem) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *ListGatewayItem) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

type ListGatewayResponse struct {
	// Total number of nodes available within the result-set.
	TotalCount int32 `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	// Nodes within this result-set.
	Result []*ListGatewayItem `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListGatewayResponse) Reset()                    { *m = ListGatewayResponse{} }
func (m *ListGatewayResponse) String() string            { return proto.CompactTextString(m) }
func (*ListGatewayResponse) ProtoMessage()               {}
func (*ListGatewayResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *ListGatewayResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListGatewayResponse) GetResult() []*ListGatewayItem {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateGatewayRequest struct {
	// Hex encoded mac address.
	Mac string `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
	// A name for the gateway
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// A description for the gateway
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Latitude of the gateway -90.0 to 90.0
	Latitude float64 `protobuf:"fixed64,4,opt,name=latitude" json:"latitude,omitempty"`
	// Longitude of the gateway -180.0 to 180.0
	Longitude float64 `protobuf:"fixed64,5,opt,name=longitude" json:"longitude,omitempty"`
	// Altitude of the gateway in meters
	Altitude float64 `protobuf:"fixed64,6,opt,name=altitude" json:"altitude,omitempty"`
	// ID of the organization to which this gateway belongs.
	OrganizationID int64 `protobuf:"varint,7,opt,name=organizationID" json:"organizationID,omitempty"`
}

func (m *UpdateGatewayRequest) Reset()                    { *m = UpdateGatewayRequest{} }
func (m *UpdateGatewayRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateGatewayRequest) ProtoMessage()               {}
func (*UpdateGatewayRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *UpdateGatewayRequest) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *UpdateGatewayRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateGatewayRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateGatewayRequest) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *UpdateGatewayRequest) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *UpdateGatewayRequest) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *UpdateGatewayRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

type UpdateGatewayResponse struct {
}

func (m *UpdateGatewayResponse) Reset()                    { *m = UpdateGatewayResponse{} }
func (m *UpdateGatewayResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateGatewayResponse) ProtoMessage()               {}
func (*UpdateGatewayResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

type GatewayStats struct {
	// Timestamp of the (aggregated) measurement.
	Timestamp string `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// Packets received by the gateway.
	RxPacketsReceived int32 `protobuf:"varint,2,opt,name=rxPacketsReceived" json:"rxPacketsReceived,omitempty"`
	// Packets received by the gateway that passed the CRC check.
	RxPacketsReceivedOK int32 `protobuf:"varint,3,opt,name=rxPacketsReceivedOK" json:"rxPacketsReceivedOK,omitempty"`
	// Packets received by the gateway for transmission.
	TxPacketsReceived int32 `protobuf:"varint,4,opt,name=txPacketsReceived" json:"txPacketsReceived,omitempty"`
	// Packets transmitted by the gateway.
	TxPacketsEmitted int32 `protobuf:"varint,5,opt,name=txPacketsEmitted" json:"txPacketsEmitted,omitempty"`
}

func (m *GatewayStats) Reset()                    { *m = GatewayStats{} }
func (m *GatewayStats) String() string            { return proto.CompactTextString(m) }
func (*GatewayStats) ProtoMessage()               {}
func (*GatewayStats) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{11} }

func (m *GatewayStats) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *GatewayStats) GetRxPacketsReceived() int32 {
	if m != nil {
		return m.RxPacketsReceived
	}
	return 0
}

func (m *GatewayStats) GetRxPacketsReceivedOK() int32 {
	if m != nil {
		return m.RxPacketsReceivedOK
	}
	return 0
}

func (m *GatewayStats) GetTxPacketsReceived() int32 {
	if m != nil {
		return m.TxPacketsReceived
	}
	return 0
}

func (m *GatewayStats) GetTxPacketsEmitted() int32 {
	if m != nil {
		return m.TxPacketsEmitted
	}
	return 0
}

type GetGatewayStatsRequest struct {
	// Hex encoded mac address.
	Mac string `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
	// Aggregation interval.  One of "second", "minute", "hour", "day", "week",
	// "month", "quarter", "year".  Case insensitive.
	Interval string `protobuf:"bytes,2,opt,name=interval" json:"interval,omitempty"`
	// Timestamp to start from.
	StartTimestamp string `protobuf:"bytes,3,opt,name=startTimestamp" json:"startTimestamp,omitempty"`
	// Timestamp until to get from.
	EndTimestamp string `protobuf:"bytes,4,opt,name=endTimestamp" json:"endTimestamp,omitempty"`
}

func (m *GetGatewayStatsRequest) Reset()                    { *m = GetGatewayStatsRequest{} }
func (m *GetGatewayStatsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetGatewayStatsRequest) ProtoMessage()               {}
func (*GetGatewayStatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{12} }

func (m *GetGatewayStatsRequest) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *GetGatewayStatsRequest) GetInterval() string {
	if m != nil {
		return m.Interval
	}
	return ""
}

func (m *GetGatewayStatsRequest) GetStartTimestamp() string {
	if m != nil {
		return m.StartTimestamp
	}
	return ""
}

func (m *GetGatewayStatsRequest) GetEndTimestamp() string {
	if m != nil {
		return m.EndTimestamp
	}
	return ""
}

type GetGatewayStatsResponse struct {
	Result []*GatewayStats `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
}

func (m *GetGatewayStatsResponse) Reset()                    { *m = GetGatewayStatsResponse{} }
func (m *GetGatewayStatsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetGatewayStatsResponse) ProtoMessage()               {}
func (*GetGatewayStatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{13} }

func (m *GetGatewayStatsResponse) GetResult() []*GatewayStats {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateGatewayRequest)(nil), "api.CreateGatewayRequest")
	proto.RegisterType((*CreateGatewayResponse)(nil), "api.CreateGatewayResponse")
	proto.RegisterType((*GetGatewayRequest)(nil), "api.GetGatewayRequest")
	proto.RegisterType((*GetGatewayResponse)(nil), "api.GetGatewayResponse")
	proto.RegisterType((*DeleteGatewayRequest)(nil), "api.DeleteGatewayRequest")
	proto.RegisterType((*DeleteGatewayResponse)(nil), "api.DeleteGatewayResponse")
	proto.RegisterType((*ListGatewayRequest)(nil), "api.ListGatewayRequest")
	proto.RegisterType((*ListGatewayItem)(nil), "api.ListGatewayItem")
	proto.RegisterType((*ListGatewayResponse)(nil), "api.ListGatewayResponse")
	proto.RegisterType((*UpdateGatewayRequest)(nil), "api.UpdateGatewayRequest")
	proto.RegisterType((*UpdateGatewayResponse)(nil), "api.UpdateGatewayResponse")
	proto.RegisterType((*GatewayStats)(nil), "api.GatewayStats")
	proto.RegisterType((*GetGatewayStatsRequest)(nil), "api.GetGatewayStatsRequest")
	proto.RegisterType((*GetGatewayStatsResponse)(nil), "api.GetGatewayStatsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Gateway service

type GatewayClient interface {
	// Create creates the given gateway.
	Create(ctx context.Context, in *CreateGatewayRequest, opts ...grpc.CallOption) (*CreateGatewayResponse, error)
	// Get returns the gateway for the requested mac address.
	Get(ctx context.Context, in *GetGatewayRequest, opts ...grpc.CallOption) (*GetGatewayResponse, error)
	// Update updates the gateway matching the given mac address.
	Update(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*UpdateGatewayResponse, error)
	// Delete deletes the gateway matching the given mac address.
	Delete(ctx context.Context, in *DeleteGatewayRequest, opts ...grpc.CallOption) (*DeleteGatewayResponse, error)
	// List lists the gateways.
	List(ctx context.Context, in *ListGatewayRequest, opts ...grpc.CallOption) (*ListGatewayResponse, error)
	// GetStats lists the gateway stats given the query parameters.
	GetStats(ctx context.Context, in *GetGatewayStatsRequest, opts ...grpc.CallOption) (*GetGatewayStatsResponse, error)
}

type gatewayClient struct {
	cc *grpc.ClientConn
}

func NewGatewayClient(cc *grpc.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Create(ctx context.Context, in *CreateGatewayRequest, opts ...grpc.CallOption) (*CreateGatewayResponse, error) {
	out := new(CreateGatewayResponse)
	err := grpc.Invoke(ctx, "/api.Gateway/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Get(ctx context.Context, in *GetGatewayRequest, opts ...grpc.CallOption) (*GetGatewayResponse, error) {
	out := new(GetGatewayResponse)
	err := grpc.Invoke(ctx, "/api.Gateway/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Update(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*UpdateGatewayResponse, error) {
	out := new(UpdateGatewayResponse)
	err := grpc.Invoke(ctx, "/api.Gateway/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Delete(ctx context.Context, in *DeleteGatewayRequest, opts ...grpc.CallOption) (*DeleteGatewayResponse, error) {
	out := new(DeleteGatewayResponse)
	err := grpc.Invoke(ctx, "/api.Gateway/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) List(ctx context.Context, in *ListGatewayRequest, opts ...grpc.CallOption) (*ListGatewayResponse, error) {
	out := new(ListGatewayResponse)
	err := grpc.Invoke(ctx, "/api.Gateway/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetStats(ctx context.Context, in *GetGatewayStatsRequest, opts ...grpc.CallOption) (*GetGatewayStatsResponse, error) {
	out := new(GetGatewayStatsResponse)
	err := grpc.Invoke(ctx, "/api.Gateway/GetStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gateway service

type GatewayServer interface {
	// Create creates the given gateway.
	Create(context.Context, *CreateGatewayRequest) (*CreateGatewayResponse, error)
	// Get returns the gateway for the requested mac address.
	Get(context.Context, *GetGatewayRequest) (*GetGatewayResponse, error)
	// Update updates the gateway matching the given mac address.
	Update(context.Context, *UpdateGatewayRequest) (*UpdateGatewayResponse, error)
	// Delete deletes the gateway matching the given mac address.
	Delete(context.Context, *DeleteGatewayRequest) (*DeleteGatewayResponse, error)
	// List lists the gateways.
	List(context.Context, *ListGatewayRequest) (*ListGatewayResponse, error)
	// GetStats lists the gateway stats given the query parameters.
	GetStats(context.Context, *GetGatewayStatsRequest) (*GetGatewayStatsResponse, error)
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gateway/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Create(ctx, req.(*CreateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gateway/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Get(ctx, req.(*GetGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gateway/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Update(ctx, req.(*UpdateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gateway/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Delete(ctx, req.(*DeleteGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gateway/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).List(ctx, req.(*ListGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gateway/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetStats(ctx, req.(*GetGatewayStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Gateway_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Gateway_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Gateway_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Gateway_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Gateway_List_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _Gateway_GetStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0x36, 0xc9, 0x36, 0x99, 0xb6, 0xb4, 0x75, 0xd3, 0x66, 0xbb, 0x8d, 0xaa, 0xb0, 0x12,
	0x28, 0x54, 0x55, 0x83, 0xca, 0x8d, 0x5b, 0xd5, 0xa2, 0xa8, 0x02, 0x89, 0x6a, 0x0b, 0x07, 0x24,
	0x24, 0x64, 0x12, 0x37, 0x18, 0xf6, 0x8f, 0xf5, 0xa4, 0xfc, 0x89, 0x0b, 0x8f, 0x00, 0xaf, 0xc2,
	0x95, 0xa7, 0xe0, 0xc2, 0x81, 0x23, 0xe2, 0x39, 0xd0, 0xda, 0x6e, 0xb2, 0xd9, 0x35, 0xf4, 0xc2,
	0x01, 0x71, 0x8b, 0xbf, 0x19, 0x7f, 0xfb, 0x79, 0xe6, 0xf3, 0x38, 0xb0, 0x34, 0xa2, 0xc8, 0x5e,
	0xd1, 0x37, 0x7b, 0x49, 0x1a, 0x63, 0x4c, 0x2a, 0x34, 0xe1, 0x6e, 0x7b, 0x14, 0xc7, 0xa3, 0x80,
	0xf5, 0x68, 0xc2, 0x7b, 0x34, 0x8a, 0x62, 0xa4, 0xc8, 0xe3, 0x48, 0xa8, 0x14, 0xef, 0xbb, 0x05,
	0xcd, 0xc3, 0x94, 0x51, 0x64, 0x7d, 0xb5, 0xd5, 0x67, 0x2f, 0xc7, 0x4c, 0x20, 0x59, 0x81, 0x4a,
	0x48, 0x07, 0x8e, 0xd5, 0xb1, 0xba, 0x0d, 0x3f, 0xfb, 0x49, 0x08, 0x54, 0x23, 0x1a, 0x32, 0x67,
	0x4e, 0x42, 0xf2, 0x37, 0xe9, 0xc0, 0xc2, 0x90, 0x89, 0x41, 0xca, 0x93, 0x8c, 0xd4, 0xa9, 0xc8,
	0x50, 0x1e, 0x22, 0x2e, 0xd4, 0x03, 0x8a, 0x1c, 0xc7, 0x43, 0xe6, 0x54, 0x3b, 0x56, 0xd7, 0xf2,
	0x27, 0x6b, 0xd2, 0x86, 0x46, 0x10, 0x47, 0x23, 0x15, 0xac, 0xc9, 0xe0, 0x14, 0xc8, 0x76, 0xd2,
	0x40, 0xef, 0xb4, 0xd5, 0xce, 0x8b, 0x35, 0xb9, 0x0e, 0x57, 0xe2, 0x74, 0x44, 0x23, 0xfe, 0x56,
	0x9e, 0xe6, 0xf8, 0xc8, 0x99, 0xef, 0x58, 0xdd, 0x8a, 0x5f, 0x40, 0xbd, 0x16, 0xac, 0x17, 0x4e,
	0x27, 0x92, 0x38, 0x12, 0xcc, 0xbb, 0x06, 0xab, 0x7d, 0x86, 0x97, 0x9d, 0xd9, 0xfb, 0x36, 0x07,
	0x24, 0x9f, 0xa7, 0x76, 0xff, 0xe3, 0xc5, 0x69, 0x43, 0x63, 0x20, 0x0f, 0x3d, 0x3c, 0x40, 0x59,
	0x97, 0x86, 0x3f, 0x05, 0xb2, 0xe8, 0x38, 0x19, 0xea, 0x68, 0x5d, 0x45, 0x27, 0x40, 0xa6, 0xf9,
	0x8c, 0xa7, 0x02, 0x4f, 0x19, 0x8b, 0x0e, 0xd0, 0x69, 0x28, 0xcd, 0x39, 0x88, 0x6c, 0x03, 0x04,
	0x74, 0x92, 0x00, 0x32, 0x21, 0x87, 0x18, 0x5a, 0xb3, 0x60, 0x6c, 0x4d, 0x17, 0x9a, 0x47, 0x2c,
	0x60, 0x97, 0x1b, 0x2f, 0x6b, 0x62, 0x21, 0x53, 0x37, 0xf1, 0x39, 0x90, 0x7b, 0x5c, 0x14, 0xbb,
	0xd8, 0x84, 0x5a, 0xc0, 0x43, 0x8e, 0x92, 0xa2, 0xe6, 0xab, 0x05, 0xd9, 0x00, 0x3b, 0x3e, 0x3b,
	0x13, 0x0c, 0x65, 0x8b, 0x6a, 0xbe, 0x5e, 0x19, 0xe4, 0x56, 0x8c, 0x72, 0xbf, 0x58, 0xb0, 0x9c,
	0xfb, 0xd8, 0x31, 0xb2, 0xf0, 0xaf, 0xd9, 0x60, 0xa6, 0x61, 0xd5, 0x3f, 0x36, 0xac, 0x56, 0x6c,
	0x58, 0x59, 0xbf, 0x6d, 0xd4, 0x3f, 0x80, 0xb5, 0x99, 0x5a, 0x69, 0x27, 0x6f, 0x03, 0x60, 0x8c,
	0x34, 0x38, 0x8c, 0xc7, 0xd1, 0x45, 0xc5, 0x72, 0x08, 0xd9, 0x05, 0x3b, 0x65, 0x62, 0x1c, 0x64,
	0x65, 0xab, 0x74, 0x17, 0xf6, 0x9b, 0x7b, 0x34, 0xe1, 0x7b, 0x85, 0x42, 0xf8, 0x3a, 0x47, 0x4e,
	0x93, 0x87, 0x52, 0xda, 0xff, 0x3a, 0x4d, 0x0a, 0xa7, 0xd3, 0x46, 0xfc, 0x69, 0xc1, 0xa2, 0xc6,
	0x4e, 0x91, 0xa2, 0xc8, 0xb4, 0x20, 0x0f, 0x99, 0x40, 0x1a, 0x26, 0xfa, 0xd4, 0x53, 0x80, 0xec,
	0xc2, 0x6a, 0xfa, 0xfa, 0x84, 0x0e, 0x5e, 0x30, 0x14, 0x3e, 0x1b, 0x30, 0x7e, 0xce, 0x86, 0xda,
	0x96, 0xe5, 0x00, 0xb9, 0x09, 0x6b, 0x25, 0xf0, 0xfe, 0x5d, 0x59, 0x9d, 0x9a, 0x6f, 0x0a, 0x65,
	0xfc, 0x58, 0xe2, 0xaf, 0x2a, 0xfe, 0x52, 0x80, 0xec, 0xc0, 0xca, 0x04, 0xbc, 0x13, 0x72, 0x44,
	0x36, 0x94, 0xe5, 0xab, 0xf9, 0x25, 0xdc, 0xfb, 0x68, 0xc1, 0xc6, 0x74, 0x1e, 0xca, 0xb3, 0xfe,
	0xbe, 0xc5, 0x2e, 0xd4, 0x79, 0x84, 0x2c, 0x3d, 0xa7, 0x81, 0x6e, 0xf3, 0x64, 0x9d, 0x95, 0x5c,
	0x20, 0x4d, 0xf1, 0xc1, 0xa4, 0x4a, 0xaa, 0xdb, 0x05, 0x94, 0x78, 0xb0, 0xc8, 0xa2, 0xe1, 0x34,
	0x4b, 0xdd, 0x8e, 0x19, 0xcc, 0x3b, 0x82, 0x56, 0x49, 0x93, 0xb6, 0xf7, 0x8d, 0x89, 0x7d, 0x2d,
	0x69, 0xdf, 0x55, 0x69, 0xdf, 0x99, 0x54, 0x9d, 0xb0, 0xff, 0xb9, 0x0a, 0xf3, 0x3a, 0x40, 0x1e,
	0x81, 0xad, 0x9e, 0x0d, 0xb2, 0x29, 0x37, 0x98, 0x5e, 0x48, 0xd7, 0x35, 0x85, 0xb4, 0x21, 0x9c,
	0x0f, 0x5f, 0x7f, 0x7c, 0x9a, 0x23, 0xde, 0x92, 0x7c, 0x76, 0xf5, 0xab, 0x2c, 0x6e, 0x5b, 0x3b,
	0xe4, 0x14, 0x2a, 0x7d, 0x86, 0x64, 0x43, 0x09, 0x29, 0x3e, 0x41, 0x6e, 0xab, 0x84, 0x6b, 0xc6,
	0x2d, 0xc9, 0xb8, 0x4e, 0xd6, 0x66, 0x18, 0x7b, 0xef, 0x42, 0x3a, 0x78, 0x4f, 0x9e, 0x80, 0xad,
	0x8c, 0xa9, 0xf5, 0x9a, 0xee, 0xa0, 0xd6, 0x6b, 0x36, 0xf0, 0xb6, 0x64, 0x77, 0x5c, 0x13, 0x7b,
	0xa6, 0xfa, 0x31, 0xd8, 0x6a, 0x04, 0xeb, 0x0f, 0x98, 0x26, 0xb7, 0xfe, 0x80, 0x79, 0x54, 0x6b,
	0xf9, 0x3b, 0x46, 0xf9, 0x27, 0x50, 0xcd, 0x26, 0x0a, 0x69, 0x15, 0x87, 0xcb, 0x05, 0xb3, 0x53,
	0x0e, 0x68, 0xde, 0x75, 0xc9, 0xbb, 0x4c, 0x66, 0x0b, 0x4d, 0x9e, 0x41, 0xbd, 0xcf, 0x50, 0xdd,
	0xc5, 0xad, 0x42, 0x49, 0xf3, 0xae, 0x75, 0xdb, 0xe6, 0xa0, 0x66, 0xbf, 0x2a, 0xd9, 0xb7, 0xc8,
	0xa6, 0x41, 0x75, 0x4f, 0x64, 0xa9, 0x4f, 0x6d, 0xf9, 0x3f, 0xea, 0xd6, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2b, 0xb2, 0x8f, 0xc3, 0x7b, 0x09, 0x00, 0x00,
}
