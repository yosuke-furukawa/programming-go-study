package fetch

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"golang.org/x/net/html"
)

func Fetch(urlPath, savePath string) (filename string, n int64, err error) {
	log.Println(savePath)
	os.MkdirAll(savePath, 0777)
	resp, err := http.Get(urlPath)

	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	u, _ := url.Parse(urlPath)
	saveFilePath := path.Join(savePath, u.Path)
	base := path.Join(savePath, path.Base(urlPath))
	os.MkdirAll(base, 0777)
	if path.Ext(u.Path) == "" {
		os.MkdirAll(path.Join(savePath, u.Path), 0777)
		saveFilePath = path.Join(savePath, u.Path, "index.html")
	}
	log.Println("saveFile " + saveFilePath)

	f, err := os.Create(saveFilePath)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	reader := bufio.NewReader(resp.Body)
	in, err := html.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}
	ReplaceTag(in, u.Hostname(), savePath)
	writer := bufio.NewWriter(f)
	html.Render(writer, in)

	return saveFilePath, n, err

}
