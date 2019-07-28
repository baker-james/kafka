package kafka

import (
	"bytes"
	"encoding/binary"
)

func buildPayload(clientId string) []byte {
	var body bytes.Buffer
	var corrId int

	key, version, detail := buildApiVersions()

	body.Write(key)
	body.Write(version)

	bCorrId := intToBigEndianSlice(corrId, 32)
	body.Write(bCorrId)

	bLenClientId := intToBigEndianSlice(len(clientId), 16)
	body.Write(bLenClientId)
	body.WriteString(clientId)

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
	case 64:
		binary.BigEndian.PutUint64(slice, uint64(val))
	}

	return slice
}