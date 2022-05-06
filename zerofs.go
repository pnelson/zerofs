// Package zerofs implements a zero io/fs.FS value.
package zerofs

import (
	"errors"
	"io"
	"io/fs"
	"time"
)

const dot = "."

type zeroFS struct{}

// New returns a new zero io/fs.FS value.
func New() fs.FS {
	return zeroFS{}
}

// Open implements the io/fs.FS interface.
func (f zeroFS) Open(name string) (fs.File, error) {
	if name != dot {
		return nil, newErrNotFound(name)
	}
	return &openDir{}, nil
}

// ReadDir implements the io/fs.ReadDirFS interface.
func (f zeroFS) ReadDir(name string) ([]fs.DirEntry, error) {
	if name != dot {
		return nil, newErrNotFound(name)
	}
	return make([]fs.DirEntry, 0), nil
}

// ReadFile implements the io/fs.ReadFileFS interface.
func (f zeroFS) ReadFile(name string) ([]byte, error) {
	if name != dot {
		return nil, newErrNotFound(name)
	}
	return nil, newErrIsDir(name)
}

type openDir struct{}

// The below implement the io/fs.File interface.
func (d *openDir) Stat() (fs.FileInfo, error) { return d, nil }
func (d *openDir) Read(b []byte) (int, error) { return 0, newErrIsDir(dot) }
func (d *openDir) Close() error               { return nil }

// The below implement the io/fs.FileInfo interface.
func (d *openDir) Name() string       { return dot }
func (d *openDir) Size() int64        { return 0 }
func (d *openDir) Mode() fs.FileMode  { return fs.ModeDir | 0555 }
func (d *openDir) ModTime() time.Time { return time.Time{} }
func (d *openDir) IsDir() bool        { return true }
func (d *openDir) Sys() any           { return nil }

// ReadDir implements the io/fs.ReadDirFile interface.
//
// The io/fs.FS documentation states that every
// directory should implement this interface.
func (d *openDir) ReadDir(count int) ([]fs.DirEntry, error) { return nil, io.EOF }

func newErrIsDir(name string) error {
	return &fs.PathError{Op: "read", Path: name, Err: errors.New("is a directory")}
}

func newErrNotFound(name string) error {
	return &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
}
