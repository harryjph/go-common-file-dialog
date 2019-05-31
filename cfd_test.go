package cfd

import (
	"testing"
)

func TestOpen(t *testing.T) {
	openDialog, err := NewOpenFileDialog()
	if err != nil {
		t.Fatal(err)
	}
	/*go func() {
		time.Sleep(1 * time.Second)
		if err := openDialog.Close(); err != nil {
			t.Fatal(err)
		}
	}()*/
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
	saveDialog, err := NewSaveFileDialog()
	if err != nil {
		t.Fatal(err)
	}
	/*go func() {
		time.Sleep(1 * time.Second)
		if err := saveDialog.Close(); err != nil {
			t.Fatal(err)
		}
	}()*/
	if err := saveDialog.Show(); err != nil {
		t.Fatal(err)
	}
	result, err := saveDialog.GetResult()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got result: %s", result)
}
