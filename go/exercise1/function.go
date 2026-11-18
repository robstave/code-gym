package util

type fn func(int)

func combine(f1 fn, f2 fn) fn {
	return func(value int) {
		f1(value)
		f2(value)
	}
}

func combineGenerics[A, B any](f1 func(A), f2 func(B)) func(A, B) {
	return func(first A, second B) {
		f1(first)
		f2(second)
	}
}
