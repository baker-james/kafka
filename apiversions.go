package kafka

func buildApiVersions() (KafkaInt16, KafkaInt16, Byter) {
	var apiKey KafkaInt16 = 18
	var version KafkaInt16 = 2
	var detail KafkaNull

	return apiKey, version, detail
}
