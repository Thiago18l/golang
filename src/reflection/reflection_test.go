package main

import (
	"reflect"
	"testing"
)

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
