// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/metrics.proto

package metrics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MetricType int32

const (
	MetricType_COUNTER MetricType = 0
	MetricType_GAUGE   MetricType = 1
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0: "COUNTER",
		1: "GAUGE",
	}
	MetricType_value = map[string]int32{
		"COUNTER": 0,
		"GAUGE":   1,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_metrics_proto_enumTypes[0].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_proto_metrics_proto_enumTypes[0]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_proto_metrics_proto_rawDescGZIP(), []int{0}
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricType MetricType `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=proto.MetricType" json:"metric_type,omitempty"`
	// Types that are assignable to Value:
	//
	//	*Metric_Counter
	//	*Metric_Gauge
	Value isMetric_Value `protobuf_oneof:"value"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_proto_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metric) GetMetricType() MetricType {
	if x != nil {
		return x.MetricType
	}
	return MetricType_COUNTER
}

func (m *Metric) GetValue() isMetric_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Metric) GetCounter() int64 {
	if x, ok := x.GetValue().(*Metric_Counter); ok {
		return x.Counter
	}
	return 0
}

func (x *Metric) GetGauge() float64 {
	if x, ok := x.GetValue().(*Metric_Gauge); ok {
		return x.Gauge
	}
	return 0
}

type isMetric_Value interface {
	isMetric_Value()
}

type Metric_Counter struct {
	Counter int64 `protobuf:"varint,3,opt,name=counter,proto3,oneof"`
}

type Metric_Gauge struct {
	Gauge float64 `protobuf:"fixed64,4,opt,name=gauge,proto3,oneof"`
}

func (*Metric_Counter) isMetric_Value() {}

func (*Metric_Gauge) isMetric_Value() {}

type GetMetricRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type MetricType `protobuf:"varint,2,opt,name=type,proto3,enum=proto.MetricType" json:"type,omitempty"`
}

func (x *GetMetricRequest) Reset() {
	*x = GetMetricRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricRequest) ProtoMessage() {}

func (x *GetMetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricRequest.ProtoReflect.Descriptor instead.
func (*GetMetricRequest) Descriptor() ([]byte, []int) {
	return file_proto_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *GetMetricRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetMetricRequest) GetType() MetricType {
	if x != nil {
		return x.Type
	}
	return MetricType_COUNTER
}

type GetMetricResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric *Metric `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
}

func (x *GetMetricResponse) Reset() {
	*x = GetMetricResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricResponse) ProtoMessage() {}

func (x *GetMetricResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricResponse.ProtoReflect.Descriptor instead.
func (*GetMetricResponse) Descriptor() ([]byte, []int) {
	return file_proto_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *GetMetricResponse) GetMetric() *Metric {
	if x != nil {
		return x.Metric
	}
	return nil
}

type GetMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMetricsRequest) Reset() {
	*x = GetMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsRequest) ProtoMessage() {}

func (x *GetMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsRequest.ProtoReflect.Descriptor instead.
func (*GetMetricsRequest) Descriptor() ([]byte, []int) {
	return file_proto_metrics_proto_rawDescGZIP(), []int{3}
}

type GetMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *GetMetricsResponse) Reset() {
	*x = GetMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsResponse) ProtoMessage() {}

func (x *GetMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsResponse.ProtoReflect.Descriptor instead.
func (*GetMetricsResponse) Descriptor() ([]byte, []int) {
	return file_proto_metrics_proto_rawDescGZIP(), []int{4}
}

func (x *GetMetricsResponse) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type AddMetricRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric *Metric `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
}

func (x *AddMetricRequest) Reset() {
	*x = AddMetricRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metrics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMetricRequest) ProtoMessage() {}

func (x *AddMetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metrics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMetricRequest.ProtoReflect.Descriptor instead.
func (*AddMetricRequest) Descriptor() ([]byte, []int) {
	return file_proto_metrics_proto_rawDescGZIP(), []int{5}
}

func (x *AddMetricRequest) GetMetric() *Metric {
	if x != nil {
		return x.Metric
	}
	return nil
}

type AddMetricResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric *Metric `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
}

func (x *AddMetricResponse) Reset() {
	*x = AddMetricResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metrics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMetricResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMetricResponse) ProtoMessage() {}

func (x *AddMetricResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metrics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMetricResponse.ProtoReflect.Descriptor instead.
func (*AddMetricResponse) Descriptor() ([]byte, []int) {
	return file_proto_metrics_proto_rawDescGZIP(), []int{6}
}

func (x *AddMetricResponse) GetMetric() *Metric {
	if x != nil {
		return x.Metric
	}
	return nil
}

type AddMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *AddMetricsRequest) Reset() {
	*x = AddMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metrics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMetricsRequest) ProtoMessage() {}

func (x *AddMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metrics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMetricsRequest.ProtoReflect.Descriptor instead.
func (*AddMetricsRequest) Descriptor() ([]byte, []int) {
	return file_proto_metrics_proto_rawDescGZIP(), []int{7}
}

func (x *AddMetricsRequest) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type AddMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *AddMetricsResponse) Reset() {
	*x = AddMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metrics_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMetricsResponse) ProtoMessage() {}

func (x *AddMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metrics_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMetricsResponse.ProtoReflect.Descriptor instead.
func (*AddMetricsResponse) Descriptor() ([]byte, []int) {
	return file_proto_metrics_proto_rawDescGZIP(), []int{8}
}

func (x *AddMetricsResponse) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_proto_metrics_proto protoreflect.FileDescriptor

var file_proto_metrics_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a,
	0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x05, 0x67,
	0x61, 0x75, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x05, 0x67, 0x61,
	0x75, 0x67, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4d, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x39, 0x0a, 0x10, 0x41,
	0x64, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x22, 0x3a, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x22, 0x3c, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x22, 0x3d, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2a,
	0x24, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x41,
	0x55, 0x47, 0x45, 0x10, 0x01, 0x32, 0x98, 0x02, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x09, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64,
	0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_metrics_proto_rawDescOnce sync.Once
	file_proto_metrics_proto_rawDescData = file_proto_metrics_proto_rawDesc
)

func file_proto_metrics_proto_rawDescGZIP() []byte {
	file_proto_metrics_proto_rawDescOnce.Do(func() {
		file_proto_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_metrics_proto_rawDescData)
	})
	return file_proto_metrics_proto_rawDescData
}

var file_proto_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_metrics_proto_goTypes = []any{
	(MetricType)(0),            // 0: proto.MetricType
	(*Metric)(nil),             // 1: proto.Metric
	(*GetMetricRequest)(nil),   // 2: proto.GetMetricRequest
	(*GetMetricResponse)(nil),  // 3: proto.GetMetricResponse
	(*GetMetricsRequest)(nil),  // 4: proto.GetMetricsRequest
	(*GetMetricsResponse)(nil), // 5: proto.GetMetricsResponse
	(*AddMetricRequest)(nil),   // 6: proto.AddMetricRequest
	(*AddMetricResponse)(nil),  // 7: proto.AddMetricResponse
	(*AddMetricsRequest)(nil),  // 8: proto.AddMetricsRequest
	(*AddMetricsResponse)(nil), // 9: proto.AddMetricsResponse
}
var file_proto_metrics_proto_depIdxs = []int32{
	0,  // 0: proto.Metric.metric_type:type_name -> proto.MetricType
	0,  // 1: proto.GetMetricRequest.type:type_name -> proto.MetricType
	1,  // 2: proto.GetMetricResponse.metric:type_name -> proto.Metric
	1,  // 3: proto.GetMetricsResponse.metrics:type_name -> proto.Metric
	1,  // 4: proto.AddMetricRequest.metric:type_name -> proto.Metric
	1,  // 5: proto.AddMetricResponse.metric:type_name -> proto.Metric
	1,  // 6: proto.AddMetricsRequest.metrics:type_name -> proto.Metric
	1,  // 7: proto.AddMetricsResponse.metrics:type_name -> proto.Metric
	2,  // 8: proto.MetricsCollector.GetMetric:input_type -> proto.GetMetricRequest
	4,  // 9: proto.MetricsCollector.GetMetrics:input_type -> proto.GetMetricsRequest
	6,  // 10: proto.MetricsCollector.AddMetric:input_type -> proto.AddMetricRequest
	8,  // 11: proto.MetricsCollector.AddMetrics:input_type -> proto.AddMetricsRequest
	3,  // 12: proto.MetricsCollector.GetMetric:output_type -> proto.GetMetricResponse
	5,  // 13: proto.MetricsCollector.GetMetrics:output_type -> proto.GetMetricsResponse
	7,  // 14: proto.MetricsCollector.AddMetric:output_type -> proto.AddMetricResponse
	9,  // 15: proto.MetricsCollector.AddMetrics:output_type -> proto.AddMetricsResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_metrics_proto_init() }
func file_proto_metrics_proto_init() {
	if File_proto_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_metrics_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Metric); i {
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
		file_proto_metrics_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetMetricRequest); i {
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
		file_proto_metrics_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetMetricResponse); i {
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
		file_proto_metrics_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetMetricsRequest); i {
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
		file_proto_metrics_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetMetricsResponse); i {
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
		file_proto_metrics_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AddMetricRequest); i {
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
		file_proto_metrics_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*AddMetricResponse); i {
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
		file_proto_metrics_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*AddMetricsRequest); i {
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
		file_proto_metrics_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*AddMetricsResponse); i {
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
	file_proto_metrics_proto_msgTypes[0].OneofWrappers = []any{
		(*Metric_Counter)(nil),
		(*Metric_Gauge)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_metrics_proto_goTypes,
		DependencyIndexes: file_proto_metrics_proto_depIdxs,
		EnumInfos:         file_proto_metrics_proto_enumTypes,
		MessageInfos:      file_proto_metrics_proto_msgTypes,
	}.Build()
	File_proto_metrics_proto = out.File
	file_proto_metrics_proto_rawDesc = nil
	file_proto_metrics_proto_goTypes = nil
	file_proto_metrics_proto_depIdxs = nil
}
