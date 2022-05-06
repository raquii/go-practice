package processor

type Processor interface {
	UseProcessors(string) map[string]string
}

type processor struct {
	manipulations map[string]stringProcessor
}

func (p processor) UseProcessors(s string) map[string]string {
	ms := make(map[string]string)
	for key, manipulator := range p.manipulations {
		ms[key] = manipulator.manipulate(s)
	}
	return ms
}

func NewProcessor() Processor {
	m := generateManipulations()

	return &processor{
		manipulations: m,
	}
}

func generateManipulations() map[string]stringProcessor {
	e := newEcho()
	r := newReverse()
	v := newVowelRemove()
	m := newMirror()

	return map[string]stringProcessor{
		"echo":        e,
		"reverse":     r,
		"vowelremove": v,
		"mirror":      m,
	}
}
