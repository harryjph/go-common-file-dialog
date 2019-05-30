package cfd

import (
	"testing"
	"time"
)

func TestOpen(t *testing.T) {
	openDialog, err := newIFileOpenDialog()
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		time.Sleep(1 * time.Second)
		if err := openDialog.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	if err := openDialog.Show(); err != nil {
		t.Fatal(err)
	}
}

func TestSave(t *testing.T) {
	saveDialog, err := newIFileSaveDialog()
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		time.Sleep(1 * time.Second)
		if err := saveDialog.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	if err := saveDialog.Show(); err != nil {
		t.Fatal(err)
	}
}
