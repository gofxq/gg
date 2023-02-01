package math

type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 //| ~string
}

func Min[T Number](v1, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}

func Max[T Number](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMax[T Number](v ...T) T {
	if len(v) == 0 {
		var zero T
		return zero
	}
	max := v[0]
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMin[T Number](v ...T) T {
	if len(v) == 0 {
		var zero T
		return zero
	}
	min := v[0]
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvg[T Number](v ...T) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

// SliceSum careful: sum might overflow
func SliceSum[T Number](v ...T) T {
	var sum T
	for _, i := range v {
		sum += i
	}
	return sum
}
