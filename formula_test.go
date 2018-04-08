package brewer_test

import (
	"testing"

	"github.com/tombell/brewer"
)

func TestFormulaSHA(t *testing.T) {
	formula := &brewer.Formula{Contents: "Hello World"}

	expected := "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"
	actual := formula.SHA()

	if expected != actual {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}

func TestFormulaUpdateTag(t *testing.T) {
	t.Skip()
}

func TestFormulaUpdateRevision(t *testing.T) {
	t.Skip()
}

func TestFormulaUpdateURL(t *testing.T) {
	t.Skip()
}

func TestFormulaUpdateSHA256(t *testing.T) {
	t.Skip()
}
