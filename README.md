# OLX talk - #3 Introduction to Golang Interfaces and Dependency Inversion Principle

Code examples are inside `cmd` folder.

Run `go mod vendor` to install dependencies.

Some unit test for the examples `cmd/06-test-gomock` and `cmd/07-test-aws-dynamo` are using [gomock](https://github.com/golang/mock), read the 
docs to know how to 
install.
Commands to generate the `gomock` files for the code examples:
```
mockgen -source=cmd/06-test-gomock/repo_database.go -destination=repo_mock.go -package=main
```
```
 mockgen -source=./vendor/github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface/interface.go -destination=dynamodb_mock.go -package=main
```
You need to move the generated go mock files to the correct folder, alternatively you can also change the `-destination` 
flag to the correct folder

Examples:

* `cmd/01-no-interface`: code with a database repository without interface
  * run: `go run cmd/01-no-interface/*.go`

* `cmd/02-interface`: code with a database repository with interface
  * run: `go run cmd/02-interface/*.go` 
  
* `cmd/03-database-map`: code with a database repository and map repository
    * run: `go run cmd/03-database-map/*.go` 

* `cmd/04-test-database/`: code with an handler endpoint that uses the database repository, creates a server on port 8080. It has unit tests for the handler using the database repository.
  * run: `go run cmd/04-test-database/main.go cmd/04-test-database/repo_database.go`
  * http request example: `127.0.0.1:8080/person?name=joe`
  * run unit test: `go test ./cmd/04-test-database/`
  
* `cmd/05-test-map`: handler unit test with the map repository
  * run unit test: `go test ./cmd/05-test-map`

* `cmd/06-test-gomock`: handler unit test using `gomock`
  * run unit test: `go test ./cmd/05-test-map`

* `cmd/07-test-aws-dynamo`: handler now uses an aws dynamo table, using `gomock` for unit testing
  * run unit test: `go test ./cmd/07-test-aws-dynamo`