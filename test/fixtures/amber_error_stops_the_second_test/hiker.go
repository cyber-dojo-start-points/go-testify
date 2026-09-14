package hiker

func answer() int {
    // The learner meant to divide the answer up and has not set the
    // divisor yet, so this divides by zero and panics.
    divisor := 0
    return 42 / divisor
}
