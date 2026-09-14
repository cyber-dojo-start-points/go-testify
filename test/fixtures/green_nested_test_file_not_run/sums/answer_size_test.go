package hiker

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// A Go package is one directory, and cyber-dojo.sh runs go test on this
// directory only, so a test file in a sub-directory is a different package,
// is never built, and never runs. The assertion is one that would fail, so
// a green says it really did not run rather than that it ran and passed.
func Test_the_answer_is_three_digits_long(t *testing.T) {
    assert.Equal(t, 3, len("42"))
}
