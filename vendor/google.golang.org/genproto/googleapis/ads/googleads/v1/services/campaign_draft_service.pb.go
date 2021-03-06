// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_draft_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [CampaignDraftService.GetCampaignDraft][google.ads.googleads.v1.services.CampaignDraftService.GetCampaignDraft].
type GetCampaignDraftRequest struct {
	// The resource name of the campaign draft to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignDraftRequest) Reset()         { *m = GetCampaignDraftRequest{} }
func (m *GetCampaignDraftRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignDraftRequest) ProtoMessage()    {}
func (*GetCampaignDraftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f742bb1c3712c22e, []int{0}
}

func (m *GetCampaignDraftRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignDraftRequest.Unmarshal(m, b)
}
func (m *GetCampaignDraftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignDraftRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignDraftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignDraftRequest.Merge(m, src)
}
func (m *GetCampaignDraftRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignDraftRequest.Size(m)
}
func (m *GetCampaignDraftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignDraftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignDraftRequest proto.InternalMessageInfo

func (m *GetCampaignDraftRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignDraftService.MutateCampaignDrafts][google.ads.googleads.v1.services.CampaignDraftService.MutateCampaignDrafts].
type MutateCampaignDraftsRequest struct {
	// The ID of the customer whose campaign drafts are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual campaign drafts.
	Operations []*CampaignDraftOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignDraftsRequest) Reset()         { *m = MutateCampaignDraftsRequest{} }
func (m *MutateCampaignDraftsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignDraftsRequest) ProtoMessage()    {}
func (*MutateCampaignDraftsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f742bb1c3712c22e, []int{1}
}

func (m *MutateCampaignDraftsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignDraftsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignDraftsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignDraftsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignDraftsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignDraftsRequest.Merge(m, src)
}
func (m *MutateCampaignDraftsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignDraftsRequest.Size(m)
}
func (m *MutateCampaignDraftsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignDraftsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignDraftsRequest proto.InternalMessageInfo

func (m *MutateCampaignDraftsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignDraftsRequest) GetOperations() []*CampaignDraftOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignDraftsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignDraftsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// Request message for [CampaignDraftService.PromoteCampaignDraft][google.ads.googleads.v1.services.CampaignDraftService.PromoteCampaignDraft].
type PromoteCampaignDraftRequest struct {
	// The resource name of the campaign draft to promote.
	CampaignDraft        string   `protobuf:"bytes,1,opt,name=campaign_draft,json=campaignDraft,proto3" json:"campaign_draft,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromoteCampaignDraftRequest) Reset()         { *m = PromoteCampaignDraftRequest{} }
func (m *PromoteCampaignDraftRequest) String() string { return proto.CompactTextString(m) }
func (*PromoteCampaignDraftRequest) ProtoMessage()    {}
func (*PromoteCampaignDraftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f742bb1c3712c22e, []int{2}
}

func (m *PromoteCampaignDraftRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromoteCampaignDraftRequest.Unmarshal(m, b)
}
func (m *PromoteCampaignDraftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromoteCampaignDraftRequest.Marshal(b, m, deterministic)
}
func (m *PromoteCampaignDraftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromoteCampaignDraftRequest.Merge(m, src)
}
func (m *PromoteCampaignDraftRequest) XXX_Size() int {
	return xxx_messageInfo_PromoteCampaignDraftRequest.Size(m)
}
func (m *PromoteCampaignDraftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PromoteCampaignDraftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PromoteCampaignDraftRequest proto.InternalMessageInfo

func (m *PromoteCampaignDraftRequest) GetCampaignDraft() string {
	if m != nil {
		return m.CampaignDraft
	}
	return ""
}

// A single operation (create, update, remove) on a campaign draft.
type CampaignDraftOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignDraftOperation_Create
	//	*CampaignDraftOperation_Update
	//	*CampaignDraftOperation_Remove
	Operation            isCampaignDraftOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CampaignDraftOperation) Reset()         { *m = CampaignDraftOperation{} }
func (m *CampaignDraftOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignDraftOperation) ProtoMessage()    {}
func (*CampaignDraftOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f742bb1c3712c22e, []int{3}
}

func (m *CampaignDraftOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignDraftOperation.Unmarshal(m, b)
}
func (m *CampaignDraftOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignDraftOperation.Marshal(b, m, deterministic)
}
func (m *CampaignDraftOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignDraftOperation.Merge(m, src)
}
func (m *CampaignDraftOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignDraftOperation.Size(m)
}
func (m *CampaignDraftOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignDraftOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignDraftOperation proto.InternalMessageInfo

func (m *CampaignDraftOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignDraftOperation_Operation interface {
	isCampaignDraftOperation_Operation()
}

type CampaignDraftOperation_Create struct {
	Create *resources.CampaignDraft `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignDraftOperation_Update struct {
	Update *resources.CampaignDraft `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignDraftOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignDraftOperation_Create) isCampaignDraftOperation_Operation() {}

func (*CampaignDraftOperation_Update) isCampaignDraftOperation_Operation() {}

func (*CampaignDraftOperation_Remove) isCampaignDraftOperation_Operation() {}

func (m *CampaignDraftOperation) GetOperation() isCampaignDraftOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignDraftOperation) GetCreate() *resources.CampaignDraft {
	if x, ok := m.GetOperation().(*CampaignDraftOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignDraftOperation) GetUpdate() *resources.CampaignDraft {
	if x, ok := m.GetOperation().(*CampaignDraftOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignDraftOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignDraftOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignDraftOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignDraftOperation_Create)(nil),
		(*CampaignDraftOperation_Update)(nil),
		(*CampaignDraftOperation_Remove)(nil),
	}
}

// Response message for campaign draft mutate.
type MutateCampaignDraftsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignDraftResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateCampaignDraftsResponse) Reset()         { *m = MutateCampaignDraftsResponse{} }
func (m *MutateCampaignDraftsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignDraftsResponse) ProtoMessage()    {}
func (*MutateCampaignDraftsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f742bb1c3712c22e, []int{4}
}

func (m *MutateCampaignDraftsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignDraftsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignDraftsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignDraftsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignDraftsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignDraftsResponse.Merge(m, src)
}
func (m *MutateCampaignDraftsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignDraftsResponse.Size(m)
}
func (m *MutateCampaignDraftsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignDraftsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignDraftsResponse proto.InternalMessageInfo

func (m *MutateCampaignDraftsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignDraftsResponse) GetResults() []*MutateCampaignDraftResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the campaign draft mutate.
type MutateCampaignDraftResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignDraftResult) Reset()         { *m = MutateCampaignDraftResult{} }
func (m *MutateCampaignDraftResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignDraftResult) ProtoMessage()    {}
func (*MutateCampaignDraftResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f742bb1c3712c22e, []int{5}
}

func (m *MutateCampaignDraftResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignDraftResult.Unmarshal(m, b)
}
func (m *MutateCampaignDraftResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignDraftResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignDraftResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignDraftResult.Merge(m, src)
}
func (m *MutateCampaignDraftResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignDraftResult.Size(m)
}
func (m *MutateCampaignDraftResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignDraftResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignDraftResult proto.InternalMessageInfo

func (m *MutateCampaignDraftResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignDraftService.ListCampaignDraftAsyncErrors][google.ads.googleads.v1.services.CampaignDraftService.ListCampaignDraftAsyncErrors].
type ListCampaignDraftAsyncErrorsRequest struct {
	// The name of the campaign draft from which to retrieve the async errors.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Token of the page to retrieve. If not specified, the first
	// page of results will be returned. Use the value obtained from
	// `next_page_token` in the previous response in order to request
	// the next page of results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Number of elements to retrieve in a single page.
	// When a page request is too large, the server may decide to
	// further limit the number of returned resources.
	PageSize             int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCampaignDraftAsyncErrorsRequest) Reset()         { *m = ListCampaignDraftAsyncErrorsRequest{} }
func (m *ListCampaignDraftAsyncErrorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCampaignDraftAsyncErrorsRequest) ProtoMessage()    {}
func (*ListCampaignDraftAsyncErrorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f742bb1c3712c22e, []int{6}
}

func (m *ListCampaignDraftAsyncErrorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest.Unmarshal(m, b)
}
func (m *ListCampaignDraftAsyncErrorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest.Marshal(b, m, deterministic)
}
func (m *ListCampaignDraftAsyncErrorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest.Merge(m, src)
}
func (m *ListCampaignDraftAsyncErrorsRequest) XXX_Size() int {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest.Size(m)
}
func (m *ListCampaignDraftAsyncErrorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest proto.InternalMessageInfo

func (m *ListCampaignDraftAsyncErrorsRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ListCampaignDraftAsyncErrorsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListCampaignDraftAsyncErrorsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Response message for [CampaignDraftService.ListCampaignDraftAsyncErrors][google.ads.googleads.v1.services.CampaignDraftService.ListCampaignDraftAsyncErrors].
type ListCampaignDraftAsyncErrorsResponse struct {
	// Details of the errors when performing the asynchronous operation.
	Errors []*status.Status `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	// Pagination token used to retrieve the next page of results.
	// Pass the content of this string as the `page_token` attribute of
	// the next request. `next_page_token` is not returned for the last
	// page.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCampaignDraftAsyncErrorsResponse) Reset()         { *m = ListCampaignDraftAsyncErrorsResponse{} }
func (m *ListCampaignDraftAsyncErrorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListCampaignDraftAsyncErrorsResponse) ProtoMessage()    {}
func (*ListCampaignDraftAsyncErrorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f742bb1c3712c22e, []int{7}
}

func (m *ListCampaignDraftAsyncErrorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse.Unmarshal(m, b)
}
func (m *ListCampaignDraftAsyncErrorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse.Marshal(b, m, deterministic)
}
func (m *ListCampaignDraftAsyncErrorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse.Merge(m, src)
}
func (m *ListCampaignDraftAsyncErrorsResponse) XXX_Size() int {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse.Size(m)
}
func (m *ListCampaignDraftAsyncErrorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse proto.InternalMessageInfo

func (m *ListCampaignDraftAsyncErrorsResponse) GetErrors() []*status.Status {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *ListCampaignDraftAsyncErrorsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignDraftRequest)(nil), "google.ads.googleads.v1.services.GetCampaignDraftRequest")
	proto.RegisterType((*MutateCampaignDraftsRequest)(nil), "google.ads.googleads.v1.services.MutateCampaignDraftsRequest")
	proto.RegisterType((*PromoteCampaignDraftRequest)(nil), "google.ads.googleads.v1.services.PromoteCampaignDraftRequest")
	proto.RegisterType((*CampaignDraftOperation)(nil), "google.ads.googleads.v1.services.CampaignDraftOperation")
	proto.RegisterType((*MutateCampaignDraftsResponse)(nil), "google.ads.googleads.v1.services.MutateCampaignDraftsResponse")
	proto.RegisterType((*MutateCampaignDraftResult)(nil), "google.ads.googleads.v1.services.MutateCampaignDraftResult")
	proto.RegisterType((*ListCampaignDraftAsyncErrorsRequest)(nil), "google.ads.googleads.v1.services.ListCampaignDraftAsyncErrorsRequest")
	proto.RegisterType((*ListCampaignDraftAsyncErrorsResponse)(nil), "google.ads.googleads.v1.services.ListCampaignDraftAsyncErrorsResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_draft_service.proto", fileDescriptor_f742bb1c3712c22e)
}

var fileDescriptor_f742bb1c3712c22e = []byte{
	// 933 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xd1, 0x6e, 0xdc, 0x44,
	0x14, 0xc5, 0x1b, 0x58, 0xba, 0x77, 0x9b, 0x16, 0x0d, 0x81, 0x6e, 0x37, 0x29, 0xac, 0x9c, 0x00,
	0xd1, 0x3e, 0xd8, 0xdd, 0xad, 0x84, 0x52, 0x47, 0x89, 0xba, 0x21, 0x4d, 0x02, 0xa2, 0x34, 0x72,
	0x20, 0x42, 0x28, 0x92, 0x35, 0x5d, 0x4f, 0x2c, 0x2b, 0xb6, 0xc7, 0xcc, 0xd8, 0x0b, 0x49, 0x55,
	0x09, 0xf1, 0xc0, 0x0f, 0xc0, 0x17, 0xf0, 0xc8, 0x1f, 0x80, 0xc4, 0x0f, 0xf0, 0x8a, 0xf8, 0x00,
	0x24, 0xc4, 0x03, 0x3f, 0xc0, 0x2b, 0x9a, 0xf1, 0xcc, 0x76, 0xbd, 0xdd, 0xcd, 0x26, 0xe1, 0x6d,
	0x7c, 0xe7, 0xcc, 0x99, 0x73, 0xe7, 0xdc, 0xb9, 0x63, 0xd8, 0x08, 0x28, 0x0d, 0x22, 0x62, 0x63,
	0x9f, 0xdb, 0xc5, 0x50, 0x8c, 0x06, 0x1d, 0x9b, 0x13, 0x36, 0x08, 0xfb, 0x84, 0xdb, 0x7d, 0x1c,
	0xa7, 0x38, 0x0c, 0x12, 0xcf, 0x67, 0xf8, 0x38, 0xf3, 0x54, 0xdc, 0x4a, 0x19, 0xcd, 0x28, 0x6a,
	0x15, 0x6b, 0x2c, 0xec, 0x73, 0x6b, 0xb8, 0xdc, 0x1a, 0x74, 0x2c, 0xbd, 0xbc, 0xf9, 0xfe, 0xb4,
	0x0d, 0x18, 0xe1, 0x34, 0x67, 0x2f, 0xee, 0x50, 0x30, 0x37, 0x97, 0xf4, 0xba, 0x34, 0xb4, 0x71,
	0x92, 0xd0, 0x0c, 0x67, 0x21, 0x4d, 0xb8, 0x9a, 0x5d, 0x56, 0xb3, 0x11, 0x4d, 0x02, 0x96, 0x27,
	0x49, 0x98, 0x04, 0x36, 0x4d, 0x09, 0x2b, 0x81, 0x94, 0x38, 0x5b, 0x7e, 0x3d, 0xc9, 0x8f, 0xed,
	0xe3, 0x90, 0x44, 0xbe, 0x17, 0x63, 0x7e, 0xa2, 0x10, 0x6f, 0x8d, 0x23, 0xbe, 0x62, 0x38, 0x4d,
	0x09, 0xd3, 0x0c, 0xb7, 0xd4, 0x3c, 0x4b, 0xfb, 0x36, 0xcf, 0x70, 0x96, 0xab, 0x09, 0x73, 0x13,
	0x6e, 0xed, 0x92, 0xec, 0x03, 0x25, 0x7c, 0x5b, 0xe8, 0x76, 0xc9, 0x97, 0x39, 0xe1, 0x19, 0x5a,
	0x86, 0x79, 0x9d, 0x9a, 0x97, 0xe0, 0x98, 0x34, 0x8c, 0x96, 0xb1, 0x5a, 0x73, 0xaf, 0xeb, 0xe0,
	0x27, 0x38, 0x26, 0xe6, 0xdf, 0x06, 0x2c, 0x3e, 0xca, 0x33, 0x9c, 0x91, 0x12, 0x07, 0xd7, 0x24,
	0x6f, 0x43, 0xbd, 0x9f, 0xf3, 0x8c, 0xc6, 0x84, 0x79, 0xa1, 0xaf, 0x28, 0x40, 0x87, 0x3e, 0xf4,
	0xd1, 0xe7, 0x00, 0xcf, 0xf3, 0x6d, 0x54, 0x5a, 0x73, 0xab, 0xf5, 0xee, 0x9a, 0x35, 0xcb, 0x0d,
	0xab, 0xb4, 0xdb, 0x63, 0x4d, 0xe0, 0x8e, 0x70, 0xa1, 0xf7, 0xe0, 0x66, 0x8a, 0x59, 0x16, 0xe2,
	0xc8, 0x3b, 0xc6, 0x61, 0x94, 0x33, 0xd2, 0x98, 0x6b, 0x19, 0xab, 0xd7, 0xdc, 0x1b, 0x2a, 0xbc,
	0x53, 0x44, 0x45, 0xa2, 0x03, 0x1c, 0x85, 0x3e, 0xce, 0x88, 0x47, 0x93, 0xe8, 0xb4, 0xf1, 0xb2,
	0x84, 0x5d, 0xd7, 0xc1, 0xc7, 0x49, 0x74, 0x6a, 0x6e, 0xc3, 0xe2, 0x3e, 0xa3, 0x31, 0x1d, 0x4b,
	0x54, 0xe7, 0xf9, 0x0e, 0xdc, 0x28, 0xbb, 0xaf, 0x52, 0x9d, 0xef, 0x8f, 0xa2, 0xcd, 0x1f, 0x2a,
	0xf0, 0xe6, 0x64, 0xe9, 0x68, 0x1d, 0xea, 0x79, 0x2a, 0x35, 0x08, 0x5f, 0xa5, 0x86, 0x7a, 0xb7,
	0xa9, 0x4f, 0x42, 0x1b, 0x6b, 0xed, 0x08, 0xeb, 0x1f, 0x61, 0x7e, 0xe2, 0x42, 0x01, 0x17, 0x63,
	0xf4, 0x11, 0x54, 0xfb, 0x8c, 0xe0, 0xac, 0x30, 0xa9, 0xde, 0xbd, 0x3b, 0xf5, 0x04, 0x87, 0xd5,
	0x5a, 0x3e, 0xc2, 0xbd, 0x97, 0x5c, 0xc5, 0x20, 0xb8, 0x0a, 0xe6, 0x46, 0xe5, 0xea, 0x5c, 0x05,
	0x03, 0x6a, 0x40, 0x95, 0x91, 0x98, 0x0e, 0x8a, 0xa3, 0xaf, 0x89, 0x99, 0xe2, 0x7b, 0xab, 0x0e,
	0xb5, 0xa1, 0x57, 0xe6, 0xaf, 0x06, 0x2c, 0x4d, 0xae, 0x22, 0x9e, 0xd2, 0x84, 0x13, 0xb4, 0x03,
	0x6f, 0x8c, 0x79, 0xe9, 0x11, 0xc6, 0x28, 0x93, 0xb4, 0xf5, 0x2e, 0xd2, 0x12, 0x59, 0xda, 0xb7,
	0x0e, 0x64, 0x7d, 0xbb, 0xaf, 0x97, 0x5d, 0x7e, 0x28, 0xe0, 0xe8, 0x33, 0x78, 0x95, 0x11, 0x9e,
	0x47, 0x99, 0x2e, 0xb5, 0xf5, 0xd9, 0xa5, 0x36, 0x41, 0x98, 0x2b, 0x39, 0x5c, 0xcd, 0x65, 0x3e,
	0x80, 0xdb, 0x53, 0x51, 0x17, 0xbb, 0x47, 0xdf, 0x19, 0xb0, 0xfc, 0x71, 0xc8, 0xcb, 0x37, 0xb1,
	0xc7, 0x4f, 0x93, 0xbe, 0x14, 0xce, 0x2f, 0x73, 0x29, 0xd1, 0x1d, 0x80, 0x14, 0x07, 0xc4, 0xcb,
	0xe8, 0x09, 0x49, 0xa4, 0x8b, 0x35, 0xb7, 0x26, 0x22, 0x9f, 0x8a, 0x00, 0x5a, 0x04, 0xf9, 0xe1,
	0xf1, 0xf0, 0xac, 0xf0, 0xe5, 0x15, 0xf7, 0x9a, 0x08, 0x1c, 0x84, 0x67, 0xc4, 0x3c, 0x83, 0x95,
	0xf3, 0x75, 0x28, 0x47, 0xda, 0x50, 0x95, 0x0e, 0xf0, 0x86, 0x21, 0x0f, 0x72, 0x92, 0x05, 0x0a,
	0x81, 0xde, 0x85, 0x9b, 0x09, 0xf9, 0x3a, 0xf3, 0x5e, 0x10, 0x35, 0x2f, 0xc2, 0xfb, 0x5a, 0x58,
	0xf7, 0xcf, 0x2a, 0x2c, 0x94, 0x36, 0x3e, 0x28, 0x3c, 0x40, 0x3f, 0x1b, 0xf0, 0xda, 0x78, 0x9b,
	0x42, 0xf7, 0x67, 0x5b, 0x37, 0xa5, 0xb5, 0x35, 0x2f, 0x5d, 0xd2, 0xe6, 0xda, 0xb7, 0xbf, 0xff,
	0xf5, 0x7d, 0xa5, 0x8b, 0xee, 0x8a, 0x8e, 0xff, 0xb4, 0x64, 0xc1, 0x86, 0xee, 0x66, 0xdc, 0x6e,
	0xdb, 0xa5, 0xeb, 0xce, 0xed, 0xf6, 0x33, 0xf4, 0x87, 0x01, 0x0b, 0x93, 0x6a, 0x1b, 0x6d, 0x5c,
	0xa9, 0xf4, 0x74, 0x25, 0x34, 0x37, 0xaf, 0xba, 0xbc, 0x30, 0xd0, 0xdc, 0x94, 0x19, 0xad, 0x99,
	0xf7, 0x44, 0x46, 0xcf, 0x53, 0x78, 0x3a, 0xd2, 0xae, 0x37, 0xda, 0xcf, 0xc6, 0x12, 0x72, 0x62,
	0x49, 0xe9, 0x18, 0x6d, 0xf4, 0x8b, 0x01, 0x0b, 0x93, 0x3a, 0xe2, 0x45, 0xf2, 0x3a, 0xa7, 0x93,
	0x36, 0xef, 0xe8, 0xe5, 0x23, 0x4f, 0xa2, 0x35, 0x6c, 0x93, 0xe6, 0xb6, 0x94, 0xbd, 0x69, 0xde,
	0x97, 0x46, 0x94, 0x7b, 0xee, 0xf9, 0x4e, 0x38, 0x69, 0xb1, 0xab, 0x10, 0xff, 0xaf, 0x01, 0x4b,
	0xe7, 0x95, 0x39, 0x7a, 0x38, 0x3b, 0x89, 0x0b, 0x5c, 0xd7, 0xe6, 0xce, 0xff, 0xa5, 0x51, 0x66,
	0xed, 0xc9, 0xac, 0xb7, 0xd0, 0x83, 0xcb, 0x96, 0x9f, 0x13, 0x85, 0x7c, 0x94, 0x71, 0xeb, 0x9b,
	0x0a, 0xac, 0xf4, 0x69, 0x3c, 0x53, 0xd7, 0xd6, 0xed, 0x49, 0x37, 0x71, 0x5f, 0x3c, 0x43, 0xfb,
	0xc6, 0x17, 0x7b, 0x6a, 0x79, 0x40, 0x23, 0x9c, 0x04, 0x16, 0x65, 0x81, 0x1d, 0x90, 0x44, 0x3e,
	0x52, 0xfa, 0xe7, 0x28, 0x0d, 0xf9, 0xf4, 0x9f, 0xb1, 0x75, 0x3d, 0xf8, 0xb1, 0x32, 0xb7, 0xdb,
	0xeb, 0xfd, 0x54, 0x69, 0xed, 0x16, 0x84, 0x3d, 0x9f, 0x5b, 0xc5, 0x50, 0x8c, 0x0e, 0x3b, 0x96,
	0xda, 0x98, 0xff, 0xa6, 0x21, 0x47, 0x3d, 0x9f, 0x1f, 0x0d, 0x21, 0x47, 0x87, 0x9d, 0x23, 0x0d,
	0xf9, 0xa7, 0xb2, 0x52, 0xc4, 0x1d, 0xa7, 0xe7, 0x73, 0xc7, 0x19, 0x82, 0x1c, 0xe7, 0xb0, 0xe3,
	0x38, 0x1a, 0xf6, 0xa4, 0x2a, 0x75, 0xde, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x22, 0x0d, 0x9b,
	0x3a, 0x33, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignDraftServiceClient is the client API for CampaignDraftService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignDraftServiceClient interface {
	// Returns the requested campaign draft in full detail.
	GetCampaignDraft(ctx context.Context, in *GetCampaignDraftRequest, opts ...grpc.CallOption) (*resources.CampaignDraft, error)
	// Creates, updates, or removes campaign drafts. Operation statuses are
	// returned.
	MutateCampaignDrafts(ctx context.Context, in *MutateCampaignDraftsRequest, opts ...grpc.CallOption) (*MutateCampaignDraftsResponse, error)
	// Promotes the changes in a draft back to the base campaign.
	//
	// This method returns a Long Running Operation (LRO) indicating if the
	// Promote is done. Use [Operations.GetOperation] to poll the LRO until it
	// is done. Only a done status is returned in the response. See the status
	// in the Campaign Draft resource to determine if the promotion was
	// successful. If the LRO failed, use
	// [CampaignDraftService.ListCampaignDraftAsyncErrors][google.ads.googleads.v1.services.CampaignDraftService.ListCampaignDraftAsyncErrors] to view the list of
	// error reasons.
	PromoteCampaignDraft(ctx context.Context, in *PromoteCampaignDraftRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Returns all errors that occurred during CampaignDraft promote. Throws an
	// error if called before campaign draft is promoted.
	// Supports standard list paging.
	ListCampaignDraftAsyncErrors(ctx context.Context, in *ListCampaignDraftAsyncErrorsRequest, opts ...grpc.CallOption) (*ListCampaignDraftAsyncErrorsResponse, error)
}

type campaignDraftServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignDraftServiceClient(cc *grpc.ClientConn) CampaignDraftServiceClient {
	return &campaignDraftServiceClient{cc}
}

func (c *campaignDraftServiceClient) GetCampaignDraft(ctx context.Context, in *GetCampaignDraftRequest, opts ...grpc.CallOption) (*resources.CampaignDraft, error) {
	out := new(resources.CampaignDraft)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignDraftService/GetCampaignDraft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignDraftServiceClient) MutateCampaignDrafts(ctx context.Context, in *MutateCampaignDraftsRequest, opts ...grpc.CallOption) (*MutateCampaignDraftsResponse, error) {
	out := new(MutateCampaignDraftsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignDraftService/MutateCampaignDrafts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignDraftServiceClient) PromoteCampaignDraft(ctx context.Context, in *PromoteCampaignDraftRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignDraftService/PromoteCampaignDraft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignDraftServiceClient) ListCampaignDraftAsyncErrors(ctx context.Context, in *ListCampaignDraftAsyncErrorsRequest, opts ...grpc.CallOption) (*ListCampaignDraftAsyncErrorsResponse, error) {
	out := new(ListCampaignDraftAsyncErrorsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignDraftService/ListCampaignDraftAsyncErrors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignDraftServiceServer is the server API for CampaignDraftService service.
type CampaignDraftServiceServer interface {
	// Returns the requested campaign draft in full detail.
	GetCampaignDraft(context.Context, *GetCampaignDraftRequest) (*resources.CampaignDraft, error)
	// Creates, updates, or removes campaign drafts. Operation statuses are
	// returned.
	MutateCampaignDrafts(context.Context, *MutateCampaignDraftsRequest) (*MutateCampaignDraftsResponse, error)
	// Promotes the changes in a draft back to the base campaign.
	//
	// This method returns a Long Running Operation (LRO) indicating if the
	// Promote is done. Use [Operations.GetOperation] to poll the LRO until it
	// is done. Only a done status is returned in the response. See the status
	// in the Campaign Draft resource to determine if the promotion was
	// successful. If the LRO failed, use
	// [CampaignDraftService.ListCampaignDraftAsyncErrors][google.ads.googleads.v1.services.CampaignDraftService.ListCampaignDraftAsyncErrors] to view the list of
	// error reasons.
	PromoteCampaignDraft(context.Context, *PromoteCampaignDraftRequest) (*longrunning.Operation, error)
	// Returns all errors that occurred during CampaignDraft promote. Throws an
	// error if called before campaign draft is promoted.
	// Supports standard list paging.
	ListCampaignDraftAsyncErrors(context.Context, *ListCampaignDraftAsyncErrorsRequest) (*ListCampaignDraftAsyncErrorsResponse, error)
}

func RegisterCampaignDraftServiceServer(s *grpc.Server, srv CampaignDraftServiceServer) {
	s.RegisterService(&_CampaignDraftService_serviceDesc, srv)
}

func _CampaignDraftService_GetCampaignDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).GetCampaignDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignDraftService/GetCampaignDraft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).GetCampaignDraft(ctx, req.(*GetCampaignDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignDraftService_MutateCampaignDrafts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignDraftsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).MutateCampaignDrafts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignDraftService/MutateCampaignDrafts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).MutateCampaignDrafts(ctx, req.(*MutateCampaignDraftsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignDraftService_PromoteCampaignDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteCampaignDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).PromoteCampaignDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignDraftService/PromoteCampaignDraft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).PromoteCampaignDraft(ctx, req.(*PromoteCampaignDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignDraftService_ListCampaignDraftAsyncErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCampaignDraftAsyncErrorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).ListCampaignDraftAsyncErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignDraftService/ListCampaignDraftAsyncErrors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).ListCampaignDraftAsyncErrors(ctx, req.(*ListCampaignDraftAsyncErrorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignDraftService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignDraftService",
	HandlerType: (*CampaignDraftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignDraft",
			Handler:    _CampaignDraftService_GetCampaignDraft_Handler,
		},
		{
			MethodName: "MutateCampaignDrafts",
			Handler:    _CampaignDraftService_MutateCampaignDrafts_Handler,
		},
		{
			MethodName: "PromoteCampaignDraft",
			Handler:    _CampaignDraftService_PromoteCampaignDraft_Handler,
		},
		{
			MethodName: "ListCampaignDraftAsyncErrors",
			Handler:    _CampaignDraftService_ListCampaignDraftAsyncErrors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_draft_service.proto",
}
