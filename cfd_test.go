package cfd

import "testing"

func Test(t *testing.T) {
	t.Errorf("%x", newIFileOpenDialog().Show())
}
