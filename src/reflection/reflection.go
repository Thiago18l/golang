package main

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	value := reflect.ValueOf(x)

	for i := 0; i < value.NumField(); i++ {
		camp := value.Field(i)
		if camp.Kind() == reflect.String {
			fn(camp.String())
		}
	}
}
