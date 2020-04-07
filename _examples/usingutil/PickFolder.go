package main

import (
	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"log"
)

func main() {
	result, err := cfdutil.ShowPickFolderDialog(cfd.DialogConfig{
		Title:         "Pick Folder",
		Role:          "PickFolderExample",
		InitialFolder: "C:\\",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen folder: %s\n", result)
}
