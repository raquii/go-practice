package duck

import "fmt"

type DuckTypes interface {
	ActLikeDuck()
}

type duck struct {
	name   string
	quack  quack
	flight fly
}

func (d *duck) learnToFly(f fly) {
	d.flight = f
}

func (d *duck) learnToQuack(q quack) {
	d.quack = q
}

func (d *duck) ActLikeDuck() {
	fmt.Println("---------🦆---------")
	fmt.Printf("%s ", d.name)
	d.flight.useFly()
	fmt.Printf("%s ", d.name)
	d.quack.useQuack()
	fmt.Println("--------------------")
}

func NewDuck(name string, fly bool, quack bool) DuckTypes {
	d := duck{name: name}

	if fly {
		f := newFlyWithWings()
		d.learnToFly(f)
	} else {
		f := newCantFly()
		d.learnToFly(f)
	}

	if quack {
		q := newDuckCall()
		d.learnToQuack(q)
	} else {
		q := newSilentCall()
		d.learnToQuack(q)
	}

	return &d
}
