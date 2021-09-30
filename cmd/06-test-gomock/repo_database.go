package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library

)

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
	for rows.Next() {
		person := Person{}
		rows.Scan(&person.Name, &person.Age)
		people = append(people, person)
	}
	return people
}

func NewRepoDatabase() PeopleRepositoryDatabase {
	sqliteDatabase, _ := sql.Open("sqlite3", "./database.sqlite")
	return PeopleRepositoryDatabase{
		db: sqliteDatabase,
	}
}