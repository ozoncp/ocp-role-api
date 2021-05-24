package main

import (
	"os"
	"path"
	"testing"
)

func TestOpenAndPrintFileInCycle(t *testing.T) {
	fname := "test_file"
	f, err := os.CreateTemp("", fname)
	if err != nil {
		t.Fatal(err)
	}
	fpath := path.Join(os.TempDir(), fname)

	data := []byte("1 2 3 4 5 6 7 8 9 10")
	err = os.WriteFile(fpath, data, 0666)
	if err != nil {
		t.Fatal(err)
	}

	f.Close()
	defer os.Remove(fpath)

	err = OpenAndPrintFileInCycle(fpath, 3)
	if err != nil {
		t.Error(err)
	}
}
