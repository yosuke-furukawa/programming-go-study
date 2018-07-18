package ex02

import (
	"io"
)

type Counter struct {
	count  int64
	writer io.Writer
}

func (c *Counter) Write(p []byte) (int, error) {
	c.count += int64(len(p))
	return c.writer.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counter := &Counter{
		0,
		w,
	}

	return counter, &(counter.count)
}
