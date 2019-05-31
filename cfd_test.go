// +build windows

package cfd

import (
	"testing"
)

func TestOpen(t *testing.T) {
	Initialize()
	defer UnInitialize()
	openDialog, err := NewOpenFileDialog(DialogConfig{
		Title:      "Test Open",
		Role:       "TestOpen",
		FileFilter: "Text Files (*.txt)|*.txt",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := openDialog.Show(); err != nil {
		t.Fatal(err)
	}
	result, err := openDialog.GetResult()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got result: %s", result)
}

func TestPickFolder(t *testing.T) {
	Initialize()
	defer UnInitialize()
	openDialog, err := NewPickFolderDialog(DialogConfig{
		Title: "Test Pick Folder",
		Role:  "TestPickFolder",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := openDialog.Show(); err != nil {
		t.Fatal(err)
	}
	result, err := openDialog.GetResult()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got result: %s", result)
}

func TestSave(t *testing.T) {
	Initialize()
	defer UnInitialize()
	saveDialog, err := NewSaveFileDialog(DialogConfig{
		Title:      "Test Save",
		Role:       "TestSave",
		FileFilter: "Text Files (*.txt)|*.txt",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := saveDialog.Show(); err != nil {
		t.Fatal(err)
	}
	result, err := saveDialog.GetResult()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got result: %s", result)
}
