package hiker

func answer() int {
    n := 0
    // The learner meant to count up to 42 and never moves n. go test has a
    // timeout of its own, ten minutes by default, which is far longer than
    // the manifest's max_seconds, so the runner is what stops this.
    for n != 42 {
    }
    return n
}
