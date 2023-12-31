package arraysandstrings_test

import (
	"fmt"
	"testing"

	arraysandstrings "github.com/atakanzen/ctci-examples/arrays-and-strings"
	"github.com/stretchr/testify/assert"
)

// 1.1 --------------------------------------------------------------------------------------------------------

func TestExampleOnePointOne(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"asdfghjkl", true},
		{"a", true},
		{"qwertyuiopq", false},
		{"qawsedrftgyhujikolpzxcvbndm", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.input)
		t.Run(testName, func(t *testing.T) {
			actual := arraysandstrings.HasAllUniqueChars(tt.input)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func BenchmarkOnePointOne(b *testing.B) {

	tests := []struct {
		input string
		want  bool
	}{
		{"asdfghjkl", true},
		{"a", true},
		{"qwertyuiopq", false},
		{"qawsedrftgyhujikolpzxcvbndm", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.input)
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arraysandstrings.HasAllUniqueChars(tt.input)
			}
		})
	}
}

// 1.2 --------------------------------------------------------------------------------------------------------

func TestExampleOnePointTwoWithHashmap(t *testing.T) {
	tests := []struct {
		inputOne, inputTwo string
		want               bool
	}{
		{"qwert", "trewq", true},
		{"abba", "baba", true},
		{"dafgq", "fqgad", true},
		{"dog", "god    ", false},
		{"a", "ba", false},
		{"abc", "abb", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s and %s", tt.inputOne, tt.inputTwo)
		t.Run(testName, func(t *testing.T) {
			actual := arraysandstrings.CheckIfPermutation(tt.inputOne, tt.inputTwo)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestExampleOnePointTwoWithRunes(t *testing.T) {
	tests := []struct {
		inputOne, inputTwo string
		want               bool
	}{
		{"qwert", "trewq", true},
		{"abba", "baba", true},
		{"dafgq", "fqgad", true},
		{"dog", "god    ", false},
		{"a", "ba", false},
		{"abc", "abb", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s and %s", tt.inputOne, tt.inputTwo)
		t.Run(testName, func(t *testing.T) {
			actual := arraysandstrings.CheckIfPermutationWithRunes(tt.inputOne, tt.inputTwo)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func FuzzOnePointTwo(f *testing.F) {
	f.Fuzz(func(t *testing.T, inputOne, inputTwo string) {
		arraysandstrings.CheckIfPermutation(inputOne, inputTwo)
	})
}

func BenchmarkOnePointTwoWithHashmap(b *testing.B) {

	tests := []struct {
		inputOne, inputTwo string
		want               bool
	}{
		{"qwert", "trewq", true},
		{"abba", "baba", true},
		{"dafgq", "fqgad", true},
		{"a", "ba", false},
		{"abc", "abb", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s and %s", tt.inputOne, tt.inputTwo)
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arraysandstrings.CheckIfPermutation(tt.inputOne, tt.inputTwo)
			}
		})
	}
}

func BenchmarkOnePointTwoWithRunes(b *testing.B) {

	tests := []struct {
		inputOne, inputTwo string
		want               bool
	}{
		{"qwert", "trewq", true},
		{"abba", "baba", true},
		{"dafgq", "fqgad", true},
		{"a", "ba", false},
		{"abc", "abb", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s and %s", tt.inputOne, tt.inputTwo)
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arraysandstrings.CheckIfPermutationWithRunes(tt.inputOne, tt.inputTwo)
			}
		})
	}
}

// 1.3 --------------------------------------------------------------------------------------------------------

func TestURLify(t *testing.T) {
	tests := []struct {
		input  string
		length int
		want   string
	}{
		{
			"Mr John Smith      ",
			13,
			"Mr%20John%20Smith  ",
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.input)
		t.Run(testName, func(t *testing.T) {
			actual := arraysandstrings.URLify(tt.input, tt.length)
			assert.Equal(t, tt.want, actual)
		})
	}
}

// 1.4 --------------------------------------------------------------------------------------------------------

func TestCheckIfPermutationPalindrome(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  bool
	}{
		{
			desc:  "Tact Coa Should Return True",
			input: "Tact Coa",
			want:  true,
		},
		{
			desc:  "Qeur rique Should Return True",
			input: "Quer rique",
			want:  true,
		},
		{
			desc:  "UTQY YYTQYY Should Return False",
			input: "UTQY YYTQYY",
			want:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := arraysandstrings.CheckIfPalindromePermutation(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}

// 1.5 --------------------------------------------------------------------------------------------------------

func TestIsOneAwayEdited(t *testing.T) {
	testCases := []struct {
		desc               string
		inputOne, inputTwo string
		want               bool
	}{
		{
			desc:     "pale and Pale should return true",
			inputOne: "pale",
			inputTwo: "Pale",
			want:     true,
		},
		{
			desc:     "pale and pales should return true",
			inputOne: "pale",
			inputTwo: "pales",
			want:     true,
		},
		{
			desc:     "pale and pale should return true",
			inputOne: "pale",
			inputTwo: "pale",
			want:     true,
		},
		{
			desc:     "pale and bake should return false",
			inputOne: "pale",
			inputTwo: "bake",
			want:     false,
		},
		{
			desc:     "pale and baked too much should return false",
			inputOne: "pale",
			inputTwo: "baked too much",
			want:     false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := arraysandstrings.IsOneAwayEdited(tC.inputOne, tC.inputTwo)
			assert.Equal(t, tC.want, actual)
		})
	}
}

// 1.6 --------------------------------------------------------------------------------------------------------

func TestStringCompression(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "should return a2B1c5a3",
			input: "aabcccccaaa",
			want:  "a2b1c5a3",
		},
		{
			desc:  "should return a2B1c5A3",
			input: "aaBcccccAAA",
			want:  "a2B1c5A3",
		},
		{
			desc:  "should return q7w2e5r3T1t2y32",
			input: "qqqqqqqwweeeeerrrTttyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
			want:  "q7w2e5r3T1t2y32",
		},
		{
			desc:  "should return kubernetes",
			input: "kubernetes",
			want:  "kubernetes",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := arraysandstrings.StringCompression(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}

func BenchmarkStringCompression(b *testing.B) {

	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "should return a2B1c5a3",
			input: "aabcccccaaa",
			want:  "a2b1c5a3",
		},
		{
			desc:  "should return a2B1c5A3",
			input: "aaBcccccAAA",
			want:  "a2B1c5A3",
		},
		{
			desc:  "should return q7w2e5r3T1t2y32",
			input: "qqqqqqqwweeeeerrrTttyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
			want:  "q7w2e5r3T1t2y32",
		},
		{
			desc:  "should return kubernetes",
			input: "kubernetes",
			want:  "kubernetes",
		},
	}

	for _, tC := range testCases {
		b.Run(tC.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arraysandstrings.StringCompression(tC.input)
			}
		})
	}
}

func BenchmarkStringCompressionWithPreliminaryCheck(b *testing.B) {

	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "should return a2B1c5a3",
			input: "aabcccccaaa",
			want:  "a2b1c5a3",
		},
		{
			desc:  "should return a2B1c5A3",
			input: "aaBcccccAAA",
			want:  "a2B1c5A3",
		},
		{
			desc:  "should return q7w2e5r3T1t2y32",
			input: "qqqqqqqwweeeeerrrTttyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
			want:  "q7w2e5r3T1t2y32",
		},
		{
			desc:  "should return kubernetes",
			input: "kubernetes",
			want:  "kubernetes",
		},
	}

	for _, tC := range testCases {
		b.Run(tC.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arraysandstrings.StringCompressionPreliminaryCheck(tC.input)
			}
		})
	}
}

// 1.7 --------------------------------------------------------------------------------------------------------

func TestRotateImageBy90Degrees(t *testing.T) {
	testCases := []struct {
		desc        string
		input, want [][]int
	}{
		{
			desc: "with 3x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			desc: "with 6x6 matrix",
			input: [][]int{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 11, 12},
				{13, 14, 15, 16, 17, 18},
				{19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 36},
			},
			want: [][]int{
				{31, 25, 19, 13, 7, 1},
				{32, 26, 20, 14, 8, 2},
				{33, 27, 21, 15, 9, 3},
				{34, 28, 22, 16, 10, 4},
				{35, 29, 23, 17, 11, 5},
				{36, 30, 24, 18, 12, 6},
			},
		},
		{
			desc:  "with nil matrix",
			input: [][]int{},
			want:  [][]int{},
		},
		{
			desc: "with 3x2 matrix",
			input: [][]int{
				{1, 2},
				{4, 5},
				{7, 8},
			},
			want: [][]int{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := arraysandstrings.RotateImageBy90Degrees(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}

// 1.8 --------------------------------------------------------------------------------------------------------

func TestZeroMatrix(t *testing.T) {
	testCases := []struct {
		desc        string
		input, want [][]int
	}{
		{
			desc: "with no zero value",
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
			},
			want: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
			},
		},
		{
			desc: "with zero on (1,2)",
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 0, 9, 10},
				{11, 12, 13, 14, 15},
			},
			want: [][]int{
				{1, 2, 0, 4, 5},
				{0, 0, 0, 0, 0},
				{11, 12, 0, 14, 15},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := arraysandstrings.ZeroMatrix(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}

func BenchmarkZeroMatrix(b *testing.B) {

	testCases := []struct {
		desc        string
		input, want [][]int
	}{
		{
			desc: "with no zero value",
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
			},
			want: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
			},
		},
		{
			desc: "with zero on (1,2)",
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 0, 9, 10},
				{11, 12, 13, 14, 15},
			},
			want: [][]int{
				{1, 2, 0, 4, 5},
				{0, 0, 0, 0, 0},
				{11, 12, 0, 14, 15},
			},
		},
	}

	for _, tC := range testCases {
		b.Run(tC.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arraysandstrings.ZeroMatrix(tC.input)
			}
		})
	}
}

// 1.9 --------------------------------------------------------------------------------------------------------

func TestIsStringRotation(t *testing.T) {
	testCases := []struct {
		desc   string
		s1, s2 string
		want   bool
	}{
		{
			desc: "waterbottle and erbottlewat",
			s1:   "erbottlewat",
			s2:   "waterbottle",
			want: true,
		},
		{
			desc: "atakanzengin and nginatakanze",
			s1:   "nginatakanze",
			s2:   "atakanzengin",
			want: true,
		},
		{
			desc: "apple and peach",
			s1:   "peach",
			s2:   "apple",
			want: false,
		},
		{
			desc: "apple and google",
			s1:   "google",
			s2:   "apple",
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := arraysandstrings.IsStringRotation(tC.s1, tC.s2)
			assert.Equal(t, tC.want, actual)
		})
	}
}
