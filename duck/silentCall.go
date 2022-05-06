package duck

import "fmt"

type silentCall struct{}

func newSilentCall() *silentCall {
	return &silentCall{}
}

func (s *silentCall) useQuack() {
	fmt.Println("is silent.")
}
