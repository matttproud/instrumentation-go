// Code generated by protoc-gen-go.
// source: census.proto
// DO NOT EDIT!

/*
Package generated_proto is a generated protocol buffer package.

It is generated from these files:
	census.proto
	stats_context.proto

It has these top-level messages:
	Duration
	Timestamp
	MeasurementDescriptor
	DistributionAggregation
	DistributionAggregationDescriptor
	IntervalAggregation
	IntervalAggregationDescriptor
	Tag
	ViewDescriptor
	DistributionView
	IntervalView
	View
	StatsContext
*/
package generated_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Fundamental units of measurement supported by Census
// TODO(aveitch): expand this to include other S.I. units?
type MeasurementDescriptor_BasicUnit int32

const (
	MeasurementDescriptor_UNKNOWN   MeasurementDescriptor_BasicUnit = 0
	MeasurementDescriptor_SCALAR    MeasurementDescriptor_BasicUnit = 1
	MeasurementDescriptor_BITS      MeasurementDescriptor_BasicUnit = 2
	MeasurementDescriptor_BYTES     MeasurementDescriptor_BasicUnit = 3
	MeasurementDescriptor_SECONDS   MeasurementDescriptor_BasicUnit = 4
	MeasurementDescriptor_CORES     MeasurementDescriptor_BasicUnit = 5
	MeasurementDescriptor_MAX_UNITS MeasurementDescriptor_BasicUnit = 6
)

var MeasurementDescriptor_BasicUnit_name = map[int32]string{
	0: "UNKNOWN",
	1: "SCALAR",
	2: "BITS",
	3: "BYTES",
	4: "SECONDS",
	5: "CORES",
	6: "MAX_UNITS",
}
var MeasurementDescriptor_BasicUnit_value = map[string]int32{
	"UNKNOWN":   0,
	"SCALAR":    1,
	"BITS":      2,
	"BYTES":     3,
	"SECONDS":   4,
	"CORES":     5,
	"MAX_UNITS": 6,
}

func (x MeasurementDescriptor_BasicUnit) String() string {
	return proto.EnumName(MeasurementDescriptor_BasicUnit_name, int32(x))
}
func (MeasurementDescriptor_BasicUnit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

type Duration struct {
	// Signed seconds of the span of time. Must be from -315,576,000,000
	// to +315,576,000,000 inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Signed fractions of a second at nanosecond resolution of the span
	// of time. Durations less than one second are represented with a 0
	// `seconds` field and a positive or negative `nanos` field. For durations
	// of one second or more, a non-zero value for the `nanos` field must be
	// of the same sign as the `seconds` field. Must be from -999,999,999
	// to +999,999,999 inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *Duration) Reset()                    { *m = Duration{} }
func (m *Duration) String() string            { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()               {}
func (*Duration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// MeasurementDescriptor describes a data point (measurement) type.
type MeasurementDescriptor struct {
	// A descriptive name, e.g. rpc_latency, cpu. Must be unique.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// More detailed description of the resource, used in documentation.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// The units used by this type of measurement.
	Unit *MeasurementDescriptor_MeasurementUnit `protobuf:"bytes,3,opt,name=unit" json:"unit,omitempty"`
}

func (m *MeasurementDescriptor) Reset()                    { *m = MeasurementDescriptor{} }
func (m *MeasurementDescriptor) String() string            { return proto.CompactTextString(m) }
func (*MeasurementDescriptor) ProtoMessage()               {}
func (*MeasurementDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MeasurementDescriptor) GetUnit() *MeasurementDescriptor_MeasurementUnit {
	if m != nil {
		return m.Unit
	}
	return nil
}

// MeasurementUnit lets you build compound units of the form
//   10^n * (A * B * ...) / (X * Y * ...),
// where the elements in the numerator and denominator are all BasicUnits.  A
// MeasurementUnit must have at least one BasicUnit in its numerator.
//
// To specify multiplication in the numerator or denominator, simply specify
// multiple numerator or denominator fields.  For example:
//
// - byte-seconds (i.e. bytes * seconds):
//     numerator: BYTES
//     numerator: SECS
//
// - events/sec^2 (i.e. rate of change of events/sec):
//     numerator: SCALAR
//     denominator: SECS
//     denominator: SECS
//
// To specify multiples (in power of 10) of units, specify a non-zero
// 'power10' value, for example:
//
// - MB/s (i.e. megabytes / s):
//     power10: 6
//     numerator: BYTES
//     denominator: SECS
//
// - nanoseconds
//     power10: -9
//     numerator: SECS
type MeasurementDescriptor_MeasurementUnit struct {
	Power10      int32                             `protobuf:"varint,1,opt,name=power10" json:"power10,omitempty"`
	Numerators   []MeasurementDescriptor_BasicUnit `protobuf:"varint,2,rep,name=numerators,enum=google.instrumentation.MeasurementDescriptor_BasicUnit" json:"numerators,omitempty"`
	Denominators []MeasurementDescriptor_BasicUnit `protobuf:"varint,3,rep,name=denominators,enum=google.instrumentation.MeasurementDescriptor_BasicUnit" json:"denominators,omitempty"`
}

func (m *MeasurementDescriptor_MeasurementUnit) Reset()         { *m = MeasurementDescriptor_MeasurementUnit{} }
func (m *MeasurementDescriptor_MeasurementUnit) String() string { return proto.CompactTextString(m) }
func (*MeasurementDescriptor_MeasurementUnit) ProtoMessage()    {}
func (*MeasurementDescriptor_MeasurementUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

// DistributionAggregation contains summary statistics for a population of
// values and, optionally, a histogram representing the distribution of those
// values across a specified set of histogram buckets, as defined in
// DistributionAggregationDescriptor.bucket_bounds.
//
// The summary statistics are the count, mean, minimum, and the maximum of the
// set of population of values.
//
// Although it is not forbidden, it is generally a bad idea to include
// non-finite values (infinities or NaNs) in the population of values, as this
// will render the `mean` field meaningless.
type DistributionAggregation struct {
	// The number of values in the population. Must be non-negative.
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	// The arithmetic mean of the values in the population. If `count` is zero
	// then this field must be zero.
	Mean float64 `protobuf:"fixed64,2,opt,name=mean" json:"mean,omitempty"`
	// The sum of the values in the population.  If `count` is zero then this
	// field must be zero.
	Sum float64 `protobuf:"fixed64,3,opt,name=sum" json:"sum,omitempty"`
	// The range of the population values. If `count` is zero, this field will not
	// be defined.
	Range *DistributionAggregation_Range `protobuf:"bytes,4,opt,name=range" json:"range,omitempty"`
	// A Distribution may optionally contain a histogram of the values in the
	// population. The histogram is given in `bucket_count` as counts of values
	// that fall into one of a sequence of non-overlapping buckets, as described
	// by `DistributionAggregationDescriptor.bucket_boundaries`. The sum of the
	// values in `bucket_counts` must equal the value in `count`.
	//
	// Bucket counts are given in order under the numbering scheme described
	// above (the underflow bucket has number 0; the finite buckets, if any,
	// have numbers 1 through N-2; the overflow bucket has number N-1).
	//
	// The size of `bucket_count` must be no greater than N as defined in
	// `bucket_boundaries`.
	//
	// Any suffix of trailing zero bucket_count fields may be omitted.
	BucketCounts []int64 `protobuf:"varint,5,rep,name=bucket_counts,json=bucketCounts" json:"bucket_counts,omitempty"`
	// Tags associated with this DistributionAggregation. These will be filled
	// in based on the View specification.
	Tags []*Tag `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
}

func (m *DistributionAggregation) Reset()                    { *m = DistributionAggregation{} }
func (m *DistributionAggregation) String() string            { return proto.CompactTextString(m) }
func (*DistributionAggregation) ProtoMessage()               {}
func (*DistributionAggregation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DistributionAggregation) GetRange() *DistributionAggregation_Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *DistributionAggregation) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Describes a range of population values.
type DistributionAggregation_Range struct {
	// The minimum of the population values.
	Min float64 `protobuf:"fixed64,1,opt,name=min" json:"min,omitempty"`
	// The maximum of the population values.
	Max float64 `protobuf:"fixed64,2,opt,name=max" json:"max,omitempty"`
}

func (m *DistributionAggregation_Range) Reset()         { *m = DistributionAggregation_Range{} }
func (m *DistributionAggregation_Range) String() string { return proto.CompactTextString(m) }
func (*DistributionAggregation_Range) ProtoMessage()    {}
func (*DistributionAggregation_Range) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type DistributionAggregationDescriptor struct {
	// A Distribution may optionally contain a histogram of the values in the
	// population. The bucket boundaries for that histogram are described by
	// `bucket_bounds`. This defines `size(bucket_bounds) + 1` (= N)
	// buckets. The boundaries for bucket index i are:
	//
	// [-infinity, bucket_bounds[i]) for i == 0
	// [bucket_bounds[i-1], bucket_bounds[i]) for 0 < i < N-2
	// [bucket_bounds[i-1], +infinity) for i == N-1
	//
	// i.e. an underflow bucket (number 0), zero or more finite buckets (1
	// through N - 2, and an overflow bucket (N - 1), with inclusive lower
	// bounds and exclusive upper bounds.
	//
	// If `bucket_bounds` has no elements (zero size), then there is no
	// histogram associated with the Distribution. If `bucket_bounds` has only
	// one element, there are no finite buckets, and that single element is the
	// common boundary of the overflow and underflow buckets. The values must
	// be monotonically increasing.
	BucketBounds []float64 `protobuf:"fixed64,1,rep,name=bucket_bounds,json=bucketBounds" json:"bucket_bounds,omitempty"`
}

func (m *DistributionAggregationDescriptor) Reset()         { *m = DistributionAggregationDescriptor{} }
func (m *DistributionAggregationDescriptor) String() string { return proto.CompactTextString(m) }
func (*DistributionAggregationDescriptor) ProtoMessage()    {}
func (*DistributionAggregationDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4}
}

// An IntervalAggreation records summary stats over various time
// windows. These stats are approximate, with the degree of accuracy
// controlled by setting the n_sub_intervals parameter in the
// IntervalAggregationDescriptor.
type IntervalAggregation struct {
	// Full set of intervals for this aggregation.
	Intervals []*IntervalAggregation_Interval `protobuf:"bytes,1,rep,name=intervals" json:"intervals,omitempty"`
	// Tags associated with this IntervalAggregation. These will be filled in
	// based on the View specification.
	Tags []*Tag `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
}

func (m *IntervalAggregation) Reset()                    { *m = IntervalAggregation{} }
func (m *IntervalAggregation) String() string            { return proto.CompactTextString(m) }
func (*IntervalAggregation) ProtoMessage()               {}
func (*IntervalAggregation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IntervalAggregation) GetIntervals() []*IntervalAggregation_Interval {
	if m != nil {
		return m.Intervals
	}
	return nil
}

func (m *IntervalAggregation) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Summary statistic over a single time interval.
type IntervalAggregation_Interval struct {
	// The interval duration. Must be positive.
	IntervalSize *Duration `protobuf:"bytes,1,opt,name=interval_size,json=intervalSize" json:"interval_size,omitempty"`
	// Approximate number of measurements recorded in this interval.
	Count float64 `protobuf:"fixed64,2,opt,name=count" json:"count,omitempty"`
	// The cumulative sum of measurements in this interval.
	Sum float64 `protobuf:"fixed64,3,opt,name=sum" json:"sum,omitempty"`
}

func (m *IntervalAggregation_Interval) Reset()                    { *m = IntervalAggregation_Interval{} }
func (m *IntervalAggregation_Interval) String() string            { return proto.CompactTextString(m) }
func (*IntervalAggregation_Interval) ProtoMessage()               {}
func (*IntervalAggregation_Interval) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *IntervalAggregation_Interval) GetIntervalSize() *Duration {
	if m != nil {
		return m.IntervalSize
	}
	return nil
}

// An IntervalAggreationDescriptor specifies time intervals for an
// IntervalAggregation.
type IntervalAggregationDescriptor struct {
	// Number of internal sub-intervals to use when collecting stats for each
	// interval. The max error in interval measurements will be approximately
	// 1/n_sub_intervals (although in practice, this will only be approached in
	// the presence of very large and bursty workload changes), and underlying
	// memory usage will be roughly proportional to the value of this
	// field. Must be in the range [2, 20]. A value of 5 will be used if this is
	// unspecified.
	NSubIntervals int32 `protobuf:"varint,1,opt,name=n_sub_intervals,json=nSubIntervals" json:"n_sub_intervals,omitempty"`
	// The size of each interval, as a time duration. Must have at least one
	// element.
	IntervalSizes []*Duration `protobuf:"bytes,2,rep,name=interval_sizes,json=intervalSizes" json:"interval_sizes,omitempty"`
}

func (m *IntervalAggregationDescriptor) Reset()                    { *m = IntervalAggregationDescriptor{} }
func (m *IntervalAggregationDescriptor) String() string            { return proto.CompactTextString(m) }
func (*IntervalAggregationDescriptor) ProtoMessage()               {}
func (*IntervalAggregationDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IntervalAggregationDescriptor) GetIntervalSizes() []*Duration {
	if m != nil {
		return m.IntervalSizes
	}
	return nil
}

// A Tag: key-value pair.
// Both strings must be printable ASCII.
type Tag struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// A ViewDescriptor specifies an AggregationDescriptor and a set of tag
// keys. Views instantiated from this descriptor will contain Aggregations
// broken down by the unique set of matching tag values for each measurement.
type ViewDescriptor struct {
	// Name of view. Must be unique.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// More detailed description, for documentation purposes.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Name of a MeasurementDescriptor to be used for this view.
	MeasurementDescriptorName string `protobuf:"bytes,3,opt,name=measurement_descriptor_name,json=measurementDescriptorName" json:"measurement_descriptor_name,omitempty"`
	// Aggregation type to associate with View.
	//
	// Types that are valid to be assigned to Aggregation:
	//	*ViewDescriptor_IntervalAggregation
	//	*ViewDescriptor_DistributionAggregation
	Aggregation isViewDescriptor_Aggregation `protobuf_oneof:"aggregation"`
	// Tag keys to match with a given measurement. If no keys are specified,
	// then all stats are recorded. Keys must be unique.
	TagKeys []string `protobuf:"bytes,6,rep,name=tag_keys,json=tagKeys" json:"tag_keys,omitempty"`
}

func (m *ViewDescriptor) Reset()                    { *m = ViewDescriptor{} }
func (m *ViewDescriptor) String() string            { return proto.CompactTextString(m) }
func (*ViewDescriptor) ProtoMessage()               {}
func (*ViewDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type isViewDescriptor_Aggregation interface {
	isViewDescriptor_Aggregation()
}

type ViewDescriptor_IntervalAggregation struct {
	IntervalAggregation *IntervalAggregationDescriptor `protobuf:"bytes,4,opt,name=interval_aggregation,json=intervalAggregation,oneof"`
}
type ViewDescriptor_DistributionAggregation struct {
	DistributionAggregation *DistributionAggregationDescriptor `protobuf:"bytes,5,opt,name=distribution_aggregation,json=distributionAggregation,oneof"`
}

func (*ViewDescriptor_IntervalAggregation) isViewDescriptor_Aggregation()     {}
func (*ViewDescriptor_DistributionAggregation) isViewDescriptor_Aggregation() {}

func (m *ViewDescriptor) GetAggregation() isViewDescriptor_Aggregation {
	if m != nil {
		return m.Aggregation
	}
	return nil
}

func (m *ViewDescriptor) GetIntervalAggregation() *IntervalAggregationDescriptor {
	if x, ok := m.GetAggregation().(*ViewDescriptor_IntervalAggregation); ok {
		return x.IntervalAggregation
	}
	return nil
}

func (m *ViewDescriptor) GetDistributionAggregation() *DistributionAggregationDescriptor {
	if x, ok := m.GetAggregation().(*ViewDescriptor_DistributionAggregation); ok {
		return x.DistributionAggregation
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ViewDescriptor) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ViewDescriptor_OneofMarshaler, _ViewDescriptor_OneofUnmarshaler, _ViewDescriptor_OneofSizer, []interface{}{
		(*ViewDescriptor_IntervalAggregation)(nil),
		(*ViewDescriptor_DistributionAggregation)(nil),
	}
}

func _ViewDescriptor_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ViewDescriptor)
	// aggregation
	switch x := m.Aggregation.(type) {
	case *ViewDescriptor_IntervalAggregation:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IntervalAggregation); err != nil {
			return err
		}
	case *ViewDescriptor_DistributionAggregation:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DistributionAggregation); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ViewDescriptor.Aggregation has unexpected type %T", x)
	}
	return nil
}

func _ViewDescriptor_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ViewDescriptor)
	switch tag {
	case 4: // aggregation.interval_aggregation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IntervalAggregationDescriptor)
		err := b.DecodeMessage(msg)
		m.Aggregation = &ViewDescriptor_IntervalAggregation{msg}
		return true, err
	case 5: // aggregation.distribution_aggregation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DistributionAggregationDescriptor)
		err := b.DecodeMessage(msg)
		m.Aggregation = &ViewDescriptor_DistributionAggregation{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ViewDescriptor_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ViewDescriptor)
	// aggregation
	switch x := m.Aggregation.(type) {
	case *ViewDescriptor_IntervalAggregation:
		s := proto.Size(x.IntervalAggregation)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ViewDescriptor_DistributionAggregation:
		s := proto.Size(x.DistributionAggregation)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// DistributionView contains all aggregations for a view specified using a
// DistributionAggregationDescriptor.
type DistributionView struct {
	// Aggregations - each will have a unique set of tag values for the tag_keys
	// associated with the corresponding View.
	Aggregations []*DistributionAggregation `protobuf:"bytes,1,rep,name=aggregations" json:"aggregations,omitempty"`
	// Start and end timestamps over which aggregations was accumulated.
	Start *Timestamp `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	End   *Timestamp `protobuf:"bytes,3,opt,name=end" json:"end,omitempty"`
}

func (m *DistributionView) Reset()                    { *m = DistributionView{} }
func (m *DistributionView) String() string            { return proto.CompactTextString(m) }
func (*DistributionView) ProtoMessage()               {}
func (*DistributionView) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DistributionView) GetAggregations() []*DistributionAggregation {
	if m != nil {
		return m.Aggregations
	}
	return nil
}

func (m *DistributionView) GetStart() *Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *DistributionView) GetEnd() *Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

// IntervalView contains all aggregations for a view specified using a
// IntervalAggregationDescriptor.
type IntervalView struct {
	// Aggregations - each will have a unique set of tag values for the tag_keys
	// associated with the corresponding View.
	Aggregations []*IntervalAggregation `protobuf:"bytes,1,rep,name=aggregations" json:"aggregations,omitempty"`
}

func (m *IntervalView) Reset()                    { *m = IntervalView{} }
func (m *IntervalView) String() string            { return proto.CompactTextString(m) }
func (*IntervalView) ProtoMessage()               {}
func (*IntervalView) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IntervalView) GetAggregations() []*IntervalAggregation {
	if m != nil {
		return m.Aggregations
	}
	return nil
}

// A View contains the aggregations based on a ViewDescriptor.
type View struct {
	// ViewDescriptor name associated with this set of View.
	ViewName string `protobuf:"bytes,1,opt,name=view_name,json=viewName" json:"view_name,omitempty"`
	// Types that are valid to be assigned to View:
	//	*View_DistributionView
	//	*View_IntervalView
	View isView_View `protobuf_oneof:"view"`
}

func (m *View) Reset()                    { *m = View{} }
func (m *View) String() string            { return proto.CompactTextString(m) }
func (*View) ProtoMessage()               {}
func (*View) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type isView_View interface {
	isView_View()
}

type View_DistributionView struct {
	DistributionView *DistributionView `protobuf:"bytes,2,opt,name=distribution_view,json=distributionView,oneof"`
}
type View_IntervalView struct {
	IntervalView *IntervalView `protobuf:"bytes,3,opt,name=interval_view,json=intervalView,oneof"`
}

func (*View_DistributionView) isView_View() {}
func (*View_IntervalView) isView_View()     {}

func (m *View) GetView() isView_View {
	if m != nil {
		return m.View
	}
	return nil
}

func (m *View) GetDistributionView() *DistributionView {
	if x, ok := m.GetView().(*View_DistributionView); ok {
		return x.DistributionView
	}
	return nil
}

func (m *View) GetIntervalView() *IntervalView {
	if x, ok := m.GetView().(*View_IntervalView); ok {
		return x.IntervalView
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*View) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _View_OneofMarshaler, _View_OneofUnmarshaler, _View_OneofSizer, []interface{}{
		(*View_DistributionView)(nil),
		(*View_IntervalView)(nil),
	}
}

func _View_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*View)
	// view
	switch x := m.View.(type) {
	case *View_DistributionView:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DistributionView); err != nil {
			return err
		}
	case *View_IntervalView:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IntervalView); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("View.View has unexpected type %T", x)
	}
	return nil
}

func _View_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*View)
	switch tag {
	case 2: // view.distribution_view
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DistributionView)
		err := b.DecodeMessage(msg)
		m.View = &View_DistributionView{msg}
		return true, err
	case 3: // view.interval_view
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IntervalView)
		err := b.DecodeMessage(msg)
		m.View = &View_IntervalView{msg}
		return true, err
	default:
		return false, nil
	}
}

func _View_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*View)
	// view
	switch x := m.View.(type) {
	case *View_DistributionView:
		s := proto.Size(x.DistributionView)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *View_IntervalView:
		s := proto.Size(x.IntervalView)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Duration)(nil), "google.instrumentation.Duration")
	proto.RegisterType((*Timestamp)(nil), "google.instrumentation.Timestamp")
	proto.RegisterType((*MeasurementDescriptor)(nil), "google.instrumentation.MeasurementDescriptor")
	proto.RegisterType((*MeasurementDescriptor_MeasurementUnit)(nil), "google.instrumentation.MeasurementDescriptor.MeasurementUnit")
	proto.RegisterType((*DistributionAggregation)(nil), "google.instrumentation.DistributionAggregation")
	proto.RegisterType((*DistributionAggregation_Range)(nil), "google.instrumentation.DistributionAggregation.Range")
	proto.RegisterType((*DistributionAggregationDescriptor)(nil), "google.instrumentation.DistributionAggregationDescriptor")
	proto.RegisterType((*IntervalAggregation)(nil), "google.instrumentation.IntervalAggregation")
	proto.RegisterType((*IntervalAggregation_Interval)(nil), "google.instrumentation.IntervalAggregation.Interval")
	proto.RegisterType((*IntervalAggregationDescriptor)(nil), "google.instrumentation.IntervalAggregationDescriptor")
	proto.RegisterType((*Tag)(nil), "google.instrumentation.Tag")
	proto.RegisterType((*ViewDescriptor)(nil), "google.instrumentation.ViewDescriptor")
	proto.RegisterType((*DistributionView)(nil), "google.instrumentation.DistributionView")
	proto.RegisterType((*IntervalView)(nil), "google.instrumentation.IntervalView")
	proto.RegisterType((*View)(nil), "google.instrumentation.View")
	proto.RegisterEnum("google.instrumentation.MeasurementDescriptor_BasicUnit", MeasurementDescriptor_BasicUnit_name, MeasurementDescriptor_BasicUnit_value)
}

func init() { proto.RegisterFile("census.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 931 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0xe3, 0x38, 0x4d, 0x4e, 0x92, 0xd6, 0xcc, 0x16, 0x36, 0xbb, 0x15, 0x52, 0xd6, 0xa0,
	0x55, 0xa4, 0x15, 0x59, 0xc8, 0x82, 0x56, 0x80, 0x40, 0x6a, 0xd2, 0x42, 0xab, 0xb2, 0x29, 0x8c,
	0x53, 0xca, 0xcf, 0x85, 0x99, 0xc4, 0x23, 0x6b, 0x68, 0x3d, 0xae, 0x3c, 0xe3, 0x96, 0xee, 0x53,
	0x70, 0x01, 0x8f, 0xc0, 0x03, 0x71, 0xc1, 0x03, 0x70, 0xc3, 0x73, 0xa0, 0x19, 0xdb, 0x89, 0xb3,
	0xc4, 0xb4, 0xd1, 0x5e, 0x65, 0xce, 0x99, 0x39, 0xdf, 0xf9, 0xce, 0xaf, 0x03, 0xad, 0x19, 0xe5,
	0x22, 0x11, 0xfd, 0xcb, 0x38, 0x92, 0x11, 0x7a, 0x2b, 0x88, 0xa2, 0xe0, 0x82, 0xf6, 0x19, 0x17,
	0x32, 0x4e, 0x42, 0xca, 0x25, 0x91, 0x2c, 0xe2, 0xce, 0x27, 0x50, 0xdf, 0x4f, 0x62, 0x7d, 0x46,
	0x1d, 0xd8, 0x14, 0x74, 0x16, 0x71, 0x5f, 0x74, 0x8c, 0xae, 0xd1, 0x33, 0x71, 0x2e, 0xa2, 0x1d,
	0xb0, 0x38, 0xe1, 0x91, 0xe8, 0x54, 0xba, 0x46, 0xcf, 0xc2, 0xa9, 0xe0, 0x7c, 0x0a, 0x8d, 0x09,
	0x0b, 0xa9, 0x90, 0x24, 0xbc, 0x5c, 0xdb, 0xf8, 0x1f, 0x13, 0xde, 0x7c, 0x41, 0x89, 0x48, 0x62,
	0xaa, 0xd8, 0xec, 0x53, 0x31, 0x8b, 0xd9, 0xa5, 0x8c, 0x62, 0x84, 0xa0, 0xca, 0x49, 0x48, 0x35,
	0x4c, 0x03, 0xeb, 0x33, 0xea, 0x42, 0xd3, 0xcf, 0x5e, 0xb0, 0x88, 0x6b, 0xa4, 0x06, 0x2e, 0xaa,
	0xd0, 0x37, 0x50, 0x4d, 0x38, 0x93, 0x1d, 0xb3, 0x6b, 0xf4, 0x9a, 0x83, 0xcf, 0xfa, 0xab, 0xe3,
	0xed, 0xaf, 0x74, 0x59, 0xd4, 0x9e, 0x72, 0x26, 0xb1, 0x86, 0x7a, 0xf8, 0xb7, 0x01, 0xdb, 0xaf,
	0xdc, 0xa8, 0x30, 0x2f, 0xa3, 0x6b, 0x1a, 0x7f, 0xf0, 0xbe, 0xe6, 0x67, 0xe1, 0x5c, 0x44, 0x67,
	0x00, 0x3c, 0x09, 0x69, 0x4c, 0x64, 0x14, 0xab, 0x58, 0xcd, 0xde, 0xd6, 0xe0, 0xf9, 0x7a, 0x34,
	0x86, 0x44, 0xb0, 0x99, 0x26, 0x50, 0x80, 0x42, 0x3f, 0x42, 0xcb, 0xa7, 0x3c, 0x0a, 0x19, 0x4f,
	0xa1, 0xcd, 0xd7, 0x83, 0x5e, 0x02, 0x73, 0x7e, 0x82, 0xc6, 0xfc, 0x0a, 0x35, 0x61, 0xf3, 0x74,
	0x7c, 0x3c, 0x3e, 0x39, 0x1b, 0xdb, 0x1b, 0x08, 0xa0, 0xe6, 0x8e, 0xf6, 0xbe, 0xda, 0xc3, 0xb6,
	0x81, 0xea, 0x50, 0x1d, 0x1e, 0x4d, 0x5c, 0xbb, 0x82, 0x1a, 0x60, 0x0d, 0xbf, 0x9f, 0x1c, 0xb8,
	0xb6, 0xa9, 0x5e, 0xbb, 0x07, 0xa3, 0x93, 0xf1, 0xbe, 0x6b, 0x57, 0x95, 0x7e, 0x74, 0x82, 0x0f,
	0x5c, 0xdb, 0x42, 0x6d, 0x68, 0xbc, 0xd8, 0xfb, 0xce, 0x3b, 0x1d, 0x2b, 0x8b, 0x9a, 0xf3, 0x47,
	0x05, 0xee, 0xef, 0x33, 0x21, 0x63, 0x36, 0x4d, 0x14, 0xc1, 0xbd, 0x20, 0x88, 0x69, 0x90, 0x76,
	0xdc, 0x0e, 0x58, 0xb3, 0x28, 0xe1, 0x32, 0x6b, 0x99, 0x54, 0x50, 0x0d, 0x10, 0x52, 0x92, 0x56,
	0xd9, 0xc0, 0xfa, 0x8c, 0x6c, 0x30, 0x45, 0x12, 0xea, 0xea, 0x1a, 0x58, 0x1d, 0xd1, 0x31, 0x58,
	0x31, 0xe1, 0x01, 0xed, 0x54, 0x75, 0xc5, 0x3f, 0x2a, 0xcb, 0x47, 0x89, 0xef, 0x3e, 0x56, 0xc6,
	0x38, 0xc5, 0x40, 0xef, 0x40, 0x7b, 0x9a, 0xcc, 0xce, 0xa9, 0xf4, 0x34, 0x05, 0xd1, 0xb1, 0xba,
	0x66, 0xcf, 0xc4, 0xad, 0x54, 0x39, 0xd2, 0x3a, 0xf4, 0x14, 0xaa, 0x92, 0x04, 0xa2, 0x53, 0xeb,
	0x9a, 0xbd, 0xe6, 0x60, 0xb7, 0xcc, 0xe1, 0x84, 0x04, 0x58, 0x3f, 0x7c, 0xf8, 0x04, 0x2c, 0xed,
	0x45, 0xb1, 0x0f, 0x19, 0xd7, 0x51, 0x1a, 0x58, 0x1d, 0xb5, 0x86, 0xfc, 0x92, 0x85, 0xa8, 0x8e,
	0xce, 0x21, 0x3c, 0x2a, 0xa1, 0x5a, 0x98, 0x8d, 0x05, 0xcf, 0x69, 0x94, 0xa4, 0xb3, 0x66, 0xf6,
	0x8c, 0x9c, 0xe7, 0x50, 0xeb, 0x9c, 0xdf, 0x2b, 0x70, 0xef, 0x88, 0x4b, 0x1a, 0x5f, 0x91, 0x8b,
	0x62, 0xb6, 0x31, 0x34, 0x58, 0xa6, 0x4e, 0x0d, 0x9b, 0x83, 0x0f, 0xcb, 0x82, 0x58, 0x61, 0x3f,
	0xd7, 0xe1, 0x05, 0xcc, 0x3c, 0x27, 0x95, 0xbb, 0xe6, 0xe4, 0x06, 0xea, 0x39, 0x0e, 0x3a, 0x80,
	0x76, 0x8e, 0xe4, 0x09, 0xf6, 0x32, 0x1d, 0xf9, 0xe6, 0xa0, 0x5b, 0x5a, 0xca, 0x6c, 0x53, 0xe1,
	0x56, 0x6e, 0xe6, 0xb2, 0x97, 0x74, 0xd1, 0x45, 0x69, 0x36, 0xb3, 0x2e, 0xfa, 0x4f, 0xc7, 0x38,
	0xbf, 0x1a, 0xf0, 0xf6, 0x8a, 0xb8, 0x0a, 0xe9, 0x7d, 0x0c, 0xdb, 0xdc, 0x13, 0xc9, 0xd4, 0x2b,
	0xe6, 0x49, 0x4d, 0x79, 0x9b, 0xbb, 0xc9, 0xf4, 0x68, 0x1e, 0xf5, 0x97, 0xb0, 0xb5, 0x44, 0x3c,
	0x8f, 0xff, 0x76, 0xe6, 0xed, 0x22, 0x73, 0xe1, 0xbc, 0x07, 0xe6, 0x84, 0x04, 0x8a, 0xeb, 0x39,
	0xbd, 0xc9, 0x36, 0x9e, 0x3a, 0xaa, 0x98, 0xae, 0xc8, 0x45, 0x42, 0xb3, 0x55, 0x97, 0x0a, 0xce,
	0x6f, 0x26, 0x6c, 0x7d, 0xcb, 0xe8, 0xf5, 0x6b, 0x6f, 0xcb, 0xcf, 0x61, 0x37, 0x5c, 0xec, 0x09,
	0xcf, 0x9f, 0xe3, 0x79, 0x1a, 0xcc, 0xd4, 0x16, 0x0f, 0xc2, 0x55, 0xab, 0x64, 0xac, 0x3c, 0xfc,
	0x0c, 0x3b, 0xf3, 0x04, 0x90, 0x45, 0x2a, 0x6f, 0x9b, 0xc5, 0xff, 0xcd, 0xfe, 0xe1, 0x06, 0xbe,
	0xc7, 0x56, 0xb4, 0xed, 0x15, 0x74, 0xfc, 0xc2, 0x60, 0x2c, 0xf9, 0xb3, 0xb4, 0xbf, 0x8f, 0xd7,
	0x9c, 0xfd, 0x25, 0x9f, 0xf7, 0xfd, 0x92, 0xe5, 0xf4, 0x00, 0xea, 0x92, 0x04, 0xde, 0x39, 0xbd,
	0x49, 0x47, 0xbe, 0x81, 0x37, 0x25, 0x09, 0x8e, 0xe9, 0x8d, 0x18, 0xb6, 0xa1, 0x59, 0x60, 0xe1,
	0xfc, 0x65, 0x80, 0x5d, 0x74, 0xa5, 0x4a, 0x84, 0x5c, 0x68, 0x15, 0xde, 0xe4, 0x03, 0xf7, 0x74,
	0x4d, 0xaa, 0x78, 0x09, 0x04, 0x3d, 0x07, 0x4b, 0x48, 0x12, 0xa7, 0xad, 0xde, 0x1c, 0x3c, 0x2a,
	0x9d, 0xb7, 0xfc, 0xbb, 0x8c, 0xd3, 0xf7, 0xe8, 0x19, 0x98, 0x94, 0xfb, 0xd9, 0xd7, 0xf1, 0x0e,
	0x66, 0xea, 0xb5, 0xe3, 0x41, 0x2b, 0xaf, 0x98, 0x0e, 0xe9, 0x64, 0x65, 0x48, 0x4f, 0xd6, 0xa8,
	0xf6, 0x72, 0x38, 0xce, 0x9f, 0x06, 0x54, 0x35, 0xf2, 0x2e, 0x34, 0xae, 0x18, 0xbd, 0xf6, 0x0a,
	0xad, 0x5c, 0x57, 0x0a, 0xdd, 0x6c, 0x67, 0xf0, 0xc6, 0x52, 0x03, 0xa8, 0x8b, 0x2c, 0x01, 0xbd,
	0xbb, 0xa4, 0x53, 0x79, 0x38, 0xdc, 0xc0, 0xb6, 0xff, 0x6a, 0x89, 0x8e, 0x0b, 0xfb, 0x47, 0x83,
	0xa6, 0xe9, 0x79, 0xf7, 0xb6, 0x80, 0x32, 0xc0, 0xf9, 0x16, 0x52, 0xf2, 0xb0, 0x06, 0x55, 0x85,
	0x31, 0xfc, 0x02, 0x1e, 0xcf, 0xa2, 0xb0, 0x0c, 0x42, 0x48, 0x22, 0xb3, 0xff, 0x64, 0xc3, 0xe6,
	0x48, 0xff, 0x43, 0xfb, 0x5a, 0x09, 0x3f, 0x6c, 0x07, 0x94, 0xab, 0x2f, 0x3e, 0xf5, 0x3d, 0x7d,
	0x3b, 0xad, 0xe9, 0x9f, 0x67, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x04, 0x30, 0xe2, 0x99, 0xc8,
	0x09, 0x00, 0x00,
}