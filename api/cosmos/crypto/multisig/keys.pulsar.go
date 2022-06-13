// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package multisig

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_LegacyAminoPubKey_2_list)(nil)

type _LegacyAminoPubKey_2_list struct {
	list *[]*anypb.Any
}

func (x *_LegacyAminoPubKey_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_LegacyAminoPubKey_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_LegacyAminoPubKey_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	(*x.list)[i] = concreteValue
}

func (x *_LegacyAminoPubKey_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	*x.list = append(*x.list, concreteValue)
}

func (x *_LegacyAminoPubKey_2_list) AppendMutable() protoreflect.Value {
	v := new(anypb.Any)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_LegacyAminoPubKey_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_LegacyAminoPubKey_2_list) NewElement() protoreflect.Value {
	v := new(anypb.Any)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_LegacyAminoPubKey_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_LegacyAminoPubKey             protoreflect.MessageDescriptor
	fd_LegacyAminoPubKey_threshold   protoreflect.FieldDescriptor
	fd_LegacyAminoPubKey_public_keys protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_crypto_multisig_keys_proto_init()
	md_LegacyAminoPubKey = File_cosmos_crypto_multisig_keys_proto.Messages().ByName("LegacyAminoPubKey")
	fd_LegacyAminoPubKey_threshold = md_LegacyAminoPubKey.Fields().ByName("threshold")
	fd_LegacyAminoPubKey_public_keys = md_LegacyAminoPubKey.Fields().ByName("public_keys")
}

var _ protoreflect.Message = (*fastReflection_LegacyAminoPubKey)(nil)

type fastReflection_LegacyAminoPubKey LegacyAminoPubKey

func (x *LegacyAminoPubKey) ProtoReflect() protoreflect.Message {
	return (*fastReflection_LegacyAminoPubKey)(x)
}

func (x *LegacyAminoPubKey) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_multisig_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_LegacyAminoPubKey_messageType fastReflection_LegacyAminoPubKey_messageType
var _ protoreflect.MessageType = fastReflection_LegacyAminoPubKey_messageType{}

type fastReflection_LegacyAminoPubKey_messageType struct{}

func (x fastReflection_LegacyAminoPubKey_messageType) Zero() protoreflect.Message {
	return (*fastReflection_LegacyAminoPubKey)(nil)
}
func (x fastReflection_LegacyAminoPubKey_messageType) New() protoreflect.Message {
	return new(fastReflection_LegacyAminoPubKey)
}
func (x fastReflection_LegacyAminoPubKey_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_LegacyAminoPubKey
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_LegacyAminoPubKey) Descriptor() protoreflect.MessageDescriptor {
	return md_LegacyAminoPubKey
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_LegacyAminoPubKey) Type() protoreflect.MessageType {
	return _fastReflection_LegacyAminoPubKey_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_LegacyAminoPubKey) New() protoreflect.Message {
	return new(fastReflection_LegacyAminoPubKey)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_LegacyAminoPubKey) Interface() protoreflect.ProtoMessage {
	return (*LegacyAminoPubKey)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_LegacyAminoPubKey) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Threshold != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Threshold)
		if !f(fd_LegacyAminoPubKey_threshold, value) {
			return
		}
	}
	if len(x.PublicKeys) != 0 {
		value := protoreflect.ValueOfList(&_LegacyAminoPubKey_2_list{list: &x.PublicKeys})
		if !f(fd_LegacyAminoPubKey_public_keys, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_LegacyAminoPubKey) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.LegacyAminoPubKey.threshold":
		return x.Threshold != uint32(0)
	case "cosmos.crypto.multisig.LegacyAminoPubKey.public_keys":
		return len(x.PublicKeys) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.LegacyAminoPubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.LegacyAminoPubKey does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LegacyAminoPubKey) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.LegacyAminoPubKey.threshold":
		x.Threshold = uint32(0)
	case "cosmos.crypto.multisig.LegacyAminoPubKey.public_keys":
		x.PublicKeys = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.LegacyAminoPubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.LegacyAminoPubKey does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_LegacyAminoPubKey) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.crypto.multisig.LegacyAminoPubKey.threshold":
		value := x.Threshold
		return protoreflect.ValueOfUint32(value)
	case "cosmos.crypto.multisig.LegacyAminoPubKey.public_keys":
		if len(x.PublicKeys) == 0 {
			return protoreflect.ValueOfList(&_LegacyAminoPubKey_2_list{})
		}
		listValue := &_LegacyAminoPubKey_2_list{list: &x.PublicKeys}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.LegacyAminoPubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.LegacyAminoPubKey does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LegacyAminoPubKey) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.LegacyAminoPubKey.threshold":
		x.Threshold = uint32(value.Uint())
	case "cosmos.crypto.multisig.LegacyAminoPubKey.public_keys":
		lv := value.List()
		clv := lv.(*_LegacyAminoPubKey_2_list)
		x.PublicKeys = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.LegacyAminoPubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.LegacyAminoPubKey does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LegacyAminoPubKey) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.LegacyAminoPubKey.public_keys":
		if x.PublicKeys == nil {
			x.PublicKeys = []*anypb.Any{}
		}
		value := &_LegacyAminoPubKey_2_list{list: &x.PublicKeys}
		return protoreflect.ValueOfList(value)
	case "cosmos.crypto.multisig.LegacyAminoPubKey.threshold":
		panic(fmt.Errorf("field threshold of message cosmos.crypto.multisig.LegacyAminoPubKey is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.LegacyAminoPubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.LegacyAminoPubKey does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_LegacyAminoPubKey) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.LegacyAminoPubKey.threshold":
		return protoreflect.ValueOfUint32(uint32(0))
	case "cosmos.crypto.multisig.LegacyAminoPubKey.public_keys":
		list := []*anypb.Any{}
		return protoreflect.ValueOfList(&_LegacyAminoPubKey_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.LegacyAminoPubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.LegacyAminoPubKey does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_LegacyAminoPubKey) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.multisig.LegacyAminoPubKey", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_LegacyAminoPubKey) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LegacyAminoPubKey) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_LegacyAminoPubKey) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_LegacyAminoPubKey) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*LegacyAminoPubKey)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Threshold != 0 {
			n += 1 + runtime.Sov(uint64(x.Threshold))
		}
		if len(x.PublicKeys) > 0 {
			for _, e := range x.PublicKeys {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*LegacyAminoPubKey)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.PublicKeys) > 0 {
			for iNdEx := len(x.PublicKeys) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.PublicKeys[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.Threshold != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Threshold))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*LegacyAminoPubKey)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LegacyAminoPubKey: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LegacyAminoPubKey: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
				}
				x.Threshold = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Threshold |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PublicKeys", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.PublicKeys = append(x.PublicKeys, &anypb.Any{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PublicKeys[len(x.PublicKeys)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/crypto/multisig/keys.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// LegacyAminoPubKey specifies a public key type
// which nests multiple public keys and a threshold,
// it uses legacy amino address rules.
type LegacyAminoPubKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Threshold  uint32       `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	PublicKeys []*anypb.Any `protobuf:"bytes,2,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
}

func (x *LegacyAminoPubKey) Reset() {
	*x = LegacyAminoPubKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_multisig_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LegacyAminoPubKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LegacyAminoPubKey) ProtoMessage() {}

// Deprecated: Use LegacyAminoPubKey.ProtoReflect.Descriptor instead.
func (*LegacyAminoPubKey) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_multisig_keys_proto_rawDescGZIP(), []int{0}
}

func (x *LegacyAminoPubKey) GetThreshold() uint32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *LegacyAminoPubKey) GetPublicKeys() []*anypb.Any {
	if x != nil {
		return x.PublicKeys
	}
	return nil
}

var File_cosmos_crypto_multisig_keys_proto protoreflect.FileDescriptor

var file_cosmos_crypto_multisig_keys_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x1a, 0x14, 0x67, 0x6f, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x11,
	0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x41, 0x6d, 0x69, 0x6e, 0x6f, 0x50, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12,
	0x42, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x0b, 0xe2, 0xde, 0x1f, 0x07,
	0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x73, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x42, 0xca, 0x01, 0x0a, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x42, 0x09, 0x4b, 0x65, 0x79, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0xa2, 0x02,
	0x03, 0x43, 0x43, 0x4d, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0xca, 0x02, 0x16,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5c, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0xe2, 0x02, 0x22, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x3a, 0x3a, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_crypto_multisig_keys_proto_rawDescOnce sync.Once
	file_cosmos_crypto_multisig_keys_proto_rawDescData = file_cosmos_crypto_multisig_keys_proto_rawDesc
)

func file_cosmos_crypto_multisig_keys_proto_rawDescGZIP() []byte {
	file_cosmos_crypto_multisig_keys_proto_rawDescOnce.Do(func() {
		file_cosmos_crypto_multisig_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_crypto_multisig_keys_proto_rawDescData)
	})
	return file_cosmos_crypto_multisig_keys_proto_rawDescData
}

var file_cosmos_crypto_multisig_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cosmos_crypto_multisig_keys_proto_goTypes = []interface{}{
	(*LegacyAminoPubKey)(nil), // 0: cosmos.crypto.multisig.LegacyAminoPubKey
	(*anypb.Any)(nil),         // 1: google.protobuf.Any
}
var file_cosmos_crypto_multisig_keys_proto_depIdxs = []int32{
	1, // 0: cosmos.crypto.multisig.LegacyAminoPubKey.public_keys:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_crypto_multisig_keys_proto_init() }
func file_cosmos_crypto_multisig_keys_proto_init() {
	if File_cosmos_crypto_multisig_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_crypto_multisig_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LegacyAminoPubKey); i {
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
			RawDescriptor: file_cosmos_crypto_multisig_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_crypto_multisig_keys_proto_goTypes,
		DependencyIndexes: file_cosmos_crypto_multisig_keys_proto_depIdxs,
		MessageInfos:      file_cosmos_crypto_multisig_keys_proto_msgTypes,
	}.Build()
	File_cosmos_crypto_multisig_keys_proto = out.File
	file_cosmos_crypto_multisig_keys_proto_rawDesc = nil
	file_cosmos_crypto_multisig_keys_proto_goTypes = nil
	file_cosmos_crypto_multisig_keys_proto_depIdxs = nil
}
