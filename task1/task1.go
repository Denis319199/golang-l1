package task1

import "fmt"

type EntityWithName interface {
	PrintName()
}

type Human struct {
	FirstName string
	LastName  string
}

func (human *Human) PrintName() {
	fmt.Println(human.FirstName, human.LastName)
}

func (human *Human) PrintReverseName() {
	fmt.Println(human.LastName, human.FirstName)
}

type Action struct {
	Name string
	Human
}

// PrintName - method shadowing (overriding)
func (action *Action) PrintName() {
	fmt.Println(action.Name)
}

func Task1() {
	action := Action{"pulling", Human{"James", "Bond"}}

	action.PrintName()        // Action -> PrintName()
	action.PrintReverseName() // Human -> PrintReverseName()
	action.Human.PrintName()  // explicit call: Human -> PrintName()

	var entityAction EntityWithName = &action
	entityAction.PrintName() // Action -> PrintName()

	var entityHuman EntityWithName = &action.Human
	entityHuman.PrintName() // Human -> PrintName()
}
