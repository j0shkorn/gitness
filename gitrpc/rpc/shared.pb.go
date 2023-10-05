// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: shared.proto

package rpc

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

type SortOrder int32

const (
	SortOrder_Default SortOrder = 0
	SortOrder_Asc     SortOrder = 1
	SortOrder_Desc    SortOrder = 2
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "Default",
		1: "Asc",
		2: "Desc",
	}
	SortOrder_value = map[string]int32{
		"Default": 0,
		"Asc":     1,
		"Desc":    2,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_shared_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_shared_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{0}
}

type RefType int32

const (
	RefType_Undefined       RefType = 0
	RefType_RefRaw          RefType = 1
	RefType_RefBranch       RefType = 2
	RefType_RefTag          RefType = 3
	RefType_RefPullReqHead  RefType = 4
	RefType_RefPullReqMerge RefType = 5
)

// Enum value maps for RefType.
var (
	RefType_name = map[int32]string{
		0: "Undefined",
		1: "RefRaw",
		2: "RefBranch",
		3: "RefTag",
		4: "RefPullReqHead",
		5: "RefPullReqMerge",
	}
	RefType_value = map[string]int32{
		"Undefined":       0,
		"RefRaw":          1,
		"RefBranch":       2,
		"RefTag":          3,
		"RefPullReqHead":  4,
		"RefPullReqMerge": 5,
	}
)

func (x RefType) Enum() *RefType {
	p := new(RefType)
	*p = x
	return p
}

func (x RefType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RefType) Descriptor() protoreflect.EnumDescriptor {
	return file_shared_proto_enumTypes[1].Descriptor()
}

func (RefType) Type() protoreflect.EnumType {
	return &file_shared_proto_enumTypes[1]
}

func (x RefType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RefType.Descriptor instead.
func (RefType) EnumDescriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{1}
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoUid string `protobuf:"bytes,1,opt,name=repo_uid,json=repoUid,proto3" json:"repo_uid,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRequest) GetRepoUid() string {
	if x != nil {
		return x.RepoUid
	}
	return ""
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoUid string    `protobuf:"bytes,1,opt,name=repo_uid,json=repoUid,proto3" json:"repo_uid,omitempty"`
	EnvVars []*EnvVar `protobuf:"bytes,2,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty"`
	Actor   *Identity `protobuf:"bytes,3,opt,name=actor,proto3" json:"actor,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{1}
}

func (x *WriteRequest) GetRepoUid() string {
	if x != nil {
		return x.RepoUid
	}
	return ""
}

func (x *WriteRequest) GetEnvVars() []*EnvVar {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

func (x *WriteRequest) GetActor() *Identity {
	if x != nil {
		return x.Actor
	}
	return nil
}

type EnvVar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EnvVar) Reset() {
	*x = EnvVar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvVar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvVar) ProtoMessage() {}

func (x *EnvVar) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvVar.ProtoReflect.Descriptor instead.
func (*EnvVar) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{2}
}

func (x *EnvVar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvVar) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type FileUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*FileUpload_Header
	//	*FileUpload_Chunk
	Data isFileUpload_Data `protobuf_oneof:"data"`
}

func (x *FileUpload) Reset() {
	*x = FileUpload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUpload) ProtoMessage() {}

func (x *FileUpload) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUpload.ProtoReflect.Descriptor instead.
func (*FileUpload) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{3}
}

func (m *FileUpload) GetData() isFileUpload_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *FileUpload) GetHeader() *FileUploadHeader {
	if x, ok := x.GetData().(*FileUpload_Header); ok {
		return x.Header
	}
	return nil
}

func (x *FileUpload) GetChunk() *Chunk {
	if x, ok := x.GetData().(*FileUpload_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isFileUpload_Data interface {
	isFileUpload_Data()
}

type FileUpload_Header struct {
	Header *FileUploadHeader `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type FileUpload_Chunk struct {
	Chunk *Chunk `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*FileUpload_Header) isFileUpload_Data() {}

func (*FileUpload_Chunk) isFileUpload_Data() {}

type FileUploadHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FileUploadHeader) Reset() {
	*x = FileUploadHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadHeader) ProtoMessage() {}

func (x *FileUploadHeader) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadHeader.ProtoReflect.Descriptor instead.
func (*FileUploadHeader) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{4}
}

func (x *FileUploadHeader) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eof  bool   `protobuf:"varint,1,opt,name=eof,proto3" json:"eof,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{5}
}

func (x *Chunk) GetEof() bool {
	if x != nil {
		return x.Eof
	}
	return false
}

func (x *Chunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha       string     `protobuf:"bytes,1,opt,name=sha,proto3" json:"sha,omitempty"`
	Title     string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Message   string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Author    *Signature `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Committer *Signature `protobuf:"bytes,5,opt,name=committer,proto3" json:"committer,omitempty"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{6}
}

func (x *Commit) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

func (x *Commit) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Commit) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Commit) GetAuthor() *Signature {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Commit) GetCommitter() *Signature {
	if x != nil {
		return x.Committer
	}
	return nil
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *Identity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	When     int64     `protobuf:"varint,2,opt,name=when,proto3" json:"when,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{7}
}

func (x *Signature) GetIdentity() *Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Signature) GetWhen() int64 {
	if x != nil {
		return x.When
	}
	return 0
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{8}
}

func (x *Identity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Identity) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// PathNotFoundError is an error returned in the case a provided path is not found in the repo.
type PathNotFoundError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// path is the path that wasn't found in the repo.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PathNotFoundError) Reset() {
	*x = PathNotFoundError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathNotFoundError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathNotFoundError) ProtoMessage() {}

func (x *PathNotFoundError) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathNotFoundError.ProtoReflect.Descriptor instead.
func (*PathNotFoundError) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{9}
}

func (x *PathNotFoundError) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_shared_proto protoreflect.FileDescriptor

var file_shared_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x72, 0x70, 0x63, 0x22, 0x28, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x69, 0x64, 0x22, 0x76, 0x0a,
	0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f,
	0x76, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x32, 0x0a, 0x06, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x69, 0x0a, 0x0a, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2d, 0x0a, 0x05,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x65, 0x6f, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa0, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x68, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x2c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x22, 0x4a,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x22, 0x34, 0x0a, 0x08, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x27, 0x0a, 0x11, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x2a, 0x2b, 0x0a, 0x09, 0x53, 0x6f, 0x72,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x73, 0x63, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x44, 0x65, 0x73, 0x63, 0x10, 0x02, 0x2a, 0x68, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x66, 0x52, 0x61, 0x77, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x52, 0x65, 0x66, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x65, 0x66, 0x54, 0x61, 0x67, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x50, 0x75,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x48, 0x65, 0x61, 0x64, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x52,
	0x65, 0x66, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x10, 0x05,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x67,
	0x69, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_shared_proto_rawDescOnce sync.Once
	file_shared_proto_rawDescData = file_shared_proto_rawDesc
)

func file_shared_proto_rawDescGZIP() []byte {
	file_shared_proto_rawDescOnce.Do(func() {
		file_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_proto_rawDescData)
	})
	return file_shared_proto_rawDescData
}

var file_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_shared_proto_goTypes = []interface{}{
	(SortOrder)(0),            // 0: rpc.SortOrder
	(RefType)(0),              // 1: rpc.RefType
	(*ReadRequest)(nil),       // 2: rpc.ReadRequest
	(*WriteRequest)(nil),      // 3: rpc.WriteRequest
	(*EnvVar)(nil),            // 4: rpc.EnvVar
	(*FileUpload)(nil),        // 5: rpc.FileUpload
	(*FileUploadHeader)(nil),  // 6: rpc.FileUploadHeader
	(*Chunk)(nil),             // 7: rpc.Chunk
	(*Commit)(nil),            // 8: rpc.Commit
	(*Signature)(nil),         // 9: rpc.Signature
	(*Identity)(nil),          // 10: rpc.Identity
	(*PathNotFoundError)(nil), // 11: rpc.PathNotFoundError
}
var file_shared_proto_depIdxs = []int32{
	4,  // 0: rpc.WriteRequest.env_vars:type_name -> rpc.EnvVar
	10, // 1: rpc.WriteRequest.actor:type_name -> rpc.Identity
	6,  // 2: rpc.FileUpload.header:type_name -> rpc.FileUploadHeader
	7,  // 3: rpc.FileUpload.chunk:type_name -> rpc.Chunk
	9,  // 4: rpc.Commit.author:type_name -> rpc.Signature
	9,  // 5: rpc.Commit.committer:type_name -> rpc.Signature
	10, // 6: rpc.Signature.identity:type_name -> rpc.Identity
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_shared_proto_init() }
func file_shared_proto_init() {
	if File_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvVar); i {
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
		file_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUpload); i {
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
		file_shared_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadHeader); i {
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
		file_shared_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_shared_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit); i {
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
		file_shared_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_shared_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_shared_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathNotFoundError); i {
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
	file_shared_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*FileUpload_Header)(nil),
		(*FileUpload_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_proto_goTypes,
		DependencyIndexes: file_shared_proto_depIdxs,
		EnumInfos:         file_shared_proto_enumTypes,
		MessageInfos:      file_shared_proto_msgTypes,
	}.Build()
	File_shared_proto = out.File
	file_shared_proto_rawDesc = nil
	file_shared_proto_goTypes = nil
	file_shared_proto_depIdxs = nil
}
