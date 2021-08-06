package main

import "reflect"

func Walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
        for i:=0; i< val.Len(); i++ {
            Walk(val.Index(i).Interface(), fn)
        }
        return
    }

    if val.Kind() == reflect.Map {
        for _, key := range val.MapKeys() {
            Walk(val.MapIndex(key).Interface(), fn)
        }
        return
    }

    if val.Kind() == reflect.String {
        fn(val.String())
        return 
    }

    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)

        switch field.Kind() {
        case reflect.String:
            fn(field.String())
        case reflect.Struct:
            Walk(field.Interface(), fn)
        }
    }
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}