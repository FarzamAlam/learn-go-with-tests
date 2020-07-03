package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := &Wallet{}
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := Bitcoin(11)
	if got != want {
		t.Errorf("got %s while want %s", got, want)
	}
}
