package hiker

import (
    "strconv"
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_life_the_universe_and_everything(t *testing.T) {
    assert.Equal(t, 42, answer())
}

func Test_the_answer_is_three_digits_long(t *testing.T) {
    assert.Equal(t, 3, len(strconv.Itoa(answer())))
}

func Test_the_answer_is_the_question(t *testing.T) {
    assert.Equal(t, "6 * 7", strconv.Itoa(answer()))
}
