package duck

import "fmt"

type duck struct {
	name   string
	quack  quack
	flight fly
}

func newDuck(name string) *duck {
	return &duck{
		name: name,
	}
}

func (d *duck) learnToFly(f fly) {
	d.flight = f
}

func (d *duck) learnToQuack(q quack) {
	d.quack = q
}

func (d *duck) actLikeDuck() {
	fmt.Printf("%s ", d.name)
	d.flight.useFly()
	fmt.Printf("%s ", d.name)
	d.quack.useQuack()
	fmt.Println("--------------------")
}

func DuckTime() {
	fmt.Println("🦆 It's time for ducks. 🦆")

	fly := newFlyWithWings()
	noFly := newCantFly()

	duckQuack := newDuckCall()
	silentQuack := newSilentCall()

	mallard := newDuck("Mallard Duck")
	mallard.learnToFly(fly)
	mallard.learnToQuack(duckQuack)
	mallard.actLikeDuck()

	pekin := newDuck("Pekin Duck")
	pekin.learnToFly(noFly)
	pekin.learnToQuack(duckQuack)
	pekin.actLikeDuck()

	rubber := newDuck("Rubber Duck")
	rubber.learnToFly(noFly)
	rubber.learnToQuack(silentQuack)
	rubber.actLikeDuck()
}
