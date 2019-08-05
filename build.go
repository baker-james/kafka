package kafka

import (
	"bytes"
	"github.com/baker-james/kafka/formats"
)

type builder struct {
	corrilationId formats.KafkaInt32
	clientId      formats.KafkaString
}

func NewRequestBuilder(clientId string) *builder {
	return &builder{
		clientId: formats.KafkaString(clientId),
	}
}

func (builder *builder) BuildPayload(req formats.Request) []byte {
	defer func() {
		builder.corrilationId++
	}()

	var payload bytes.Buffer

	payload.Write(req.Key())
	payload.Write(req.Version())
	payload.Write(builder.corrilationId.Bytes())
	payload.Write(builder.clientId.Bytes())

	payload.Write(req.Body())

	var length = formats.KafkaInt32(payload.Len())

	return append(length.Bytes(), payload.Bytes()...)
}