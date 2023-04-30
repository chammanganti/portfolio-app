// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: proto/project_status.proto

package project_status

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

type ProjectStatusProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectStatusId string `protobuf:"bytes,1,opt,name=project_status_id,json=projectStatusId,proto3" json:"project_status_id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsHealthy       bool   `protobuf:"varint,3,opt,name=is_healthy,json=isHealthy,proto3" json:"is_healthy,omitempty"`
	ProjectId       string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ProjectStatusProto) Reset() {
	*x = ProjectStatusProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_project_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectStatusProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectStatusProto) ProtoMessage() {}

func (x *ProjectStatusProto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectStatusProto.ProtoReflect.Descriptor instead.
func (*ProjectStatusProto) Descriptor() ([]byte, []int) {
	return file_proto_project_status_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectStatusProto) GetProjectStatusId() string {
	if x != nil {
		return x.ProjectStatusId
	}
	return ""
}

func (x *ProjectStatusProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectStatusProto) GetIsHealthy() bool {
	if x != nil {
		return x.IsHealthy
	}
	return false
}

func (x *ProjectStatusProto) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_project_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.
func (*FindRequest) Descriptor() ([]byte, []int) {
	return file_proto_project_status_proto_rawDescGZIP(), []int{1}
}

type FindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectStatuses []*ProjectStatusProto `protobuf:"bytes,1,rep,name=project_statuses,json=projectStatuses,proto3" json:"project_statuses,omitempty"`
}

func (x *FindResponse) Reset() {
	*x = FindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_project_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindResponse) ProtoMessage() {}

func (x *FindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindResponse.ProtoReflect.Descriptor instead.
func (*FindResponse) Descriptor() ([]byte, []int) {
	return file_proto_project_status_proto_rawDescGZIP(), []int{2}
}

func (x *FindResponse) GetProjectStatuses() []*ProjectStatusProto {
	if x != nil {
		return x.ProjectStatuses
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectStatusId string `protobuf:"bytes,1,opt,name=project_status_id,json=projectStatusId,proto3" json:"project_status_id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsHealthy       bool   `protobuf:"varint,3,opt,name=is_healthy,json=isHealthy,proto3" json:"is_healthy,omitempty"`
	ProjectId       string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_project_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_project_status_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetProjectStatusId() string {
	if x != nil {
		return x.ProjectStatusId
	}
	return ""
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetIsHealthy() bool {
	if x != nil {
		return x.IsHealthy
	}
	return false
}

func (x *UpdateRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectStatusId string `protobuf:"bytes,1,opt,name=project_status_id,json=projectStatusId,proto3" json:"project_status_id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsHealthy       bool   `protobuf:"varint,3,opt,name=is_healthy,json=isHealthy,proto3" json:"is_healthy,omitempty"`
	ProjectId       string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_project_status_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_status_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_project_status_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateResponse) GetProjectStatusId() string {
	if x != nil {
		return x.ProjectStatusId
	}
	return ""
}

func (x *UpdateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateResponse) GetIsHealthy() bool {
	if x != nil {
		return x.IsHealthy
	}
	return false
}

func (x *UpdateResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

var File_proto_project_status_proto protoreflect.FileDescriptor

var file_proto_project_status_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x92, 0x01, 0x0a,
	0x12, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x5d, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4d, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22,
	0x8d, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22,
	0x8e, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x32, 0x9b, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x41, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18,
	0x5a, 0x16, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_project_status_proto_rawDescOnce sync.Once
	file_proto_project_status_proto_rawDescData = file_proto_project_status_proto_rawDesc
)

func file_proto_project_status_proto_rawDescGZIP() []byte {
	file_proto_project_status_proto_rawDescOnce.Do(func() {
		file_proto_project_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_project_status_proto_rawDescData)
	})
	return file_proto_project_status_proto_rawDescData
}

var file_proto_project_status_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_project_status_proto_goTypes = []interface{}{
	(*ProjectStatusProto)(nil), // 0: project_status.ProjectStatusProto
	(*FindRequest)(nil),        // 1: project_status.FindRequest
	(*FindResponse)(nil),       // 2: project_status.FindResponse
	(*UpdateRequest)(nil),      // 3: project_status.UpdateRequest
	(*UpdateResponse)(nil),     // 4: project_status.UpdateResponse
}
var file_proto_project_status_proto_depIdxs = []int32{
	0, // 0: project_status.FindResponse.project_statuses:type_name -> project_status.ProjectStatusProto
	1, // 1: project_status.ProjectStatus.Find:input_type -> project_status.FindRequest
	3, // 2: project_status.ProjectStatus.Update:input_type -> project_status.UpdateRequest
	2, // 3: project_status.ProjectStatus.Find:output_type -> project_status.FindResponse
	4, // 4: project_status.ProjectStatus.Update:output_type -> project_status.UpdateResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_project_status_proto_init() }
func file_proto_project_status_proto_init() {
	if File_proto_project_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_project_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectStatusProto); i {
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
		file_proto_project_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRequest); i {
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
		file_proto_project_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindResponse); i {
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
		file_proto_project_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_project_status_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_proto_project_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_project_status_proto_goTypes,
		DependencyIndexes: file_proto_project_status_proto_depIdxs,
		MessageInfos:      file_proto_project_status_proto_msgTypes,
	}.Build()
	File_proto_project_status_proto = out.File
	file_proto_project_status_proto_rawDesc = nil
	file_proto_project_status_proto_goTypes = nil
	file_proto_project_status_proto_depIdxs = nil
}
