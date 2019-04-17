package TempClass

import "fmt"

type PeopleBase struct {
}

func (this *PeopleBase) PrintName() {
	fmt.Println(123)
}

func (this *PeopleBase) TT() {
	this.PrintName()
}

func NewPeopleBase() *PeopleBase {
	return &PeopleBase{}
}
