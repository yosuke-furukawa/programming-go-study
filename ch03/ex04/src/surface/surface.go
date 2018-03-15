package surface

import (
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/yosuke-furukawa/programming-go-study/ch03/ex04/src/types"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	topc          = "#ff0000"
	bottomc       = "#0000ff"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// Surface write svg file to out using query
func Surface(out io.Writer, query types.Query) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", query.Stroke, query.Fill, query.Width, query.Height)

	for i := 0; i < query.Cells; i++ {
		for j := 0; j < query.Cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			z := top(i, j)
			if z > 0.3 {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, query.Top)
			} else if z < 0.04 {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, query.Bottom)
			} else {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, query.Fill)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func top(i, j int) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	return f(x, y)
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, errors.New("Invalid value")
	}
	if math.IsNaN(z) {
		return 0, 0, errors.New("Invalid value")
	}
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	return math.Pow(2.0, math.Sin(y)) * math.Pow(2.0, math.Sin(x)) / 12
}
