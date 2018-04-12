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

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    (buildpath/"src/github.com/tombell/releasekit").install buildpath.children
    cd "src/github.com/tombell/releasekit" do
      commit = Utils.popen_read("git rev-parse --short HEAD").chomp
      ldflags = [
        "-X main.Version=#{version}",
        "-X main.Commit=#{commit}",
      ]
      system "go", "build",
             "-o", bin/"releasekit",
             "-ldflags", ldflags.join(" "),
             "github.com/tombell/releasekit/cmd/releasekit"
      prefix.install_metafiles
    end
  end

  test do
    # TODO
    # system "#{bin}/releasekit"
  end
end
`

	formulaWithoutTag = `
class Lock < Formula
  desc "Command-line app to quickly lock macOS"
  homepage "https://github.com/tombell/lock"
  url "https://github.com/tombell/lock/archive/v1.0.0.tar.gz"
  sha256 "5c8a518829a40193c805ff85f3c799f8755e2f81c7a00b9ab32698c801897a17"
  head "https://github.com/tombell/lock.git"

  def install
    system "clang", "-framework", "login", "-F", "/System/Library/PrivateFrameworks", "--output=lock", "lock.c"
    bin.install "lock"
  end

  test do
    # how to test this? :joy:
  end
end
`
)

func TestFormulaContentsSHA(t *testing.T) {
	tt := []struct {
		name     string
		contents string
		expected string
	}{
		{
			name:     "Empty",
			contents: "",
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "Simple",
			contents: "Hello World",
			expected: "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.contents}

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
		contents string
		expected string
	}{
		{
			name:     "FormulaWithTag",
			contents: formulaWithTag,
			expected: "v0.1.1",
		},
		{
			name:     "FormulaWithoutTag",
			contents: formulaWithoutTag,
			expected: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.contents}

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
		contents string
		expected string
	}{
		{
			name:     "FormulaWithTag",
			contents: formulaWithTag,
			expected: "v2.0.0",
		},
		{
			name:     "FormulaWithoutTag",
			contents: formulaWithoutTag,
			expected: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.contents}
			formula.UpdateTag(tc.expected)

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
		contents string
		expected string
	}{
		{
			name:     "FormulaWithTag",
			contents: formulaWithTag,
			expected: "6c4e8a83b3632c8a5670261c657d8a01c5f0680b",
		},
		{
			name:     "FormulaWithoutTag",
			contents: formulaWithoutTag,
			expected: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.contents}

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
		contents string
		expected string
	}{
		{
			name:     "FormulaWithTag",
			contents: formulaWithTag,
			expected: "b2cd2dc739420df385a3bc996fc5335bb7cdf444",
		},
		{
			name:     "FormulaWithoutTag",
			contents: formulaWithoutTag,
			expected: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.contents}
			formula.UpdateRevision(tc.expected)

			expected := tc.expected
			actual := formula.Revision()

			if expected != actual {
				t.Errorf("expected %s, but got %s", expected, actual)
			}
		})
	}
}

func TestFormulaURL(t *testing.T) {
	tt := []struct {
		name     string
		contents string
		expected string
	}{
		{
			name:     "FormulaWithTag",
			contents: formulaWithTag,
			expected: "https://github.com/tombell/releasekit.git",
		},
		{
			name:     "FormulaWithoutTag",
			contents: formulaWithoutTag,
			expected: "https://github.com/tombell/lock/archive/v1.0.0.tar.gz",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.contents}

			expected := tc.expected
			actual := formula.URL()

			if expected != actual {
				t.Errorf("expected %s, but got %s", expected, actual)
			}
		})
	}
}

func TestFormulaUpdateURL(t *testing.T) {
	tt := []struct {
		name     string
		contents string
		expected string
	}{
		{
			name:     "FormulaWithTag",
			contents: formulaWithTag,
			expected: "https://github.com/releasekit/releasekit.git",
		},
		{
			name:     "FormulaWithoutTag",
			contents: formulaWithoutTag,
			expected: "https://github.com/tombell/lock/archive/v2.0.0.tar.gz",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formula := brewer.Formula{Contents: tc.contents}
			formula.UpdateURL(tc.expected)

			expected := tc.expected
			actual := formula.URL()

			if expected != actual {
				t.Errorf("expected %s, but got %s", expected, actual)
			}
		})
	}
}
