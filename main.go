package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	baseDir = "./protos"
)

func main() {
	GetEn(baseDir)
}
func Gen(dirEntries []os.DirEntry, path string) {
	for _, entry := range dirEntries {
		tmp := path + "/" + entry.Name()
		if !entry.IsDir() {
			GetFiles(tmp)
		} else {
			GetEn(tmp)
		}
	}
}

func GetEn(base string) {
	dirEntries, err := os.ReadDir(base)
	if err != nil {
		panic(err)
	}
	Gen(dirEntries, base)
}

func GetFiles(path string) {
	fmt.Println(path)
	cmd := exec.Command("./protoc.exe", " --proto_path=./proto ",
		"--validate_out=\"lang=go:./\" ",
		"--go_out=./",
		"  --doc_out=./doc",
		" --doc_opt=html,index.html", path)
	cmd.Run()
}
