package formats

import (
	"bytes"
	"encoding/binary"
)

const (
	int16Size = 2
	int32Size = 4
)

type Byter interface {
	Bytes() []byte
}

type ByterGroup []Byter

func (com ByterGroup) Bytes() []byte {
	var buf bytes.Buffer

	for i := range com {
		buf.Write(com[i].Bytes())
	}

	return buf.Bytes()
}

type KafkaBoolean bool

func (b KafkaBoolean) Bytes() []byte {
	if !b {
		return []byte{0x00}
	}
	return []byte{0x01}
}

type KafkaInt16 int16

func (i KafkaInt16) Bytes() []byte {
	slice := make([]byte, int16Size)
	binary.BigEndian.PutUint16(slice, uint16(i))
	return slice
}

type KafkaInt32 int32

func (i KafkaInt32) Bytes() []byte {
	slice := make([]byte, int32Size)
	binary.BigEndian.PutUint32(slice, uint32(i))
	return slice
}

type KafkaString string

func (str KafkaString) Bytes() []byte {
	length := KafkaInt16(len(str))
	b := []byte(str)
	return append(length.Bytes(), b...)
}

type KafkaArray []ByterGroup

func (arr KafkaArray) Bytes() []byte {
	var buf bytes.Buffer
	var length = KafkaInt32(len(arr))

	buf.Write(length.Bytes())
	for i := range arr {
		buf.Write(arr[i].Bytes())
	}

	return buf.Bytes()
}

type KafkaNull struct{}

func (k KafkaNull) Bytes() []byte {
	return nil
}