package composite

import (
	"os"
	"path/filepath"
)

// Tree は Fileと Dirを抽象化して同一視するためのインターフェイス
type Tree interface {
	Open() ([]string, error)
}

// NewTree は特定のパスから Treeを構成する
func NewTree(path string) (Tree, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return newDir(path)
	} else {
		return newFile(path)
	}
}

func newDir(path string) (Tree, error) {
	des, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var entries []Tree
	for _, de := range des {
		if de.IsDir() {
			dir, err := newDir(filepath.Join(path, de.Name()))
			if err != nil {
				return nil, err
			}
			entries = append(entries, dir)
		} else {
			f, err := newFile(filepath.Join(path, de.Name()))
			if err != nil {
				return nil, err
			}
			entries = append(entries, f)
		}
	}
	return &treeDir{
		entries: entries,
	}, nil
}

func newFile(path string) (Tree, error) {
	return &treeFile{path: path}, nil
}
