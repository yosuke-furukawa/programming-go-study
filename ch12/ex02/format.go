package ex01

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

var out io.Writer = os.Stdout
var maxCount = 100

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		fieldPath := ""
		for i := 0; i < v.NumField(); i++ {
			fieldPath += fmt.Sprintf("%s: %s, ", v.Type().Field(i).Name, formatAtom(v.Field(i)))
		}
		return fieldPath

	case reflect.Array:
		result := ""
		for i := 0; i < v.Len(); i++ {
			result += fmt.Sprintf("%s[%d] = %s ", "array", i, v.Index(i))
		}
		return result
	default:
		return v.Type().String() + " value"
	}
}

func Display(name string, x interface{}) {
	fmt.Fprintf(out, "Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

var count = 0

func display(path string, v reflect.Value) {
	count++
	if count > maxCount {
		fmt.Fprintf(out, "count is out of max %d, maxCount is $d", count, maxCount)
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprintf(out, "%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Fprintf(out, "%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprintf(out, "%s = nil\n", path)
		} else {
			fmt.Fprintf(out, "%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Fprintf(out, "%s = %s\n", path, formatAtom(v))
	}
}
