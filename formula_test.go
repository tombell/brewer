package brewer_test

import (
	"testing"

	"github.com/tombell/brewer"
)

const (
	formulaTagRevision = `
class Releasekit < Formula
  desc "Create GitHub releases from closed issues and pull requests"
  homepage "https://github.com/tombell/releasekit"
  url "https://github.com/tombell/releasekit.git",
    :tag => "v0.1.1",
    :revision => "6c4e8a83b3632c8a5670261c657d8a01c5f0680b"

  depends_on "go" => :build`

	formulaUrlSha = `
class Lock < Formula
  desc "Command-line app to quickly lock macOS"
  homepage "https://github.com/tombell/lock"
  url "https://github.com/tombell/lock/archive/v1.0.0.tar.gz"
  sha256 "5c8a518829a40193c805ff85f3c799f8755e2f81c7a00b9ab32698c801897a17"`
)

func TestFormulaContentsSHA(t *testing.T) {
	formula := &brewer.Formula{Contents: "Hello World"}

	expected := "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"
	actual := formula.ContentsSHA()

	if expected != actual {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}

func TestFormulaTag(t *testing.T) {
	formula := &brewer.Formula{Contents: formulaTagRevision}

	expected := "v0.1.1"
	actual := formula.Tag()

	if expected != actual {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}

func TestFormulaTagNotFound(t *testing.T) {
	formula := &brewer.Formula{Contents: formulaUrlSha}

	expected := ""
	actual := formula.Tag()

	if expected != actual {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}

func TestFormulaUpdateTag(t *testing.T) {
	formula := &brewer.Formula{Contents: formulaTagRevision}
	formula.UpdateTag("v2.0.0")

	expected := "v2.0.0"
	actual := formula.Tag()

	if expected != actual {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}
