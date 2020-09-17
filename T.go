package gollections

// T replace parameterized type waiting for golang 1.18 with generics
type T string

// Equal compare T with a string (make it simpler for unit tests)
func (t T) Equal(s string) bool {
	return string(t) == s
}
