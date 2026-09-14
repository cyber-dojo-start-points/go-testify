package hiker

import (
    "strconv"
    "testing"
    "github.com/stretchr/testify/assert"
)

// go test prints nothing about a test that passes, so this file leaves no
// mark of its own in the output. Its sibling case,
// red_extra_test_file_one_failing, names a test from an extra file in the
// output, and that is what shows such a file is picked up.
func Test_the_answer_is_two_digits_long(t *testing.T) {
    assert.Equal(t, 2, len(strconv.Itoa(answer())))
}
