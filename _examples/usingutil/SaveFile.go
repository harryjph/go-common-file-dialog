package main

import (
	cfd "github.com/harry1453/go-common-file-dialog"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"log"
)

func main() {
	result, err := cfdutil.ShowSaveFileDialog(cfd.DialogConfig{
		Title: "Save Text File",
		Role:  "SaveTextExample",
		FileFilter: []cfd.FileFilter{
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen file: %s\n", result)
}
