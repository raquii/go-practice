package duck

import "fmt"

type duckCall struct{}

func newDuckCall() *duckCall {
	return &duckCall{}
}

func (d *duckCall) useQuack() {
	fmt.Println("is quacking like a duck.")
}
