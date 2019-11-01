package main

import (
	"github.com/harry1453/go-common-file-dialog/cfd"
	"log"
)

func main() {
	openMultiDialog, err := cfd.NewOpenMultipleFileDialog(cfd.DialogConfig{
		Title: "Open Text File",
		Role:  "OpenTextExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "Image Files (*.jpg, *.png)",
				Pattern:     "*.jpg;*.png",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := openMultiDialog.Show(); err != nil {
		log.Fatal(err)
	}
	results, err := openMultiDialog.GetResults()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Chosen files", results)
}
