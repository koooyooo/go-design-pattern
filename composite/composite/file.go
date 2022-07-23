package composite

import "os"

type treeFile struct {
	path string
}

// Open はFile実装では中身を収集する
func (t *treeFile) Open() ([]string, error) {
	b, err := os.ReadFile(t.path)
	if err != nil {
		return nil, err
	}
	return []string{string(b)}, nil
}
