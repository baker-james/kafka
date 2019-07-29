package kafka

import (
	"github.com/baker-james/kafka/formats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildApiVersions(t *testing.T) {
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
	req, isSupported := formats.GetApiVersions(2, 2)
	assert.True(t, isSupported, "expect to be supported")

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
	req, isSupported := formats.GetApiVersions(2, 2)
	assert.True(t, isSupported, "expect to be supported")

	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "Should be equal")
}

func Test_CorrilationIdIncrements(t *testing.T) {
	builder := NewRequestBuilder("")

	expected := []byte{
		0x00, 0x00, 0x00, 0x0a,
		0x00, 0x12, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x01, // Should increment after every call
		0x00, 0x00,
	}

	req, isSupported := formats.GetApiVersions(2, 2)
	assert.True(t, isSupported, "expect to be supported")

	builder.BuildPayload(req)

	actual := builder.BuildPayload(req)
	assert.Equal(t, expected, actual, "CorrilationID should have incremented")
}
