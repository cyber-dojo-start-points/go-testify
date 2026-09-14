package hiker

import (
    "strconv"
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_the_answer_is_two_digits_long(t *testing.T) {
    assert.Equal(t, 2, len(strconv.Itoa(answer())
}
