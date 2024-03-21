package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {

	for i := range s {
		initial = fn(initial, s[i])
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for _, val := range s.Reverse() {
		initial = fn(val, initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	a := IntList{}
	for i := range s {
		if fn(s[i]) {
			a = a.Append(IntList{s[i]})
		}
	}
	return a
}

func (s IntList) Length() int {
	n := 0
	for range s {
		n++
	}
	return n
}

func (s IntList) Map(fn func(int) int) IntList {
	n := make(IntList, s.Length())
	for i := range s {
		n[i] = fn(s[i])
	}
	return n
}

func (s IntList) Reverse() IntList {
	n := make([]int, s.Length())
	for i := 0; i < s.Length(); i++ {
		n[i] = s[s.Length()-1-i]
	}
	return n
}

func (s IntList) Append(lst IntList) IntList {
	n := make([]int, s.Length()+lst.Length())
	for i := range s {
		n[i] = s[i]
	}
	for j := range lst {
		n[j+s.Length()] = lst[j]
	}
	return n
}

func (s IntList) Concat(lists []IntList) IntList {
	l := IntList{}.Append(s)
	for i := range lists {
		l = l.Append(lists[i])
	}
	return l
}
