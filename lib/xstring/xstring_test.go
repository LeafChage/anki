package xstring_test

import (
	"anki/lib/xstring"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExcludeWhiteSpaces(t *testing.T) {
	a := `hi!_ spaced_    tabed_
newlined_`
	assert.Equal(t, xstring.ExcludeWhiteSpaces(a), "hi!_spaced_tabed_newlined_")
}
