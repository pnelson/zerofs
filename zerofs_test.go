package zerofs

import (
	"errors"
	"io/fs"
	"testing"
)

var (
	_ fs.FS         = zeroFS{}
	_ fs.ReadDirFS  = zeroFS{}
	_ fs.ReadFileFS = zeroFS{}
)

var (
	_ fs.File        = (*openDir)(nil)
	_ fs.FileInfo    = (*openDir)(nil)
	_ fs.ReadDirFile = (*openDir)(nil)
)

func TestZeroFS(t *testing.T) {
	zero := New()
	_, err := zero.Open("")
	if !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("unexpected error: %v", err)
	}
}
