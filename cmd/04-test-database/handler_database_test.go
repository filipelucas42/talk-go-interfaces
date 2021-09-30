package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerDatabase(t *testing.T){
	tests :=  []struct{
		url string
		expectedCode int
	}{
		{
			"/person?name=joe",
			200,
		},
		{
			"/person?name=anna",
			404,
		},
	}

	handler := HandlerStruct{
		repo: NewRepoDatabase(),
	}

	for _, testCase := range tests {
		request := httptest.NewRequest(http.MethodGet, testCase.url, nil)
		response := httptest.NewRecorder()
		handler.GetPersonByName(response, request)
		if response.Code != testCase.expectedCode {
			t.Log(fmt.Printf("url: %s\n", testCase.url))
			t.Error(fmt.Printf("expected code %d got %d", testCase.expectedCode, response.Code))
		}
	}
}