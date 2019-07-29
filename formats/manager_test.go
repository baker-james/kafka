package formats

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGetUnsupported(t *testing.T) {
	expReq := Request{}
	expSupported := false

	actReq, actSupported := GetApiVersions(math.MaxInt64, math.MaxInt64)

	assert.Equalf(t, expReq, actReq, "unsupported versions should return empty struct")
	assert.Equal(t, expSupported, actSupported, "unsupported versions should return false boolean")
}

func TestGetLargeMax(t *testing.T) {
	// Overwrite map for test consistency

	expReq := Request{
		key:     18,
		version: 2,
		details: apiVersions(),
	}
	expSupported := true

	actReq, actSupported := GetApiVersions(0, math.MaxInt32)

	assert.Equalf(t, expReq, actReq, "should return the highest possible request struct")
	assert.Equal(t, expSupported, actSupported, "should be true when version is supported")
}

func TestGetFixedVersion(t *testing.T) {
	expReq := Request{
		key:     18,
		version: 2,
		details: apiVersions(),
	}
	expSupported := true

	actReq, actSupported := GetApiVersions(2, 2)

	assert.Equalf(t, expReq, actReq, "should return the correct version")
	assert.Equal(t, expSupported, actSupported, "should be true when version is supported")
}