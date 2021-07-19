// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: chaos/experimentation/v1/experiment_run_details.proto

package experimentationv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ExperimentRunDetails frontend-renderable details of an experiment run and experiment config associated with it.
type ExperimentRunDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of an experiment run.
	RunId string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	// The status of an experiment run.
	Status Experiment_Status `protobuf:"varint,2,opt,name=status,proto3,enum=clutch.chaos.experimentation.v1.Experiment_Status" json:"status,omitempty"`
	// The list of properties associated with the receiver that's created as a combination
	// of properties for a given experiment run and an experiment config for the run.
	Properties *PropertiesList `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
	// The raw experiment configuration associated with an experiment run.
	Config *anypb.Any `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ExperimentRunDetails) Reset() {
	*x = ExperimentRunDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_experimentation_v1_experiment_run_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentRunDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentRunDetails) ProtoMessage() {}

func (x *ExperimentRunDetails) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_experimentation_v1_experiment_run_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentRunDetails.ProtoReflect.Descriptor instead.
func (*ExperimentRunDetails) Descriptor() ([]byte, []int) {
	return file_chaos_experimentation_v1_experiment_run_details_proto_rawDescGZIP(), []int{0}
}

func (x *ExperimentRunDetails) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *ExperimentRunDetails) GetStatus() Experiment_Status {
	if x != nil {
		return x.Status
	}
	return Experiment_STATUS_UNSPECIFIED
}

func (x *ExperimentRunDetails) GetProperties() *PropertiesList {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *ExperimentRunDetails) GetConfig() *anypb.Any {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_chaos_experimentation_v1_experiment_run_details_proto protoreflect.FileDescriptor

var file_chaos_experimentation_v1_experiment_run_details_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x14, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4f, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6f, 0x73,
	0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaos_experimentation_v1_experiment_run_details_proto_rawDescOnce sync.Once
	file_chaos_experimentation_v1_experiment_run_details_proto_rawDescData = file_chaos_experimentation_v1_experiment_run_details_proto_rawDesc
)

func file_chaos_experimentation_v1_experiment_run_details_proto_rawDescGZIP() []byte {
	file_chaos_experimentation_v1_experiment_run_details_proto_rawDescOnce.Do(func() {
		file_chaos_experimentation_v1_experiment_run_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaos_experimentation_v1_experiment_run_details_proto_rawDescData)
	})
	return file_chaos_experimentation_v1_experiment_run_details_proto_rawDescData
}

var file_chaos_experimentation_v1_experiment_run_details_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chaos_experimentation_v1_experiment_run_details_proto_goTypes = []interface{}{
	(*ExperimentRunDetails)(nil), // 0: clutch.chaos.experimentation.v1.ExperimentRunDetails
	(Experiment_Status)(0),       // 1: clutch.chaos.experimentation.v1.Experiment.Status
	(*PropertiesList)(nil),       // 2: clutch.chaos.experimentation.v1.PropertiesList
	(*anypb.Any)(nil),            // 3: google.protobuf.Any
}
var file_chaos_experimentation_v1_experiment_run_details_proto_depIdxs = []int32{
	1, // 0: clutch.chaos.experimentation.v1.ExperimentRunDetails.status:type_name -> clutch.chaos.experimentation.v1.Experiment.Status
	2, // 1: clutch.chaos.experimentation.v1.ExperimentRunDetails.properties:type_name -> clutch.chaos.experimentation.v1.PropertiesList
	3, // 2: clutch.chaos.experimentation.v1.ExperimentRunDetails.config:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chaos_experimentation_v1_experiment_run_details_proto_init() }
func file_chaos_experimentation_v1_experiment_run_details_proto_init() {
	if File_chaos_experimentation_v1_experiment_run_details_proto != nil {
		return
	}
	file_chaos_experimentation_v1_properties_proto_init()
	file_chaos_experimentation_v1_experiment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chaos_experimentation_v1_experiment_run_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentRunDetails); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chaos_experimentation_v1_experiment_run_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chaos_experimentation_v1_experiment_run_details_proto_goTypes,
		DependencyIndexes: file_chaos_experimentation_v1_experiment_run_details_proto_depIdxs,
		MessageInfos:      file_chaos_experimentation_v1_experiment_run_details_proto_msgTypes,
	}.Build()
	File_chaos_experimentation_v1_experiment_run_details_proto = out.File
	file_chaos_experimentation_v1_experiment_run_details_proto_rawDesc = nil
	file_chaos_experimentation_v1_experiment_run_details_proto_goTypes = nil
	file_chaos_experimentation_v1_experiment_run_details_proto_depIdxs = nil
}
