package math

import "testing"

var v1 = 1
var v2 = 3

func TestMax(t *testing.T) {
	if got := Max(v1, v2); got != 3 {
		t.Errorf("Max() = %v, want %v", got, 3)
	}
}

func TestMin(t *testing.T) {
	if got := Min(v1, v2); got != 1 {
		t.Errorf("Min() = %v, want %v", got, 1)
	}
}

var v = []int32{1, 2, 3, 4}

func TestSliceAvg(t *testing.T) {
	if got := SliceAvg(v...); got != 2.5 {
		t.Errorf("SliceAvg() = %v, want %v", got, 2.5)
	}
}

func TestSliceMax(t *testing.T) {
	if got := SliceMax(v...); got != 4 {
		t.Errorf("SliceMax() = %v, want %v", got, 4)
	}
}

func TestSliceMin(t *testing.T) {
	if got := SliceMin(v...); got != 1 {
		t.Errorf("SliceMin() = %v, want %v", got, 1)
	}
}

func TestSliceSum(t *testing.T) {
	if got := SliceSum(v...); got != 10 {
		t.Errorf("SliceSum() = %v, want %v", got, 10)
	}
}
