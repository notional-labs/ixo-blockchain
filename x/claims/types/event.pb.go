// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/claims/v1beta1/event.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CollectionCreatedEvent is an event triggered on a Collection creation
type CollectionCreatedEvent struct {
	Collection *Collection `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
}

func (m *CollectionCreatedEvent) Reset()         { *m = CollectionCreatedEvent{} }
func (m *CollectionCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*CollectionCreatedEvent) ProtoMessage()    {}
func (*CollectionCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b125bce36ddb45, []int{0}
}
func (m *CollectionCreatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectionCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollectionCreatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollectionCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionCreatedEvent.Merge(m, src)
}
func (m *CollectionCreatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *CollectionCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionCreatedEvent proto.InternalMessageInfo

func (m *CollectionCreatedEvent) GetCollection() *Collection {
	if m != nil {
		return m.Collection
	}
	return nil
}

// CollectionUpdatedEvent is an event triggered on a Collection update
type CollectionUpdatedEvent struct {
	Collection *Collection `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
}

func (m *CollectionUpdatedEvent) Reset()         { *m = CollectionUpdatedEvent{} }
func (m *CollectionUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*CollectionUpdatedEvent) ProtoMessage()    {}
func (*CollectionUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b125bce36ddb45, []int{1}
}
func (m *CollectionUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectionUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollectionUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollectionUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionUpdatedEvent.Merge(m, src)
}
func (m *CollectionUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *CollectionUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionUpdatedEvent proto.InternalMessageInfo

func (m *CollectionUpdatedEvent) GetCollection() *Collection {
	if m != nil {
		return m.Collection
	}
	return nil
}

// CollectionCreatedEvent is an event triggered on a Claim submission
type ClaimSubmittedEvent struct {
	Claim *Claim `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (m *ClaimSubmittedEvent) Reset()         { *m = ClaimSubmittedEvent{} }
func (m *ClaimSubmittedEvent) String() string { return proto.CompactTextString(m) }
func (*ClaimSubmittedEvent) ProtoMessage()    {}
func (*ClaimSubmittedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b125bce36ddb45, []int{2}
}
func (m *ClaimSubmittedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimSubmittedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimSubmittedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimSubmittedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimSubmittedEvent.Merge(m, src)
}
func (m *ClaimSubmittedEvent) XXX_Size() int {
	return m.Size()
}
func (m *ClaimSubmittedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimSubmittedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimSubmittedEvent proto.InternalMessageInfo

func (m *ClaimSubmittedEvent) GetClaim() *Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

// ClaimUpdatedEvent is an event triggered on a Claim update
type ClaimUpdatedEvent struct {
	Claim *Claim `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (m *ClaimUpdatedEvent) Reset()         { *m = ClaimUpdatedEvent{} }
func (m *ClaimUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*ClaimUpdatedEvent) ProtoMessage()    {}
func (*ClaimUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b125bce36ddb45, []int{3}
}
func (m *ClaimUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimUpdatedEvent.Merge(m, src)
}
func (m *ClaimUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *ClaimUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimUpdatedEvent proto.InternalMessageInfo

func (m *ClaimUpdatedEvent) GetClaim() *Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

// ClaimEvaluatedEvent is an event triggered on a Claim evaluation
type ClaimEvaluatedEvent struct {
	Evaluation *Evaluation `protobuf:"bytes,1,opt,name=evaluation,proto3" json:"evaluation,omitempty"`
}

func (m *ClaimEvaluatedEvent) Reset()         { *m = ClaimEvaluatedEvent{} }
func (m *ClaimEvaluatedEvent) String() string { return proto.CompactTextString(m) }
func (*ClaimEvaluatedEvent) ProtoMessage()    {}
func (*ClaimEvaluatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b125bce36ddb45, []int{4}
}
func (m *ClaimEvaluatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimEvaluatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimEvaluatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimEvaluatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimEvaluatedEvent.Merge(m, src)
}
func (m *ClaimEvaluatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *ClaimEvaluatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimEvaluatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimEvaluatedEvent proto.InternalMessageInfo

func (m *ClaimEvaluatedEvent) GetEvaluation() *Evaluation {
	if m != nil {
		return m.Evaluation
	}
	return nil
}

// ClaimDisputedEvent is an event triggered on a Claim dispute
type ClaimDisputedEvent struct {
	Dispute *Dispute `protobuf:"bytes,1,opt,name=dispute,proto3" json:"dispute,omitempty"`
}

func (m *ClaimDisputedEvent) Reset()         { *m = ClaimDisputedEvent{} }
func (m *ClaimDisputedEvent) String() string { return proto.CompactTextString(m) }
func (*ClaimDisputedEvent) ProtoMessage()    {}
func (*ClaimDisputedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b125bce36ddb45, []int{5}
}
func (m *ClaimDisputedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimDisputedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimDisputedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimDisputedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimDisputedEvent.Merge(m, src)
}
func (m *ClaimDisputedEvent) XXX_Size() int {
	return m.Size()
}
func (m *ClaimDisputedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimDisputedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimDisputedEvent proto.InternalMessageInfo

func (m *ClaimDisputedEvent) GetDispute() *Dispute {
	if m != nil {
		return m.Dispute
	}
	return nil
}

// ClaimDisputedEvent is an event triggered on a Claim dispute
type PaymentWithdrawnEvent struct {
	Withdraw *WithdrawPaymentConstraints `protobuf:"bytes,1,opt,name=withdraw,proto3" json:"withdraw,omitempty"`
}

func (m *PaymentWithdrawnEvent) Reset()         { *m = PaymentWithdrawnEvent{} }
func (m *PaymentWithdrawnEvent) String() string { return proto.CompactTextString(m) }
func (*PaymentWithdrawnEvent) ProtoMessage()    {}
func (*PaymentWithdrawnEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b125bce36ddb45, []int{6}
}
func (m *PaymentWithdrawnEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PaymentWithdrawnEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PaymentWithdrawnEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PaymentWithdrawnEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentWithdrawnEvent.Merge(m, src)
}
func (m *PaymentWithdrawnEvent) XXX_Size() int {
	return m.Size()
}
func (m *PaymentWithdrawnEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentWithdrawnEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentWithdrawnEvent proto.InternalMessageInfo

func (m *PaymentWithdrawnEvent) GetWithdraw() *WithdrawPaymentConstraints {
	if m != nil {
		return m.Withdraw
	}
	return nil
}

// ClaimDisputedEvent is an event triggered on a Claim dispute
type PaymentWithdrawCreatedEvent struct {
	Withdraw *WithdrawPaymentConstraints `protobuf:"bytes,1,opt,name=withdraw,proto3" json:"withdraw,omitempty"`
}

func (m *PaymentWithdrawCreatedEvent) Reset()         { *m = PaymentWithdrawCreatedEvent{} }
func (m *PaymentWithdrawCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*PaymentWithdrawCreatedEvent) ProtoMessage()    {}
func (*PaymentWithdrawCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b125bce36ddb45, []int{7}
}
func (m *PaymentWithdrawCreatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PaymentWithdrawCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PaymentWithdrawCreatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PaymentWithdrawCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentWithdrawCreatedEvent.Merge(m, src)
}
func (m *PaymentWithdrawCreatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *PaymentWithdrawCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentWithdrawCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentWithdrawCreatedEvent proto.InternalMessageInfo

func (m *PaymentWithdrawCreatedEvent) GetWithdraw() *WithdrawPaymentConstraints {
	if m != nil {
		return m.Withdraw
	}
	return nil
}

func init() {
	proto.RegisterType((*CollectionCreatedEvent)(nil), "ixo.claims.v1beta1.CollectionCreatedEvent")
	proto.RegisterType((*CollectionUpdatedEvent)(nil), "ixo.claims.v1beta1.CollectionUpdatedEvent")
	proto.RegisterType((*ClaimSubmittedEvent)(nil), "ixo.claims.v1beta1.ClaimSubmittedEvent")
	proto.RegisterType((*ClaimUpdatedEvent)(nil), "ixo.claims.v1beta1.ClaimUpdatedEvent")
	proto.RegisterType((*ClaimEvaluatedEvent)(nil), "ixo.claims.v1beta1.ClaimEvaluatedEvent")
	proto.RegisterType((*ClaimDisputedEvent)(nil), "ixo.claims.v1beta1.ClaimDisputedEvent")
	proto.RegisterType((*PaymentWithdrawnEvent)(nil), "ixo.claims.v1beta1.PaymentWithdrawnEvent")
	proto.RegisterType((*PaymentWithdrawCreatedEvent)(nil), "ixo.claims.v1beta1.PaymentWithdrawCreatedEvent")
}

func init() { proto.RegisterFile("ixo/claims/v1beta1/event.proto", fileDescriptor_26b125bce36ddb45) }

var fileDescriptor_26b125bce36ddb45 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0x9b, 0xe1, 0xff, 0x21, 0x33, 0x11, 0x3e, 0x04, 0xad, 0x64, 0x50, 0x26, 0x16, 0x62,
	0x15, 0x04, 0x23, 0x03, 0x69, 0x19, 0x60, 0xa9, 0x82, 0x2a, 0x10, 0x9b, 0xe3, 0x98, 0xc6, 0x22,
	0xb1, 0xa3, 0xc4, 0x69, 0x53, 0x9e, 0x82, 0xc7, 0x62, 0xec, 0xc8, 0x88, 0xda, 0x17, 0x41, 0x89,
	0xdd, 0xd0, 0x40, 0x24, 0x84, 0xd4, 0x2d, 0xf7, 0x9e, 0x73, 0x7e, 0xf7, 0x3a, 0xb2, 0x01, 0x64,
	0xb9, 0x40, 0x24, 0xc4, 0x2c, 0x4a, 0xd1, 0xb8, 0xeb, 0x51, 0x89, 0xbb, 0x88, 0x8e, 0x29, 0x97,
	0x76, 0x9c, 0x08, 0x29, 0x4c, 0x93, 0xe5, 0xc2, 0x56, 0xba, 0xad, 0xf5, 0xf6, 0xf6, 0x48, 0x8c,
	0x44, 0x29, 0xa3, 0xe2, 0x4b, 0x39, 0xdb, 0x07, 0x0d, 0x24, 0x1d, 0x54, 0x86, 0xa6, 0x51, 0x38,
	0x93, 0xc1, 0xb3, 0xd2, 0xad, 0x7b, 0xb0, 0xeb, 0x88, 0x30, 0xa4, 0x44, 0x32, 0xc1, 0x9d, 0x84,
	0x62, 0x49, 0xfd, 0x7e, 0xb1, 0x8a, 0x79, 0x01, 0x00, 0xa9, 0x94, 0x3d, 0xe3, 0xd0, 0x38, 0xda,
	0x38, 0x81, 0xf6, 0xf7, 0xcd, 0xec, 0xcf, 0xbc, 0xbb, 0x92, 0xa8, 0x93, 0x87, 0xb1, 0xbf, 0x3e,
	0xf2, 0x15, 0xd8, 0x72, 0x0a, 0xe3, 0x6d, 0xe6, 0x45, 0x4c, 0x56, 0x58, 0x04, 0xfe, 0x94, 0x79,
	0x4d, 0xdc, 0x6f, 0x24, 0x16, 0xa5, 0xab, 0x7c, 0x56, 0x0f, 0x6c, 0x96, 0x75, 0x6d, 0xb9, 0x5f,
	0x53, 0x86, 0x7a, 0x9b, 0xfe, 0x18, 0x87, 0x59, 0xed, 0x90, 0x54, 0x75, 0x7e, 0x38, 0x64, 0xbf,
	0x72, 0xb9, 0x2b, 0x09, 0xeb, 0x06, 0x98, 0x25, 0xb6, 0xc7, 0xd2, 0x38, 0xab, 0xa8, 0x67, 0xe0,
	0x9f, 0xaf, 0x1a, 0x1a, 0xd9, 0x69, 0x42, 0xea, 0x8c, 0xbb, 0xf4, 0x5a, 0x04, 0xec, 0x0c, 0xf0,
	0x34, 0xa2, 0x5c, 0xde, 0x31, 0x19, 0xf8, 0x09, 0x9e, 0x70, 0xc5, 0xbb, 0x06, 0xff, 0x27, 0xba,
	0xa3, 0x81, 0x76, 0x13, 0x70, 0x99, 0xd2, 0x10, 0x47, 0xf0, 0x54, 0x26, 0x98, 0x71, 0x99, 0xba,
	0x55, 0xde, 0x62, 0xa0, 0xf3, 0x65, 0x48, 0xed, 0x3e, 0xad, 0x71, 0xd4, 0xe5, 0xe0, 0x75, 0x0e,
	0x8d, 0xd9, 0x1c, 0x1a, 0xef, 0x73, 0x68, 0xbc, 0x2c, 0x60, 0x6b, 0xb6, 0x80, 0xad, 0xb7, 0x05,
	0x6c, 0x3d, 0x9c, 0x8f, 0x98, 0x0c, 0x32, 0xcf, 0x26, 0x22, 0x42, 0x2c, 0x17, 0x8f, 0x22, 0xe3,
	0x7e, 0xf9, 0x43, 0x8b, 0xea, 0xd8, 0x0b, 0x05, 0x79, 0x22, 0x01, 0x66, 0x1c, 0xe5, 0xcb, 0x57,
	0x21, 0xa7, 0x31, 0x4d, 0xbd, 0xbf, 0xe5, 0x73, 0x38, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7d,
	0x27, 0x8b, 0xa3, 0x9b, 0x03, 0x00, 0x00,
}

func (m *CollectionCreatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectionCreatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectionCreatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Collection != nil {
		{
			size, err := m.Collection.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CollectionUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectionUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectionUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Collection != nil {
		{
			size, err := m.Collection.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimSubmittedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimSubmittedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimSubmittedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimEvaluatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimEvaluatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimEvaluatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Evaluation != nil {
		{
			size, err := m.Evaluation.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimDisputedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimDisputedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimDisputedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Dispute != nil {
		{
			size, err := m.Dispute.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PaymentWithdrawnEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PaymentWithdrawnEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PaymentWithdrawnEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Withdraw != nil {
		{
			size, err := m.Withdraw.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PaymentWithdrawCreatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PaymentWithdrawCreatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PaymentWithdrawCreatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Withdraw != nil {
		{
			size, err := m.Withdraw.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CollectionCreatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Collection != nil {
		l = m.Collection.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *CollectionUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Collection != nil {
		l = m.Collection.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *ClaimSubmittedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *ClaimUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *ClaimEvaluatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Evaluation != nil {
		l = m.Evaluation.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *ClaimDisputedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dispute != nil {
		l = m.Dispute.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *PaymentWithdrawnEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Withdraw != nil {
		l = m.Withdraw.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *PaymentWithdrawCreatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Withdraw != nil {
		l = m.Withdraw.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CollectionCreatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CollectionCreatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectionCreatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Collection == nil {
				m.Collection = &Collection{}
			}
			if err := m.Collection.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CollectionUpdatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CollectionUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectionUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Collection == nil {
				m.Collection = &Collection{}
			}
			if err := m.Collection.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimSubmittedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimSubmittedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimSubmittedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimUpdatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimEvaluatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimEvaluatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimEvaluatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evaluation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Evaluation == nil {
				m.Evaluation = &Evaluation{}
			}
			if err := m.Evaluation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimDisputedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimDisputedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimDisputedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dispute", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dispute == nil {
				m.Dispute = &Dispute{}
			}
			if err := m.Dispute.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PaymentWithdrawnEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PaymentWithdrawnEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PaymentWithdrawnEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdraw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Withdraw == nil {
				m.Withdraw = &WithdrawPaymentConstraints{}
			}
			if err := m.Withdraw.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PaymentWithdrawCreatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PaymentWithdrawCreatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PaymentWithdrawCreatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdraw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Withdraw == nil {
				m.Withdraw = &WithdrawPaymentConstraints{}
			}
			if err := m.Withdraw.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
