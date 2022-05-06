package duck

import "fmt"

type flyWithWings struct{}

func newFlyWithWings() *flyWithWings {
	return &flyWithWings{}
}

func (f *flyWithWings) useFly() {
	fmt.Println("is flying like a duck.")
}
