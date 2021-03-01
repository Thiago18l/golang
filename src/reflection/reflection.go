package main

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	value := reflect.ValueOf(x)
	camp := value.Field(0)
	fn(camp.String())
}
