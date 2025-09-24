// Package sparkline handles creation of string sparklines.
package sparkline

import "math"

// Style is the general definition for a sparkline rendering
// style.
type Style interface {
	Render([]uint8) string
	Levels() int
}

type blocks struct{}

func (b blocks) Levels() int { return 8 }

func (b blocks) Render(quantizedData []uint8) string {
	output := make([]rune, len(quantizedData))
	for i, q := range quantizedData {
		output[i] = blockRunes[q]
	}
	return string(output)
}

type dots struct{}

// Levels returns the number of levels in this sparkline style.
func (d dots) Levels() int { return 5 }

// Render renders the provided quantized data datapoints
// using braille characters.
func (d dots) Render(quantizedData []uint8) string {
	data := quantizedData
	if len(quantizedData)%2 != 0 {
		data = append(quantizedData, 0)
	}
	output := make([]rune, len(data)/2)

	for i := 0; i < len(data); i += 2 {
		a, b := data[i], data[i+1]
		output[i/2] = dotRunes[a][b]
	}

	return string(output)
}

var blockRunes = []rune("▁▂▃▄▅▆▇█")

var dotRunes = [][]rune{
	{' ', '⢀', '⢠', '⢰', '⢸'},
	{'⡀', '⣀', '⣠', '⣰', '⣸'},
	{'⡄', '⣄', '⣤', '⣴', '⣼'},
	{'⡆', '⣆', '⣦', '⣶', '⣾'},
	{'⡇', '⣇', '⣧', '⣷', '⣿'},
}

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type sparklineOptions struct {
	style    Style
	max      any
	min      any
	hasRange bool
}

func defaultSparklineOptions() sparklineOptions {
	return sparklineOptions{
		style: blocks{},
	}
}

// Sparkline generates a sparkline for the provided values.
func Sparkline[T Number](values []T, opts ...Option) string {
	if len(values) == 0 {
		return ""
	}

	options := defaultSparklineOptions()
	for _, opt := range opts {
		opt(&options)
	}

	if !options.hasRange {
		options.min, options.max = minMax[T](values)
	}

	span := float64(options.max.(T) - options.min.(T))
	levelCount := float64(options.style.Levels() - 1)
	quantizedData := make([]uint8, len(values))
	for i, value := range values {
		quantizedData[i] = uint8(math.Floor(levelCount * float64(value-options.min.(T)) / span))
	}
	return options.style.Render(quantizedData)
}

func minMax[T Number](values []T) (T, T) {
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

	return minValue, maxValue
}

// Option defines a function for passing options to the sparkline command.
type Option func(*sparklineOptions)

// DotsStyle specifies that the sparkline should be rendered using braille
// unicode characters.
func DotsStyle(opts *sparklineOptions) {
	opts.style = dots{}
}

// BlocksStyle specifieses that the sparkline should be rendered using blocks
// unicode characters.
func BlocksStyle(opts *sparklineOptions) {
	opts.style = blocks{}
}

// WithRange specifies a range (min to max) that
// the datapoints can have.
func WithRange(min, max any) Option {
	return func(opts *sparklineOptions) {
		opts.min = min
		opts.max = max
		opts.hasRange = true
	}
}
