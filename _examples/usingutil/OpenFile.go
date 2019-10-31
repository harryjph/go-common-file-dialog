package main

import (
	cfd "github.com/harry1453/go-common-file-dialog"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"log"
)

func main() {
	result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
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
	log.Printf("Chosen file: %s\n", result)
}
