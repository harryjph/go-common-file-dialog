package main

import (
	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"log"
)

func main() {
	result, err := cfdutil.ShowPickFolderDialog(cfd.DialogConfig{
		Title:  "Pick Folder",
		Role:   "PickFolderExample",
		Folder: "C:\\",
	})
	if err == cfd.ErrorCancelled {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen folder: %s\n", result)
}
