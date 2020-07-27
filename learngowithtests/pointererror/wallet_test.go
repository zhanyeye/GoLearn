package pointererror

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		// To make Bitcoin you just use the syntax Bitcoin(10).
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError (t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		// We've introduced t.Fatal which will stop the test if it is called.
		// This is because we don't want to make any more assertions on the error returned if there isn't one around.
		// Without this the test would carry on to the next step and panic because of a nil pointer.
		t.Fatal("didn't get an error but wanted one")
	}
	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError (t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

/**
Unchecked errors
Whilst the Go compiler helps you a lot, sometimes there are things you can still miss and error handling can sometimes be tricky.
To find it, run the following in a terminal to install errcheck, one of many linters available for Go.
go get -u github.com/kisielk/errcheck
Then, inside the directory with your code run errcheck .
*/