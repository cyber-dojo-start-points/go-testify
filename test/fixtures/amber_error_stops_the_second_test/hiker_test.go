package hiker

import (
    "strconv"
    "testing"
    "github.com/stretchr/testify/assert"
)

// Both tests call the same panicking code, and only the first one is ever
// seen: a panic takes the whole test binary down with it, so the second
// test is never started and its name never reaches the output.
func Test_life_the_universe_and_everything(t *testing.T) {
    assert.Equal(t, 42, answer())
}

func Test_the_answer_is_two_digits_long(t *testing.T) {
    assert.Equal(t, 2, len(strconv.Itoa(answer())))
}
