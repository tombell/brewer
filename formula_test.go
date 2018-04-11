package brewer_test

import (
	"testing"

	"github.com/tombell/brewer"
)

const (
	formulaWithTag = `
class Releasekit < Formula
  desc "Create GitHub releases from closed issues and pull requests"
  homepage "https://github.com/tombell/releasekit"
  url "https://github.com/tombell/releasekit.git",
    :tag => "v0.1.1",
    :revision => "6c4e8a83b3632c8a5670261c657d8a01c5f0680b"

  depends_on "go" => :build`

	formulaWithoutTag = `
class Lock < Formula
  desc "Command-line app to quickly lock macOS"
  homepage "https://github.com/tombell/lock"
  url "https://github.com/tombell/lock/archive/v1.0.0.tar.gz"
  sha256 "5c8a518829a40193c805ff85f3c799f8755e2f81c7a00b9ab32698c801897a17"

  depends_on "go" => :build`
)

func TestFormulaContentsSHA(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty",
			input:    "",
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "Simple",
			input:    "Hello World",
			expected: "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.input}

			expected := tc.expected
			actual := formula.ContentsSHA()

			if expected != actual {
				t.Errorf("expected %s, but got %s", expected, actual)
			}
		})
	}
}

func TestFormulaTag(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "FormulaWithTag",
			input:    formulaWithTag,
			expected: "v0.1.1",
		},
		{
			name:     "FormulaWithoutTag",
			input:    formulaWithoutTag,
			expected: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.input}

			expected := tc.expected
			actual := formula.Tag()

			if expected != actual {
				t.Errorf("expected %s, but got %s", expected, actual)
			}
		})
	}
}

func TestFormulaUpdateTag(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		tag      string
		expected string
	}{
		{
			name:     "FormulaWithTag",
			input:    formulaWithTag,
			tag:      "v2.0.0",
			expected: "v2.0.0",
		},
		{
			name:     "FormulaWithoutTag",
			input:    formulaWithoutTag,
			tag:      "v2.0.0",
			expected: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.input}
			formula.UpdateTag(tc.tag)

			expected := tc.expected
			actual := formula.Tag()

			if expected != actual {
				t.Errorf("expected %s, but got %s", expected, actual)
			}
		})
	}
}

func TestFormulaRevision(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "FormulaWithTag",
			input:    formulaWithTag,
			expected: "6c4e8a83b3632c8a5670261c657d8a01c5f0680b",
		},
		{
			name:     "FormulaWithoutTag",
			input:    formulaWithoutTag,
			expected: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.input}

			expected := tc.expected
			actual := formula.Revision()

			if expected != actual {
				t.Errorf("expected %s, but got %s", expected, actual)
			}
		})
	}
}

func TestFormulaUpdateRevision(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		revision string
		expected string
	}{
		{
			name:     "FormulaWithTag",
			input:    formulaWithTag,
			revision: "b2cd2dc739420df385a3bc996fc5335bb7cdf444",
			expected: "b2cd2dc739420df385a3bc996fc5335bb7cdf444",
		},
		{
			name:     "FormulaWithoutTag",
			input:    formulaWithoutTag,
			revision: "",
			expected: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.input}
			formula.UpdateRevision(tc.revision)

			expected := tc.expected
			actual := formula.Revision()

			if expected != actual {
				t.Errorf("expected %s, but got %s", expected, actual)
			}
		})
	}
}
