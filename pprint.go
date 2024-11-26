package pprint

import (
	"fmt"
	json "github.com/goccy/go-json"
	"reflect"
)

func PrettyFormat(v interface{}) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return fmt.Sprintf("Error: %v\n", err)
		}
		return string(b)
	default:
		return fmt.Sprintf("%+v\n", v)
	}
}

func PrettyPrint(v interface{}) {
	fmt.Print(PrettyFormat(v))
}
