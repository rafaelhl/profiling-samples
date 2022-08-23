package samples_test

import (
	"testing"

	"github.com/rafaelhl/profiling-samples/samples"
	"github.com/stretchr/testify/assert"
)

func TestTimezone(t *testing.T) {
	tz := samples.Timezone(10)

	assert.NotNil(t, tz)
}
