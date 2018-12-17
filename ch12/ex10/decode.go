package ex10

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

func Unmarshal(data []byte, out interface{}) (err error) {
	decoder := NewDecoder(bytes.NewReader(data))
	return decoder.Decode(out)
}

type Decoder struct {
	reader io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r}
}

func (d *Decoder) Decode(out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(d.reader)
	lex.next() // get the first token
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

const StartList = '('
const Symbol = '#'

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		} else if lex.text() == "t" {
			v.SetBool(true)
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		v.SetInt(int64(i))
		lex.next()
		return
	case StartList:
		lex.next()
		readList(lex, v)
		lex.next() // consume ')'
		return
	case scanner.Float:
		switch v.Kind() {
		case reflect.Float32:
			f, _ := strconv.ParseFloat(lex.text(), 32) // NOTE: ignoring erros
			v.SetFloat(f)
		case reflect.Float64:
			f, _ := strconv.ParseFloat(lex.text(), 64) // NOTE: ignoring erros
			v.SetFloat(f)
		default:
			panic(fmt.Sprintf("unexpected type: %d", v.Kind()))
		}
		lex.next()
		return

	case Symbol:
		lex.next() // Ident
		lex.next() // '('
		lex.next() // Float
		r := lex.text()
		lex.next() // Float
		i := lex.text()
		lex.next() // ')'
		lex.consume(')')

		var bitSize int
		switch v.Kind() {
		case reflect.Complex64:
			bitSize = 32
		case reflect.Complex128:
			bitSize = 64
		default:
			panic(fmt.Sprintf("unexpected type: %d", v.Kind()))
		}
		fr, _ := strconv.ParseFloat(r, bitSize)
		fi, _ := strconv.ParseFloat(i, bitSize)
		v.SetComplex(complex(fr, fi))
		return

	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	case reflect.Interface:
		t, _ := strconv.Unquote(lex.text())
		value := reflect.New(typeOf(t)).Elem()
		lex.next()
		read(lex, value)
		v.Set(value)
	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

func typeOf(tName string) reflect.Type {
	switch tName {
	case "int":
		var x int
		return reflect.TypeOf(x)
	case "[]int":
		var x []int
		return reflect.TypeOf(x)
	default:
		panic(fmt.Sprintf("%s not supported yet\n", tName))
	}
}
