package composite

import (
	"fmt"
	"os"
	"strings"
)

type (
	// Tree は Fileと Dirを抽象化して同一視するためのインターフェイス
	Tree interface {
		Open() (string, error)
	}
	treeFile struct {
		path string
	}
	treeDir struct {
		path    string
		entries []Tree
	}
)

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

func newFile(path string) (Tree, error) {
	return &treeFile{path: path}, nil
}

func newDir(path string) (Tree, error) {
	des, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var entries []Tree
	for _, de := range des {
		if de.IsDir() {
			dir, err := newDir(path + "/" + de.Name())
			if err != nil {
				return nil, err
			}
			entries = append(entries, dir)
		} else {
			f, err := newFile(path + "/" + de.Name())
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

func (t *treeFile) Open() (string, error) {
	b, err := os.ReadFile(t.path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`"%s"`, string(b)), nil
}

func (t *treeDir) Open() (string, error) {
	var sb strings.Builder
	sb.WriteString("[")
	for _, e := range t.entries {
		s, err := e.Open()
		if err != nil {
			return "", err
		}
		sb.WriteString(s)
	}
	sb.WriteString("]")
	return sb.String(), nil
}
