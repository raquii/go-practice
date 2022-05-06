package processor

type echo struct {
	num int
}

func newEcho() *echo {
	return &echo{
		num: 3,
	}
}

// echoes a string
func (e *echo) manipulate(s string) (result string) {
	x := e.num
	result = s
	for i := 1; i < x; i++ {
		result = result + " " + s
	}
	return
}
