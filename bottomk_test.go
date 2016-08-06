package minhash

import (
	"github.com/dgryski/go-spooky"
	"testing"
)

func TestBottomK(t *testing.T) {

	tests := []struct {
		s1 []string
		s2 []string
	}{
		{
			[]string{"hello", "world", "foo", "baz", "bar", "zomg"},
			[]string{"goodbye", "world", "foo", "qux", "bar", "zomg"},
		},
	}

	for _, tt := range tests {
		m1 := NewBottomK(spooky.Hash64, 4)
		m2 := NewBottomK(spooky.Hash64, 4)

		for _, s := range tt.s1 {
			m1.Push([]byte(s))
		}

		for _, s := range tt.s2 {
			m2.Push([]byte(s))
		}

		t.Log(m1.Similarity(m2))
	}
}

func TestBottomKMerge(t *testing.T) {

	s1 := []string{"hello", "world", "foo", "baz"}
	s2 := []string{"goodbye", "world", "foo", "qux", "bar", "zomg"}

	s1a := []string{"bar", "zomg"}

	m1 := NewBottomK(spooky.Hash64, 4)
	m2 := NewBottomK(spooky.Hash64, 4)

	for _, s := range s1 {
		m1.Push([]byte(s))
	}

	for _, s := range s2 {
		m2.Push([]byte(s))
	}

	t.Log(m1.Similarity(m2))

	m1a := NewBottomK(spooky.Hash64, 4)
	for _, s := range s1a {
		m1a.Push([]byte(s))
	}

	m1.Merge(m1a)

	t.Log(m1.Similarity(m2))
}
