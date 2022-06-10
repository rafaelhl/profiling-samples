package samples_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rafaelhl/profiling-samples/samples"
)

func TestUnmarshall(t *testing.T) {
	samples.Path = ""
	result := samples.Unmarshall[map[string]int](10)

	assert.Equal(t, map[string]int{
		"profiling": 1,
	}, result)
}
