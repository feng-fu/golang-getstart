package p

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		wallet.Withdraw(Bitcoin(90))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdrwa insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, InsufficientFundsError.Error())
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want string) {
	if got == nil {
		t.Errorf("wanted an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}