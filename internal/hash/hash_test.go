package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexSha256(t *testing.T) {
	expectedhashabc := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	expectedhashdef := "cb8379ac2098aa165029e3938a51da0bcecfc008fd6795f401178647f96c5b34"
	expectedhashxyz := "3608bca1e44ea6c4d268eb6db02260269892c0b42b86bbf1e77a6fa16c3c9282"

	assert.Equal(t, expectedhashabc, hexSha256("abc"))
	assert.Equal(t, expectedhashdef, hexSha256("def"))
	assert.Equal(t, expectedhashxyz, hexSha256("xyz"))

	// req := httptest.NewRequest(http.MethodGet, "abc", nil)
	// w := httptest.NewRecorder()

}
