package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type Person struct {
	Name string
	Age  int
}

type PeopleRepository struct {
	db *sql.DB
}

func (repo PeopleRepository) GetPersonByName(name string) Person {
	row := repo.db.QueryRow("select name, age from people where name = ?", name)
	person := Person{}
	row.Scan(&person.Name, &person.Age)
	return person
}

// GetAllPeople is a method for type PeopleRepository struct
func (repo PeopleRepository) GetAllPeople() []Person {
	rows, _ := repo.db.Query("select name, age from people")
	var people []Person
	for rows.Next() {
		person := Person{}
		rows.Scan(&person.Name, &person.Age)
		people = append(people, person)
	}
	return people
}

func NewRepo() PeopleRepository {
	sqliteDatabase, _ := sql.Open("sqlite3", "./database.sqlite")
	return PeopleRepository{
		db: sqliteDatabase,
	}
}

type PeopleRepositoryInterface interface {
	GetPersonByName(name string) Person
	GetAllPeople() []Person
}

func main() {
	var repo PeopleRepositoryInterface

	repo = NewRepo()

	person := repo.GetPersonByName("joe")

	fmt.Println(person.Name, person.Age)

	fmt.Println(repo.GetAllPeople())
}
