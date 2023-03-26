// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: edholm/profzf/projects/v1beta1/projects_service.proto

package projectsv1beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProjectsRequest) Reset() {
	*x = ListProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsRequest) ProtoMessage() {}

func (x *ListProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescGZIP(), []int{0}
}

type ListProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *ListProjectsResponse) Reset() {
	*x = ListProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse) ProtoMessage() {}

func (x *ListProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListProjectsResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

type GetProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IncreaseUsageCount bool   `protobuf:"varint,2,opt,name=increase_usage_count,json=increaseUsageCount,proto3" json:"increase_usage_count,omitempty"`
}

func (x *GetProjectRequest) Reset() {
	*x = GetProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectRequest) ProtoMessage() {}

func (x *GetProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectRequest.ProtoReflect.Descriptor instead.
func (*GetProjectRequest) Descriptor() ([]byte, []int) {
	return file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProjectRequest) GetIncreaseUsageCount() bool {
	if x != nil {
		return x.IncreaseUsageCount
	}
	return false
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path        string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	GitStatus   *GitStatus             `protobuf:"bytes,3,opt,name=git_status,json=gitStatus,proto3" json:"git_status,omitempty"`
	UsageCount  int32                  `protobuf:"varint,4,opt,name=usage_count,json=usageCount,proto3" json:"usage_count,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescGZIP(), []int{3}
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Project) GetGitStatus() *GitStatus {
	if x != nil {
		return x.GitStatus
	}
	return nil
}

func (x *Project) GetUsageCount() int32 {
	if x != nil {
		return x.UsageCount
	}
	return 0
}

func (x *Project) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type GitStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Branch string `protobuf:"bytes,1,opt,name=branch,proto3" json:"branch,omitempty"`
	Dirty  bool   `protobuf:"varint,2,opt,name=dirty,proto3" json:"dirty,omitempty"`
	Action string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"` // (rebase/merge)
}

func (x *GitStatus) Reset() {
	*x = GitStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitStatus) ProtoMessage() {}

func (x *GitStatus) ProtoReflect() protoreflect.Message {
	mi := &file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitStatus.ProtoReflect.Descriptor instead.
func (*GitStatus) Descriptor() ([]byte, []int) {
	return file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescGZIP(), []int{4}
}

func (x *GitStatus) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *GitStatus) GetDirty() bool {
	if x != nil {
		return x.Dirty
	}
	return false
}

func (x *GitStatus) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

var File_edholm_profzf_projects_v1beta1_projects_service_proto protoreflect.FileDescriptor

var file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x65, 0x64, 0x68, 0x6f, 0x6c, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x65, 0x64, 0x68, 0x6f, 0x6c, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5b, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x64, 0x68, 0x6f, 0x6c, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x59, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xdb, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x48, 0x0a, 0x0a, 0x67, 0x69, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x65, 0x64, 0x68, 0x6f, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x67, 0x69, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x22, 0x51, 0x0a, 0x09, 0x47, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x69, 0x72, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x69, 0x72, 0x74, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xc0, 0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x33, 0x2e, 0x65, 0x64,
	0x68, 0x6f, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x65, 0x64, 0x68, 0x6f, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x7a, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x12, 0x95, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x31, 0x2e, 0x65, 0x64, 0x68, 0x6f, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x7a, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x64, 0x68, 0x6f, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x7a, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0xa1, 0x02, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x65, 0x64, 0x68, 0x6f, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x14, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x65, 0x64, 0x68, 0x6f, 0x6c, 0x6d,
	0x2e, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x64, 0x68, 0x6f, 0x6c, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x7a, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x50, 0x50, 0xaa, 0x02, 0x1e, 0x45, 0x64, 0x68,
	0x6f, 0x6c, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1e, 0x45, 0x64,
	0x68, 0x6f, 0x6c, 0x6d, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x5c, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2a, 0x45,
	0x64, 0x68, 0x6f, 0x6c, 0x6d, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x5c, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x45, 0x64, 0x68, 0x6f,
	0x6c, 0x6d, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x66, 0x7a, 0x66, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescOnce sync.Once
	file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescData = file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDesc
)

func file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescGZIP() []byte {
	file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescOnce.Do(func() {
		file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescData)
	})
	return file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDescData
}

var file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_edholm_profzf_projects_v1beta1_projects_service_proto_goTypes = []interface{}{
	(*ListProjectsRequest)(nil),   // 0: edholm.profzf.projects.v1beta1.ListProjectsRequest
	(*ListProjectsResponse)(nil),  // 1: edholm.profzf.projects.v1beta1.ListProjectsResponse
	(*GetProjectRequest)(nil),     // 2: edholm.profzf.projects.v1beta1.GetProjectRequest
	(*Project)(nil),               // 3: edholm.profzf.projects.v1beta1.Project
	(*GitStatus)(nil),             // 4: edholm.profzf.projects.v1beta1.GitStatus
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_edholm_profzf_projects_v1beta1_projects_service_proto_depIdxs = []int32{
	3, // 0: edholm.profzf.projects.v1beta1.ListProjectsResponse.projects:type_name -> edholm.profzf.projects.v1beta1.Project
	4, // 1: edholm.profzf.projects.v1beta1.Project.git_status:type_name -> edholm.profzf.projects.v1beta1.GitStatus
	5, // 2: edholm.profzf.projects.v1beta1.Project.last_updated:type_name -> google.protobuf.Timestamp
	0, // 3: edholm.profzf.projects.v1beta1.ProjectsService.ListProjects:input_type -> edholm.profzf.projects.v1beta1.ListProjectsRequest
	2, // 4: edholm.profzf.projects.v1beta1.ProjectsService.GetProject:input_type -> edholm.profzf.projects.v1beta1.GetProjectRequest
	1, // 5: edholm.profzf.projects.v1beta1.ProjectsService.ListProjects:output_type -> edholm.profzf.projects.v1beta1.ListProjectsResponse
	3, // 6: edholm.profzf.projects.v1beta1.ProjectsService.GetProject:output_type -> edholm.profzf.projects.v1beta1.Project
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_edholm_profzf_projects_v1beta1_projects_service_proto_init() }
func file_edholm_profzf_projects_v1beta1_projects_service_proto_init() {
	if File_edholm_profzf_projects_v1beta1_projects_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsRequest); i {
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
		file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsResponse); i {
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
		file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectRequest); i {
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
		file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitStatus); i {
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
			RawDescriptor: file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_edholm_profzf_projects_v1beta1_projects_service_proto_goTypes,
		DependencyIndexes: file_edholm_profzf_projects_v1beta1_projects_service_proto_depIdxs,
		MessageInfos:      file_edholm_profzf_projects_v1beta1_projects_service_proto_msgTypes,
	}.Build()
	File_edholm_profzf_projects_v1beta1_projects_service_proto = out.File
	file_edholm_profzf_projects_v1beta1_projects_service_proto_rawDesc = nil
	file_edholm_profzf_projects_v1beta1_projects_service_proto_goTypes = nil
	file_edholm_profzf_projects_v1beta1_projects_service_proto_depIdxs = nil
}
