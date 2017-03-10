// Code generated by protoc-gen-go.
// source: service/monitoring.proto
// DO NOT EDIT!

/*
Package generated_proto is a generated protocol buffer package.

It is generated from these files:
	service/monitoring.proto

It has these top-level messages:
	CanonicalRpcStats
	StatsRequest
	StatsResponse
	TraceRequest
	TraceResponse
	MonitoringDataGroup
	CustomMonitoringData
	ViewResponse
*/
package generated_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_instrumentation "stats"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import _ "github.com/golang/protobuf/ptypes/empty"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Canonical RPC stats exported by gRPC.
type CanonicalRpcStats struct {
	RpcClientErrors            *CanonicalRpcStats_View `protobuf:"bytes,1,opt,name=rpc_client_errors,json=rpcClientErrors" json:"rpc_client_errors,omitempty"`
	RpcClientCompletedRpcs     *CanonicalRpcStats_View `protobuf:"bytes,2,opt,name=rpc_client_completed_rpcs,json=rpcClientCompletedRpcs" json:"rpc_client_completed_rpcs,omitempty"`
	RpcClientStartedRpcs       *CanonicalRpcStats_View `protobuf:"bytes,3,opt,name=rpc_client_started_rpcs,json=rpcClientStartedRpcs" json:"rpc_client_started_rpcs,omitempty"`
	RpcClientElapsedTime       *CanonicalRpcStats_View `protobuf:"bytes,4,opt,name=rpc_client_elapsed_time,json=rpcClientElapsedTime" json:"rpc_client_elapsed_time,omitempty"`
	RpcClientServerElapsedTime *CanonicalRpcStats_View `protobuf:"bytes,5,opt,name=rpc_client_server_elapsed_time,json=rpcClientServerElapsedTime" json:"rpc_client_server_elapsed_time,omitempty"`
	RpcClientRequestBytes      *CanonicalRpcStats_View `protobuf:"bytes,6,opt,name=rpc_client_request_bytes,json=rpcClientRequestBytes" json:"rpc_client_request_bytes,omitempty"`
	RpcClientResponseBytes     *CanonicalRpcStats_View `protobuf:"bytes,7,opt,name=rpc_client_response_bytes,json=rpcClientResponseBytes" json:"rpc_client_response_bytes,omitempty"`
	RpcClientRequestCount      *CanonicalRpcStats_View `protobuf:"bytes,8,opt,name=rpc_client_request_count,json=rpcClientRequestCount" json:"rpc_client_request_count,omitempty"`
	RpcClientResponseCount     *CanonicalRpcStats_View `protobuf:"bytes,9,opt,name=rpc_client_response_count,json=rpcClientResponseCount" json:"rpc_client_response_count,omitempty"`
	RpcServerErrors            *CanonicalRpcStats_View `protobuf:"bytes,10,opt,name=rpc_server_errors,json=rpcServerErrors" json:"rpc_server_errors,omitempty"`
	RpcServerCompletedRpcs     *CanonicalRpcStats_View `protobuf:"bytes,11,opt,name=rpc_server_completed_rpcs,json=rpcServerCompletedRpcs" json:"rpc_server_completed_rpcs,omitempty"`
	RpcServerServerElapsedTime *CanonicalRpcStats_View `protobuf:"bytes,12,opt,name=rpc_server_server_elapsed_time,json=rpcServerServerElapsedTime" json:"rpc_server_server_elapsed_time,omitempty"`
	RpcServerRequestBytes      *CanonicalRpcStats_View `protobuf:"bytes,13,opt,name=rpc_server_request_bytes,json=rpcServerRequestBytes" json:"rpc_server_request_bytes,omitempty"`
	RpcServerResponseBytes     *CanonicalRpcStats_View `protobuf:"bytes,14,opt,name=rpc_server_response_bytes,json=rpcServerResponseBytes" json:"rpc_server_response_bytes,omitempty"`
	RpcServerRequestCount      *CanonicalRpcStats_View `protobuf:"bytes,15,opt,name=rpc_server_request_count,json=rpcServerRequestCount" json:"rpc_server_request_count,omitempty"`
	RpcServerResponseCount     *CanonicalRpcStats_View `protobuf:"bytes,16,opt,name=rpc_server_response_count,json=rpcServerResponseCount" json:"rpc_server_response_count,omitempty"`
	RpcServerElapsedTime       *CanonicalRpcStats_View `protobuf:"bytes,17,opt,name=rpc_server_elapsed_time,json=rpcServerElapsedTime" json:"rpc_server_elapsed_time,omitempty"`
}

func (m *CanonicalRpcStats) Reset()                    { *m = CanonicalRpcStats{} }
func (m *CanonicalRpcStats) String() string            { return proto.CompactTextString(m) }
func (*CanonicalRpcStats) ProtoMessage()               {}
func (*CanonicalRpcStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CanonicalRpcStats) GetRpcClientErrors() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientErrors
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientCompletedRpcs() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientCompletedRpcs
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientStartedRpcs() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientStartedRpcs
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientElapsedTime() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientElapsedTime
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientServerElapsedTime() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientServerElapsedTime
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientRequestBytes() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientRequestBytes
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientResponseBytes() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientResponseBytes
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientRequestCount() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientRequestCount
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientResponseCount() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientResponseCount
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerErrors() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerErrors
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerCompletedRpcs() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerCompletedRpcs
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerServerElapsedTime() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerServerElapsedTime
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerRequestBytes() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerRequestBytes
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerResponseBytes() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerResponseBytes
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerRequestCount() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerRequestCount
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerResponseCount() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerResponseCount
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerElapsedTime() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerElapsedTime
	}
	return nil
}

// Wrapper combining View and ViewDescriptor.
type CanonicalRpcStats_View struct {
	MeasurementDescriptor *google_instrumentation.MeasurementDescriptor `protobuf:"bytes,1,opt,name=measurement_descriptor,json=measurementDescriptor" json:"measurement_descriptor,omitempty"`
	ViewDescriptor        *google_instrumentation.ViewDescriptor        `protobuf:"bytes,2,opt,name=view_descriptor,json=viewDescriptor" json:"view_descriptor,omitempty"`
	View                  *google_instrumentation.View                  `protobuf:"bytes,3,opt,name=view" json:"view,omitempty"`
}

func (m *CanonicalRpcStats_View) Reset()                    { *m = CanonicalRpcStats_View{} }
func (m *CanonicalRpcStats_View) String() string            { return proto.CompactTextString(m) }
func (*CanonicalRpcStats_View) ProtoMessage()               {}
func (*CanonicalRpcStats_View) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CanonicalRpcStats_View) GetMeasurementDescriptor() *google_instrumentation.MeasurementDescriptor {
	if m != nil {
		return m.MeasurementDescriptor
	}
	return nil
}

func (m *CanonicalRpcStats_View) GetViewDescriptor() *google_instrumentation.ViewDescriptor {
	if m != nil {
		return m.ViewDescriptor
	}
	return nil
}

func (m *CanonicalRpcStats_View) GetView() *google_instrumentation.View {
	if m != nil {
		return m.View
	}
	return nil
}

// This message is sent when requesting a set of stats (Census Views) from
// a client system, as part of the MonitoringService API's.
// TODO(aveitch): should this be named ViewRequest?
type StatsRequest struct {
	// An optional set of ViewDescriptor names. Only Views using these
	// descriptors will be sent back in the response. If no names are provided,
	// then all Views present in the client system will be included in every
	// response. If measurement_names is also provided, then Views matching the
	// intersection of the two are returned.
	// TODO(aveitch): Consider making this a list of regexes or prefix matches in
	// the future.
	ViewNames []string `protobuf:"bytes,1,rep,name=view_names,json=viewNames" json:"view_names,omitempty"`
	// An optional set of MeasurementDescriptor names. Only Views using these
	// descriptors will be sent back in the response. If no names are provided,
	// then all Views present in the client system will be included in every
	// response. If view_names is also provided, then Views matching the
	// intersection of the two are returned.
	// TODO(aveitch): Consider making this a list of regexes or prefix matches in
	// the future.
	MeasurementNames []string `protobuf:"bytes,2,rep,name=measurement_names,json=measurementNames" json:"measurement_names,omitempty"`
	// By default, the MeasurementDescriptors and ViewDescriptors corresponding to
	// the Views that can returned in a StatsResponse will be included in the
	// first such response. Set this to true to have them not sent.
	DontIncludeDescriptorsInFirstResponse bool `protobuf:"varint,3,opt,name=dont_include_descriptors_in_first_response,json=dontIncludeDescriptorsInFirstResponse" json:"dont_include_descriptors_in_first_response,omitempty"`
}

func (m *StatsRequest) Reset()                    { *m = StatsRequest{} }
func (m *StatsRequest) String() string            { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()               {}
func (*StatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// This message is sent in response to a GetStats or WatchStats call.
type StatsResponse struct {
	// The set of Views corresponding to the StatsRequest.
	ViewResponses []*ViewResponse `protobuf:"bytes,1,rep,name=view_responses,json=viewResponses" json:"view_responses,omitempty"`
}

func (m *StatsResponse) Reset()                    { *m = StatsResponse{} }
func (m *StatsResponse) String() string            { return proto.CompactTextString(m) }
func (*StatsResponse) ProtoMessage()               {}
func (*StatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StatsResponse) GetViewResponses() []*ViewResponse {
	if m != nil {
		return m.ViewResponses
	}
	return nil
}

type TraceRequest struct {
}

func (m *TraceRequest) Reset()                    { *m = TraceRequest{} }
func (m *TraceRequest) String() string            { return proto.CompactTextString(m) }
func (*TraceRequest) ProtoMessage()               {}
func (*TraceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type TraceResponse struct {
}

func (m *TraceResponse) Reset()                    { *m = TraceResponse{} }
func (m *TraceResponse) String() string            { return proto.CompactTextString(m) }
func (*TraceResponse) ProtoMessage()               {}
func (*TraceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type MonitoringDataGroup struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *MonitoringDataGroup) Reset()                    { *m = MonitoringDataGroup{} }
func (m *MonitoringDataGroup) String() string            { return proto.CompactTextString(m) }
func (*MonitoringDataGroup) ProtoMessage()               {}
func (*MonitoringDataGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// The wrapper for custom monitoring data.
type CustomMonitoringData struct {
	// can be any application specific monitoring data
	Contents *google_protobuf.Any `protobuf:"bytes,1,opt,name=contents" json:"contents,omitempty"`
}

func (m *CustomMonitoringData) Reset()                    { *m = CustomMonitoringData{} }
func (m *CustomMonitoringData) String() string            { return proto.CompactTextString(m) }
func (*CustomMonitoringData) ProtoMessage()               {}
func (*CustomMonitoringData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CustomMonitoringData) GetContents() *google_protobuf.Any {
	if m != nil {
		return m.Contents
	}
	return nil
}

// This message contains all information relevant to a single View. It is used
// in StatsResponse.
type ViewResponse struct {
	// Each ViewResponse can optionally contain the MeasurementDescriptor and
	// ViewDescriptor for the View. These will be sent in the first WatchStats
	// response, or all GetStats responses. This is disabled if
	// dont_include_descriptors_in_first_response is set to true in the
	// StatsRequest.
	MeasurementDescriptor *google_instrumentation.MeasurementDescriptor `protobuf:"bytes,1,opt,name=measurement_descriptor,json=measurementDescriptor" json:"measurement_descriptor,omitempty"`
	ViewDescriptor        *google_instrumentation.ViewDescriptor        `protobuf:"bytes,2,opt,name=view_descriptor,json=viewDescriptor" json:"view_descriptor,omitempty"`
	// The View data.
	View *google_instrumentation.View `protobuf:"bytes,3,opt,name=view" json:"view,omitempty"`
}

func (m *ViewResponse) Reset()                    { *m = ViewResponse{} }
func (m *ViewResponse) String() string            { return proto.CompactTextString(m) }
func (*ViewResponse) ProtoMessage()               {}
func (*ViewResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ViewResponse) GetMeasurementDescriptor() *google_instrumentation.MeasurementDescriptor {
	if m != nil {
		return m.MeasurementDescriptor
	}
	return nil
}

func (m *ViewResponse) GetViewDescriptor() *google_instrumentation.ViewDescriptor {
	if m != nil {
		return m.ViewDescriptor
	}
	return nil
}

func (m *ViewResponse) GetView() *google_instrumentation.View {
	if m != nil {
		return m.View
	}
	return nil
}

func init() {
	proto.RegisterType((*CanonicalRpcStats)(nil), "google.instrumentation.CanonicalRpcStats")
	proto.RegisterType((*CanonicalRpcStats_View)(nil), "google.instrumentation.CanonicalRpcStats.View")
	proto.RegisterType((*StatsRequest)(nil), "google.instrumentation.StatsRequest")
	proto.RegisterType((*StatsResponse)(nil), "google.instrumentation.StatsResponse")
	proto.RegisterType((*TraceRequest)(nil), "google.instrumentation.TraceRequest")
	proto.RegisterType((*TraceResponse)(nil), "google.instrumentation.TraceResponse")
	proto.RegisterType((*MonitoringDataGroup)(nil), "google.instrumentation.MonitoringDataGroup")
	proto.RegisterType((*CustomMonitoringData)(nil), "google.instrumentation.CustomMonitoringData")
	proto.RegisterType((*ViewResponse)(nil), "google.instrumentation.ViewResponse")
}

func init() { proto.RegisterFile("service/monitoring.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 864 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x97, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xc7, 0xe3, 0x26, 0x94, 0xe4, 0xe4, 0xc3, 0xc9, 0xe0, 0xba, 0xae, 0xf9, 0x10, 0x5a, 0xd1,
	0xaa, 0xa5, 0x60, 0x47, 0xe1, 0x09, 0x88, 0x9b, 0x9a, 0x0a, 0x15, 0xd0, 0xb6, 0x02, 0xb5, 0x80,
	0x56, 0xdb, 0xf1, 0xa9, 0x19, 0xc9, 0x3b, 0xb3, 0xcc, 0xcc, 0x3a, 0xca, 0x1d, 0x2f, 0xc5, 0x73,
	0x71, 0xc9, 0x1d, 0x42, 0xf3, 0xe1, 0xf5, 0xae, 0xbd, 0x1b, 0xd0, 0xc6, 0x5c, 0x71, 0x95, 0xcc,
	0xcc, 0x99, 0xff, 0x6f, 0xce, 0xd9, 0xff, 0xd9, 0x1d, 0x43, 0x4f, 0xa1, 0x9c, 0x33, 0x8a, 0xc3,
	0x44, 0x70, 0xa6, 0x85, 0x64, 0x7c, 0x3a, 0x48, 0xa5, 0xd0, 0x82, 0x74, 0xa7, 0x42, 0x4c, 0x67,
	0x38, 0x60, 0x5c, 0x69, 0x99, 0x25, 0xc8, 0x75, 0xac, 0x99, 0xe0, 0x7d, 0xa2, 0x74, 0xac, 0xd5,
	0x90, 0x22, 0x57, 0x99, 0x72, 0xb1, 0xfd, 0x7b, 0x2e, 0x76, 0x68, 0x47, 0x6f, 0xb2, 0xb7, 0xc3,
	0x98, 0x5f, 0xf9, 0xa5, 0xf7, 0x57, 0x97, 0x30, 0x49, 0xb5, 0x5f, 0x0c, 0x7e, 0x3b, 0x86, 0x93,
	0x51, 0xcc, 0x05, 0x67, 0x34, 0x9e, 0x85, 0x29, 0x7d, 0x61, 0xa4, 0xc9, 0x6b, 0x38, 0x91, 0x29,
	0x8d, 0xe8, 0x8c, 0x21, 0xd7, 0x11, 0x4a, 0x29, 0xa4, 0xea, 0xb5, 0x3e, 0x6e, 0x3d, 0xdc, 0x3f,
	0x1b, 0x0c, 0xaa, 0x4f, 0x35, 0x58, 0x53, 0x19, 0x7c, 0xcf, 0xf0, 0x32, 0x6c, 0xcb, 0x94, 0x8e,
	0xac, 0xce, 0x85, 0x95, 0x21, 0x0c, 0xee, 0x15, 0xb4, 0xa9, 0x48, 0xd2, 0x19, 0x6a, 0x9c, 0x44,
	0x32, 0xa5, 0xaa, 0x77, 0xab, 0x11, 0xa3, 0x9b, 0x33, 0x46, 0x0b, 0xb9, 0x30, 0xa5, 0x8a, 0x20,
	0xdc, 0x2d, 0xa0, 0x94, 0x8e, 0x65, 0x0e, 0xda, 0x6e, 0x04, 0xea, 0xe4, 0xa0, 0x17, 0x4e, 0xac,
	0x02, 0x83, 0xb3, 0x38, 0x55, 0x38, 0x89, 0x34, 0x4b, 0xb0, 0xb7, 0x73, 0x43, 0xcc, 0x85, 0x13,
	0x7b, 0xc9, 0x12, 0x24, 0x12, 0x3e, 0x2a, 0x66, 0x83, 0x72, 0x8e, 0xb2, 0x4c, 0x7b, 0xa7, 0x11,
	0xad, 0xbf, 0x4c, 0xca, 0x6a, 0x16, 0x99, 0x53, 0xe8, 0x15, 0x98, 0x12, 0x7f, 0xcd, 0x50, 0xe9,
	0xe8, 0xcd, 0x95, 0x46, 0xd5, 0xbb, 0xdd, 0x88, 0x76, 0x27, 0xa7, 0x85, 0x4e, 0xed, 0xdc, 0x88,
	0xad, 0xb8, 0x42, 0xa2, 0x4a, 0x05, 0x57, 0xe8, 0x49, 0xef, 0xde, 0xd0, 0x15, 0xa1, 0x97, 0x73,
	0xa8, 0xea, 0x9c, 0xa8, 0xc8, 0xb8, 0xee, 0xed, 0x6e, 0x26, 0xa7, 0x91, 0x11, 0xab, 0xcb, 0xc9,
	0x91, 0xf6, 0x36, 0x94, 0x93, 0x43, 0xf9, 0x86, 0x5d, 0x98, 0xc2, 0x35, 0x2c, 0x34, 0x6e, 0x58,
	0x6f, 0x84, 0x52, 0xc3, 0x7a, 0xed, 0x95, 0x86, 0xdd, 0x6f, 0x9c, 0x86, 0x63, 0x94, 0x1b, 0xd6,
	0x5b, 0xdc, 0xa3, 0xaa, 0x2c, 0x7e, 0xd0, 0xd8, 0xe2, 0x8e, 0x57, 0x6b, 0x71, 0x0f, 0x2b, 0x5b,
	0xfc, 0xb0, 0xb1, 0x1d, 0x1c, 0xa7, 0xca, 0xe2, 0x39, 0xa8, 0x64, 0xf1, 0xa3, 0x1b, 0xd6, 0xb1,
	0xd2, 0xe2, 0x2b, 0x39, 0x39, 0xe3, 0xb5, 0x37, 0x93, 0x53, 0xc9, 0xe2, 0xab, 0x39, 0x39, 0xd2,
	0xf1, 0x86, 0x72, 0x72, 0x28, 0xff, 0x96, 0xad, 0x32, 0xc5, 0x49, 0xe3, 0xb7, 0xec, 0x9a, 0x1d,
	0xfa, 0x7f, 0xb4, 0x60, 0xc7, 0x2c, 0x93, 0x09, 0x74, 0x13, 0x8c, 0x55, 0x26, 0xd1, 0x28, 0x45,
	0x13, 0x54, 0x54, 0xb2, 0x54, 0x0b, 0xe9, 0x3f, 0x84, 0x9f, 0xd7, 0xe1, 0x9e, 0x2f, 0x77, 0x3d,
	0xc9, 0x37, 0x85, 0x77, 0x92, 0xaa, 0x69, 0xf2, 0x2d, 0xb4, 0xe7, 0x0c, 0x2f, 0x8b, 0xf2, 0xee,
	0x1b, 0xf8, 0xa0, 0x4e, 0xde, 0x1c, 0xae, 0xa0, 0x7b, 0x34, 0x2f, 0x8d, 0xc9, 0x29, 0xec, 0x98,
	0x19, 0xff, 0x81, 0xfb, 0xe0, 0x3a, 0x95, 0xd0, 0x46, 0x06, 0xbf, 0xb7, 0xe0, 0xc0, 0x96, 0xc5,
	0x3f, 0x59, 0xf2, 0x21, 0x80, 0x3d, 0x13, 0x8f, 0x13, 0x34, 0x9f, 0xfd, 0xed, 0x87, 0x7b, 0xe1,
	0x9e, 0x99, 0xf9, 0xc6, 0x4c, 0x90, 0xc7, 0x70, 0x52, 0x2c, 0x8c, 0x8b, 0xba, 0x65, 0xa3, 0x8e,
	0x0b, 0x0b, 0x2e, 0xf8, 0x15, 0x7c, 0x3a, 0x11, 0x5c, 0x47, 0x8c, 0xd3, 0x59, 0x36, 0xc1, 0x42,
	0x9e, 0x2a, 0x62, 0x3c, 0x7a, 0xcb, 0xa4, 0x5a, 0xbe, 0x1a, 0xed, 0xa1, 0x77, 0xc3, 0xfb, 0x66,
	0xc7, 0x33, 0xb7, 0x61, 0x99, 0x99, 0x7a, 0xc6, 0x9f, 0x9a, 0xe8, 0x85, 0x2b, 0x82, 0x9f, 0xe0,
	0xd0, 0x1f, 0xdb, 0x4d, 0x90, 0xaf, 0xc1, 0x16, 0x23, 0x97, 0x73, 0x67, 0xdf, 0x3f, 0xfb, 0xe4,
	0xda, 0x22, 0xf8, 0xe0, 0xf0, 0x70, 0x5e, 0x18, 0xa9, 0xe0, 0x08, 0x0e, 0x5e, 0xca, 0x98, 0xa2,
	0x2f, 0x4a, 0xd0, 0x86, 0x43, 0x3f, 0xf6, 0xf8, 0x47, 0xf0, 0xde, 0xf3, 0xfc, 0xc6, 0xf6, 0x24,
	0xd6, 0xf1, 0x58, 0x8a, 0x2c, 0x25, 0x04, 0x76, 0x4c, 0x45, 0xac, 0x49, 0xf6, 0x42, 0xfb, 0x7f,
	0xf0, 0x15, 0x74, 0x46, 0x99, 0xd2, 0x22, 0x29, 0x6f, 0x20, 0xa7, 0xb0, 0x4b, 0x05, 0xd7, 0xc8,
	0xf5, 0xe2, 0x76, 0xd5, 0x59, 0x1c, 0x75, 0x71, 0x59, 0x1b, 0x7c, 0xc9, 0xaf, 0xc2, 0x3c, 0x2a,
	0xf8, 0xb3, 0x05, 0x07, 0xc5, 0x53, 0xff, 0x6f, 0x5c, 0x7a, 0xf6, 0xd7, 0x36, 0xc0, 0xb2, 0x7c,
	0xe4, 0x47, 0xe8, 0x8c, 0x51, 0xaf, 0xdf, 0x5c, 0xbb, 0x6b, 0x05, 0xbc, 0x30, 0xb7, 0xdd, 0xfe,
	0xa3, 0x7f, 0xfd, 0x72, 0x08, 0xb6, 0xc8, 0x2b, 0xd8, 0x1d, 0xa3, 0x76, 0x82, 0xb5, 0xe6, 0x29,
	0xb6, 0x4c, 0xff, 0xfe, 0x3f, 0x44, 0x79, 0xcf, 0x6c, 0x91, 0x9f, 0x01, 0x7e, 0x88, 0x35, 0xfd,
	0xe5, 0xbf, 0x10, 0x3f, 0x6d, 0x91, 0x18, 0x8e, 0xc7, 0xb8, 0xb8, 0x85, 0x58, 0xbf, 0x5e, 0x03,
	0x29, 0xfa, 0xbb, 0x1e, 0x52, 0x76, 0xfd, 0x16, 0x99, 0xc3, 0x5d, 0x53, 0xf9, 0x2a, 0x3f, 0x3f,
	0xae, 0x35, 0xdb, 0x7a, 0xa3, 0xf4, 0x3f, 0xab, 0x7d, 0x22, 0x15, 0xd2, 0xc1, 0xd6, 0xf9, 0x53,
	0x78, 0x40, 0x45, 0x52, 0xb7, 0xc9, 0xfe, 0x24, 0x72, 0x0f, 0xfd, 0x7c, 0x7f, 0x64, 0x7f, 0x19,
	0x7d, 0x67, 0x06, 0xaf, 0xdb, 0x53, 0xe4, 0x28, 0x63, 0x73, 0x61, 0x71, 0x96, 0xb8, 0x6d, 0xff,
	0x7c, 0xf1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x8d, 0xbd, 0x66, 0x78, 0x0d, 0x00, 0x00,
}