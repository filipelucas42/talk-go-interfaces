package main

import (
	"fmt"
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
	{
		"/person?name=anna",
		404,
	},
	}
	ctrl := gomock.NewController(t)
	repoMock := NewMockPeopleRepositoryInterface(ctrl)

	repoMock.EXPECT().
		GetPersonByName("joe").
		Return(Person{Name: "joe", Age: 18}).AnyTimes()

	repoMock.EXPECT().
		GetPersonByName("smith").
		Return(Person{}).AnyTimes()

	repoMock.EXPECT().GetPersonByName(gomock.Any()).Return(Person{}).AnyTimes()

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