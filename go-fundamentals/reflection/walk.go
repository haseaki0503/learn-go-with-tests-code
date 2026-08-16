package reflection

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for _, field := range val.Fields() {
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
