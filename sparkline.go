// Package sparkline handles creation of string sparklines.
package sparkline

import "math"

var blocks = []rune("▁▂▃▄▅▆▇█")

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Sparkline generates a sparkline for the provided values.
func Sparkline[T Number](values []T, opts ...Option) string {
	if len(values) == 0 {
		return ""
	}

	output := make([]rune, len(values))

	minValue := values[0]
	maxValue := values[0]

	for _, value := range values {
		if minValue > value {
			minValue = value
		}
		if maxValue < value {
			maxValue = value
		}
	}

	span := float64(maxValue - minValue)
	levelCount := float64(len(blocks) - 1)
	for i, value := range values {
		level := int(math.Floor(levelCount * float64(value-minValue) / span))
		output[i] = blocks[level]
	}
	return string(output)
}

type Option int
