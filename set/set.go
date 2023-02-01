package set

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](keys ...T) Set[T] {
	s := Set[T]{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func NewSetK[T comparable](k T, keys ...T) Set[T] {
	s := Set[T]{k: struct{}{}}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Set[T]) Exist(key T) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func (s Set[T]) Add(key T) {
	s[key] = struct{}{}
}

func (s Set[T]) AddAll(keys ...T) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Set[T]) Remove(key T) {
	delete(s, key)
}

func (s Set[T]) Contains(key T) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Set[T]) ToList() []T {
	var ret []T

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Set[T]) Intersect(s2 Set[T]) Set[T] {
	newSet := Set[T]{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	newSet := Set[T]{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Set[T]) Minus(s2 Set[T]) Set[T] {
	newSet := Set[T]{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Set[T]) MinusInPlace(s2 Set[T]) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Equal(s2 Set[T]) bool {
	if len(s) != len(s2) {
		return false
	}
	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}
	return true
}

func (s Set[T]) Size() int {
	return len(s)
}
