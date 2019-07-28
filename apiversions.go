package kafka

func buildApiVersions() ([]byte, []byte, []byte) {
	apiKey := 18
	version := 2
	detail := []byte(nil)

	return intToBigEndianSlice(apiKey, 16),
		intToBigEndianSlice(version, 16),
		detail
}
