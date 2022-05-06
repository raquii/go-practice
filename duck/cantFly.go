package duck

import "fmt"

type cantFly struct{}

func newCantFly() *cantFly {
	return &cantFly{}
}

func (n *cantFly) useFly() {
	fmt.Println("doesn't seem to know how to fly.")
}
