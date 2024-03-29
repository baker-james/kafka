package kafka

import (
	"github.com/baker-james/kafka/formats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildApiVersionsV0(t *testing.T) {
	builder := NewRequestBuilder("console-producer")

	expected := []byte {
		0x00, 0x00, 0x00, 0x1a,
		0x00, 0x12, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x63, 0x6f,
		0x6e, 0x73, 0x6f, 0x6c,
		0x65, 0x2d, 0x70, 0x72,
		0x6f, 0x64, 0x75, 0x63,
		0x65, 0x72,
	}

	req := formats.GetApiVersionsV0()
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Should be equal")
}

func Test_buildApiVersionsV1(t *testing.T) {
	builder := NewRequestBuilder("console-producer")

	expected := []byte {
		0x00, 0x00, 0x00, 0x1a,
		0x00, 0x12, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x63, 0x6f,
		0x6e, 0x73, 0x6f, 0x6c,
		0x65, 0x2d, 0x70, 0x72,
		0x6f, 0x64, 0x75, 0x63,
		0x65, 0x72,
	}

	req := formats.GetApiVersionsV1()
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Should be equal")
}

func Test_buildApiVersionsV2(t *testing.T) {
	builder := NewRequestBuilder("console-producer")

	expected := []byte {
		0x00, 0x00, 0x00, 0x1a,
		0x00, 0x12, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x63, 0x6f,
		0x6e, 0x73, 0x6f, 0x6c,
		0x65, 0x2d, 0x70, 0x72,
		0x6f, 0x64, 0x75, 0x63,
		0x65, 0x72,
	}

	req := formats.GetApiVersionsV2()
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Should be equal")
}

func Test_buildPayloadWithEmptyClientId(t *testing.T) {
	builder := NewRequestBuilder("")

	expected := []byte{
		0x00, 0x00, 0x00, 0x0a,
		0x00, 0x12, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
	}
	req := formats.GetApiVersionsV2()
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Should be equal")
}

func TestBuilder_MetadataV8_WithNil_TopicCreation(t *testing.T) {
	builder := NewRequestBuilder("console-producer")

	expected := []byte{
		0x00, 0x00, 0x00, 0x21, 0x00, 0x03,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x63, 0x6f, 0x6e, 0x73,
		0x6f, 0x6c, 0x65, 0x2d, 0x70, 0x72,
		0x6f, 0x64, 0x75, 0x63, 0x65, 0x72,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
		0x00,
	}
	req := formats.GetMetadataV8([]string{}, true, false, false)
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "????")
}

func TestBuilder_MetadataV1_SingleTopic(t *testing.T) {
	builder := NewRequestBuilder("test")

	expected := []byte{
		0x00, 0x00, 0x00, 0x17, 0x00,
		0x03, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x04, 0x74,
		0x65, 0x73, 0x74, 0x00, 0x00,
		0x00, 0x01, 0x00, 0x03, 0x62,
		0x6f, 0x6f,
	}
	req := formats.GetMetadataV1([]string{"boo"})
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Array length and items should be correctly represented")
}

func TestBuilder_MetadataV2_SingleTopic(t *testing.T) {
	builder := NewRequestBuilder("test")

	expected := []byte{
		0x00, 0x00, 0x00, 0x17, 0x00,
		0x03, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x04, 0x74,
		0x65, 0x73, 0x74, 0x00, 0x00,
		0x00, 0x01, 0x00, 0x03, 0x62,
		0x6f, 0x6f,
	}
	req := formats.GetMetadataV2([]string{"boo"})
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Array length and items should be correctly represented")
}

func TestBuilder_MetadataV3_SingleTopic(t *testing.T) {
	builder := NewRequestBuilder("test")

	expected := []byte{
		0x00, 0x00, 0x00, 0x17, 0x00,
		0x03, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x04, 0x74,
		0x65, 0x73, 0x74, 0x00, 0x00,
		0x00, 0x01, 0x00, 0x03, 0x62,
		0x6f, 0x6f,
	}
	req := formats.GetMetadataV3([]string{"boo"})
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Array length and items should be correctly represented")
}

func TestBuilder_MetadataV0_SingleTopic(t *testing.T) {
	builder := NewRequestBuilder("test")

	expected := []byte{
		0x00, 0x00, 0x00, 0x17, 0x00,
		0x03, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x04, 0x74,
		0x65, 0x73, 0x74, 0x00, 0x00,
		0x00, 0x01, 0x00, 0x03, 0x62,
		0x6f, 0x6f,
	}
	req := formats.GetMetadataV0([]string{"boo"})
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Array length and items should be correctly represented")
}

func TestBuilder_MetadataV4_SingleTopic_WithTopicCreation(t *testing.T) {
	builder := NewRequestBuilder("test")

	expected := []byte{
		0x00, 0x00, 0x00, 0x1a, 0x00, 0x03,
		0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x04, 0x74, 0x65, 0x73, 0x74,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x05,
		0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x01,
	}
	req := formats.GetMetadataV4([]string{"yello"}, true)
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Boolean value should be correctly represented")
}

func TestBuilder_MetadataV5_SingleTopic_WithTopicCreation(t *testing.T) {
	builder := NewRequestBuilder("test")

	expected := []byte{
		0x00, 0x00, 0x00, 0x1a, 0x00, 0x03,
		0x00, 0x05, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x04, 0x74, 0x65, 0x73, 0x74,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x05,
		0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x01,
	}
	req := formats.GetMetadataV5([]string{"yello"}, true)
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Boolean value should be correctly represented")
}

func TestBuilder_MetadataV6_SingleTopic_WithTopicCreation(t *testing.T) {
	builder := NewRequestBuilder("test")

	expected := []byte{
		0x00, 0x00, 0x00, 0x1a, 0x00, 0x03,
		0x00, 0x06, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x04, 0x74, 0x65, 0x73, 0x74,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x05,
		0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x01,
	}
	req := formats.GetMetadataV6([]string{"yello"}, true)
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Boolean value should be correctly represented")
}


func TestBuilder_MetadataV7_SingleTopic_WithTopicCreation(t *testing.T) {
	builder := NewRequestBuilder("test")

	expected := []byte{
		0x00, 0x00, 0x00, 0x1a, 0x00, 0x03,
		0x00, 0x07, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x04, 0x74, 0x65, 0x73, 0x74,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x05,
		0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x01,
	}
	req := formats.GetMetadataV7([]string{"yello"}, true)
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Boolean value should be correctly represented")
}

func Test_CorrilationIdIncrements(t *testing.T) {
	builder := NewRequestBuilder("")

	expected := []byte{
		0x00, 0x00, 0x00, 0x0a,
		0x00, 0x12, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x01, // Should increment after every call
		0x00, 0x00,
	}

	req := formats.GetApiVersionsV2()
	builder.BuildPayload(req)
	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "CorrilationID should have incremented")
}
