package main

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	value := getValue(x)

	countOfValues := 0
	var getCamp func(int) reflect.Value

	switch value.Kind() {
	case reflect.Struct:
		countOfValues = value.NumField()
		getCamp = value.Field
	case reflect.Slice:
		countOfValues = value.Len()
		getCamp = value.Index
	case reflect.String:
		fn(value.String())
	}

	for i := 0; i < countOfValues; i++ {
		percorre(getCamp(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}
