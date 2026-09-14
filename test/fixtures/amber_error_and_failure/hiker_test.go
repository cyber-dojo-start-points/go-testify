package hiker

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_life_the_universe_and_everything(t *testing.T) {
    assert.Equal(t, 42, answer())
}

func Test_the_checksum_of_the_answer(t *testing.T) {
    assert.Equal(t, 0, checksum())
}
