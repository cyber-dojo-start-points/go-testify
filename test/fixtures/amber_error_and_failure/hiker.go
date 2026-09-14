package hiker

func answer() int {
    return 6 * 9
}

func checksum() int {
    // The learner meant to divide by the number of digits and has not set
    // the divisor yet, so this divides by zero and panics.
    divisor := 0
    return 42 / divisor
}
