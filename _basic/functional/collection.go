package functional

import "github.com/google/go-cmp/cmp"

func Contains[T any](l []T, t T) bool {
	for _, v := range l {
		if cmp.Equal(v, t) {
			return true
		}
	}
	return false
}

func Filter[T any](l []T, f func(t T) bool) []T {
	var result list[T]
	for _, v := range l {
		if !f(v) {
			continue
		}
		result = append(result, v)
	}
	return result
}

func Map[T, O any](l []T, f func(t T) O) []O {
	var result = make(list[O], len(l))
	for i, v := range l {
		result[i] = f(v)
	}
	return result
}

func FoldLeft[T any](l []T, zero T, f func(t1, t2 T) T) T {
	var tmp = zero
	for _, v := range l {
		tmp = f(tmp, v)
	}
	return tmp
}

func FoldRight[T any](l []T, zero T, f func(t1, t2 T) T) T {
	return FoldLeft(Reverse(l), zero, f)
}

func Reverse[T any](l []T) []T {
	var tmp = l
	for i := 0; i < len(l)/2; i++ {
		tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
	}
	return tmp
}
