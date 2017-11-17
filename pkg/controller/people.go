package controller

import (
	"fmt"

	"github.com/n3wscott/cmd-service-broker/pkg/apis/cmdbroker/v1"
)

type PeopleController struct {
	People []v1.Person
}

func (pc *PeopleController) GetPeople() ([]v1.Person, error) {
	return pc.People, nil
}

func (pc *PeopleController) GetPerson(id string) (*v1.Person, error) {
	for _, item := range pc.People {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("person not found")
}

func (pc *PeopleController) CreatePerson(id string, person *v1.CreatePersonRequest) (*v1.Person, error) {
	// TODO: validate request

	if _, err := pc.GetPerson(id); err == nil {
		return nil, fmt.Errorf("person found")
	}

	p := v1.Person{
		ID:        id,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Address:   person.Address,
	}
	pc.People = append(pc.People, p)
	return &p, nil
}

func (pc *PeopleController) DeletePerson(id string) error {
	for index, item := range pc.People {
		if item.ID == id {
			pc.People = append(pc.People[:index], pc.People[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("person not found")
}
