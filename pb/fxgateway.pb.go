// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fxgateway.proto

package pb

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_fxgateway_7f71f224744a556d, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Message struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_fxgateway_7f71f224744a556d, []int{1}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type CreateFunctionRequest struct {
	Service              string             `protobuf:"bytes,1,opt,name=Service,proto3" json:"Service,omitempty"`
	Image                string             `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
	EnvVars              map[string]string  `protobuf:"bytes,3,rep,name=EnvVars,proto3" json:"EnvVars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string  `protobuf:"bytes,4,rep,name=Labels,proto3" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations          map[string]string  `protobuf:"bytes,5,rep,name=Annotations,proto3" json:"Annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Constraints          []string           `protobuf:"bytes,6,rep,name=Constraints,proto3" json:"Constraints,omitempty"`
	Secrets              []string           `protobuf:"bytes,7,rep,name=Secrets,proto3" json:"Secrets,omitempty"`
	RegistryAuth         string             `protobuf:"bytes,8,opt,name=RegistryAuth,proto3" json:"RegistryAuth,omitempty"`
	Limits               *FunctionResources `protobuf:"bytes,9,opt,name=Limits,proto3" json:"Limits,omitempty"`
	Requests             *FunctionResources `protobuf:"bytes,10,opt,name=Requests,proto3" json:"Requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateFunctionRequest) Reset()         { *m = CreateFunctionRequest{} }
func (m *CreateFunctionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFunctionRequest) ProtoMessage()    {}
func (*CreateFunctionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fxgateway_7f71f224744a556d, []int{2}
}
func (m *CreateFunctionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFunctionRequest.Unmarshal(m, b)
}
func (m *CreateFunctionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFunctionRequest.Marshal(b, m, deterministic)
}
func (dst *CreateFunctionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFunctionRequest.Merge(dst, src)
}
func (m *CreateFunctionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFunctionRequest.Size(m)
}
func (m *CreateFunctionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFunctionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFunctionRequest proto.InternalMessageInfo

func (m *CreateFunctionRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *CreateFunctionRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *CreateFunctionRequest) GetEnvVars() map[string]string {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

func (m *CreateFunctionRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateFunctionRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *CreateFunctionRequest) GetConstraints() []string {
	if m != nil {
		return m.Constraints
	}
	return nil
}

func (m *CreateFunctionRequest) GetSecrets() []string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *CreateFunctionRequest) GetRegistryAuth() string {
	if m != nil {
		return m.RegistryAuth
	}
	return ""
}

func (m *CreateFunctionRequest) GetLimits() *FunctionResources {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *CreateFunctionRequest) GetRequests() *FunctionResources {
	if m != nil {
		return m.Requests
	}
	return nil
}

type DeleteFunctionRequest struct {
	FunctionName         string   `protobuf:"bytes,1,opt,name=FunctionName,proto3" json:"FunctionName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFunctionRequest) Reset()         { *m = DeleteFunctionRequest{} }
func (m *DeleteFunctionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFunctionRequest) ProtoMessage()    {}
func (*DeleteFunctionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fxgateway_7f71f224744a556d, []int{3}
}
func (m *DeleteFunctionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFunctionRequest.Unmarshal(m, b)
}
func (m *DeleteFunctionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFunctionRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteFunctionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFunctionRequest.Merge(dst, src)
}
func (m *DeleteFunctionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFunctionRequest.Size(m)
}
func (m *DeleteFunctionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFunctionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFunctionRequest proto.InternalMessageInfo

func (m *DeleteFunctionRequest) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

type ScaleServiceRequest struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	Replicas             uint64   `protobuf:"varint,2,opt,name=Replicas,proto3" json:"Replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleServiceRequest) Reset()         { *m = ScaleServiceRequest{} }
func (m *ScaleServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ScaleServiceRequest) ProtoMessage()    {}
func (*ScaleServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fxgateway_7f71f224744a556d, []int{4}
}
func (m *ScaleServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleServiceRequest.Unmarshal(m, b)
}
func (m *ScaleServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleServiceRequest.Marshal(b, m, deterministic)
}
func (dst *ScaleServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleServiceRequest.Merge(dst, src)
}
func (m *ScaleServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ScaleServiceRequest.Size(m)
}
func (m *ScaleServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleServiceRequest proto.InternalMessageInfo

func (m *ScaleServiceRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ScaleServiceRequest) GetReplicas() uint64 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

type InvokeServiceRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=Service,proto3" json:"Service,omitempty"`
	Input                []byte   `protobuf:"bytes,2,opt,name=Input,proto3" json:"Input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeServiceRequest) Reset()         { *m = InvokeServiceRequest{} }
func (m *InvokeServiceRequest) String() string { return proto.CompactTextString(m) }
func (*InvokeServiceRequest) ProtoMessage()    {}
func (*InvokeServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fxgateway_7f71f224744a556d, []int{5}
}
func (m *InvokeServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeServiceRequest.Unmarshal(m, b)
}
func (m *InvokeServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeServiceRequest.Marshal(b, m, deterministic)
}
func (dst *InvokeServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeServiceRequest.Merge(dst, src)
}
func (m *InvokeServiceRequest) XXX_Size() int {
	return xxx_messageInfo_InvokeServiceRequest.Size(m)
}
func (m *InvokeServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeServiceRequest proto.InternalMessageInfo

func (m *InvokeServiceRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *InvokeServiceRequest) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

type FunctionRequest struct {
	FunctionName         string   `protobuf:"bytes,1,opt,name=FunctionName,proto3" json:"FunctionName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FunctionRequest) Reset()         { *m = FunctionRequest{} }
func (m *FunctionRequest) String() string { return proto.CompactTextString(m) }
func (*FunctionRequest) ProtoMessage()    {}
func (*FunctionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fxgateway_7f71f224744a556d, []int{6}
}
func (m *FunctionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionRequest.Unmarshal(m, b)
}
func (m *FunctionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionRequest.Marshal(b, m, deterministic)
}
func (dst *FunctionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionRequest.Merge(dst, src)
}
func (m *FunctionRequest) XXX_Size() int {
	return xxx_messageInfo_FunctionRequest.Size(m)
}
func (m *FunctionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionRequest proto.InternalMessageInfo

func (m *FunctionRequest) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

type Functions struct {
	Functions            []*Function `protobuf:"bytes,1,rep,name=Functions,proto3" json:"Functions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Functions) Reset()         { *m = Functions{} }
func (m *Functions) String() string { return proto.CompactTextString(m) }
func (*Functions) ProtoMessage()    {}
func (*Functions) Descriptor() ([]byte, []int) {
	return fileDescriptor_fxgateway_7f71f224744a556d, []int{7}
}
func (m *Functions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Functions.Unmarshal(m, b)
}
func (m *Functions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Functions.Marshal(b, m, deterministic)
}
func (dst *Functions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Functions.Merge(dst, src)
}
func (m *Functions) XXX_Size() int {
	return xxx_messageInfo_Functions.Size(m)
}
func (m *Functions) XXX_DiscardUnknown() {
	xxx_messageInfo_Functions.DiscardUnknown(m)
}

var xxx_messageInfo_Functions proto.InternalMessageInfo

func (m *Functions) GetFunctions() []*Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

type Function struct {
	Name                 string            `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Image                string            `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
	InvocationCount      uint64            `protobuf:"varint,3,opt,name=InvocationCount,proto3" json:"InvocationCount,omitempty"`
	Replicas             uint64            `protobuf:"varint,4,opt,name=Replicas,proto3" json:"Replicas,omitempty"`
	AvailableReplicas    uint64            `protobuf:"varint,5,opt,name=AvailableReplicas,proto3" json:"AvailableReplicas,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,6,rep,name=Annotations,proto3" json:"Annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,7,rep,name=Labels,proto3" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_fxgateway_7f71f224744a556d, []int{8}
}
func (m *Function) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Function.Unmarshal(m, b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Function.Marshal(b, m, deterministic)
}
func (dst *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(dst, src)
}
func (m *Function) XXX_Size() int {
	return xxx_messageInfo_Function.Size(m)
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Function) GetInvocationCount() uint64 {
	if m != nil {
		return m.InvocationCount
	}
	return 0
}

func (m *Function) GetReplicas() uint64 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *Function) GetAvailableReplicas() uint64 {
	if m != nil {
		return m.AvailableReplicas
	}
	return 0
}

func (m *Function) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Function) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type FunctionResources struct {
	Memory               string   `protobuf:"bytes,1,opt,name=Memory,proto3" json:"Memory,omitempty"`
	CPU                  string   `protobuf:"bytes,2,opt,name=CPU,proto3" json:"CPU,omitempty"`
	GPU                  string   `protobuf:"bytes,3,opt,name=GPU,proto3" json:"GPU,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FunctionResources) Reset()         { *m = FunctionResources{} }
func (m *FunctionResources) String() string { return proto.CompactTextString(m) }
func (*FunctionResources) ProtoMessage()    {}
func (*FunctionResources) Descriptor() ([]byte, []int) {
	return fileDescriptor_fxgateway_7f71f224744a556d, []int{9}
}
func (m *FunctionResources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionResources.Unmarshal(m, b)
}
func (m *FunctionResources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionResources.Marshal(b, m, deterministic)
}
func (dst *FunctionResources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionResources.Merge(dst, src)
}
func (m *FunctionResources) XXX_Size() int {
	return xxx_messageInfo_FunctionResources.Size(m)
}
func (m *FunctionResources) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionResources.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionResources proto.InternalMessageInfo

func (m *FunctionResources) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *FunctionResources) GetCPU() string {
	if m != nil {
		return m.CPU
	}
	return ""
}

func (m *FunctionResources) GetGPU() string {
	if m != nil {
		return m.GPU
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*CreateFunctionRequest)(nil), "pb.CreateFunctionRequest")
	proto.RegisterMapType((map[string]string)(nil), "pb.CreateFunctionRequest.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "pb.CreateFunctionRequest.EnvVarsEntry")
	proto.RegisterMapType((map[string]string)(nil), "pb.CreateFunctionRequest.LabelsEntry")
	proto.RegisterType((*DeleteFunctionRequest)(nil), "pb.DeleteFunctionRequest")
	proto.RegisterType((*ScaleServiceRequest)(nil), "pb.ScaleServiceRequest")
	proto.RegisterType((*InvokeServiceRequest)(nil), "pb.InvokeServiceRequest")
	proto.RegisterType((*FunctionRequest)(nil), "pb.FunctionRequest")
	proto.RegisterType((*Functions)(nil), "pb.Functions")
	proto.RegisterType((*Function)(nil), "pb.Function")
	proto.RegisterMapType((map[string]string)(nil), "pb.Function.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "pb.Function.LabelsEntry")
	proto.RegisterType((*FunctionResources)(nil), "pb.FunctionResources")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FxGatewayClient is the client API for FxGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FxGatewayClient interface {
	Invoke(ctx context.Context, in *InvokeServiceRequest, opts ...grpc.CallOption) (*Message, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Functions, error)
	Deploy(ctx context.Context, in *CreateFunctionRequest, opts ...grpc.CallOption) (*Message, error)
	Delete(ctx context.Context, in *DeleteFunctionRequest, opts ...grpc.CallOption) (*Message, error)
	Update(ctx context.Context, in *CreateFunctionRequest, opts ...grpc.CallOption) (*Message, error)
	GetMeta(ctx context.Context, in *FunctionRequest, opts ...grpc.CallOption) (*Function, error)
	GetLog(ctx context.Context, in *FunctionRequest, opts ...grpc.CallOption) (*Message, error)
	ReplicaUpdate(ctx context.Context, in *ScaleServiceRequest, opts ...grpc.CallOption) (*Message, error)
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error)
	HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error)
}

type fxGatewayClient struct {
	cc *grpc.ClientConn
}

func NewFxGatewayClient(cc *grpc.ClientConn) FxGatewayClient {
	return &fxGatewayClient{cc}
}

func (c *fxGatewayClient) Invoke(ctx context.Context, in *InvokeServiceRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.FxGateway/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxGatewayClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Functions, error) {
	out := new(Functions)
	err := c.cc.Invoke(ctx, "/pb.FxGateway/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxGatewayClient) Deploy(ctx context.Context, in *CreateFunctionRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.FxGateway/Deploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxGatewayClient) Delete(ctx context.Context, in *DeleteFunctionRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.FxGateway/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxGatewayClient) Update(ctx context.Context, in *CreateFunctionRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.FxGateway/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxGatewayClient) GetMeta(ctx context.Context, in *FunctionRequest, opts ...grpc.CallOption) (*Function, error) {
	out := new(Function)
	err := c.cc.Invoke(ctx, "/pb.FxGateway/GetMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxGatewayClient) GetLog(ctx context.Context, in *FunctionRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.FxGateway/GetLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxGatewayClient) ReplicaUpdate(ctx context.Context, in *ScaleServiceRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.FxGateway/ReplicaUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxGatewayClient) Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.FxGateway/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxGatewayClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.FxGateway/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FxGatewayServer is the server API for FxGateway service.
type FxGatewayServer interface {
	Invoke(context.Context, *InvokeServiceRequest) (*Message, error)
	List(context.Context, *Empty) (*Functions, error)
	Deploy(context.Context, *CreateFunctionRequest) (*Message, error)
	Delete(context.Context, *DeleteFunctionRequest) (*Message, error)
	Update(context.Context, *CreateFunctionRequest) (*Message, error)
	GetMeta(context.Context, *FunctionRequest) (*Function, error)
	GetLog(context.Context, *FunctionRequest) (*Message, error)
	ReplicaUpdate(context.Context, *ScaleServiceRequest) (*Message, error)
	Info(context.Context, *Empty) (*Message, error)
	HealthCheck(context.Context, *Empty) (*Message, error)
}

func RegisterFxGatewayServer(s *grpc.Server, srv FxGatewayServer) {
	s.RegisterService(&_FxGateway_serviceDesc, srv)
}

func _FxGateway_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxGatewayServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FxGateway/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxGatewayServer).Invoke(ctx, req.(*InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxGateway_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxGatewayServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FxGateway/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxGatewayServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxGateway_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxGatewayServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FxGateway/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxGatewayServer).Deploy(ctx, req.(*CreateFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxGateway_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxGatewayServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FxGateway/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxGatewayServer).Delete(ctx, req.(*DeleteFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxGateway_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxGatewayServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FxGateway/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxGatewayServer).Update(ctx, req.(*CreateFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxGateway_GetMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxGatewayServer).GetMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FxGateway/GetMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxGatewayServer).GetMeta(ctx, req.(*FunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxGateway_GetLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxGatewayServer).GetLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FxGateway/GetLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxGatewayServer).GetLog(ctx, req.(*FunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxGateway_ReplicaUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxGatewayServer).ReplicaUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FxGateway/ReplicaUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxGatewayServer).ReplicaUpdate(ctx, req.(*ScaleServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxGateway_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxGatewayServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FxGateway/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxGatewayServer).Info(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxGateway_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxGatewayServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FxGateway/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxGatewayServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _FxGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FxGateway",
	HandlerType: (*FxGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _FxGateway_Invoke_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FxGateway_List_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _FxGateway_Deploy_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FxGateway_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FxGateway_Update_Handler,
		},
		{
			MethodName: "GetMeta",
			Handler:    _FxGateway_GetMeta_Handler,
		},
		{
			MethodName: "GetLog",
			Handler:    _FxGateway_GetLog_Handler,
		},
		{
			MethodName: "ReplicaUpdate",
			Handler:    _FxGateway_ReplicaUpdate_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _FxGateway_Info_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _FxGateway_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fxgateway.proto",
}

func init() { proto.RegisterFile("fxgateway.proto", fileDescriptor_fxgateway_7f71f224744a556d) }

var fileDescriptor_fxgateway_7f71f224744a556d = []byte{
	// 855 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xeb, 0x44,
	0x10, 0x26, 0x7f, 0x4e, 0x32, 0xce, 0xa1, 0xed, 0xf6, 0x07, 0x1f, 0x9f, 0x83, 0x48, 0x8d, 0x0a,
	0x51, 0xa0, 0x09, 0x2d, 0xa0, 0x42, 0x2b, 0x7e, 0xaa, 0xb4, 0x0d, 0x15, 0x49, 0xa9, 0x5c, 0x05,
	0xb8, 0xdd, 0x98, 0xad, 0x63, 0xc5, 0xb1, 0x8d, 0x77, 0x13, 0x6a, 0x10, 0x37, 0xbc, 0x00, 0x17,
	0x48, 0xbc, 0x18, 0x4f, 0x80, 0xc4, 0x83, 0xa0, 0x5d, 0xdb, 0xa9, 0x9d, 0x38, 0x85, 0x9e, 0xde,
	0xed, 0xcc, 0xce, 0xf7, 0xed, 0xec, 0xb7, 0x33, 0x63, 0xc3, 0xda, 0xed, 0x9d, 0x89, 0x19, 0xf9,
	0x09, 0x07, 0x2d, 0xcf, 0x77, 0x99, 0x8b, 0xf2, 0xde, 0x50, 0x7d, 0x69, 0xba, 0xae, 0x69, 0x93,
	0x36, 0xf6, 0xac, 0x36, 0x76, 0x1c, 0x97, 0x61, 0x66, 0xb9, 0x0e, 0x0d, 0x23, 0xb4, 0x32, 0x94,
	0xce, 0x27, 0x1e, 0x0b, 0xb4, 0x17, 0x50, 0xee, 0x13, 0x4a, 0xb1, 0x49, 0xd0, 0x3a, 0x14, 0xfa,
	0xd4, 0x54, 0x72, 0xf5, 0x5c, 0xa3, 0xaa, 0xf3, 0xa5, 0xf6, 0x7b, 0x09, 0xb6, 0x3b, 0x3e, 0xc1,
	0x8c, 0x5c, 0x4c, 0x1d, 0x83, 0xe3, 0x75, 0xf2, 0xe3, 0x94, 0x50, 0x86, 0x14, 0x28, 0xdf, 0x10,
	0x7f, 0x66, 0x19, 0x24, 0x8a, 0x8f, 0x4d, 0xb4, 0x05, 0xa5, 0xcb, 0x09, 0x36, 0x89, 0x92, 0x17,
	0xfe, 0xd0, 0x40, 0x5f, 0x42, 0xf9, 0xdc, 0x99, 0x7d, 0x8b, 0x7d, 0xaa, 0x14, 0xea, 0x85, 0x86,
	0x7c, 0xf8, 0x4e, 0xcb, 0x1b, 0xb6, 0x32, 0xb9, 0x5b, 0x51, 0xe0, 0xb9, 0xc3, 0xfc, 0x40, 0x8f,
	0x61, 0xe8, 0x33, 0x90, 0x7a, 0x78, 0x48, 0x6c, 0xaa, 0x14, 0x05, 0xc1, 0xde, 0x6a, 0x82, 0x30,
	0x2e, 0xc4, 0x47, 0x20, 0xd4, 0x03, 0xf9, 0xf4, 0x5e, 0x05, 0xa5, 0x24, 0x38, 0x9a, 0xab, 0x39,
	0x12, 0xc1, 0x21, 0x51, 0x12, 0x8e, 0xea, 0x20, 0x77, 0x5c, 0x87, 0x32, 0x1f, 0x5b, 0x0e, 0xa3,
	0x8a, 0x54, 0x2f, 0x34, 0xaa, 0x7a, 0xd2, 0x15, 0x0a, 0x64, 0xf8, 0x84, 0x51, 0xa5, 0x2c, 0x76,
	0x63, 0x13, 0x69, 0x50, 0xd3, 0x89, 0x69, 0x51, 0xe6, 0x07, 0xa7, 0x53, 0x36, 0x52, 0x2a, 0x42,
	0xa7, 0x94, 0x0f, 0xed, 0x83, 0xd4, 0xb3, 0x26, 0x16, 0xa3, 0x4a, 0xb5, 0x9e, 0x6b, 0xc8, 0x87,
	0xdb, 0x3c, 0xd1, 0xfb, 0x14, 0xa9, 0x3b, 0xf5, 0x0d, 0x42, 0xf5, 0x28, 0x08, 0x1d, 0x40, 0x25,
	0xca, 0x9b, 0x2a, 0xf0, 0x10, 0x60, 0x1e, 0xa6, 0x1e, 0x43, 0x2d, 0xa9, 0x33, 0x7f, 0xfc, 0x31,
	0x09, 0xe2, 0xc7, 0x1f, 0x93, 0x80, 0x3f, 0xe4, 0x0c, 0xdb, 0xd3, 0xf9, 0x43, 0x0a, 0xe3, 0x38,
	0xff, 0x49, 0x4e, 0xfd, 0x14, 0xe4, 0x84, 0xc4, 0x8f, 0x82, 0x7e, 0x0e, 0xeb, 0x8b, 0xca, 0x3e,
	0x06, 0xaf, 0x9d, 0xc0, 0xf6, 0x19, 0xb1, 0xc9, 0x72, 0x41, 0x6a, 0x50, 0x8b, 0x5d, 0x57, 0x78,
	0x12, 0x57, 0x65, 0xca, 0xa7, 0xdd, 0xc0, 0xe6, 0x8d, 0x81, 0x6d, 0x12, 0x95, 0x6a, 0x0c, 0xad,
	0x83, 0x1c, 0x79, 0x12, 0xc8, 0xa4, 0x0b, 0xa9, 0x5c, 0x5f, 0xcf, 0xb6, 0x0c, 0x4c, 0x45, 0x4a,
	0x45, 0x7d, 0x6e, 0x6b, 0x17, 0xb0, 0x75, 0xe9, 0xcc, 0xdc, 0xf1, 0x22, 0xeb, 0xc3, 0x1d, 0xe2,
	0x78, 0x53, 0x26, 0xa8, 0x6a, 0x7a, 0x68, 0x68, 0x1f, 0xc3, 0xda, 0xab, 0xdc, 0xe9, 0x08, 0xaa,
	0xb1, 0x4d, 0x51, 0x33, 0x61, 0x28, 0x39, 0x51, 0xe2, 0xb5, 0x54, 0x21, 0xdc, 0x6f, 0x6b, 0x7f,
	0x16, 0xa0, 0x12, 0x5b, 0x08, 0x41, 0x31, 0x71, 0x82, 0x58, 0xaf, 0x68, 0xe4, 0x06, 0xac, 0xf1,
	0xeb, 0x1a, 0xe2, 0x01, 0x3b, 0xee, 0xd4, 0x61, 0x4a, 0x41, 0x28, 0xb2, 0xe8, 0x4e, 0x89, 0x56,
	0x4c, 0x8b, 0x86, 0xde, 0x87, 0x8d, 0xd3, 0x19, 0xb6, 0x6c, 0x3c, 0xb4, 0xc9, 0x3c, 0xa8, 0x24,
	0x82, 0x96, 0x37, 0xd0, 0x17, 0xe9, 0xde, 0x95, 0xc4, 0xc5, 0xde, 0x4c, 0x5e, 0xec, 0x3f, 0xda,
	0xf5, 0x83, 0xf9, 0xec, 0x28, 0x0b, 0xac, 0x92, 0xc2, 0x66, 0x8c, 0x8b, 0xa7, 0xd6, 0xe9, 0x13,
	0x5a, 0x44, 0xfb, 0x06, 0x36, 0x96, 0x1a, 0x17, 0xed, 0x80, 0xd4, 0x27, 0x13, 0xd7, 0x8f, 0x39,
	0x22, 0x8b, 0x13, 0x77, 0xae, 0x07, 0x11, 0x09, 0x5f, 0x72, 0x4f, 0xf7, 0x7a, 0x20, 0x1e, 0xa5,
	0xaa, 0xf3, 0xe5, 0xe1, 0xdf, 0x25, 0xa8, 0x5e, 0xdc, 0x75, 0xc3, 0x2f, 0x04, 0x3a, 0x00, 0x29,
	0xac, 0x57, 0x24, 0x54, 0xc8, 0xaa, 0x5d, 0x55, 0xe6, 0x3b, 0xd1, 0x67, 0x41, 0x7b, 0x0d, 0x9d,
	0x40, 0xb1, 0x67, 0x51, 0x86, 0xaa, 0xdc, 0x2d, 0x3e, 0x1b, 0xea, 0xb3, 0xa4, 0x82, 0x54, 0x7b,
	0xfe, 0xdb, 0x5f, 0xff, 0xfc, 0x91, 0xdf, 0x44, 0x1b, 0x6d, 0x1a, 0x50, 0x46, 0x26, 0xed, 0xdb,
	0x79, 0x4d, 0x5e, 0x81, 0x74, 0x46, 0x3c, 0xdb, 0x0d, 0xd0, 0xf3, 0x95, 0xd3, 0x36, 0x7d, 0xe0,
	0x4b, 0x41, 0xb6, 0xa3, 0x2d, 0x93, 0x1d, 0xe7, 0x9a, 0xe8, 0x6b, 0xce, 0xc7, 0x27, 0x40, 0xc8,
	0x97, 0x39, 0x0d, 0xd2, 0x7c, 0x51, 0x72, 0xcd, 0xec, 0xe4, 0x06, 0xde, 0x0f, 0x38, 0x26, 0xfb,
	0xff, 0xc9, 0xa9, 0xd9, 0xc9, 0x0d, 0xa0, 0xdc, 0x25, 0xac, 0x4f, 0x18, 0x46, 0x9b, 0xe9, 0x09,
	0x1c, 0x52, 0xa5, 0xba, 0x51, 0x7b, 0x57, 0x70, 0xed, 0xa2, 0xb7, 0x16, 0xb9, 0xda, 0xbf, 0x24,
	0x7b, 0xfc, 0x57, 0xf4, 0x1d, 0x48, 0x5d, 0xc2, 0x7a, 0xae, 0x99, 0xcd, 0x9a, 0x4a, 0xf0, 0x3d,
	0x41, 0xba, 0x87, 0xde, 0x5e, 0x24, 0xdd, 0xb7, 0x5d, 0x73, 0x91, 0xf8, 0x7b, 0x78, 0x16, 0x75,
	0x59, 0x24, 0xc3, 0x1b, 0x9c, 0x2a, 0x63, 0x48, 0xa6, 0xcf, 0xd8, 0x15, 0x67, 0xbc, 0x50, 0x77,
	0xe2, 0x33, 0x28, 0x47, 0xec, 0xc7, 0x27, 0x71, 0x25, 0x3e, 0x82, 0xe2, 0xa5, 0x73, 0xeb, 0x26,
	0x6b, 0x26, 0x45, 0xb1, 0x25, 0x28, 0x5e, 0x47, 0xb5, 0x98, 0xc2, 0xe2, 0xd1, 0x47, 0x20, 0x7f,
	0x45, 0xb0, 0xcd, 0x46, 0x9d, 0x11, 0x31, 0xc6, 0x2b, 0xc1, 0xeb, 0x02, 0x0c, 0xa8, 0xd2, 0x1e,
	0x89, 0xe8, 0x9f, 0x87, 0x92, 0xf8, 0xad, 0xf9, 0xf0, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40,
	0x63, 0x84, 0xda, 0x0b, 0x09, 0x00, 0x00,
}