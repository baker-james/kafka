package kafka

import (
	"encoding/binary"
)

const (
	int16Size = 2
	int32Size = 4
)

type Byter interface {
	Bytes() []byte
}

type KafkaInt16 int16
type KafkaInt32 int32
type KafkaString string
type KafkaNull struct{}

func (k KafkaInt16) Bytes() []byte {
	slice := make([]byte, int16Size)
	binary.BigEndian.PutUint16(slice, uint16(k))
	return slice
}

func (k KafkaInt32) Bytes() []byte {
	slice := make([]byte, int32Size)
	binary.BigEndian.PutUint32(slice, uint32(k))
	return slice
}

func (k KafkaString) Bytes() []byte {
	length := KafkaInt16(len(k))
	b := []byte(k)
	return append(length.Bytes(), b...)
}

func (k KafkaNull) Bytes() []byte {
	return nil
}