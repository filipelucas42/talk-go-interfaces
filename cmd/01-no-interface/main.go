package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

// Person type
type Person struct {
	Name string
	Age int
}

// PeopleRepositoryDatabase depends on a external database to store people
type PeopleRepositoryDatabase struct {
	db *sql.DB
}

// GetPersonByName is a method for type PeopleRepositoryDatabase struct
func (repo PeopleRepositoryDatabase) GetPersonByName(name string) Person {
	row := repo.db.QueryRow("select name, age from people where name = ?", name)
	person := Person{}
	row.Scan(&person.Name, &person.Age)
	return person
}

// GetAllPeople is a method for type PeopleRepositoryDatabase struct
func (repo PeopleRepositoryDatabase) GetAllPeople() []Person {
	rows, _ := repo.db.Query("select name, age from people")
	var people []Person
	for rows.Next(){
		person := Person{}
		rows.Scan(&person.Name, &person.Age)
		people = append(people, person)
	}
	return people
}


func NewRepo() PeopleRepositoryDatabase {
	sqliteDatabase, _ := sql.Open("sqlite3", "./database.sqlite")
	return PeopleRepositoryDatabase{
		db: sqliteDatabase,
	}
}

func main(){
	var repo PeopleRepositoryDatabase

	repo = NewRepo()

	person := repo.GetPersonByName("joe")

	fmt.Println(person.Name, person.Age)

	fmt.Println(repo.GetAllPeople())
}