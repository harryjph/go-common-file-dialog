package main

import (
	cfd2 "github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"log"
)

func main() {
	results, err := cfdutil.ShowOpenMultipleFilesDialog(cfd2.DialogConfig{
		Title: "Open Text File",
		Role:  "OpenTextExample",
		FileFilters: []cfd2.FileFilter{
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
	log.Printf("Chosen file: %s\n", results)
}
