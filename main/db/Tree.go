package db

type Tree struct {
	Left 		*Tree
	Right		*Tree
	Name 		string
	documents 	[]Document
	Level 		int
}

func NewNode(_name string, _level int, _docs []Document, _left *Tree, _right *Tree) *Tree {
	return &Tree{
		Left:      _left,
		Right:     _right,
		Name:      _name,
		documents: _docs,
		Level:     _level,
	}
}