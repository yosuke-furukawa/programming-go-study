package ex08

import (
	"bytes"
	"fmt"
	"reflect"
)

func MarshalJSON(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encodeJson(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encodeJson(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encodeJson(buf, v.Elem())

	case reflect.Array, reflect.Slice:
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encodeJson(buf, v.Index(i)); err != nil {
				return err
			}
			if i < v.Len()-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(']')

	case reflect.Bool:
		if v.Bool() == true {
			fmt.Fprintf(buf, "true")
		} else {
			fmt.Fprintf(buf, "false")
		}

	case reflect.Struct:
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "%q: ", v.Type().Field(i).Name)
			if err := encodeJson(buf, v.Field(i)); err != nil {
				return err
			}

			if i < v.NumField()-1 {
				buf.WriteByte(',')
				buf.WriteByte(' ')
			}
		}
		buf.WriteByte('}')

	case reflect.Map:
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encodeJson(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			buf.WriteByte(' ')
			if err := encodeJson(buf, v.MapIndex(key)); err != nil {
				return err
			}
			if i < len(v.MapKeys())-1 {
				buf.WriteByte(',')
				buf.WriteByte(' ')
			}
		}
		buf.WriteByte('}')

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
