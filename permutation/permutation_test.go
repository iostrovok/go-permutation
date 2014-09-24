package permutation

import (
	"testing"
)

func TestPermutation(t *testing.T) {

	s4 := []string{"a", "b", "c", "d"}
	r4 := [][]string{
		{"a", "b", "c", "d"}, {"a", "b", "d", "c"}, {"a", "c", "b", "d"}, {"a", "c", "d", "b"},
		{"a", "d", "b", "c"}, {"a", "d", "c", "b"}, {"b", "a", "c", "d"}, {"b", "a", "d", "c"},
		{"b", "c", "a", "d"}, {"b", "c", "d", "a"}, {"b", "d", "a", "c"}, {"b", "d", "c", "a"},
		{"c", "a", "b", "d"}, {"c", "a", "d", "b"}, {"c", "b", "a", "d"}, {"c", "b", "d", "a"},
		{"c", "d", "a", "b"}, {"c", "d", "b", "a"}, {"d", "a", "b", "c"}, {"d", "a", "c", "b"},
		{"d", "b", "a", "c"}, {"d", "b", "c", "a"}, {"d", "c", "a", "b"}, {"d", "c", "b", "a"},
	}
	out4 := Strings(s4)
	_check_array(t, out4, r4)

	s3 := []string{"a", "b", "c"}
	r3 := [][]string{
		{"a", "b", "c"}, {"a", "c", "b"}, {"b", "a", "c"}, {"b", "c", "a"}, {"c", "a", "b"}, {"c", "b", "a"},
	}
	out3 := Strings(s3)
	_check_array(t, out3, r3)

	s2 := []string{"a", "b"}
	r2 := [][]string{
		{"a", "b"}, {"b", "a"},
	}
	out2 := Strings(s2)
	_check_array(t, out2, r2)

	s1 := []string{"a"}
	r1 := [][]string{
		{"a"},
	}

	out1 := Strings(s1)
	_check_array(t, out1, r1)
}

func _check_array(t *testing.T, r1 [][]string, r2 [][]string) {
	for i1, v1 := range r1 {
		for i2, v2 := range v1 {
			if r2[i1][i2] != v2 {
				t.Fatal("error Permutation")
			}
		}
	}
}
