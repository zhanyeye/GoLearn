package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spaishHelloProfix = "Hola, "
const frenchHelloPrefix = "Binjour, "
const spanish = "Spanish"
const french = "French"


// we've added another keyword string in the definition. This means this function returns a string.
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

// In our function signature we have made a named return value (prefix string).
// This will create a variable called prefix in your function.
// - It will be assigned the "zero" value.
//   - You can return whatever it's set to by just calling return rather than return prefix.
// - This will display in the Go Doc for your function so it can make the intent of your code clearer.
// The function name starts with a lowercase letter. In Go public functions start with a capital letter and private ones
// start with a lowercase. We don't want the internals of our algorithm to be exposed to the world, so we made this function private.
func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spaishHelloProfix:
        prefix = spaishHelloProfix
    default:
        prefix = englishHelloPrefix
    }
    return
}

func main()  {
    fmt.Println(Hello("Chirs", ""))
}