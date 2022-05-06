package processor

type mirror struct {
	forward bool
}

func newMirror() *mirror {
	return &mirror{
		forward: true,
	}
}

// mirrors a string
func (m *mirror) manipulate(s string) (result string) {
	if m.forward {
		for _, c := range s {
			result = string(c) + result
		}
		result = s + result
	} else {
		result = s
		for _, c := range s {
			result = string(c) + result
		}
	}
	return
}
