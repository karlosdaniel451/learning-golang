package main

import "fmt"

type BankAccount struct {
	Id             int
	BalanceInCents int
}

type Transference struct {
	OriginBankAccount      *BankAccount
	DestinationBankAccount *BankAccount
	ValueInCents           int
}

func main() {
	bankAccounts := []BankAccount{
		{Id: 112},
		{Id: 98},
		{Id: 159},
		{Id: 654},
		{Id: 842},
		{Id: 197},
		{Id: 901},
		{Id: 357},
	}

	transferences := []Transference{
		{
			OriginBankAccount:      &bankAccounts[0],
			DestinationBankAccount: &bankAccounts[1],
			ValueInCents:           500_00,
		},
		{
			OriginBankAccount:      &bankAccounts[1],
			DestinationBankAccount: &bankAccounts[0],
			ValueInCents:           100_00,
		},
		{
			OriginBankAccount:      &bankAccounts[2],
			DestinationBankAccount: &bankAccounts[3],
			ValueInCents:           1_000_00,
		},
	}

	applyTransferences(transferences)

	fmt.Println("transferences:")
	for _, bankAccount := range bankAccounts {
		fmt.Printf("%#v\n", bankAccount)
	}
}

// Call `applyTransference()` for each transference of `transferences`.
func applyTransferences(transferences []Transference) {
	for i := range transferences {
		applyTransference(&transferences[i])
	}
}

// Apply a `Transference` updating the account for the origin and destination
// `BankAccount` according to the `ValueInCents`.
func applyTransference(transference *Transference) {
	transference.OriginBankAccount.BalanceInCents -= transference.ValueInCents
	transference.DestinationBankAccount.BalanceInCents += transference.ValueInCents
}
