package ex07

type Set struct {
	data map[int]struct{}
}

func Init(words []uint) Set {
	result := Set{
		data: make(map[int]struct{}),
	}
	for _, w := range words {
		result.data[int(w)] = struct{}{}
	}

	return result
}

func (s *Set) Has(x int) bool {
	_, ok := s.data[x]
	return ok
}

func (s *Set) Add(x int) {
	if _, ok := s.data[x]; !ok {
		s.data[x] = struct{}{}
	}
}

func (s *Set) UnionWith(t *Set) {
	for k, v := range t.data {
		s.data[k] = v
	}
}

func (s *Set) Len() int {
	return len(s.data)
}

func (s *Set) Remove(x int) {
	if !s.Has(x) {
		return
	}
	delete(s.data, x)
}

func (s *Set) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *Set) IntersectWith(t *Set) {
	for k, _ := range t.data {
		if !s.Has(k) {
			s.Remove(k)
		}
	}
	for k, _ := range s.data {
		if !t.Has(k) {
			s.Remove(k)
		}
	}
}

func (s *Set) DifferenceWith(t *Set) {
	for k, _ := range t.data {
		if !s.Has(k) {
			s.Remove(k)
		}
	}
}
