package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose prgress meessage")

type rootDir struct {
	root      string
	fileBytes int64
	fileNum   int64
}

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan rootDir)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	rootDirs := make(map[string]rootDir)
	for _, root := range roots {
		rootDirs[root] = rootDir{
			root,
			0,
			0,
		}
	}

loop:
	for {
		select {
		case file, ok := <-fileSizes:
			if !ok {
				break loop
			}
			rootDir := rootDirs[file.root]
			rootDir.fileBytes += file.fileBytes
			rootDir.fileNum++
			rootDirs[file.root] = rootDir
		case <-tick:
			printDiskUsage(rootDirs)
		}
	}
	printDiskUsage(rootDirs)
}

func printDiskUsage(rootDirs map[string]rootDir) {
	for root, _ := range rootDirs {
		fmt.Printf("%s: %d files %1.f Gb\n", root, rootDirs[root].fileNum, float64(rootDirs[root].fileBytes)/1e9)
	}
}

func walkDir(root, dir string, n *sync.WaitGroup, fileSizes chan<- rootDir) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, fileSizes)
		} else {
			fileSizes <- rootDir{
				root,
				entry.Size(),
				0,
			}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
