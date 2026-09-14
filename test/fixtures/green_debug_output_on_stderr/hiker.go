package hiker

import (
    "fmt"
    "os"
)

func answer() int {
    // The learner is watching when this gets called and has not taken the
    // line out yet. go test gives the test binary one stream for both of
    // its own, so this arrives among the results rather than beside them.
    fmt.Fprintln(os.Stderr, "answer was called")
    return 6 * 7
}
