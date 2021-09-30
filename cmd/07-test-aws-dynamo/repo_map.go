package main

type PeopleRepositoryMap struct {
	PeopleMap map[string]Person
}

// GetPersonByName is a method for type PeopleRepositoryMap struct
func (repo PeopleRepositoryMap) GetPersonByName(name string) Person {
	person, find := repo.PeopleMap[name]
	if !find {
		person = Person{}
	}
	return person
}

// GetAllPeople is a method for type PeopleRepositoryMap struct
func (repo PeopleRepositoryMap) GetAllPeople() []Person {
	var people []Person
	for _, person := range repo.PeopleMap {
		people = append(people, person)
	}
	return people
}

func NewRepoMap() PeopleRepositoryMap {
	data := map[string]Person{
		"joe": {
			Name: "joe",
			Age: 18,
		},
		"john": {
			Name: "john",
			Age: 30,
		},
		"smith": {
			Name:"smith",
			Age: 32,
		},
	}
	return PeopleRepositoryMap{
		PeopleMap: data,
	}
}