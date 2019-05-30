package cfd

import "testing"

func Test(t *testing.T) {
	openDialog, err := newIFileOpenDialog()
	if err != nil {
		t.Fatal(err)
	}
	t.Error(openDialog.Show())
}
