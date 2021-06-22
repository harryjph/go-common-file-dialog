package main

import (
	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"log"
)

func main() {
	result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title: "Open A File",
		Role:  "OpenFileExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "Image Files (*.jpg, *.png)",
				Pattern:     "*.jpg;*.png",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
		SelectedFileFilterIndex: 2,
		FileName:                "file.txt",
		DefaultExtension:        "txt",
	})
	if err == cfd.ErrorCancelled {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen file: %s\n", result)
}
