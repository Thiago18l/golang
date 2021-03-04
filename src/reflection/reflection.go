package main

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	value := getValue(x)

	runValue := func(v reflect.Value) {
		percorre(v.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			runValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			runValue(value.Index(i))
		}
	case reflect.String:
		fn(value.String())
	case reflect.Map:
		for _, key := range value.MapKeys() {
			runValue(value.MapIndex(key))
		}
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}
