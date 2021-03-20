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
		file fs.FileInfo
	}
	treeDir struct {
		entries []os.DirEntry
	}
)

func NewTreeFile(path string) (*treeFile, error) {
	//os.ReadFile()
	return nil, nil // TODO
}

func NewTreeDir(path string) (*treeDir, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return &treeDir{
		entries: dirEntries,
	}, nil
}
