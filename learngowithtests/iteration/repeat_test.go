package iteration

import (
    "fmt"
    "testing"
)

func TestRepeat(t *testing.T) {
    repeated := Repeat("a", 5)
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}

// When the benchmark code is executed, it runs b.N times and measures how long it takes.
// To run the benchmarks do go test -bench=. (or if you're in Windows Powershell go test -bench=".")
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}

// Please note that the example function will not be executed if you remove the comment //Output: ???
func ExampleRepeat() {
    repeated := Repeat("a", 5)
    fmt.Println(repeated)
    // Output: aaaaa

}