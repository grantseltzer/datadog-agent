package aggregator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetEmptyFlush(t *testing.T) {
	// Initialize set
	set := NewSet()

	// Flush w/o samples: error
	_, err := set.flush(50)
	assert.NotNil(t, err)
}

func TestSetAddSample(t *testing.T) {
	// Initialize set
	set := NewSet()

	// Add samples
	sampleValues := []string{"a", "b", "c", "c", "b"}
	for _, sampleValue := range sampleValues {
		sample := MetricSample{RawValue: sampleValue}
		set.addSample(&sample, 55)
	}
	series, err := set.flush(60)
	require.Nil(t, err)

	require.Len(t, series, 1)
	require.Len(t, series[0].Points, 1)
	assert.EqualValues(t, 3, series[0].Points[0].Value)
	assert.EqualValues(t, 60, series[0].Points[0].Ts)

	// Add a few new samples and flush: the set should've been reset after the previous flush
	sampleValues = []string{"b", "b"}
	for _, sampleValue := range sampleValues {
		sample := MetricSample{RawValue: sampleValue}
		set.addSample(&sample, 65)
	}
	series, err = set.flush(70)
	require.Nil(t, err)
	require.Len(t, series, 1)
	require.Len(t, series[0].Points, 1)
	assert.EqualValues(t, 1, series[0].Points[0].Value)
	assert.EqualValues(t, 70, series[0].Points[0].Ts)

	// Flush w/o samples: error
	_, err = set.flush(80)
	assert.NotNil(t, err)
}
