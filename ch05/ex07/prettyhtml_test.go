package ex07

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettyHTML(t *testing.T) {
	tests := []struct {
		actual string
		expect string
	}{
		{
			"<html><div>hoo</div></html>",
			`<html>
  <head />
  <body>
    <div>
      hoo
    </div>
  </body>
</html>
`,
		},
		{
			"<html><div>hoo<img src=\"foo.jpg\" /></div></html>",
			`<html>
  <head />
  <body>
    <div>
      hoo
      <img src="foo.jpg" />
    </div>
  </body>
</html>
`,
		},
		{
			"<html><div>hoo<img src=\"foo.jpg\" /><a href=\"example.com/foo\">nya-n</a></div></html>",
			`<html>
  <head />
  <body>
    <div>
      hoo
      <img src="foo.jpg" />
      <a href="example.com/foo">
        nya-n
      </a>
    </div>
  </body>
</html>
`,
		},
	}

	for _, test := range tests {
		reader := strings.NewReader(test.actual)
		in, err := html.Parse(reader)
		if err != nil {
			t.Fatalf("error is thrown %s", err)
		}
		out := PrettyHTML(in)
		t.Log(out)
		if out != test.expect {
			t.Errorf("expect is %s but actual is %s", test.expect, out)
		}
	}
}
