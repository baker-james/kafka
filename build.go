package kafka

import (
	"bytes"
	"encoding/binary"
)

type requestBuilder struct {
	corrId int
	clientId string
}

func NewRequestBuilder(clientId string) *requestBuilder {
	return &requestBuilder{
		clientId: clientId,
	}
}

func (builder *requestBuilder) BuildPayload() []byte {
	var body bytes.Buffer
	defer func() {
		builder.corrId++
	}()

	key, version, detail := buildApiVersions()

	body.Write(key)
	body.Write(version)

	bCorrId := intToBigEndianSlice(builder.corrId, 32)
	body.Write(bCorrId)

	bLenClientId := intToBigEndianSlice(len(builder.clientId), 16)
	body.Write(bLenClientId)
	body.WriteString(builder.clientId)

	body.Write(detail)
	lenBody := intToBigEndianSlice(body.Len(), 32)
	payload := append(lenBody, body.Bytes()...)

	return payload
}

func intToBigEndianSlice(val, size int) []byte {
	slice := make([]byte, size/8)
	switch size {
	case 16:
		binary.BigEndian.PutUint16(slice, uint16(val))
	case 32:
		binary.BigEndian.PutUint32(slice, uint32(val))
	}

	return slice
}