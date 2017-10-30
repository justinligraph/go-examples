package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

const (
	E_Args = 1
)

func usage() {
	fmt.Printf("Usage for %s\n", os.Args[0])
	fmt.Printf("%s <src_path> <dst_path>", os.Args[0])
}

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(E_Args)
	}

	srcDir := os.Args[1]
	dstDir := os.Args[2]

	srcFiles, err := ioutil.ReadDir(srcDir)
	if err != nil {
		panic(err)
	}
	dstFiles, err := ioutil.ReadDir(dstDir)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Source has %d files while dest has %d files\n", len(srcFiles), len(dstFiles))

	srcMap := make(map[string]os.FileInfo)
	for _, f := range srcFiles {
		srcMap[f.Name()] = f
	}
	dstMap := make(map[string]os.FileInfo)
	for _, f := range dstFiles {
		dstMap[f.Name()] = f
	}

	for key, _ := range srcMap {
		if _, ok := dstMap[key]; !ok {
			copyFileContents(path.Join(srcDir, key), path.Join(dstDir, key))
		}
	}
}

func copyFileContents(src, dst string) (err error) {
	fmt.Println("Copying ", src, " to ", dst)
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
