package main

import "fmt"


// we've added another keyword string in the definition. This means this function returns a string.
func Hello() string {
	return "Hello, world"
}

func main()  {
	fmt.Println(Hello())
}