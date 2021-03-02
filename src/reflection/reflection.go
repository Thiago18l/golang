package main

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	value := reflect.ValueOf(x)

	for i := 0; i < value.NumField(); i++ {
		camp := value.Field(i)
		switch camp.Kind() {
		case reflect.String:
			fn(camp.String())
		case reflect.Struct:
			percorre(camp.Interface(), fn)
		}
	}
}
