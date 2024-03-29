package ex08

import (
	"strings"
	"testing"

	"github.com/yosuke-furukawa/programming-go-study/ch05/ex07"
	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	tests := []struct {
		actual   string
		query    string
		expected string
	}{
		{
			`
<html>
  <head>
  </head>
  <body>
    <div id="foo">bar</div>
  </body>
</html>`,
			"foo",
			`<div id="foo">
  bar
</div>
`,
		},
		{
			`
<html>
  <head>
  </head>
  <body>
    <div id="foo">bar1</div>
    <div id="foo">bar2</div>
  </body>
</html>`,
			"foo",
			`<div id="foo">
  bar1
</div>
`,
		},
	}

	for _, test := range tests {
		reader := strings.NewReader(test.actual)
		in, err := html.Parse(reader)
		if err != nil {
			t.Fatalf("error is thrown %s", err)
		}
		node := ElementByID(in, test.query)
		out := ex07.PrettyHTML(node)
		if test.expected != out {
			t.Errorf("expected result %s, but actual is %s", test.expected, out)
		}
	}
}
