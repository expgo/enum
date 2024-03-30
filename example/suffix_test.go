package example

//go:generate ag --dev-plugin=github.com/expgo/enum --dev-plugin-dir=../ --file-suffix .enum.gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// SuffixTest @ENUM{some_item}
type SuffixTest string

func TestSuffix(t *testing.T) {
	x := Suffix("")
	assert.Equal(t, "Suffix().Name", x.String())

	assert.Equal(t, Suffix("gen"), SuffixGen)
}

func TestSuffixTest(t *testing.T) {
	x := SuffixTest("")
	assert.Equal(t, "SuffixTest().Name", x.String())

	assert.Equal(t, SuffixTest("some_item"), SuffixTestSomeItem)
}
