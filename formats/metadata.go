package formats

import "bytes"

type metadataFormat1 struct {
	buf bytes.Buffer
	KafkaArray
}

type metadataFormat2 struct {
	metadataFormat1
	allowAutoTopicCreation KafkaBoolean
}

func (m2 metadataFormat2) Bytes() []byte {
	m2.buf.Reset()
	m2.buf.Write(m2.metadataFormat1.Bytes())
	m2.buf.Write(m2.allowAutoTopicCreation.Bytes())
	return m2.buf.Bytes()
}

type metadataFormat3 struct {
	metadataFormat2
	inclClusterAuth KafkaBoolean
	inclTopicAuth KafkaBoolean
}

func (m3 metadataFormat3) Bytes() []byte {
	m3.buf.Reset()
	m3.buf.Write(m3.metadataFormat2.Bytes())
	m3.buf.Write(m3.inclClusterAuth.Bytes())
	m3.buf.Write(m3.inclTopicAuth.Bytes())
	return m3.buf.Bytes()
}

func createMetadataV0_V3(topics []string) (body metadataFormat1) {
	body.KafkaArray = make(KafkaArray, len(topics))
	for i, topic := range topics {
		body.KafkaArray[i] = ByterGroup{
			KafkaString(topic),
		}
	}
	return
}

func createMetadataV4_V7(topics []string, allowAutoTopicCreation bool) (body metadataFormat2) {
	body.metadataFormat1 = createMetadataV0_V3(topics)
	body.allowAutoTopicCreation = KafkaBoolean(allowAutoTopicCreation)
	return
}

func createMetadataV8(topics []string, allowAutoTopicCreation, inclClusterAuth, inclTopicAuth bool) (body metadataFormat3) {
	body.metadataFormat2 = createMetadataV4_V7(topics, allowAutoTopicCreation)
	body.inclClusterAuth = KafkaBoolean(inclClusterAuth)
	body.inclTopicAuth = KafkaBoolean(inclTopicAuth)
	return
}

