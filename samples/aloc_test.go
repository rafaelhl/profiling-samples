package samples_test

import (
	"testing"

	"github.com/rafaelhl/profiling-samples/samples"
	"github.com/stretchr/testify/assert"
)

func TestAloc(t *testing.T) {
	result := samples.Aloc("test", 10)

	assert.Len(t, result, 11)
}
