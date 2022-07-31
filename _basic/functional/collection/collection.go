package collection

import "github.com/google/go-cmp/cmp"

func contains[T any](l []T, t ...T) bool {
	for _, tv := range t {
		var found bool
		for _, lv := range l {
			if cmp.Equal(tv, lv) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func filter[T any](l []T, f func(t T) bool) []T {
	var result List[T]
	for _, v := range l {
		if !f(v) {
			continue
		}
		result = append(result, v)
	}
	return result
}

func map1[T, O any](l []T, f func(t T) O) []O {
	var result = make(List[O], len(l))
	for i, v := range l {
		result[i] = f(v)
	}
	return result
}

func foldLeft[T any](l []T, zero T, f func(t1, t2 T) T) T {
	var tmp = zero
	for _, v := range l {
		tmp = f(tmp, v)
	}
	return tmp
}

func foldRight[T any](l []T, zero T, f func(t1, t2 T) T) T {
	return foldLeft(reverse(l), zero, f)
}

func reverse[T any](l []T) []T {
	var tmp = l
	for i := 0; i < len(l)/2; i++ {
		tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
	}
	return tmp
}

func headTail[T any](l []T) (*T, []T, bool) {
	if len(l) == 0 {
		return nil, nil, false
	}
	return &l[0], l[1:], true
}
