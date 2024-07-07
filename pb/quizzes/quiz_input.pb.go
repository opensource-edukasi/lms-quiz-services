// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: quizzes/quiz_input.proto

package quizzes

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

type QuizUpdateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	EndDate     string                 `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Question    []*QuestionUpdateInput `protobuf:"bytes,5,rep,name=question,proto3" json:"question,omitempty"`
}

func (x *QuizUpdateInput) Reset() {
	*x = QuizUpdateInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quizzes_quiz_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuizUpdateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuizUpdateInput) ProtoMessage() {}

func (x *QuizUpdateInput) ProtoReflect() protoreflect.Message {
	mi := &file_quizzes_quiz_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuizUpdateInput.ProtoReflect.Descriptor instead.
func (*QuizUpdateInput) Descriptor() ([]byte, []int) {
	return file_quizzes_quiz_input_proto_rawDescGZIP(), []int{0}
}

func (x *QuizUpdateInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QuizUpdateInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QuizUpdateInput) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *QuizUpdateInput) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *QuizUpdateInput) GetQuestion() []*QuestionUpdateInput {
	if x != nil {
		return x.Question
	}
	return nil
}

type QuizAnswerInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuizId         string                 `protobuf:"bytes,1,opt,name=quiz_id,json=quizId,proto3" json:"quiz_id,omitempty"`
	QuestionAnswer []*QuestionAnswerInput `protobuf:"bytes,2,rep,name=question_answer,json=questionAnswer,proto3" json:"question_answer,omitempty"`
}

func (x *QuizAnswerInput) Reset() {
	*x = QuizAnswerInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quizzes_quiz_input_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuizAnswerInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuizAnswerInput) ProtoMessage() {}

func (x *QuizAnswerInput) ProtoReflect() protoreflect.Message {
	mi := &file_quizzes_quiz_input_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuizAnswerInput.ProtoReflect.Descriptor instead.
func (*QuizAnswerInput) Descriptor() ([]byte, []int) {
	return file_quizzes_quiz_input_proto_rawDescGZIP(), []int{1}
}

func (x *QuizAnswerInput) GetQuizId() string {
	if x != nil {
		return x.QuizId
	}
	return ""
}

func (x *QuizAnswerInput) GetQuestionAnswer() []*QuestionAnswerInput {
	if x != nil {
		return x.QuestionAnswer
	}
	return nil
}

var File_quizzes_quiz_input_proto protoreflect.FileDescriptor

var file_quizzes_quiz_input_proto_rawDesc = []byte{
	0x0a, 0x18, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x71, 0x75, 0x69, 0x7a,
	0x7a, 0x65, 0x73, 0x1a, 0x1c, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0f, 0x51, 0x75, 0x69, 0x7a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x7a,
	0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x71, 0x0a, 0x0f, 0x51, 0x75, 0x69, 0x7a, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x71, 0x75, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x71, 0x75, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x52, 0x0e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x62, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65,
	0x73, 0x3b, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_quizzes_quiz_input_proto_rawDescOnce sync.Once
	file_quizzes_quiz_input_proto_rawDescData = file_quizzes_quiz_input_proto_rawDesc
)

func file_quizzes_quiz_input_proto_rawDescGZIP() []byte {
	file_quizzes_quiz_input_proto_rawDescOnce.Do(func() {
		file_quizzes_quiz_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_quizzes_quiz_input_proto_rawDescData)
	})
	return file_quizzes_quiz_input_proto_rawDescData
}

var file_quizzes_quiz_input_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_quizzes_quiz_input_proto_goTypes = []any{
	(*QuizUpdateInput)(nil),     // 0: quizzes.QuizUpdateInput
	(*QuizAnswerInput)(nil),     // 1: quizzes.QuizAnswerInput
	(*QuestionUpdateInput)(nil), // 2: quizzes.QuestionUpdateInput
	(*QuestionAnswerInput)(nil), // 3: quizzes.QuestionAnswerInput
}
var file_quizzes_quiz_input_proto_depIdxs = []int32{
	2, // 0: quizzes.QuizUpdateInput.question:type_name -> quizzes.QuestionUpdateInput
	3, // 1: quizzes.QuizAnswerInput.question_answer:type_name -> quizzes.QuestionAnswerInput
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_quizzes_quiz_input_proto_init() }
func file_quizzes_quiz_input_proto_init() {
	if File_quizzes_quiz_input_proto != nil {
		return
	}
	file_quizzes_question_input_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_quizzes_quiz_input_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*QuizUpdateInput); i {
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
		file_quizzes_quiz_input_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*QuizAnswerInput); i {
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
			RawDescriptor: file_quizzes_quiz_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_quizzes_quiz_input_proto_goTypes,
		DependencyIndexes: file_quizzes_quiz_input_proto_depIdxs,
		MessageInfos:      file_quizzes_quiz_input_proto_msgTypes,
	}.Build()
	File_quizzes_quiz_input_proto = out.File
	file_quizzes_quiz_input_proto_rawDesc = nil
	file_quizzes_quiz_input_proto_goTypes = nil
	file_quizzes_quiz_input_proto_depIdxs = nil
}
