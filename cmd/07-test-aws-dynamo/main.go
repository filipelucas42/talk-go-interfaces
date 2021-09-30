package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"net/http"
)

type Person struct {
	Name string
	Age int
}

type myDynamo interface {
	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)

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
	repo dynamodbiface.DynamoDBAPI
}

func (handler HandlerStruct) GetPersonByName(w http.ResponseWriter, r *http.Request) {
	personName := r.URL.Query().Get("name")
	item := make(map[string]*dynamodb.AttributeValue)
	item["name"] = &dynamodb.AttributeValue{
		S: aws.String(personName),
	}
	//pk, _ := dynamodbattribute.MarshalMap(item)

	dynamoItem, _ := handler.repo.GetItem(&dynamodb.GetItemInput{
		Key: item,
	})
	var person Person
	dynamodbattribute.UnmarshalMap(dynamoItem.Item, &person)
	fmt.Println("Inside handler", person)
	if person.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not found")
		return
	}

	json.NewEncoder(w).Encode(person)
}

func main(){

}
