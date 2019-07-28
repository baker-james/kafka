package kafka

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildPayload(t *testing.T) {
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
	actual := buildPayload("console-producer")
	assert.Equal(t, expected, actual, "Should be equal")
}

func Test_buildPayloadWithEmptyClientId(t *testing.T) {
	expected := []byte{
		0x00, 0x00, 0x00, 0x0a,
		0x00, 0x12, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
	}
	actual := buildPayload("")
	assert.Equal(t, expected, actual, "Should be equal")
}
