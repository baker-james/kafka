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

	var body bytes.Buffer

	body.Write(req.Key())
	body.Write(req.Version())
	body.Write(builder.corrilationId.Bytes())
	body.Write(builder.clientId.Bytes())

	body.Write(req.Details())

	var bodyLength = formats.KafkaInt32(body.Len())

	return append(bodyLength.Bytes(), body.Bytes()...)
}