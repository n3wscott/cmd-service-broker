package v1

type Controller interface {
	GetPeople() ([]Person, error)

	GetPerson(id string) (*Person, error)
	CreatePerson(id string, person *CreatePersonRequest) (*Person, error)
	DeletePerson(id string) error
}
