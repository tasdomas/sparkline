package sparkline_test

import (
	"testing"

	"github.com/tasdomas/sparkline"
)

var sparklineTests = []struct {
	name     string
	values   any
	call     func(any, ...sparkline.Option) string
	expected string
	opts     []sparkline.Option
}{{
	name:   "empty slice",
	values: []uint32{},
	call: func(values any, _ ...sparkline.Option) string {
		return sparkline.Sparkline(values.([]uint32))
	},
	expected: "",
}, {
	name:   "simple test",
	values: []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0},
	call: func(values any, _ ...sparkline.Option) string {
		return sparkline.Sparkline(values.([]float32))
	},
	expected: "▁▂▃▄▅▆▇█",
}, {
	name:   "simple test, uint8",
	values: []uint8{1, 2, 3, 4, 5, 6, 7, 8},
	call: func(values any, _ ...sparkline.Option) string {
		return sparkline.Sparkline(values.([]uint8))
	},
	expected: "▁▂▃▄▅▆▇█",
}, {
	name:   "simple test, uint32",
	values: []uint32{0, 1, 2, 3, 4, 5, 6, 7},
	call: func(values any, _ ...sparkline.Option) string {
		return sparkline.Sparkline(values.([]uint32))
	},
	expected: "▁▂▃▄▅▆▇█",
}, {
	name:   "equal values",
	values: []float32{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0},
	call: func(values any, _ ...sparkline.Option) string {
		return sparkline.Sparkline(values.([]float32))
	},
	expected: "▁▁▁▁▁▁▁▁",
}, {
	name:   "range provided",
	values: []uint32{0, 1, 2, 3, 4, 5, 6, 7},
	call: func(values any, opts ...sparkline.Option) string {
		return sparkline.Sparkline(values.([]uint32), opts...)
	},
	expected: "▁▁▁▁▁▁▁▁",
	opts:     []sparkline.Option{sparkline.WithRange(uint32(0), uint32(100))},
}, {
	name:   "dots style",
	values: []uint32{0, 1, 2, 3, 4, 5, 6, 7},
	call: func(values any, opts ...sparkline.Option) string {
		return sparkline.Sparkline(values.([]uint32), opts...)
	},
	expected: " ⣀⣤⣾",
	opts:     []sparkline.Option{sparkline.DotsStyle},
}}

func TestSparkline(t *testing.T) {
	t.Parallel()
	for _, test := range sparklineTests {
		t.Run(test.name, func(t *testing.T) {
			result := test.call(test.values, test.opts...)
			if result != test.expected {
				t.Fatalf("expected sparkline to be %q, got %q", test.expected, result)
			}
		})
	}
}
