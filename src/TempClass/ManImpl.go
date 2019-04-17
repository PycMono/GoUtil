package TempClass

import "fmt"

type ManImpl struct {
	*PeopleBase
}

func (this *ManImpl) PrintName() {
	fmt.Println(444)
}

func NewManImpl() *ManImpl {
	return &ManImpl{
		PeopleBase: NewPeopleBase(),
	}
}
