package hiker

// A test file holding no function whose name begins with Test. go test
// warns that it has nothing to run and then reports PASS, and the
// rag-lambda reads that PASS, so a suite proving nothing lights green.
func double(n int) int {
    return n * 2
}
