package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for _, field := range val.Fields() {
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}

