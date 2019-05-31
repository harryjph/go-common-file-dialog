package main

import (
	cfd "github.com/harry1453/go-common-file-dialog"
	"log"
)

func main() {
	cfd.Initialize()
	defer cfd.UnInitialize()
	pickFolderDialog, err := cfd.NewOpenFileDialog(cfd.DialogConfig{
		Title: "Pick Folder",
		Role:  "PickFolderExample",
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := pickFolderDialog.Show(); err != nil {
		log.Fatal(err)
	}
	result, err := pickFolderDialog.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen folder: %s\n", result)
}
