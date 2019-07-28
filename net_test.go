package kafka

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

type mockReadWriter struct {
	reader   io.Reader
	readErr error
	writeErr error
}

func(mock mockReadWriter) Read(p []byte) (n int, err error) {
	if mock.readErr != nil {
		return 0, mock.readErr
	}
	return mock.reader.Read(p)
}

func (mock mockReadWriter) Write(p []byte) (n int, err error) {
	return 0, mock.writeErr
}

func TestSend(t *testing.T) {
	var mock = mockReadWriter{
		reader:   strings.NewReader("mock"),
		readErr: nil,
		writeErr: nil,
	}

	req := []byte {
		0x00, 0x00, 0x00, 0x1a,
		0x00, 0x12, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x63, 0x6f,
		0x6e, 0x73, 0x6f, 0x6c,
		0x65, 0x2d, 0x70, 0x72,
		0x6f, 0x64, 0x75, 0x63,
		0x65, 0x72,
	}

	actRes, actErr :=  Send(mock, req)

	assert.Equal(t, []byte("mock"), actRes, "Should not be tampered")
	assert.Equal(t, error(nil), actErr, "Should not be tampered")
}

func TestSendWithWriteErr(t *testing.T) {
	var mock = mockReadWriter{
		reader:   strings.NewReader("mock"),
		readErr: nil,
		writeErr: errors.New("Failure"),
	}

	req := []byte {
		0x00, 0x00, 0x00, 0x1a,
		0x00, 0x12, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x63, 0x6f,
		0x6e, 0x73, 0x6f, 0x6c,
		0x65, 0x2d, 0x70, 0x72,
		0x6f, 0x64, 0x75, 0x63,
		0x65, 0x72,
	}

	actRes, actErr :=  Send(mock, req)

	assert.Equal(t, []byte(nil), actRes, "Should not be tampered")
	assert.Equal(t, mock.writeErr, actErr, "Should not be tampered")
}

func TestSendWithReadErr(t *testing.T) {
	var mock = mockReadWriter{
		reader:   strings.NewReader("mock"),
		readErr: errors.New("Failure"),
		writeErr: nil,
	}

	req := []byte {
		0x00, 0x00, 0x00, 0x1a,
		0x00, 0x12, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x63, 0x6f,
		0x6e, 0x73, 0x6f, 0x6c,
		0x65, 0x2d, 0x70, 0x72,
		0x6f, 0x64, 0x75, 0x63,
		0x65, 0x72,
	}

	actRes, actErr :=  Send(mock, req)

	assert.Equal(t, []byte(nil), actRes, "Should not be tampered")
	assert.Equal(t, mock.readErr, actErr, "Should not be tampered")
}