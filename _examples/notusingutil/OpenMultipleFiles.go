package main

import (
	"github.com/harry1453/go-common-file-dialog/cfd"
	"log"
)

func main() {
	openMultiDialog, err := cfd.NewOpenMultipleFilesDialog(cfd.DialogConfig{
		Title: "Open Multiple Files",
		Role:  "OpenFilesExample",
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
	if err != nil {
		log.Fatal(err)
	}
	if err := openMultiDialog.Show(); err != nil {
		log.Fatal(err)
	}
	results, err := openMultiDialog.GetResults()
	if err == cfd.ErrorCancelled {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen file(s): %s\n", results)
}
