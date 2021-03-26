package composite

import (
	"io/fs"
	"os"
)

type (
	Tree interface {
		Open() (string, error)
	}
	treeFile struct {
		path     string
		fileInfo fs.FileInfo
	}
	treeDir struct {
		path    string
		entries []os.DirEntry
	}
)

func NewTreeFile(path string) (*treeFile, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	fi.Mode().String()
	return &treeFile{
		path:     path,
		fileInfo: fi,
	}, nil
}

func NewTreeDir(path string) (*treeDir, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return &treeDir{
		path:    path,
		entries: dirEntries,
	}, nil
}
