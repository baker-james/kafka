package formats

const Metadata KafkaInt16 = 3
const ApiVersions KafkaInt16 = 18

type Request struct {
	key     KafkaInt16
	version KafkaInt16
	body    Byter
}

func (req Request) Key() []byte {
	return req.key.Bytes()
}

func (req Request) Version() []byte {
	return req.version.Bytes()
}

func (req Request) Body() []byte {
	return req.body.Bytes()
}

func GetApiVersions(min, max int) (Request, bool) {
	var arr = []Byter {
		apiversions{},
		apiversions{},
		apiversions{},
	}

	version, supported := chooseVersion(len(arr)-1, min, max)
	if !supported {
		return Request{}, supported
	}

	return Request{
		key:     ApiVersions,
		version: version,
		body:    arr[version],
	}, supported
}

func GetMetadata(min, max int, topics []string) (Request, bool) {
	var body metadataV1

	version, supported := chooseVersion(3, min, max)
	if !supported {
		return Request{}, false
	}

	body.KafkaArray = make(KafkaArray, len(topics))
	for i, topic := range topics  {
		body.KafkaArray[i] = KafkaComposite{
			KafkaString(topic),
		}
	}

	return Request{
		key:     Metadata,
		version: version,
		body:    body,
	}, supported
}

func GetMatadataV4(min, max int, topics []string, createTopics bool) (Request, bool) {
	var body metadataV4
	version, supported := chooseVersion(3, min, max)
	if !supported {
		return Request{}, false
	}

	body.KafkaArray = make(KafkaArray, len(topics))
	for i, topic := range topics  {
		body.KafkaArray[i] = KafkaComposite{
			KafkaString(topic),
		}
	}
	body.createTopics = KafkaBoolean(createTopics)

	return Request{
		key:     Metadata,
		version: version,
		body:    body,
	}, supported
}

func chooseVersion(maxSupported, min, max int) (KafkaInt16, bool) {
	switch {
	case min > maxSupported:
		return -1, false
	case max > maxSupported:
		return KafkaInt16(maxSupported), true
	default:
		return KafkaInt16(max), true
	}
}