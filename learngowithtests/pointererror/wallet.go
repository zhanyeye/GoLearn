package pointererror

import (
	"errors"
	"fmt"
)

// Go lets you create new types from existing ones
// The syntax is type MyName OriginalType
type Bitcoin int

/**
By doing this we're making a new type and we can declare methods on them.
This can be very useful when you want to add some domain specific functionality on top of existing types.
Let's implement Stringer on Bitcoin
type Stringer interface {
	String() string
}
This interface is defined in the fmt package and lets you define how your type is printed when used with the %s format string in prints.
*/

// As you can see, the syntax for creating a method on a type alias is the same as it is on a struct
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// In Go if a symbol (so variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
	balance Bitcoin
}

// In Go, when you call a function or a method the arguments are copied.
// When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
// We can fix this with pointers. Pointers let us point to some values and then let us change them
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// The difference is the receiver type is *Wallet rather than Wallet which you can read as "a pointer to a wallet".
func (w *Wallet) Balance() Bitcoin {
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

// The var keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
	    return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
