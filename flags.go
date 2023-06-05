package main

type Flags struct {
	size  bool
	words bool
	lines bool
	chars bool
}

func (f *Flags) IsNoneSet() bool {
	return (!f.chars && !f.lines && !f.size && !f.words)
}

func (f *Flags) SetAllTrue() {
	f.chars = true
	f.words = true
	f.lines = true
	f.size = true
}
