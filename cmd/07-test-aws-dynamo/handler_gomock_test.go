package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerWithGoMock(t *testing.T){
	tests :=  []struct{
		url string
		code int
	}{
		{
			"/person?name=joe",
			200,
		},
		{
			"/person?name=smith",
			404,
		},
	}
	ctrl := gomock.NewController(t)
	repoMock := NewMockDynamoDBAPI(ctrl)

	repoMock.EXPECT().
		GetItem(gomock.Any()).
		Return(&dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"Name": &dynamodb.AttributeValue{
				S: aws.String("anna"),
			},
		}}, nil).Times(1)

	repoMock.EXPECT().
		GetItem(gomock.Any()).
		Return(&dynamodb.GetItemOutput{}, nil).Times(1)


	//t.Log("Inside test", repoMock.GetPersonByName("joe"))
	handler := HandlerStruct{
		repo: repoMock,
	}

	for _, testCase := range tests {
		req := httptest.NewRequest(http.MethodGet, testCase.url, nil)
		w := httptest.NewRecorder()
		handler.GetPersonByName(w, req)
		if w.Code != testCase.code {
			t.Log(fmt.Printf("url: %s\n", testCase.url))
			t.Error(fmt.Printf("expected code %d got %d", testCase.code, w.Code))
		}
	}
}