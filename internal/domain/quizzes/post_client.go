package quizzes

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protowire"
)

// postServiceClient handles communication with lms-post-service
type postServiceClient struct {
	Log *log.Logger
}

// createPost calls lms-post-service to create a post with type QUIZ (5)
func (c *postServiceClient) createPost(ctx context.Context, subjectClassID, topicSubjectID, quizID, title, userID string) {
	postServiceAddr := os.Getenv("POST_SERVICE_ADDRESS")
	if postServiceAddr == "" {
		c.Log.Println("POST_SERVICE_ADDRESS not configured, skipping post creation")
		return
	}

	conn, err := grpc.NewClient(postServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.Log.Printf("error connecting to post service: %v", err)
		return
	}
	defer conn.Close()

	// Forward user metadata
	md := metadata.New(map[string]string{
		"user_id": userID,
	})
	outCtx := metadata.NewOutgoingContext(ctx, md)

	// Manually encode CreatePostRequest protobuf
	// Field 2: subject_class_id (string)
	// Field 3: topic_subject_id (string)
	// Field 4: type (int32) = 5 (QUIZ)
	// Field 5: type_id (string) = quizID
	// Field 6: title (string)
	// Field 12: is_published (bool) = true
	var buf []byte
	buf = protowire.AppendTag(buf, 2, protowire.BytesType)
	buf = protowire.AppendString(buf, subjectClassID)
	buf = protowire.AppendTag(buf, 3, protowire.BytesType)
	buf = protowire.AppendString(buf, topicSubjectID)
	buf = protowire.AppendTag(buf, 4, protowire.VarintType)
	buf = protowire.AppendVarint(buf, 5) // QUIZ = 5
	buf = protowire.AppendTag(buf, 5, protowire.BytesType)
	buf = protowire.AppendString(buf, quizID)
	buf = protowire.AppendTag(buf, 6, protowire.BytesType)
	buf = protowire.AppendString(buf, title)
	buf = protowire.AppendTag(buf, 12, protowire.VarintType)
	buf = protowire.AppendVarint(buf, 1) // is_published = true

	// Use raw codec to send the request
	var resp []byte
	err = conn.Invoke(outCtx, "/posts.Posts/CreatePost", &rawMessage{data: buf}, &rawMessage{data: resp}, grpc.ForceCodec(rawCodec{}))
	if err != nil {
		c.Log.Printf("error creating post for quiz: %v", err)
		return
	}

	c.Log.Printf("successfully created post for quiz %s", quizID)
}

// rawMessage wraps raw bytes for gRPC codec
type rawMessage struct {
	data []byte
}

// rawCodec is a gRPC codec that passes raw bytes
type rawCodec struct{}

func (rawCodec) Marshal(v interface{}) ([]byte, error) {
	msg, ok := v.(*rawMessage)
	if !ok {
		return nil, nil
	}
	return msg.data, nil
}

func (rawCodec) Unmarshal(data []byte, v interface{}) error {
	msg, ok := v.(*rawMessage)
	if !ok {
		return nil
	}
	msg.data = data
	return nil
}

func (rawCodec) Name() string {
	return "proto"
}
