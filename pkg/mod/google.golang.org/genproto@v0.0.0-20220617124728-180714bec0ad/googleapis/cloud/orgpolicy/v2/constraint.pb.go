// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/orgpolicy/v2/constraint.proto

package orgpolicy

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specifies the default behavior in the absence of any `Policy` for the
// `Constraint`. This must not be `CONSTRAINT_DEFAULT_UNSPECIFIED`.
//
// Immutable after creation.
type Constraint_ConstraintDefault int32

const (
	// This is only used for distinguishing unset values and should never be
	// used.
	Constraint_CONSTRAINT_DEFAULT_UNSPECIFIED Constraint_ConstraintDefault = 0
	// Indicate that all values are allowed for list constraints.
	// Indicate that enforcement is off for boolean constraints.
	Constraint_ALLOW Constraint_ConstraintDefault = 1
	// Indicate that all values are denied for list constraints.
	// Indicate that enforcement is on for boolean constraints.
	Constraint_DENY Constraint_ConstraintDefault = 2
)

// Enum value maps for Constraint_ConstraintDefault.
var (
	Constraint_ConstraintDefault_name = map[int32]string{
		0: "CONSTRAINT_DEFAULT_UNSPECIFIED",
		1: "ALLOW",
		2: "DENY",
	}
	Constraint_ConstraintDefault_value = map[string]int32{
		"CONSTRAINT_DEFAULT_UNSPECIFIED": 0,
		"ALLOW":                          1,
		"DENY":                           2,
	}
)

func (x Constraint_ConstraintDefault) Enum() *Constraint_ConstraintDefault {
	p := new(Constraint_ConstraintDefault)
	*p = x
	return p
}

func (x Constraint_ConstraintDefault) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Constraint_ConstraintDefault) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_orgpolicy_v2_constraint_proto_enumTypes[0].Descriptor()
}

func (Constraint_ConstraintDefault) Type() protoreflect.EnumType {
	return &file_google_cloud_orgpolicy_v2_constraint_proto_enumTypes[0]
}

func (x Constraint_ConstraintDefault) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Constraint_ConstraintDefault.Descriptor instead.
func (Constraint_ConstraintDefault) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_orgpolicy_v2_constraint_proto_rawDescGZIP(), []int{0, 0}
}

// A `constraint` describes a way to restrict resource's configuration. For
// example, you could enforce a constraint that controls which cloud services
// can be activated across an organization, or whether a Compute Engine instance
// can have serial port connections established. `Constraints` can be configured
// by the organization's policy administrator to fit the needs of the
// organization by setting a `policy` that includes `constraints` at different
// locations in the organization's resource hierarchy. Policies are inherited
// down the resource hierarchy from higher levels, but can also be overridden.
// For details about the inheritance rules please read about
// [`policies`][google.cloud.OrgPolicy.v2.Policy].
//
// `Constraints` have a default behavior determined by the `constraint_default`
// field, which is the enforcement behavior that is used in the absence of a
// `policy` being defined or inherited for the resource in question.
type Constraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the Constraint. Must be in one of
	// the following forms:
	// * `projects/{project_number}/constraints/{constraint_name}`
	// * `folders/{folder_id}/constraints/{constraint_name}`
	// * `organizations/{organization_id}/constraints/{constraint_name}`
	//
	// For example, "/projects/123/constraints/compute.disableSerialPortAccess".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The human readable name.
	//
	// Mutable.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Detailed description of what this `Constraint` controls as well as how and
	// where it is enforced.
	//
	// Mutable.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The evaluation behavior of this constraint in the absence of 'Policy'.
	ConstraintDefault Constraint_ConstraintDefault `protobuf:"varint,4,opt,name=constraint_default,json=constraintDefault,proto3,enum=google.cloud.orgpolicy.v2.Constraint_ConstraintDefault" json:"constraint_default,omitempty"`
	// The type of restrictions for this `Constraint`.
	//
	// Immutable after creation.
	//
	// Types that are assignable to ConstraintType:
	//	*Constraint_ListConstraint_
	//	*Constraint_BooleanConstraint_
	ConstraintType isConstraint_ConstraintType `protobuf_oneof:"constraint_type"`
}

func (x *Constraint) Reset() {
	*x = Constraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Constraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constraint) ProtoMessage() {}

func (x *Constraint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constraint.ProtoReflect.Descriptor instead.
func (*Constraint) Descriptor() ([]byte, []int) {
	return file_google_cloud_orgpolicy_v2_constraint_proto_rawDescGZIP(), []int{0}
}

func (x *Constraint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Constraint) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Constraint) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Constraint) GetConstraintDefault() Constraint_ConstraintDefault {
	if x != nil {
		return x.ConstraintDefault
	}
	return Constraint_CONSTRAINT_DEFAULT_UNSPECIFIED
}

func (m *Constraint) GetConstraintType() isConstraint_ConstraintType {
	if m != nil {
		return m.ConstraintType
	}
	return nil
}

func (x *Constraint) GetListConstraint() *Constraint_ListConstraint {
	if x, ok := x.GetConstraintType().(*Constraint_ListConstraint_); ok {
		return x.ListConstraint
	}
	return nil
}

func (x *Constraint) GetBooleanConstraint() *Constraint_BooleanConstraint {
	if x, ok := x.GetConstraintType().(*Constraint_BooleanConstraint_); ok {
		return x.BooleanConstraint
	}
	return nil
}

type isConstraint_ConstraintType interface {
	isConstraint_ConstraintType()
}

type Constraint_ListConstraint_ struct {
	// Defines this constraint as being a ListConstraint.
	ListConstraint *Constraint_ListConstraint `protobuf:"bytes,5,opt,name=list_constraint,json=listConstraint,proto3,oneof"`
}

type Constraint_BooleanConstraint_ struct {
	// Defines this constraint as being a BooleanConstraint.
	BooleanConstraint *Constraint_BooleanConstraint `protobuf:"bytes,6,opt,name=boolean_constraint,json=booleanConstraint,proto3,oneof"`
}

func (*Constraint_ListConstraint_) isConstraint_ConstraintType() {}

func (*Constraint_BooleanConstraint_) isConstraint_ConstraintType() {}

// A `Constraint` that allows or disallows a list of string values, which are
// configured by an Organization's policy administrator with a `Policy`.
type Constraint_ListConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates whether values grouped into categories can be used in
	// `Policy.allowed_values` and `Policy.denied_values`. For example,
	// `"in:Python"` would match any value in the 'Python' group.
	SupportsIn bool `protobuf:"varint,1,opt,name=supports_in,json=supportsIn,proto3" json:"supports_in,omitempty"`
	// Indicates whether subtrees of Cloud Resource Manager resource hierarchy
	// can be used in `Policy.allowed_values` and `Policy.denied_values`. For
	// example, `"under:folders/123"` would match any resource under the
	// 'folders/123' folder.
	SupportsUnder bool `protobuf:"varint,2,opt,name=supports_under,json=supportsUnder,proto3" json:"supports_under,omitempty"`
}

func (x *Constraint_ListConstraint) Reset() {
	*x = Constraint_ListConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Constraint_ListConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constraint_ListConstraint) ProtoMessage() {}

func (x *Constraint_ListConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constraint_ListConstraint.ProtoReflect.Descriptor instead.
func (*Constraint_ListConstraint) Descriptor() ([]byte, []int) {
	return file_google_cloud_orgpolicy_v2_constraint_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Constraint_ListConstraint) GetSupportsIn() bool {
	if x != nil {
		return x.SupportsIn
	}
	return false
}

func (x *Constraint_ListConstraint) GetSupportsUnder() bool {
	if x != nil {
		return x.SupportsUnder
	}
	return false
}

// A `Constraint` that is either enforced or not.
//
// For example a constraint `constraints/compute.disableSerialPortAccess`.
// If it is enforced on a VM instance, serial port connections will not be
// opened to that instance.
type Constraint_BooleanConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Constraint_BooleanConstraint) Reset() {
	*x = Constraint_BooleanConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Constraint_BooleanConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constraint_BooleanConstraint) ProtoMessage() {}

func (x *Constraint_BooleanConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constraint_BooleanConstraint.ProtoReflect.Descriptor instead.
func (*Constraint_BooleanConstraint) Descriptor() ([]byte, []int) {
	return file_google_cloud_orgpolicy_v2_constraint_proto_rawDescGZIP(), []int{0, 1}
}

var File_google_cloud_orgpolicy_v2_constraint_proto protoreflect.FileDescriptor

var file_google_cloud_orgpolicy_v2_constraint_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x06, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x66, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x5f, 0x0a, 0x0f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x68, 0x0a, 0x12, 0x62, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x11, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x74, 0x1a, 0x58, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x49, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0x13, 0x0a,
	0x11, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x4f, 0x4e, 0x53, 0x54,
	0x52, 0x41, 0x49, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x02,
	0x3a, 0xb8, 0x01, 0xea, 0x41, 0xb4, 0x01, 0x0a, 0x23, 0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x2b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x7d, 0x12, 0x29, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x74, 0x7d, 0x12, 0x35, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x7d, 0x42, 0x11, 0x0a, 0x0f, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xcd,
	0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x32,
	0x42, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x32, 0x3b, 0x6f, 0x72,
	0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4f, 0x72, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x56, 0x32, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x4f, 0x72, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c, 0x56, 0x32, 0xea,
	0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x4f, 0x72, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_orgpolicy_v2_constraint_proto_rawDescOnce sync.Once
	file_google_cloud_orgpolicy_v2_constraint_proto_rawDescData = file_google_cloud_orgpolicy_v2_constraint_proto_rawDesc
)

func file_google_cloud_orgpolicy_v2_constraint_proto_rawDescGZIP() []byte {
	file_google_cloud_orgpolicy_v2_constraint_proto_rawDescOnce.Do(func() {
		file_google_cloud_orgpolicy_v2_constraint_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_orgpolicy_v2_constraint_proto_rawDescData)
	})
	return file_google_cloud_orgpolicy_v2_constraint_proto_rawDescData
}

var file_google_cloud_orgpolicy_v2_constraint_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_orgpolicy_v2_constraint_proto_goTypes = []interface{}{
	(Constraint_ConstraintDefault)(0),    // 0: google.cloud.orgpolicy.v2.Constraint.ConstraintDefault
	(*Constraint)(nil),                   // 1: google.cloud.orgpolicy.v2.Constraint
	(*Constraint_ListConstraint)(nil),    // 2: google.cloud.orgpolicy.v2.Constraint.ListConstraint
	(*Constraint_BooleanConstraint)(nil), // 3: google.cloud.orgpolicy.v2.Constraint.BooleanConstraint
}
var file_google_cloud_orgpolicy_v2_constraint_proto_depIdxs = []int32{
	0, // 0: google.cloud.orgpolicy.v2.Constraint.constraint_default:type_name -> google.cloud.orgpolicy.v2.Constraint.ConstraintDefault
	2, // 1: google.cloud.orgpolicy.v2.Constraint.list_constraint:type_name -> google.cloud.orgpolicy.v2.Constraint.ListConstraint
	3, // 2: google.cloud.orgpolicy.v2.Constraint.boolean_constraint:type_name -> google.cloud.orgpolicy.v2.Constraint.BooleanConstraint
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_orgpolicy_v2_constraint_proto_init() }
func file_google_cloud_orgpolicy_v2_constraint_proto_init() {
	if File_google_cloud_orgpolicy_v2_constraint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Constraint); i {
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
		file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Constraint_ListConstraint); i {
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
		file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Constraint_BooleanConstraint); i {
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
	file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Constraint_ListConstraint_)(nil),
		(*Constraint_BooleanConstraint_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_orgpolicy_v2_constraint_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_orgpolicy_v2_constraint_proto_goTypes,
		DependencyIndexes: file_google_cloud_orgpolicy_v2_constraint_proto_depIdxs,
		EnumInfos:         file_google_cloud_orgpolicy_v2_constraint_proto_enumTypes,
		MessageInfos:      file_google_cloud_orgpolicy_v2_constraint_proto_msgTypes,
	}.Build()
	File_google_cloud_orgpolicy_v2_constraint_proto = out.File
	file_google_cloud_orgpolicy_v2_constraint_proto_rawDesc = nil
	file_google_cloud_orgpolicy_v2_constraint_proto_goTypes = nil
	file_google_cloud_orgpolicy_v2_constraint_proto_depIdxs = nil
}
