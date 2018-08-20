package fetch

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func Fetch(url, savePath string) (filename string, n int64, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)

	if strings.LastIndex(local, "/") >= len(local) {
		os.MkdirAll(savePath+local, 0777)
		local += "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err

}
