package hiker

import (
    "strconv"
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_the_answer_is_three_digits_long(t *testing.T) {
    assert.Equal(t, 3, len(strconv.Itoa(answer())))
}
