package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"log"
	"net/http"
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

type HandlerStruct struct {
	repo PeopleRepositoryInterface
}

func (handler HandlerStruct) GetPersonByName(w http.ResponseWriter, r *http.Request) {
	personName := r.URL.Query().Get("name")
	person := handler.repo.GetPersonByName(personName)
	if person.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not found")
		return
	}

	json.NewEncoder(w).Encode(person)
}

func main(){
	handler := HandlerStruct{
		repo: NewRepo("database"),
	}

	http.HandleFunc("/person", handler.GetPersonByName)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Println(err)
	}
}
