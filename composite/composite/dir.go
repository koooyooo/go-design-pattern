package composite

type treeDir struct {
	path    string
	entries []Tree
}

// Open はDir実装では配下のメンバに命令を伝搬させる
func (t *treeDir) Open() ([]string, error) {
	var contents []string
	for _, e := range t.entries {
		s, err := e.Open()
		if err != nil {
			return nil, err
		}
		contents = append(contents, s...)
	}
	return contents, nil
}
