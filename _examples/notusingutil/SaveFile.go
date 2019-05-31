package main

import (
	cfd "github.com/harry1453/go-common-file-dialog"
	"log"
)

func main() {
	cfd.Initialize()
	defer cfd.UnInitialize()
	saveDialog, err := cfd.NewSaveFileDialog(cfd.DialogConfig{
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
	if err := saveDialog.Show(); err != nil {
		log.Fatal(err)
	}
	result, err := saveDialog.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen file: %s\n", result)
}
