package main

import "testing"

// - It needs to be in a file with a name like xxx_test.go
// - The test function must start with the word Test
// - The test function takes one argument only t *testing.T
func TestHello(t *testing.T) {

    // In Go you can declare functions inside other functions and assign them to variables.
    assertCorrectMessage := func(t *testing.T, got, want string) {
        // t.Helper() is needed to tell the test suite that this method is a helper.
        // By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
        t.Helper()
        if got != want {
            // We are calling the Errorf method on our t which will print out a message and fail the test.
            // The f stands for format which allows us to build a string with values inserted into the placeholder values %q.
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris", "")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("say 'Hello, World' when an empty string si supplied", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })

}