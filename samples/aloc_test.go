package samples_test

import (
	"testing"
	"errors"

	"github.com/rafaelhl/profiling-samples/samples"
	"github.com/stretchr/testify/assert"
)

func TestAloc(t *testing.T) {
	result := samples.Aloc("test", 10)

	assert.Len(t, result, 11)
}

func TestAlocError(t *testing.T) {
	result := samples.Aloc(errors.New("test"), 10)

	assert.Len(t, result, 0)
}
