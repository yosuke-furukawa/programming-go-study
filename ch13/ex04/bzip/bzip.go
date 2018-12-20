package bzip

import (
	"bytes"
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	w     io.Writer
	cmd   *exec.Cmd
	mutex sync.Mutex
}

func NewWriter(out io.Writer) io.WriteCloser {
	w := &writer{
		w:   out,
		cmd: exec.Command("bzip2"),
	}
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.cmd.Stdin = bytes.NewReader(data)
	w.cmd.Stdout = &w.w
	if err := w.cmd.Run(); err != nil {
		return 0, nil
	}
	var total int

	for len(data) > 0 {
		n, err := w.w.Write(data)
		if err != nil {
			return total + n, err
		}
		total += n
		data = data[total:]
	}
	return total, nil
}

func (w *writer) Close() error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if err := w.cmd.Wait(); err != nil {
		return err
	}
	return nil
}
