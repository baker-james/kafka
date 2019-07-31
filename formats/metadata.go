package formats


type metadataV1 struct {
	KafkaArray
}

type metadataV4 struct {
	metadataV1
	createTopics KafkaBoolean
}

func (m metadataV4) Bytes() []byte {
	return append(
		m.metadataV1.Bytes(),
		m.createTopics.Bytes()...
		)
}



