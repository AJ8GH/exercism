package flatten

import (
	"reflect"
)

func Flatten(nested interface{}) []interface{} {
	out := make([]interface{}, 0)
	val := reflect.ValueOf(nested)

	switch {
	case val.IsNil():
		return out
	default:
		return getAll(val, out)
	}
}

func getAll(val reflect.Value, out []interface{}) []interface{} {
	for i := 0; i < val.Len(); i++ {
		v := reflect.ValueOf(val.Index(i).Interface())
		switch v.Kind() {
		case reflect.Array, reflect.Slice:
			out = getAll(v, out)
		case reflect.Int:
			out = append(out, int(v.Int()))
		}
	}
	return out
}
