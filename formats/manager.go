package formats

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


func GetApiVersionsV0() Request {
	return Request{
		key:     18,
		version: 0,
		body:    apiversionsFormat1{},
	}}

func GetApiVersionsV1() Request {
	return Request{
		key:     18,
		version: 1,
		body:    apiversionsFormat1{},
	}}

func GetApiVersionsV2() Request {
	return Request{
		key:     18,
		version: 2,
		body:    apiversionsFormat1{},
	}}

func GetMetadataV0(topics []string) Request {
	return Request{
		3,
		0,
		createMetadataV0_V3(topics),
	}
}

func GetMetadataV1(topics []string) Request {
	return Request{
		3,
		1,
		createMetadataV0_V3(topics),
	}
}

func GetMetadataV2(topics []string) Request {
	return Request{
		3,
		2,
		createMetadataV0_V3(topics),
	}
}

func GetMetadataV3(topics []string) Request {
	return Request{
		3,
		3,
		createMetadataV0_V3(topics),
	}
}

func GetMetadataV4(topics []string, allowAutoTopicCreation bool) Request {
	return Request{
		key:     3,
		version: 4,
		body:    createMetadataV4_V7(topics, allowAutoTopicCreation),
	}
}

func GetMetadataV5(topics []string, allowAutoTopicCreation bool) Request {
	return Request{
		key:     3,
		version: 5,
		body:    createMetadataV4_V7(topics, allowAutoTopicCreation),
	}
}

func GetMetadataV6(topics []string, allowAutoTopicCreation bool) Request {
	return Request{
		key:     3,
		version: 6,
		body:    createMetadataV4_V7(topics, allowAutoTopicCreation),
	}
}

func GetMetadataV7(topics []string, allowAutoTopicCreation bool) Request {
	return Request{
		key:     3,
		version: 7,
		body:    createMetadataV4_V7(topics, allowAutoTopicCreation),
	}
}

func GetMetadataV8(topics []string, allowAutoTopicCreation, inclClusterAuth, inclTopicAuth bool) Request {
	return Request{
		key:     3,
		version: 8,
		body:    createMetadataV8(topics, allowAutoTopicCreation, inclClusterAuth, inclTopicAuth),
	}
}