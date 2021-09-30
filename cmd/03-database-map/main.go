package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type Person struct {
	Name string
	Age int
}

type PeopleRepositoryInterface interface {
	GetPersonByName(name string) Person
	GetAllPeople() []Person
}

func NewRepo(repoType string) PeopleRepositoryInterface {
	if repoType == "database" {
		return NewRepoDatabase()
	}
	return NewRepoMap()
}

func main(){
	var repo PeopleRepositoryInterface

	repo = NewRepo("map")

	person := repo.GetPersonByName("joe")

	fmt.Println(person.Name, person.Age)

	fmt.Println(repo.GetAllPeople())
}
