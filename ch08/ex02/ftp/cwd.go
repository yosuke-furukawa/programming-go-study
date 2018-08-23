package ftp

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Cwd struct {
	dir string
}

func NewCwd() *Cwd {
	return &Cwd{
		os.Getenv("HOME"),
	}
}

func (cwd *Cwd) Pwd() string {
	log.Println(cwd.dir)
	return cwd.dir
}

func (cwd *Cwd) Chdir(path string) (string, error) {
	if !filepath.IsAbs(path) {
		cwd.dir = filepath.Join(cwd.dir, path)
		return cwd.dir, nil
	}
	cwd.dir = path
	return cwd.dir, nil
}

func (cwd *Cwd) Stat(path string) (os.FileInfo, error) {
	if filepath.IsAbs(path) {
		return os.Stat(path)
	}
	return os.Stat(filepath.Join(cwd.dir, path))
}

func (cwd *Cwd) Get(path string) ([]byte, error) {
	if filepath.IsAbs(path) {
		return ioutil.ReadFile(path)
	}
	return ioutil.ReadFile(filepath.Join(cwd.dir, path))
}

func (cwd *Cwd) Put(path string, content []byte) error {
	return ioutil.WriteFile(filepath.Join(cwd.dir, path), content, 0777)
}
