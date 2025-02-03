package flatten

import (
	"fmt"
	"reflect"
)

func Flatten(nested interface{}) (out []interface{}) {
	val := reflect.ValueOf(nested)

	switch {
	case val.IsNil():
		return []interface{}{}
	default:
		return getAllVals(val, out)
	}
}

func getAllVals(val reflect.Value, out []interface{}) []interface{} {
	if val.IsNil() {
		return out
	}

	switch val.Kind() {
	case reflect.Pointer:
		val = val.Elem()
		vk := val.Kind()
		fmt.Printf("%v", vk)
		fallthrough
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			f := val.Field(i)
			if f.IsNil() {
				continue
			}
			out = getAllVals(f, out)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			el := val.Index(i)
			k := el.Kind()
			fmt.Printf("%v", k)
			out = getAllVals(el, out)
		}
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		out = append(out, val.Int())
	case reflect.Uint16, reflect.Uint32, reflect.Uint8, reflect.Uint:
		out = append(out, val.Uint())
	default:
		v := val.Interface()
		vo := reflect.ValueOf(val)
		println(fmt.Sprintf("%v %v", v, vo))
		out = append(out, val.Interface())
	}
	return out
}
