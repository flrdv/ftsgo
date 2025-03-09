package main

type Reference = uint32

type FTS struct {
	refcount Reference
	words    map[string][]Reference
}

func New() FTS {
	return FTS{
		words: make(map[string][]Reference),
	}
}

func (f FTS) Add(content string) Reference {
	return 0
}
