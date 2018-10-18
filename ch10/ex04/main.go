package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type packageInfo struct {
	ImportPath string
	Deps       []string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) == 1 {
		return fmt.Errorf("too few arguments")
	}
	pkgs, err := execGoList(os.Args[1:]...)
	if err != nil {
		return err
	}
	allPackages, err := execGoList("...")
	if err != nil {
		return err
	}

loop:
	for _, pkg := range allPackages {
		for _, p := range pkgs {
			if !contains(pkg.Deps, p.ImportPath) {
				continue loop
			}
		}
		fmt.Println(pkg.ImportPath)
	}

	return nil
}

func execGoList(pkgs ...string) ([]packageInfo, error) {
	args := []string{"list", "-e", "-json"}
	args = append(args, pkgs...)
	cmd := exec.Command("go", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	go cmd.Wait()

	dec := json.NewDecoder(stdout)
	var ps []packageInfo
	for {
		p := packageInfo{}
		err := dec.Decode(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}

func contains(ss []string, s string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}
