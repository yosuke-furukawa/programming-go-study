package ftp

import (
	"io/ioutil"
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
	return cwd.dir
}

func (cwd *Cwd) Chdir(path string) (string, error) {
	err := os.Chdir(path)

	if err != nil {
		return "", err
	}
	if !filepath.IsAbs(path) {
		return filepath.Join(cwd.dir, path), nil
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
