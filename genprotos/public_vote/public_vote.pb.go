// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: public_vote.proto

package public_vote

import (
	election "genprotos/election"
	public "genprotos/public"
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

type PublicVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Election *election.Election `protobuf:"bytes,2,opt,name=election,proto3" json:"election,omitempty"`
	Public   *public.Public     `protobuf:"bytes,3,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *PublicVote) Reset() {
	*x = PublicVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicVote) ProtoMessage() {}

func (x *PublicVote) ProtoReflect() protoreflect.Message {
	mi := &file_public_vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicVote.ProtoReflect.Descriptor instead.
func (*PublicVote) Descriptor() ([]byte, []int) {
	return file_public_vote_proto_rawDescGZIP(), []int{0}
}

func (x *PublicVote) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublicVote) GetElection() *election.Election {
	if x != nil {
		return x.Election
	}
	return nil
}

func (x *PublicVote) GetPublic() *public.Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type GetAllPublucVotes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicVotes []*PublicVote `protobuf:"bytes,1,rep,name=publicVotes,proto3" json:"publicVotes,omitempty"`
}

func (x *GetAllPublucVotes) Reset() {
	*x = GetAllPublucVotes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_vote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPublucVotes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPublucVotes) ProtoMessage() {}

func (x *GetAllPublucVotes) ProtoReflect() protoreflect.Message {
	mi := &file_public_vote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPublucVotes.ProtoReflect.Descriptor instead.
func (*GetAllPublucVotes) Descriptor() ([]byte, []int) {
	return file_public_vote_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllPublucVotes) GetPublicVotes() []*PublicVote {
	if x != nil {
		return x.PublicVotes
	}
	return nil
}

var File_public_vote_proto protoreflect.FileDescriptor

var file_public_vote_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0e, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x0a, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x49, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x75, 0x63, 0x56, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x0b, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x32, 0xf5, 0x01, 0x0a, 0x11, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f,
	0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42,
	0x79, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x75, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_public_vote_proto_rawDescOnce sync.Once
	file_public_vote_proto_rawDescData = file_public_vote_proto_rawDesc
)

func file_public_vote_proto_rawDescGZIP() []byte {
	file_public_vote_proto_rawDescOnce.Do(func() {
		file_public_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_vote_proto_rawDescData)
	})
	return file_public_vote_proto_rawDescData
}

var file_public_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_public_vote_proto_goTypes = []interface{}{
	(*PublicVote)(nil),        // 0: protos.PublicVote
	(*GetAllPublucVotes)(nil), // 1: protos.GetAllPublucVotes
	(*election.Election)(nil), // 2: protos.Election
	(*public.Public)(nil),     // 3: protos.Public
	(*election.ById)(nil),     // 4: protos.ById
	(*election.Void)(nil),     // 5: protos.Void
}
var file_public_vote_proto_depIdxs = []int32{
	2, // 0: protos.PublicVote.election:type_name -> protos.Election
	3, // 1: protos.PublicVote.public:type_name -> protos.Public
	0, // 2: protos.GetAllPublucVotes.publicVotes:type_name -> protos.PublicVote
	0, // 3: protos.PublicVoteService.CreatePublicVote:input_type -> protos.PublicVote
	4, // 4: protos.PublicVoteService.DeletePublicVote:input_type -> protos.ById
	0, // 5: protos.PublicVoteService.UpdatePublicVote:input_type -> protos.PublicVote
	4, // 6: protos.PublicVoteService.GetByIdPublicVote:input_type -> protos.ById
	5, // 7: protos.PublicVoteService.CreatePublicVote:output_type -> protos.Void
	5, // 8: protos.PublicVoteService.DeletePublicVote:output_type -> protos.Void
	5, // 9: protos.PublicVoteService.UpdatePublicVote:output_type -> protos.Void
	1, // 10: protos.PublicVoteService.GetByIdPublicVote:output_type -> protos.GetAllPublucVotes
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_public_vote_proto_init() }
func file_public_vote_proto_init() {
	if File_public_vote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicVote); i {
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
		file_public_vote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPublucVotes); i {
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
			RawDescriptor: file_public_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_vote_proto_goTypes,
		DependencyIndexes: file_public_vote_proto_depIdxs,
		MessageInfos:      file_public_vote_proto_msgTypes,
	}.Build()
	File_public_vote_proto = out.File
	file_public_vote_proto_rawDesc = nil
	file_public_vote_proto_goTypes = nil
	file_public_vote_proto_depIdxs = nil
}