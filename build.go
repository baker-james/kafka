package kafka

import (
	"bytes"
)

type requestBuilder struct {
	corrilationId KafkaInt32
	clientId      KafkaString
}

func NewRequestBuilder(clientId string) *requestBuilder {
	return &requestBuilder{
		clientId: KafkaString(clientId),
	}
}

func (builder *requestBuilder) BuildPayload() []byte {
	var body bytes.Buffer
	defer func() {
		builder.corrilationId++
	}()

	key, version, detail := buildApiVersions()

	body.Write(key.Bytes())
	body.Write(version.Bytes())
	body.Write(builder.corrilationId.Bytes())
	body.Write(builder.clientId.Bytes())

	body.Write(detail.Bytes())

	var bodyLength = KafkaInt32(body.Len())

	return append(bodyLength.Bytes(), body.Bytes()...)
}