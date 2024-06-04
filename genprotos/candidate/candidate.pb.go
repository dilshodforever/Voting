// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: candidate.proto

package candidate

import (
	election "genprotos/election"
	party "genprotos/party"
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

type Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Election *election.Election `protobuf:"bytes,2,opt,name=election,proto3" json:"election,omitempty"`
	Party    *party.Party       `protobuf:"bytes,3,opt,name=party,proto3" json:"party,omitempty"`
	Public   *public.Public     `protobuf:"bytes,4,opt,name=public,proto3" json:"public,omitempty"`
	Date     string             `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Candidate) Reset() {
	*x = Candidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candidate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidate) ProtoMessage() {}

func (x *Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_candidate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidate.ProtoReflect.Descriptor instead.
func (*Candidate) Descriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{0}
}

func (x *Candidate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Candidate) GetElection() *election.Election {
	if x != nil {
		return x.Election
	}
	return nil
}

func (x *Candidate) GetParty() *party.Party {
	if x != nil {
		return x.Party
	}
	return nil
}

func (x *Candidate) GetPublic() *public.Public {
	if x != nil {
		return x.Public
	}
	return nil
}

func (x *Candidate) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetAllCandidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candidates []*Candidate `protobuf:"bytes,1,rep,name=Candidates,proto3" json:"Candidates,omitempty"`
}

func (x *GetAllCandidate) Reset() {
	*x = GetAllCandidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candidate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCandidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCandidate) ProtoMessage() {}

func (x *GetAllCandidate) ProtoReflect() protoreflect.Message {
	mi := &file_candidate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCandidate.ProtoReflect.Descriptor instead.
func (*GetAllCandidate) Descriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllCandidate) GetCandidates() []*Candidate {
	if x != nil {
		return x.Candidates
	}
	return nil
}

var File_candidate_proto protoreflect.FileDescriptor

var file_candidate_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0e, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x05,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x22, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x32, 0xec, 0x01, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42,
	0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_candidate_proto_rawDescOnce sync.Once
	file_candidate_proto_rawDescData = file_candidate_proto_rawDesc
)

func file_candidate_proto_rawDescGZIP() []byte {
	file_candidate_proto_rawDescOnce.Do(func() {
		file_candidate_proto_rawDescData = protoimpl.X.CompressGZIP(file_candidate_proto_rawDescData)
	})
	return file_candidate_proto_rawDescData
}

var file_candidate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_candidate_proto_goTypes = []interface{}{
	(*Candidate)(nil),         // 0: protos.Candidate
	(*GetAllCandidate)(nil),   // 1: protos.GetAllCandidate
	(*election.Election)(nil), // 2: protos.Election
	(*party.Party)(nil),       // 3: protos.Party
	(*public.Public)(nil),     // 4: protos.Public
	(*election.ById)(nil),     // 5: protos.ById
	(*election.Void)(nil),     // 6: protos.Void
}
var file_candidate_proto_depIdxs = []int32{
	2, // 0: protos.Candidate.election:type_name -> protos.Election
	3, // 1: protos.Candidate.party:type_name -> protos.Party
	4, // 2: protos.Candidate.public:type_name -> protos.Public
	0, // 3: protos.GetAllCandidate.Candidates:type_name -> protos.Candidate
	0, // 4: protos.CandidateService.CreateCandidate:input_type -> protos.Candidate
	5, // 5: protos.CandidateService.DeleteCandidate:input_type -> protos.ById
	0, // 6: protos.CandidateService.UpdateCandidate:input_type -> protos.Candidate
	5, // 7: protos.CandidateService.GetByIdCandidate:input_type -> protos.ById
	6, // 8: protos.CandidateService.CreateCandidate:output_type -> protos.Void
	6, // 9: protos.CandidateService.DeleteCandidate:output_type -> protos.Void
	6, // 10: protos.CandidateService.UpdateCandidate:output_type -> protos.Void
	1, // 11: protos.CandidateService.GetByIdCandidate:output_type -> protos.GetAllCandidate
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_candidate_proto_init() }
func file_candidate_proto_init() {
	if File_candidate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_candidate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candidate); i {
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
		file_candidate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCandidate); i {
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
			RawDescriptor: file_candidate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_candidate_proto_goTypes,
		DependencyIndexes: file_candidate_proto_depIdxs,
		MessageInfos:      file_candidate_proto_msgTypes,
	}.Build()
	File_candidate_proto = out.File
	file_candidate_proto_rawDesc = nil
	file_candidate_proto_goTypes = nil
	file_candidate_proto_depIdxs = nil
}