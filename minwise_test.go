package minhash

import (
	"testing"

	metro "github.com/dgryski/go-metro"
	spooky "github.com/dgryski/go-spooky"
)

func mhash(b []byte) uint64 { return metro.Hash64(b, 0) }

func TestMinwise(t *testing.T) {

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
		m1 := NewMinWise(spooky.Hash64, mhash, 10)
		m2 := NewMinWise(spooky.Hash64, mhash, 10)

		for _, s := range tt.s1 {
			m1.Push([]byte(s))
		}

		for _, s := range tt.s2 {
			m2.Push([]byte(s))
		}

		t.Log(m1.Similarity(m2))
	}
}
func TestMinwiseInit(t *testing.T) {

	tests := []struct {
		strings []string
		sigs    []uint64
	}{
		{
			[]string{"hello", "world", "foo", "baz", "bar", "zomg"},
			[]uint64{2080620544968867365, 8563714006720870342, 2321955019530081269, 748419707729262257, 3830358558201948858, 1245459476431551274, 719322203384083329, 12825625163505565688, 168963933333021279, 1984971849485248784},
		},
	}

	for _, tt := range tests {

		m1 := NewMinWiseFromSignatures(spooky.Hash64, mhash, tt.sigs)
		m2 := NewMinWise(spooky.Hash64, mhash, 10)

		for _, s := range tt.strings {
			m2.Push([]byte(s))
		}

		t.Log(m1.Similarity(m2))
	}
}
