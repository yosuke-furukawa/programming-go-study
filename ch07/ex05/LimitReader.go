package ex05

import (
	"io"
)

type Limit struct {
	r io.Reader
	n int64
}

func (l *Limit) Read(p []byte) (int, error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.n {
		p = p[:l.n]
	}
	num, err := l.r.Read(p)
	l.n -= int64(num)
	return num, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &Limit{
		r: r,
		n: n,
	}
}
