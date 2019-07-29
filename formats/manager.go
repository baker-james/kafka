package formats

type apikey int

const Metadata KafkaInt16 = 3
const ApiVersions KafkaInt16 = 18

type Request struct {
	key KafkaInt16
	version KafkaInt16
	details Byter
}

func (req Request) Key() []byte {
	return req.key.Bytes()
}

func (req Request) Version() []byte {
	return req.version.Bytes()
}

func (req Request) Details() []byte {
	return req.details.Bytes()
}

type versions []Byter

var requestVersions = map[KafkaInt16]versions {
	ApiVersions: {
		apiVersions(),
		apiVersions(),
		apiVersions(),
	},
	Metadata: {
		metadata(),
	},
}

func GetApiVersions(min, max int) (Request, bool) {
	var req = Request{
		key: ApiVersions,
	}

	supported := requestVersions[req.key]

	version, isSupported := chooseVersion(supported, min, max)
	if !isSupported {
		return Request{}, isSupported
	}
	req.version = version
	req.details = supported[req.version]

	return req, true
}

func chooseVersion(supported versions, min, max int) (KafkaInt16, bool) {
	maxAvailable := len(supported)-1

	switch {
	case min > len(supported):
		return 0, false
	case max > maxAvailable:
		return KafkaInt16(maxAvailable), true
	default:
		return KafkaInt16(max), true
	}
}