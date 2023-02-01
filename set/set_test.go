package set

import "testing"

func TestNewSet(t *testing.T) {
	s := NewSet(1, 2, 3)
	s2 := NewSet(1, 2, 3, 4)
	s.Add(4)
	t.Log(s.ToList())
	t.Log(s.Exist(5))
	t.Log(s.Exist(3))
	t.Log(s.Union(s2).ToList())
	t.Log(s.Minus(s2).ToList())
	t.Log(s.Size())
}
