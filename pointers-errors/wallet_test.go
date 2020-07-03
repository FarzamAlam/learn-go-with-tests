package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s while want %s", got, want)
		}
	}

	t.Run("Deposit : ", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw : ", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		_ = wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("OutOfLimit Withdraw : ", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))
		if err == nil {
			t.Errorf("Wanted an Error but didn't get one")
		} else {
			t.Log("Err : ", err)
		}
	})
}
