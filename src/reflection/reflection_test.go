package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name   string
	Perfil Perfil
}

type Perfil struct {
	Age  int
	City string
}

func TestPercorre(t *testing.T) {
	cases := []struct {
		Name        string
		Input       interface{}
		waitedCalls []string
	}{
		{
			"Struct with value string",
			struct {
				Name string
			}{"Thiago"},
			[]string{"Thiago"},
		},
		{
			"Struct with two string values",
			struct {
				Name string
				City string
			}{"Thiago", "London"},
			[]string{"Thiago", "London"},
		},
		{
			"Struct without string value",
			struct {
				Name string
				Age  int
			}{"Thiago", 23},
			[]string{"Thiago"},
		},
		{
			"Struct with other nested struct inside",
			Person{
				"Thiago",
				Perfil{33, "London"},
			},
			[]string{"Thiago", "London"},
		},
		{
			"Pointers for things",
			&Person{
				"Thiago",
				Perfil{23, "London"},
			},
			[]string{"Thiago", "London"},
		},
		{
			"Slices",
			[]Perfil{
				{33, "London"},
				{23, "Bangladesh"},
			},
			[]string{"London", "Bangladesh"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var result []string
			percorre(test.Input, func(entrada string) {
				result = append(result, entrada)
			})
			if !reflect.DeepEqual(result, test.waitedCalls) {
				t.Errorf("result '%s', waited '%s'", result, test.waitedCalls)
			}
		})
	}

}
