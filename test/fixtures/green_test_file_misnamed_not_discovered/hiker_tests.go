package hiker

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Named hiker_tests.go rather than hiker_test.go, so go test takes it for
// ordinary source: it is compiled along with the rest of the package, and
// the function below is never run as a test. The assertion is one that
// would fail, so a green says it really did not run.
func Test_the_answer_is_three_digits_long(t *testing.T) {
    assert.Equal(t, 3, len("42"))
}
