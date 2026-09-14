package hiker

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Both files are renamed. Go finds tests by the _test.go suffix, so the
// renamed test file keeps that suffix, and the package name is the one
// thing that stays put: it names the directory's package, not the file.
func Test_life_the_universe_and_everything(t *testing.T) {
    assert.Equal(t, 42, fizz_buzz())
}
