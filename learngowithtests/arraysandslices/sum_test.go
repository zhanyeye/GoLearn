package arraysandslices

import (
    "reflect"
    "testing"
)

func TestSum(t *testing.T)  {

    t.Run("collection of any size", func(t *testing.T) {
        // We can initialize an array in two ways:
        // [N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
        // [...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}
        // The slice type syntax is very similar to arrays, you just omit the size when declaring them
        numbers := []int{1, 2, 3}

        got := Sum(numbers)
        want := 6

        if got != want {
            // we are using the %v placeholder which is the "default" format, which works well for arrays.
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

}

func TestSumAll(t *testing.T) {

    got := SumAll([]int{1, 2}, []int{0, 9})
    want := []int{3, 9}

    // we can use reflect.DeepEqual which is useful for seeing if any two variables are the same.
    // reflect.DeepEqual is not "type safe"
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }

}


// Go's built-in testing toolkit features a coverage tool, which can help identify areas of your code you have not covered
// go test -cover

func TestSumAllTails(t *testing.T)  {

    checkSums := func(t *testing.T, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }

    t.Run("make the sums of tails of", func(t *testing.T) {
        got := SumAllTails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}
        checkSums(t, got, want)
    })

    t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{3, 4, 5})
        want := []int{0, 9}
        checkSums(t, got, want)
    })

}