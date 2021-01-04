package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	Key     string
	User    *User
	IsError bool
}

func TestGetUser(t *testing.T) {
	cases := []TestCase{
		TestCase{"ok", &User{ID: 28}, false},
		TestCase{"fail", nil, true},
		TestCase{"not_exist", nil, true},
	}

	for caseNum, item := range cases {
		u, err := GetUser(item.Key)

		if item.IsError && err == nil {
			t.Errorf("[%d] expected error, got nil", caseNum)
		}
		if !item.IsError && err != nil {
			t.Errorf("[%d] unexpected error %v", caseNum, err)
		}
		if !reflect.DeepEqual(u, item.User) {
			t.Errorf("[%d] wrong result: got %+v, expected %+v", caseNum, u, item.User)
		}
	}
}

/*
go test
	-v 		показывает запущенный тест и его результат
	-cover  показывает степень покрытия кода тестами

go test -coverprofile=cover.out 				экспорт профилированных результатов в файл
go tool cover -html=cover.out -o cover.html 	конвертация в html
*/
