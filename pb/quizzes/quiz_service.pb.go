// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: quizzes/quiz_service.proto

package quizzes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_quizzes_quiz_service_proto protoreflect.FileDescriptor

var file_quizzes_quiz_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x71, 0x75,
	0x69, 0x7a, 0x7a, 0x65, 0x73, 0x1a, 0x1a, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2f, 0x71,
	0x75, 0x69, 0x7a, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x5f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x71, 0x75, 0x69,
	0x7a, 0x7a, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe3, 0x01, 0x0a, 0x07, 0x51,
	0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x69, 0x7a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0d, 0x2e, 0x71, 0x75, 0x69,
	0x7a, 0x7a, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x69, 0x7a, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2e,
	0x51, 0x75, 0x69, 0x7a, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x13, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x69, 0x7a, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e,
	0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x71, 0x75, 0x69,
	0x7a, 0x7a, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x69, 0x7a, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x1b, 0x2e, 0x71,
	0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x51, 0x75, 0x69, 0x7a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x13, 0x2e, 0x71, 0x75, 0x69, 0x7a,
	0x7a, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x69, 0x7a, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x00,
	0x42, 0x14, 0x5a, 0x12, 0x70, 0x62, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x3b, 0x71,
	0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_quizzes_quiz_service_proto_goTypes = []any{
	(*QuizUpdateInput)(nil),    // 0: quizzes.QuizUpdateInput
	(*QuizAnswerInput)(nil),    // 1: quizzes.QuizAnswerInput
	(*Id)(nil),                 // 2: quizzes.Id
	(*GetResultQuizInput)(nil), // 3: quizzes.GetResultQuizInput
	(*Quiz)(nil),               // 4: quizzes.Quiz
	(*QuizAnswer)(nil),         // 5: quizzes.QuizAnswer
}
var file_quizzes_quiz_service_proto_depIdxs = []int32{
	0, // 0: quizzes.Quizzes.Update:input_type -> quizzes.QuizUpdateInput
	1, // 1: quizzes.Quizzes.Answer:input_type -> quizzes.QuizAnswerInput
	2, // 2: quizzes.Quizzes.Get:input_type -> quizzes.Id
	3, // 3: quizzes.Quizzes.GetResultQuiz:input_type -> quizzes.GetResultQuizInput
	4, // 4: quizzes.Quizzes.Update:output_type -> quizzes.Quiz
	5, // 5: quizzes.Quizzes.Answer:output_type -> quizzes.QuizAnswer
	4, // 6: quizzes.Quizzes.Get:output_type -> quizzes.Quiz
	5, // 7: quizzes.Quizzes.GetResultQuiz:output_type -> quizzes.QuizAnswer
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_quizzes_quiz_service_proto_init() }
func file_quizzes_quiz_service_proto_init() {
	if File_quizzes_quiz_service_proto != nil {
		return
	}
	file_quizzes_quiz_message_proto_init()
	file_quizzes_quiz_input_proto_init()
	file_quizzes_generic_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_quizzes_quiz_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_quizzes_quiz_service_proto_goTypes,
		DependencyIndexes: file_quizzes_quiz_service_proto_depIdxs,
	}.Build()
	File_quizzes_quiz_service_proto = out.File
	file_quizzes_quiz_service_proto_rawDesc = nil
	file_quizzes_quiz_service_proto_goTypes = nil
	file_quizzes_quiz_service_proto_depIdxs = nil
}
