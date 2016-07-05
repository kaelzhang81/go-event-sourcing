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

func NewFrequentFlierAccountFromHistory(events []interface{}) *FrequentFlierAccount {
	state := &FrequentFlierAccount{}

	for _, event := range events {
		switch e := event.(type) {
		case FrequentFlierAccountCreated:
			state.id = e.AccountId
			state.miles = e.OpeningMiles
			state.tierPoints = e.OpeningTierPoints
			state.status = StatusRed
		case StatusMatched:
			state.status = e.NewStatus
		case FlightTaken:
			state.miles = e.MilesAdded
			state.tierPoints = e.TierPointAdded
		case PromotedToGoldStatus:
			state.status = StatusGold
		}
	}

	return state
}
