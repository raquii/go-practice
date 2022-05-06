package processor

type reverse struct{}

func newReverse() *reverse {
	return &reverse{}
}

// reverses a string
func (r *reverse) manipulate(s string) (result string) {
	for _, c := range s {
		result = string(c) + result
	}
	return
}
