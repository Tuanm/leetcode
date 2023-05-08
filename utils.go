package leetcode

type Void struct{}

type Set[V comparable] struct {
	values map[V]Void
}

func (set *Set[V]) Add(v V) {
	set.values[v] = Void{}
}

func (set *Set[V]) Remove(v V) {
	delete(set.values, v)
}

func Abs[T int | float64](i T) T {
	if i >= 0 {
		return i
	}
	return -i
}
