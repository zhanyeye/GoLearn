package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T)  {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Please note that the example function will not be executed if you remove the comment //Output: 6.
// Although the function will be compiled, it won't be executed.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// go test -v
// === RUN   TestAddr
// --- PASS: TestAddr (0.00s)
// === RUN   ExampleAdd
// --- PASS: ExampleAdd (0.00s)
// PASS
// ok      learngowithtests/integers       0.172s
