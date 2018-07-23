package equal

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		x     interface{}
		y     interface{}
		equal bool
	}{
		{0, 0, true},
		{1, 2, false},
		{1, -1, false},
		{uint(1e9), int(1e9 + 1), true},
		{uint32(1e9), int32(1e9 - 1), false},
		{int32(-1e9), int64(-1e9 + 1), false},
		{int64(-1e9), float32(-1e9 - 1), true},
		{float32(1e-9), float64(1e-9 - 1e-18), false},
		{[]float64{1e-9}, []float64{1e-9 + 1e-18}, true},
	}

	for _, test := range tests {
		if Equal(test.x, test.y) != test.equal {
			t.Errorf("TestEqual inputs (%v (%[1]T) %v (%[2]T)), got %t", test.x, test.y, !test.equal)
		}
	}
}
