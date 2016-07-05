package main

import (
	"fmt"
)

// FrequentFlierAccount represents the state of an instance of the frequent flier
// account aggregate. It tracks changes on itself in the form of domain events.
type FrequentFlierAccount struct {
	id         string
	miles      int
	tierPoints int
	status     Status
}

func (account FrequentFlierAccount) String() string {
	format := `FrequentFlierAccount: %s
    Miles: %d
    TierPoints: %d
    Status: %s`

	return fmt.Sprintf(format, account.id, account.miles, account.tierPoints, account.status)
}
