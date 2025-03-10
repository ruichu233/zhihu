// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: like.proto

package like

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

type LikeActionRequest_ActionType int32

const (
	LikeActionRequest_UNKNOWN LikeActionRequest_ActionType = 0
	LikeActionRequest_LIKE    LikeActionRequest_ActionType = 1
	LikeActionRequest_UNLIKE  LikeActionRequest_ActionType = 2
)

// Enum value maps for LikeActionRequest_ActionType.
var (
	LikeActionRequest_ActionType_name = map[int32]string{
		0: "UNKNOWN",
		1: "LIKE",
		2: "UNLIKE",
	}
	LikeActionRequest_ActionType_value = map[string]int32{
		"UNKNOWN": 0,
		"LIKE":    1,
		"UNLIKE":  2,
	}
)

func (x LikeActionRequest_ActionType) Enum() *LikeActionRequest_ActionType {
	p := new(LikeActionRequest_ActionType)
	*p = x
	return p
}

func (x LikeActionRequest_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LikeActionRequest_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_like_proto_enumTypes[0].Descriptor()
}

func (LikeActionRequest_ActionType) Type() protoreflect.EnumType {
	return &file_like_proto_enumTypes[0]
}

func (x LikeActionRequest_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LikeActionRequest_ActionType.Descriptor instead.
func (LikeActionRequest_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{2, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type LikeActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId      string                       `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`                                                                     // 业务id
	ObjId      int64                        `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"`                                                                    // 点赞对象id
	UserId     int64                        `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`                                                                  // 用户id
	ActionType LikeActionRequest_ActionType `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3,enum=like.LikeActionRequest_ActionType" json:"action_type,omitempty"` // 要执行的动作类型
}

func (x *LikeActionRequest) Reset() {
	*x = LikeActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeActionRequest) ProtoMessage() {}

func (x *LikeActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeActionRequest.ProtoReflect.Descriptor instead.
func (*LikeActionRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{2}
}

func (x *LikeActionRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *LikeActionRequest) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *LikeActionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LikeActionRequest) GetActionType() LikeActionRequest_ActionType {
	if x != nil {
		return x.ActionType
	}
	return LikeActionRequest_UNKNOWN
}

type LikeActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LikeCount int64 `protobuf:"varint,1,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
}

func (x *LikeActionResponse) Reset() {
	*x = LikeActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeActionResponse) ProtoMessage() {}

func (x *LikeActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeActionResponse.ProtoReflect.Descriptor instead.
func (*LikeActionResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{3}
}

func (x *LikeActionResponse) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

type CheckLikeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId  string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`    // 业务id
	ObjId  int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"`   // 点赞对象id
	UserId int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"` // 用户id
}

func (x *CheckLikeStatusRequest) Reset() {
	*x = CheckLikeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckLikeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckLikeStatusRequest) ProtoMessage() {}

func (x *CheckLikeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckLikeStatusRequest.ProtoReflect.Descriptor instead.
func (*CheckLikeStatusRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{4}
}

func (x *CheckLikeStatusRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *CheckLikeStatusRequest) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *CheckLikeStatusRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CheckLikeStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLiked bool `protobuf:"varint,1,opt,name=is_liked,json=isLiked,proto3" json:"is_liked,omitempty"`
}

func (x *CheckLikeStatusResponse) Reset() {
	*x = CheckLikeStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckLikeStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckLikeStatusResponse) ProtoMessage() {}

func (x *CheckLikeStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckLikeStatusResponse.ProtoReflect.Descriptor instead.
func (*CheckLikeStatusResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{5}
}

func (x *CheckLikeStatusResponse) GetIsLiked() bool {
	if x != nil {
		return x.IsLiked
	}
	return false
}

type GetPostLikeCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`  // 业务id
	ObjId int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"` // 点赞对象id
}

func (x *GetPostLikeCountRequest) Reset() {
	*x = GetPostLikeCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostLikeCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostLikeCountRequest) ProtoMessage() {}

func (x *GetPostLikeCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostLikeCountRequest.ProtoReflect.Descriptor instead.
func (*GetPostLikeCountRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{6}
}

func (x *GetPostLikeCountRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *GetPostLikeCountRequest) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

type GetPostLikeCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetPostLikeCountResponse) Reset() {
	*x = GetPostLikeCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostLikeCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostLikeCountResponse) ProtoMessage() {}

func (x *GetPostLikeCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostLikeCountResponse.ProtoReflect.Descriptor instead.
func (*GetPostLikeCountResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{7}
}

func (x *GetPostLikeCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetUserLikesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserLikesRequest) Reset() {
	*x = GetUserLikesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLikesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLikesRequest) ProtoMessage() {}

func (x *GetUserLikesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLikesRequest.ProtoReflect.Descriptor instead.
func (*GetUserLikesRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserLikesRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserLikesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostIds []int64 `protobuf:"varint,1,rep,packed,name=post_ids,json=postIds,proto3" json:"post_ids,omitempty"`
}

func (x *GetUserLikesResponse) Reset() {
	*x = GetUserLikesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLikesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLikesResponse) ProtoMessage() {}

func (x *GetUserLikesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLikesResponse.ProtoReflect.Descriptor instead.
func (*GetUserLikesResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserLikesResponse) GetPostIds() []int64 {
	if x != nil {
		return x.PostIds
	}
	return nil
}

type GetPostLikersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`  // 业务id
	ObjId int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"` // 点赞对象id
}

func (x *GetPostLikersRequest) Reset() {
	*x = GetPostLikersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostLikersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostLikersRequest) ProtoMessage() {}

func (x *GetPostLikersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostLikersRequest.ProtoReflect.Descriptor instead.
func (*GetPostLikersRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{10}
}

func (x *GetPostLikersRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *GetPostLikersRequest) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

type GetPostLikersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []int64 `protobuf:"varint,1,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *GetPostLikersResponse) Reset() {
	*x = GetPostLikersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostLikersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostLikersResponse) ProtoMessage() {}

func (x *GetPostLikersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostLikersResponse.ProtoReflect.Descriptor instead.
func (*GetPostLikersResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{11}
}

func (x *GetPostLikersResponse) GetUserIds() []int64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type GetUserTotalLikesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserTotalLikesRequest) Reset() {
	*x = GetUserTotalLikesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTotalLikesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTotalLikesRequest) ProtoMessage() {}

func (x *GetUserTotalLikesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTotalLikesRequest.ProtoReflect.Descriptor instead.
func (*GetUserTotalLikesRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{12}
}

func (x *GetUserTotalLikesRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserTotalLikesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalLikes int64 `protobuf:"varint,1,opt,name=total_likes,json=totalLikes,proto3" json:"total_likes,omitempty"`
}

func (x *GetUserTotalLikesResponse) Reset() {
	*x = GetUserTotalLikesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTotalLikesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTotalLikesResponse) ProtoMessage() {}

func (x *GetUserTotalLikesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTotalLikesResponse.ProtoReflect.Descriptor instead.
func (*GetUserTotalLikesResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{13}
}

func (x *GetUserTotalLikesResponse) GetTotalLikes() int64 {
	if x != nil {
		return x.TotalLikes
	}
	return 0
}

var File_like_proto protoreflect.FileDescriptor

var file_like_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69,
	0x6b, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e,
	0x67, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x22, 0xcd, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x62,
	0x6a, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x2f, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c,
	0x49, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4c, 0x49, 0x4b, 0x45, 0x10,
	0x02, 0x22, 0x33, 0x0a, 0x12, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x69, 0x6b,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c,
	0x69, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x6b,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x22, 0x45, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49,
	0x64, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x73, 0x22, 0x42, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x4c, 0x69, 0x6b, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x69, 0x7a, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x33,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x69,
	0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6b, 0x65,
	0x73, 0x32, 0xf8, 0x03, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e,
	0x4c, 0x69, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x6b, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4c, 0x69, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e,
	0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x69, 0x6b, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x6c, 0x69,
	0x6b, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x69,
	0x6b, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_like_proto_rawDescOnce sync.Once
	file_like_proto_rawDescData = file_like_proto_rawDesc
)

func file_like_proto_rawDescGZIP() []byte {
	file_like_proto_rawDescOnce.Do(func() {
		file_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_like_proto_rawDescData)
	})
	return file_like_proto_rawDescData
}

var file_like_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_like_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_like_proto_goTypes = []any{
	(LikeActionRequest_ActionType)(0), // 0: like.LikeActionRequest.ActionType
	(*Request)(nil),                   // 1: like.Request
	(*Response)(nil),                  // 2: like.Response
	(*LikeActionRequest)(nil),         // 3: like.LikeActionRequest
	(*LikeActionResponse)(nil),        // 4: like.LikeActionResponse
	(*CheckLikeStatusRequest)(nil),    // 5: like.CheckLikeStatusRequest
	(*CheckLikeStatusResponse)(nil),   // 6: like.CheckLikeStatusResponse
	(*GetPostLikeCountRequest)(nil),   // 7: like.GetPostLikeCountRequest
	(*GetPostLikeCountResponse)(nil),  // 8: like.GetPostLikeCountResponse
	(*GetUserLikesRequest)(nil),       // 9: like.GetUserLikesRequest
	(*GetUserLikesResponse)(nil),      // 10: like.GetUserLikesResponse
	(*GetPostLikersRequest)(nil),      // 11: like.GetPostLikersRequest
	(*GetPostLikersResponse)(nil),     // 12: like.GetPostLikersResponse
	(*GetUserTotalLikesRequest)(nil),  // 13: like.GetUserTotalLikesRequest
	(*GetUserTotalLikesResponse)(nil), // 14: like.GetUserTotalLikesResponse
}
var file_like_proto_depIdxs = []int32{
	0,  // 0: like.LikeActionRequest.action_type:type_name -> like.LikeActionRequest.ActionType
	1,  // 1: like.Like.Ping:input_type -> like.Request
	3,  // 2: like.Like.LikeAction:input_type -> like.LikeActionRequest
	5,  // 3: like.Like.CheckLikeStatus:input_type -> like.CheckLikeStatusRequest
	7,  // 4: like.Like.GetPostLikeCount:input_type -> like.GetPostLikeCountRequest
	9,  // 5: like.Like.GetUserLikes:input_type -> like.GetUserLikesRequest
	11, // 6: like.Like.GetPostLikers:input_type -> like.GetPostLikersRequest
	13, // 7: like.Like.GetUserTotalLikes:input_type -> like.GetUserTotalLikesRequest
	2,  // 8: like.Like.Ping:output_type -> like.Response
	4,  // 9: like.Like.LikeAction:output_type -> like.LikeActionResponse
	6,  // 10: like.Like.CheckLikeStatus:output_type -> like.CheckLikeStatusResponse
	8,  // 11: like.Like.GetPostLikeCount:output_type -> like.GetPostLikeCountResponse
	10, // 12: like.Like.GetUserLikes:output_type -> like.GetUserLikesResponse
	12, // 13: like.Like.GetPostLikers:output_type -> like.GetPostLikersResponse
	14, // 14: like.Like.GetUserTotalLikes:output_type -> like.GetUserTotalLikesResponse
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_like_proto_init() }
func file_like_proto_init() {
	if File_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_like_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Request); i {
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
		file_like_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
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
		file_like_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LikeActionRequest); i {
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
		file_like_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LikeActionResponse); i {
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
		file_like_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CheckLikeStatusRequest); i {
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
		file_like_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CheckLikeStatusResponse); i {
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
		file_like_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetPostLikeCountRequest); i {
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
		file_like_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetPostLikeCountResponse); i {
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
		file_like_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserLikesRequest); i {
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
		file_like_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserLikesResponse); i {
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
		file_like_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetPostLikersRequest); i {
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
		file_like_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetPostLikersResponse); i {
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
		file_like_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserTotalLikesRequest); i {
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
		file_like_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserTotalLikesResponse); i {
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
			RawDescriptor: file_like_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_like_proto_goTypes,
		DependencyIndexes: file_like_proto_depIdxs,
		EnumInfos:         file_like_proto_enumTypes,
		MessageInfos:      file_like_proto_msgTypes,
	}.Build()
	File_like_proto = out.File
	file_like_proto_rawDesc = nil
	file_like_proto_goTypes = nil
	file_like_proto_depIdxs = nil
}
