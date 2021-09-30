package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerMap(t *testing.T){
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
		{
			"/person?name=anna",
			404,
		},

	}
	handler := HandlerStruct{
		repo: NewRepo("database"),
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