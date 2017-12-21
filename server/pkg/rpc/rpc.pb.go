// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	ListEntriesRequest
	GetEntryRequest
	CreateEntryRequest
	UpdateEntryRequest
	EntryList
	ListEntryCommentsRequest
	CommentList
	Entry
	Author
	CreateEntryCommentRequest
	UpdateEntryCommentRequest
	Comment
	Tag
	JWT
	AuthURL
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

type ListEntriesRequest struct {
	Page       int64  `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	NumPerPage int64  `protobuf:"varint,2,opt,name=num_per_page,json=numPerPage" json:"num_per_page,omitempty"`
	Keyword    string `protobuf:"bytes,3,opt,name=keyword" json:"keyword,omitempty"`
}

func (m *ListEntriesRequest) Reset()                    { *m = ListEntriesRequest{} }
func (m *ListEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntriesRequest) ProtoMessage()               {}
func (*ListEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListEntriesRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListEntriesRequest) GetNumPerPage() int64 {
	if m != nil {
		return m.NumPerPage
	}
	return 0
}

func (m *ListEntriesRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type GetEntryRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetEntryRequest) Reset()                    { *m = GetEntryRequest{} }
func (m *GetEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryRequest) ProtoMessage()               {}
func (*GetEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetEntryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateEntryRequest struct {
	Title    string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Content  string   `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	AuthorId string   `protobuf:"bytes,4,opt,name=author_id,json=authorId" json:"author_id,omitempty"`
	TagIds   []string `protobuf:"bytes,5,rep,name=tag_ids,json=tagIds" json:"tag_ids,omitempty"`
}

func (m *CreateEntryRequest) Reset()                    { *m = CreateEntryRequest{} }
func (m *CreateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateEntryRequest) ProtoMessage()               {}
func (*CreateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateEntryRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateEntryRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateEntryRequest) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *CreateEntryRequest) GetTagIds() []string {
	if m != nil {
		return m.TagIds
	}
	return nil
}

type UpdateEntryRequest struct {
	Id       string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Content  string   `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	AuthorId string   `protobuf:"bytes,4,opt,name=author_id,json=authorId" json:"author_id,omitempty"`
	TagIds   []string `protobuf:"bytes,5,rep,name=tag_ids,json=tagIds" json:"tag_ids,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateEntryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateEntryRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateEntryRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *UpdateEntryRequest) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *UpdateEntryRequest) GetTagIds() []string {
	if m != nil {
		return m.TagIds
	}
	return nil
}

type EntryList struct {
	Entries    []*Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	NumResults int64    `protobuf:"varint,2,opt,name=num_results,json=numResults" json:"num_results,omitempty"`
}

func (m *EntryList) Reset()                    { *m = EntryList{} }
func (m *EntryList) String() string            { return proto.CompactTextString(m) }
func (*EntryList) ProtoMessage()               {}
func (*EntryList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EntryList) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *EntryList) GetNumResults() int64 {
	if m != nil {
		return m.NumResults
	}
	return 0
}

type ListEntryCommentsRequest struct {
	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId" json:"entry_id,omitempty"`
}

func (m *ListEntryCommentsRequest) Reset()                    { *m = ListEntryCommentsRequest{} }
func (m *ListEntryCommentsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntryCommentsRequest) ProtoMessage()               {}
func (*ListEntryCommentsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListEntryCommentsRequest) GetEntryId() string {
	if m != nil {
		return m.EntryId
	}
	return ""
}

type CommentList struct {
	Comments []*Comment `protobuf:"bytes,1,rep,name=comments" json:"comments,omitempty"`
}

func (m *CommentList) Reset()                    { *m = CommentList{} }
func (m *CommentList) String() string            { return proto.CompactTextString(m) }
func (*CommentList) ProtoMessage()               {}
func (*CommentList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CommentList) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

type Entry struct {
	Id          string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title       string  `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Content     string  `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	Author      *Author `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
	Tags        []*Tag  `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
	CreatedDate string  `protobuf:"bytes,6,opt,name=created_date,json=createdDate" json:"created_date,omitempty"`
	UpdatedDate string  `protobuf:"bytes,7,opt,name=updated_date,json=updatedDate" json:"updated_date,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Entry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Entry) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Entry) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Entry) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Entry) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Entry) GetCreatedDate() string {
	if m != nil {
		return m.CreatedDate
	}
	return ""
}

func (m *Entry) GetUpdatedDate() string {
	if m != nil {
		return m.UpdatedDate
	}
	return ""
}

type Author struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *Author) Reset()                    { *m = Author{} }
func (m *Author) String() string            { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()               {}
func (*Author) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Author) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Author) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateEntryCommentRequest struct {
	EntryId  string `protobuf:"bytes,1,opt,name=entry_id,json=entryId" json:"entry_id,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId" json:"author_id,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *CreateEntryCommentRequest) Reset()                    { *m = CreateEntryCommentRequest{} }
func (m *CreateEntryCommentRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateEntryCommentRequest) ProtoMessage()               {}
func (*CreateEntryCommentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CreateEntryCommentRequest) GetEntryId() string {
	if m != nil {
		return m.EntryId
	}
	return ""
}

func (m *CreateEntryCommentRequest) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *CreateEntryCommentRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type UpdateEntryCommentRequest struct {
	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId" json:"entry_id,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *UpdateEntryCommentRequest) Reset()                    { *m = UpdateEntryCommentRequest{} }
func (m *UpdateEntryCommentRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryCommentRequest) ProtoMessage()               {}
func (*UpdateEntryCommentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateEntryCommentRequest) GetEntryId() string {
	if m != nil {
		return m.EntryId
	}
	return ""
}

func (m *UpdateEntryCommentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateEntryCommentRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Comment struct {
	Id          string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Author      *Author `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	CreatedDate string  `protobuf:"bytes,3,opt,name=created_date,json=createdDate" json:"created_date,omitempty"`
	Content     string  `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
}

func (m *Comment) Reset()                    { *m = Comment{} }
func (m *Comment) String() string            { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()               {}
func (*Comment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Comment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Comment) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Comment) GetCreatedDate() string {
	if m != nil {
		return m.CreatedDate
	}
	return ""
}

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Tag struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Tag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tag) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type JWT struct {
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
}

func (m *JWT) Reset()                    { *m = JWT{} }
func (m *JWT) String() string            { return proto.CompactTextString(m) }
func (*JWT) ProtoMessage()               {}
func (*JWT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *JWT) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type AuthURL struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *AuthURL) Reset()                    { *m = AuthURL{} }
func (m *AuthURL) String() string            { return proto.CompactTextString(m) }
func (*AuthURL) ProtoMessage()               {}
func (*AuthURL) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *AuthURL) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*ListEntriesRequest)(nil), "rpc.ListEntriesRequest")
	proto.RegisterType((*GetEntryRequest)(nil), "rpc.GetEntryRequest")
	proto.RegisterType((*CreateEntryRequest)(nil), "rpc.CreateEntryRequest")
	proto.RegisterType((*UpdateEntryRequest)(nil), "rpc.UpdateEntryRequest")
	proto.RegisterType((*EntryList)(nil), "rpc.EntryList")
	proto.RegisterType((*ListEntryCommentsRequest)(nil), "rpc.ListEntryCommentsRequest")
	proto.RegisterType((*CommentList)(nil), "rpc.CommentList")
	proto.RegisterType((*Entry)(nil), "rpc.Entry")
	proto.RegisterType((*Author)(nil), "rpc.Author")
	proto.RegisterType((*CreateEntryCommentRequest)(nil), "rpc.CreateEntryCommentRequest")
	proto.RegisterType((*UpdateEntryCommentRequest)(nil), "rpc.UpdateEntryCommentRequest")
	proto.RegisterType((*Comment)(nil), "rpc.Comment")
	proto.RegisterType((*Tag)(nil), "rpc.Tag")
	proto.RegisterType((*JWT)(nil), "rpc.JWT")
	proto.RegisterType((*AuthURL)(nil), "rpc.AuthURL")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KobService service

type KobServiceClient interface {
	ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*EntryList, error)
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*Entry, error)
	CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*Entry, error)
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*Entry, error)
	ListEntryComments(ctx context.Context, in *ListEntryCommentsRequest, opts ...grpc.CallOption) (*CommentList, error)
	CreateEntryComment(ctx context.Context, in *CreateEntryCommentRequest, opts ...grpc.CallOption) (*Comment, error)
	UpdateEntryComment(ctx context.Context, in *UpdateEntryCommentRequest, opts ...grpc.CallOption) (*Comment, error)
	CreateJWT(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*JWT, error)
	CreateAuthURL(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*AuthURL, error)
}

type kobServiceClient struct {
	cc *grpc.ClientConn
}

func NewKobServiceClient(cc *grpc.ClientConn) KobServiceClient {
	return &kobServiceClient{cc}
}

func (c *kobServiceClient) ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*EntryList, error) {
	out := new(EntryList)
	err := grpc.Invoke(ctx, "/rpc.KobService/ListEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kobServiceClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := grpc.Invoke(ctx, "/rpc.KobService/GetEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kobServiceClient) CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := grpc.Invoke(ctx, "/rpc.KobService/CreateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kobServiceClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := grpc.Invoke(ctx, "/rpc.KobService/UpdateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kobServiceClient) ListEntryComments(ctx context.Context, in *ListEntryCommentsRequest, opts ...grpc.CallOption) (*CommentList, error) {
	out := new(CommentList)
	err := grpc.Invoke(ctx, "/rpc.KobService/ListEntryComments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kobServiceClient) CreateEntryComment(ctx context.Context, in *CreateEntryCommentRequest, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := grpc.Invoke(ctx, "/rpc.KobService/CreateEntryComment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kobServiceClient) UpdateEntryComment(ctx context.Context, in *UpdateEntryCommentRequest, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := grpc.Invoke(ctx, "/rpc.KobService/UpdateEntryComment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kobServiceClient) CreateJWT(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*JWT, error) {
	out := new(JWT)
	err := grpc.Invoke(ctx, "/rpc.KobService/CreateJWT", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kobServiceClient) CreateAuthURL(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*AuthURL, error) {
	out := new(AuthURL)
	err := grpc.Invoke(ctx, "/rpc.KobService/CreateAuthURL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KobService service

type KobServiceServer interface {
	ListEntries(context.Context, *ListEntriesRequest) (*EntryList, error)
	GetEntry(context.Context, *GetEntryRequest) (*Entry, error)
	CreateEntry(context.Context, *CreateEntryRequest) (*Entry, error)
	UpdateEntry(context.Context, *UpdateEntryRequest) (*Entry, error)
	ListEntryComments(context.Context, *ListEntryCommentsRequest) (*CommentList, error)
	CreateEntryComment(context.Context, *CreateEntryCommentRequest) (*Comment, error)
	UpdateEntryComment(context.Context, *UpdateEntryCommentRequest) (*Comment, error)
	CreateJWT(context.Context, *google_protobuf1.Empty) (*JWT, error)
	CreateAuthURL(context.Context, *google_protobuf1.Empty) (*AuthURL, error)
}

func RegisterKobServiceServer(s *grpc.Server, srv KobServiceServer) {
	s.RegisterService(&_KobService_serviceDesc, srv)
}

func _KobService_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KobServiceServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KobService/ListEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KobServiceServer).ListEntries(ctx, req.(*ListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KobService_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KobServiceServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KobService/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KobServiceServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KobService_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KobServiceServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KobService/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KobServiceServer).CreateEntry(ctx, req.(*CreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KobService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KobServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KobService/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KobServiceServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KobService_ListEntryComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntryCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KobServiceServer).ListEntryComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KobService/ListEntryComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KobServiceServer).ListEntryComments(ctx, req.(*ListEntryCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KobService_CreateEntryComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KobServiceServer).CreateEntryComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KobService/CreateEntryComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KobServiceServer).CreateEntryComment(ctx, req.(*CreateEntryCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KobService_UpdateEntryComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KobServiceServer).UpdateEntryComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KobService/UpdateEntryComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KobServiceServer).UpdateEntryComment(ctx, req.(*UpdateEntryCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KobService_CreateJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KobServiceServer).CreateJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KobService/CreateJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KobServiceServer).CreateJWT(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KobService_CreateAuthURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KobServiceServer).CreateAuthURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KobService/CreateAuthURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KobServiceServer).CreateAuthURL(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _KobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.KobService",
	HandlerType: (*KobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEntries",
			Handler:    _KobService_ListEntries_Handler,
		},
		{
			MethodName: "GetEntry",
			Handler:    _KobService_GetEntry_Handler,
		},
		{
			MethodName: "CreateEntry",
			Handler:    _KobService_CreateEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _KobService_UpdateEntry_Handler,
		},
		{
			MethodName: "ListEntryComments",
			Handler:    _KobService_ListEntryComments_Handler,
		},
		{
			MethodName: "CreateEntryComment",
			Handler:    _KobService_CreateEntryComment_Handler,
		},
		{
			MethodName: "UpdateEntryComment",
			Handler:    _KobService_UpdateEntryComment_Handler,
		},
		{
			MethodName: "CreateJWT",
			Handler:    _KobService_CreateJWT_Handler,
		},
		{
			MethodName: "CreateAuthURL",
			Handler:    _KobService_CreateAuthURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 804 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4f, 0x6f, 0xfb, 0x44,
	0x10, 0x55, 0xe2, 0xfc, 0xf3, 0x38, 0x2d, 0xed, 0x52, 0x51, 0xc7, 0x29, 0x25, 0x5d, 0x10, 0x8a,
	0x84, 0x94, 0x48, 0x41, 0x88, 0x73, 0x29, 0xa5, 0xb4, 0xf4, 0x50, 0x99, 0x54, 0x3d, 0x06, 0x27,
	0x5e, 0x8c, 0x69, 0x62, 0x1b, 0x7b, 0x0d, 0x0a, 0x55, 0x2f, 0x1c, 0xb9, 0xf2, 0xc5, 0x90, 0xf8,
	0x0a, 0x9c, 0xf9, 0x0c, 0x68, 0x67, 0x77, 0x53, 0xdb, 0x69, 0x00, 0x09, 0xe9, 0x77, 0xf3, 0xee,
	0x8c, 0xe7, 0xbd, 0x9d, 0x37, 0xfb, 0x16, 0xcc, 0x34, 0x59, 0x8c, 0x92, 0x34, 0xe6, 0x31, 0x31,
	0xd2, 0x64, 0xe1, 0x9c, 0x04, 0x71, 0x1c, 0x2c, 0xd9, 0xd8, 0x4b, 0xc2, 0xb1, 0x17, 0x45, 0x31,
	0xf7, 0x78, 0x18, 0x47, 0x99, 0x4c, 0x71, 0xfa, 0x2a, 0x8a, 0xab, 0x79, 0xfe, 0xed, 0x98, 0xad,
	0x12, 0xbe, 0x96, 0x41, 0xea, 0x03, 0xb9, 0x0d, 0x33, 0x7e, 0x19, 0xf1, 0x34, 0x64, 0x99, 0xcb,
	0x7e, 0xc8, 0x59, 0xc6, 0x09, 0x81, 0x46, 0xe2, 0x05, 0xcc, 0xae, 0x0d, 0x6a, 0x43, 0xc3, 0xc5,
	0x6f, 0x32, 0x80, 0x6e, 0x94, 0xaf, 0x66, 0x09, 0x4b, 0x67, 0x18, 0xab, 0x63, 0x0c, 0xa2, 0x7c,
	0x75, 0xc7, 0xd2, 0x3b, 0x91, 0x61, 0x43, 0xfb, 0x91, 0xad, 0x7f, 0x8a, 0x53, 0xdf, 0x36, 0x06,
	0xb5, 0xa1, 0xe9, 0xea, 0x25, 0x3d, 0x83, 0xb7, 0xae, 0x18, 0x82, 0xac, 0x35, 0xc4, 0x3e, 0xd4,
	0x43, 0x1f, 0x01, 0x4c, 0xb7, 0x1e, 0xfa, 0xf4, 0x67, 0x20, 0x17, 0x29, 0xf3, 0x38, 0x2b, 0x65,
	0x1d, 0x41, 0x93, 0x87, 0x7c, 0x29, 0xd1, 0x4c, 0x57, 0x2e, 0x04, 0xd0, 0x22, 0x8e, 0x38, 0x8b,
	0xb8, 0x06, 0x52, 0x4b, 0xd2, 0x07, 0xd3, 0xcb, 0xf9, 0x77, 0x71, 0x3a, 0x0b, 0x7d, 0xbb, 0x81,
	0xb1, 0x8e, 0xdc, 0xb8, 0xf6, 0xc9, 0x31, 0xb4, 0xb9, 0x17, 0xcc, 0x42, 0x3f, 0xb3, 0x9b, 0x03,
	0x63, 0x68, 0xba, 0x2d, 0xee, 0x05, 0xd7, 0x7e, 0x46, 0x7f, 0xad, 0x01, 0xb9, 0x4f, 0xfc, 0x2a,
	0x78, 0x85, 0xe2, 0x1b, 0x22, 0xe3, 0x82, 0x89, 0x2c, 0x84, 0x2c, 0xe4, 0x03, 0x68, 0x33, 0x29,
	0x8d, 0x5d, 0x1b, 0x18, 0x43, 0x6b, 0x02, 0x23, 0xa1, 0xbd, 0xa4, 0xa9, 0x43, 0xe4, 0x3d, 0xb0,
	0x84, 0x34, 0x29, 0xcb, 0xf2, 0x25, 0xcf, 0x0a, 0xca, 0xb8, 0x72, 0x87, 0x7e, 0x02, 0xb6, 0x56,
	0x79, 0x7d, 0x11, 0xaf, 0x56, 0x2c, 0xe2, 0x1b, 0xad, 0x7b, 0xd0, 0x11, 0x75, 0xd6, 0xb3, 0xcd,
	0x59, 0xb1, 0xee, 0xfa, 0xda, 0xa7, 0x9f, 0x82, 0xa5, 0xb2, 0x91, 0xcc, 0x10, 0x3a, 0x0b, 0xf5,
	0xb3, 0x62, 0xd3, 0x45, 0x36, 0x2a, 0xc7, 0xdd, 0x44, 0xe9, 0xef, 0x35, 0x68, 0x22, 0xd8, 0xff,
	0xee, 0xe1, 0xfb, 0xd0, 0x92, 0x2d, 0xc3, 0x06, 0x5a, 0x13, 0x0b, 0x11, 0xcf, 0x71, 0xcb, 0x55,
	0x21, 0x72, 0x02, 0x0d, 0xee, 0x05, 0xb2, 0x91, 0xd6, 0xa4, 0x83, 0x29, 0x53, 0x2f, 0x70, 0x71,
	0x97, 0x9c, 0x41, 0x77, 0x81, 0x93, 0xe5, 0xcf, 0x84, 0xc4, 0x76, 0x0b, 0x11, 0x2c, 0xb5, 0xf7,
	0xb9, 0xc7, 0x99, 0x48, 0xc9, 0x51, 0x7f, 0x95, 0xd2, 0x96, 0x29, 0x6a, 0x4f, 0xa4, 0xd0, 0x11,
	0xb4, 0x24, 0xea, 0x6b, 0x47, 0x62, 0x2b, 0x2f, 0x5c, 0xea, 0x23, 0xe1, 0x82, 0xae, 0xa0, 0x57,
	0x98, 0x67, 0xdd, 0xa2, 0x7f, 0xed, 0x79, 0x79, 0x68, 0xea, 0x95, 0xa1, 0xd9, 0xd9, 0x27, 0xfa,
	0x0d, 0xf4, 0x0a, 0x13, 0xfc, 0xdf, 0xe1, 0xe4, 0x61, 0xea, 0x9b, 0xc3, 0xec, 0x46, 0x78, 0x82,
	0xb6, 0x2a, 0xbb, 0xd5, 0x81, 0x17, 0x91, 0xea, 0xbb, 0x45, 0xaa, 0xca, 0x60, 0x6c, 0xcb, 0x50,
	0x00, 0x6f, 0x94, 0xc1, 0x3f, 0x02, 0x63, 0xea, 0x05, 0xaf, 0xb5, 0x7e, 0xe9, 0xcd, 0xd9, 0xa6,
	0xf5, 0xb8, 0xa0, 0x3d, 0x30, 0x6e, 0x1e, 0xa6, 0xc2, 0xc4, 0x16, 0xb1, 0xcf, 0x54, 0x3a, 0x7e,
	0xd3, 0x3e, 0xb4, 0x05, 0xad, 0x7b, 0xf7, 0x96, 0x1c, 0x80, 0x91, 0xa7, 0x4b, 0x15, 0x15, 0x9f,
	0x93, 0xbf, 0x9a, 0x00, 0x5f, 0xc5, 0xf3, 0xaf, 0x59, 0xfa, 0x63, 0xb8, 0x60, 0xe4, 0x0b, 0xb0,
	0x0a, 0xd6, 0x48, 0x8e, 0xf1, 0x50, 0xdb, 0x66, 0xe9, 0xec, 0xbf, 0x5c, 0x49, 0x11, 0xa5, 0xfb,
	0xbf, 0xfc, 0xf1, 0xe7, 0x6f, 0xf5, 0x0e, 0x69, 0x8d, 0xb1, 0xc9, 0xe4, 0x1c, 0x3a, 0xda, 0xfc,
	0xc8, 0x11, 0xe6, 0x56, 0xbc, 0xd0, 0x29, 0x5c, 0x6a, 0xfa, 0x36, 0xfe, 0xbd, 0x47, 0x2c, 0xf9,
	0xf7, 0xf8, 0x29, 0xf4, 0x9f, 0xc9, 0x67, 0x60, 0x15, 0x86, 0x49, 0x51, 0xd9, 0xb6, 0xcb, 0x52,
	0x21, 0x45, 0x83, 0x6a, 0x1a, 0x57, 0x60, 0x15, 0x26, 0x44, 0xd5, 0xd8, 0x76, 0xbd, 0xd7, 0xc8,
	0x38, 0x25, 0x32, 0x01, 0x1c, 0x6e, 0x99, 0x09, 0x79, 0xb7, 0xd4, 0x9d, 0xaa, 0xc9, 0x38, 0x07,
	0x45, 0xa3, 0xc0, 0x2e, 0x9d, 0x61, 0xe9, 0x3e, 0xe9, 0xe9, 0xd2, 0x7a, 0x42, 0x9f, 0xc7, 0xca,
	0x46, 0x08, 0x2b, 0x3d, 0x09, 0x7a, 0xf8, 0x4e, 0xab, 0x87, 0x2f, 0x0f, 0xbb, 0x53, 0xf2, 0x24,
	0x0d, 0x43, 0xff, 0x01, 0xe6, 0xfb, 0x92, 0xf9, 0x97, 0x61, 0x76, 0xde, 0xa9, 0x0a, 0xcc, 0x87,
	0x08, 0x33, 0x70, 0x4e, 0x77, 0xc2, 0x68, 0x21, 0x4d, 0xc9, 0x5c, 0x0c, 0xe8, 0x3b, 0x23, 0xf9,
	0x32, 0x8f, 0xf4, 0xcb, 0x3c, 0xba, 0x14, 0x2f, 0xb3, 0x23, 0x0d, 0xec, 0xe6, 0x61, 0x5a, 0x18,
	0x06, 0x71, 0x89, 0xc6, 0x3c, 0x7e, 0x64, 0x11, 0xf9, 0x12, 0xf6, 0x64, 0x0d, 0x3d, 0xc9, 0xbb,
	0xea, 0x74, 0x37, 0xd7, 0xf0, 0xde, 0xbd, 0xa5, 0x87, 0x58, 0xcb, 0x22, 0xa6, 0xac, 0x95, 0xa7,
	0xcb, 0x79, 0x0b, 0x7f, 0xf8, 0xf8, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x6d, 0x04, 0x28,
	0x50, 0x08, 0x00, 0x00,
}