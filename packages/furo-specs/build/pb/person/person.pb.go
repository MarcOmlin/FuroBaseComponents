// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: person/person.proto

package person

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import furo "../furo"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Person message type
type Person struct {
	// Localized String representation of a person
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// First name of a person
	FirstName string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// Identity of a person
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of a person
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Internal phone number
	PhoneNr string `protobuf:"bytes,5,opt,name=phone_nr,json=phoneNr,proto3" json:"phone_nr,omitempty"`
	// List of main skills of a person
	Skills               []string `protobuf:"bytes,6,rep,name=skills" json:"skills,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_person_528180fa434f0f2a, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Person.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return m.Size()
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetPhoneNr() string {
	if m != nil {
		return m.PhoneNr
	}
	return ""
}

func (m *Person) GetSkills() []string {
	if m != nil {
		return m.Skills
	}
	return nil
}

// PersonCollection with repeated PersonEntity
type PersonCollection struct {
	// Contains a person.PersonEntity repeated
	Entities []*PersonEntity `protobuf:"bytes,4,rep,name=entities" json:"entities,omitempty"`
	// Hateoas links
	Links []*furo.Link `protobuf:"bytes,3,rep,name=links" json:"links,omitempty"`
	// Meta for the response
	Meta                 *furo.Meta `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PersonCollection) Reset()         { *m = PersonCollection{} }
func (m *PersonCollection) String() string { return proto.CompactTextString(m) }
func (*PersonCollection) ProtoMessage()    {}
func (*PersonCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_person_528180fa434f0f2a, []int{1}
}
func (m *PersonCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PersonCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PersonCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PersonCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonCollection.Merge(dst, src)
}
func (m *PersonCollection) XXX_Size() int {
	return m.Size()
}
func (m *PersonCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonCollection.DiscardUnknown(m)
}

var xxx_messageInfo_PersonCollection proto.InternalMessageInfo

func (m *PersonCollection) GetEntities() []*PersonEntity {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *PersonCollection) GetLinks() []*furo.Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PersonCollection) GetMeta() *furo.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

// PersonEntity with Person
type PersonEntity struct {
	// contains a person.Person
	Data *Person `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	// Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links" json:"links,omitempty"`
	// Meta for the response
	Meta                 *furo.Meta `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PersonEntity) Reset()         { *m = PersonEntity{} }
func (m *PersonEntity) String() string { return proto.CompactTextString(m) }
func (*PersonEntity) ProtoMessage()    {}
func (*PersonEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_person_528180fa434f0f2a, []int{2}
}
func (m *PersonEntity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PersonEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PersonEntity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PersonEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonEntity.Merge(dst, src)
}
func (m *PersonEntity) XXX_Size() int {
	return m.Size()
}
func (m *PersonEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonEntity.DiscardUnknown(m)
}

var xxx_messageInfo_PersonEntity proto.InternalMessageInfo

func (m *PersonEntity) GetData() *Person {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PersonEntity) GetLinks() []*furo.Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PersonEntity) GetMeta() *furo.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "person.Person")
	proto.RegisterType((*PersonCollection)(nil), "person.PersonCollection")
	proto.RegisterType((*PersonEntity)(nil), "person.PersonEntity")
}
func (m *Person) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Person) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPerson(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.DisplayName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPerson(dAtA, i, uint64(len(m.DisplayName)))
		i += copy(dAtA[i:], m.DisplayName)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPerson(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.FirstName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPerson(dAtA, i, uint64(len(m.FirstName)))
		i += copy(dAtA[i:], m.FirstName)
	}
	if len(m.PhoneNr) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintPerson(dAtA, i, uint64(len(m.PhoneNr)))
		i += copy(dAtA[i:], m.PhoneNr)
	}
	if len(m.Skills) > 0 {
		for _, s := range m.Skills {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PersonCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PersonCollection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Meta != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPerson(dAtA, i, uint64(m.Meta.Size()))
		n1, err := m.Meta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Links) > 0 {
		for _, msg := range m.Links {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintPerson(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Entities) > 0 {
		for _, msg := range m.Entities {
			dAtA[i] = 0x22
			i++
			i = encodeVarintPerson(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PersonEntity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PersonEntity) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPerson(dAtA, i, uint64(m.Data.Size()))
		n2, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Links) > 0 {
		for _, msg := range m.Links {
			dAtA[i] = 0x12
			i++
			i = encodeVarintPerson(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Meta != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPerson(dAtA, i, uint64(m.Meta.Size()))
		n3, err := m.Meta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPerson(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Person) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPerson(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovPerson(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPerson(uint64(l))
	}
	l = len(m.FirstName)
	if l > 0 {
		n += 1 + l + sovPerson(uint64(l))
	}
	l = len(m.PhoneNr)
	if l > 0 {
		n += 1 + l + sovPerson(uint64(l))
	}
	if len(m.Skills) > 0 {
		for _, s := range m.Skills {
			l = len(s)
			n += 1 + l + sovPerson(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PersonCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovPerson(uint64(l))
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovPerson(uint64(l))
		}
	}
	if len(m.Entities) > 0 {
		for _, e := range m.Entities {
			l = e.Size()
			n += 1 + l + sovPerson(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PersonEntity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovPerson(uint64(l))
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovPerson(uint64(l))
		}
	}
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovPerson(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPerson(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPerson(x uint64) (n int) {
	return sovPerson(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Person) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPerson
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Person: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Person: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhoneNr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PhoneNr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skills", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Skills = append(m.Skills, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPerson(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPerson
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PersonCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPerson
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PersonCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PersonCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &furo.Meta{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &furo.Link{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entities = append(m.Entities, &PersonEntity{})
			if err := m.Entities[len(m.Entities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPerson(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPerson
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PersonEntity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPerson
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PersonEntity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PersonEntity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &Person{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &furo.Link{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPerson
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &furo.Meta{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPerson(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPerson
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPerson(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPerson
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
					return 0, ErrIntOverflowPerson
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPerson
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPerson
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPerson
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPerson(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPerson = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPerson   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("person/person.proto", fileDescriptor_person_528180fa434f0f2a) }

var fileDescriptor_person_528180fa434f0f2a = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x45, 0xe5, 0x24, 0x0d, 0xed, 0xb4, 0x2a, 0xc8, 0x20, 0x14, 0x90, 0x88, 0x42, 0x56, 0x5d,
	0xa5, 0xa8, 0xdc, 0x00, 0xc4, 0x0e, 0x2a, 0xe4, 0x0b, 0x54, 0x86, 0xb8, 0xc2, 0xaa, 0x6b, 0x47,
	0xb6, 0x59, 0xf4, 0x02, 0xdc, 0x83, 0xdb, 0xb0, 0xe4, 0x08, 0xa8, 0x27, 0x41, 0x19, 0x07, 0x0a,
	0x12, 0xac, 0x92, 0xf9, 0xef, 0x8f, 0xf2, 0xff, 0x04, 0x0e, 0x1b, 0x61, 0x9d, 0xd1, 0xd3, 0xf0,
	0xa8, 0x1a, 0x6b, 0xbc, 0xa1, 0x69, 0x98, 0x4e, 0xf7, 0x97, 0xcf, 0xd6, 0x4c, 0xd7, 0xc2, 0xf3,
	0x00, 0x3a, 0x41, 0x49, 0xbd, 0x0a, 0x42, 0xf9, 0x4a, 0x20, 0xbd, 0x47, 0x33, 0x1d, 0x43, 0x24,
	0xeb, 0x8c, 0x14, 0x64, 0x32, 0x60, 0x91, 0xac, 0xe9, 0x39, 0x8c, 0x6a, 0xe9, 0x1a, 0xc5, 0x37,
	0x0b, 0xcd, 0xd7, 0x22, 0x8b, 0x90, 0x0c, 0x3b, 0x6d, 0xce, 0xd7, 0x82, 0x52, 0x48, 0x10, 0xc5,
	0x88, 0xf0, 0x9d, 0x9e, 0x01, 0x2c, 0xa5, 0x75, 0x3e, 0x2c, 0x25, 0x48, 0x06, 0xa8, 0xe0, 0xca,
	0x09, 0xf4, 0x9b, 0x27, 0xa3, 0xc5, 0x42, 0xdb, 0xac, 0x87, 0x70, 0x0f, 0xe7, 0xb9, 0xa5, 0xc7,
	0x90, 0xba, 0x95, 0x54, 0xca, 0x65, 0x69, 0x11, 0x4f, 0x06, 0xac, 0x9b, 0xca, 0x17, 0x02, 0x07,
	0x21, 0xe3, 0xb5, 0x51, 0x4a, 0x3c, 0x7a, 0x69, 0x34, 0xcd, 0x21, 0x69, 0x7b, 0x61, 0xaa, 0xe1,
	0x0c, 0xaa, 0xb6, 0x58, 0x75, 0x27, 0x3c, 0x67, 0xa8, 0xd3, 0x02, 0x7a, 0x6d, 0x4d, 0x97, 0xc5,
	0x45, 0xbc, 0x33, 0xdc, 0x4a, 0xbd, 0x62, 0x01, 0xd0, 0x0b, 0xe8, 0x0b, 0xed, 0xa5, 0x97, 0xc2,
	0x65, 0x09, 0x9a, 0x8e, 0xaa, 0xee, 0x8a, 0xe1, 0x6b, 0x37, 0x2d, 0xdd, 0xb0, 0x6f, 0x57, 0xe9,
	0x61, 0xf4, 0x93, 0xd0, 0x12, 0x92, 0x9a, 0x7b, 0x8e, 0x37, 0x1b, 0xce, 0xc6, 0xbf, 0xb7, 0x19,
	0xb2, 0x5d, 0x8e, 0xe8, 0xbf, 0x1c, 0x5f, 0x4d, 0xe2, 0xbf, 0x9b, 0x5c, 0x8d, 0xde, 0xb6, 0x39,
	0x79, 0xdf, 0xe6, 0xe4, 0x63, 0x9b, 0x93, 0x87, 0x14, 0xff, 0xdb, 0xe5, 0x67, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x48, 0x4a, 0xd9, 0xe4, 0xf8, 0x01, 0x00, 0x00,
}