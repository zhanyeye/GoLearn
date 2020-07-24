package pointererror

import "fmt"

type Wallet struct {
	// In Go if a symbol (so variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
	balance int
}

// In Go, when you call a function or a method the arguments are copied.
// When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
// We can fix this with pointers. Pointers let us point to some values and then let us change them
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// The difference is the receiver type is *Wallet rather than Wallet which you can read as "a pointer to a wallet".
func (w *Wallet) Balance() int {
	return w.balance
}

/**
Now you might wonder, why did they pass? We didn't dereference (解引用) the pointer in the function, like so:
func (w *Wallet) Balance() int {
	return (*w).balance
}
and seemingly addressed the object directly. In fact, the code above using (*w) is absolutely valid. However,
the makers of Go deemed this notation cumbersome, so the language permits us to write w.balance,
without explicit dereference. These pointers to structs even have their own name:
struct pointers and they are automatically dereferenced.
*/
