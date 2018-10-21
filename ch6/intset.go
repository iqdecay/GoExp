package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint64(x%64)
	if word < len(s.words) {
			s.words[word] &= ^(1 << bit)
	}
}
func (s *IntSet) Clear() {
	for i , _ := range s.words {
		s.words[i] &= 0
	}
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint64(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) ToString() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Copy() *IntSet {
	t := new(IntSet)
	for _, word := range s.words {
		t.words = append(t.words, word)
	}
	return t


}


func (s *IntSet) Len() int {
	var l int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				l++
			}
		}
	}
	return l
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(4)
	y.Add(8)
	y.Add(16)
	x.UnionWith(&y)
	y = *x.Copy()
	fmt.Println(y.ToString(), x.ToString())
}
