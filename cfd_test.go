package cfd

import "testing"

func TestOpen(t *testing.T) {
	openDialog, err := newIFileOpenDialog()
	if err != nil {
		t.Fatal(err)
	}
	t.Error(openDialog.Show())
}

func TestSave(t *testing.T) {
	openDialog, err := newIFileSaveDialog()
	if err != nil {
		t.Fatal(err)
	}
	t.Error(openDialog.Show())
}
