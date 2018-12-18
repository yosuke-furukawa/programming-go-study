package ex11

import (
	"fmt"
	"reflect"
	"strings"
)

func Pack(ptr interface{}) string {
	v := reflect.ValueOf(ptr).Elem()

	parts := []string{}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		parts = append(parts, fragment(name, v.Field(i)))
	}

	return strings.Join(parts, "&")
}

func fragment(name string, v reflect.Value) string {
	result := ""
	switch v.Kind() {
	case reflect.String:
		result = fmt.Sprintf("%s=%s", name, v.String())
	case reflect.Int:
		result = fmt.Sprintf("%s=%d", name, v.Int())
	case reflect.Bool:
		if v.Bool() {
			result = fmt.Sprintf("%s=true", name)
		} else {
			result = fmt.Sprintf("%s=false", name)
		}
	case reflect.Array, reflect.Slice:
		var parts []string
		for i := 0; i < v.Len(); i++ {
			parts = append(parts, fragment(name, v.Index(i)))
		}
		result = strings.Join(parts, "&")
	}
	return result
}
