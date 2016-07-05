package main

import (
	"fmt"
)

// FrequentFlierAccount represents the state of an instance of the frequent flier
// account aggregate. It tracks changes on itself in the form of domain events.
type FrequentFlierAccount struct {
	id              string
	miles           int
	tierPoints      int
	status          Status
	expectedVersion int
	changes         []interface{}
}

func (account FrequentFlierAccount) String() string {
	format := `FrequentFlierAccount: %s
    Miles: %d
    TierPoints: %d
    Status: %s
    (Expected Version: %d)
    (Pending Changes: %d)
`

	return fmt.Sprintf(format, account.id, account.miles, account.tierPoints,
		account.status, account.expectedVersion, len(account.changes))
}

func NewFrequentFlierAccountFromHistory(events []interface{}) *FrequentFlierAccount {
	state := &FrequentFlierAccount{}

	for _, event := range events {
		state.transition(event)
		state.expectedVersion++
	}

	return state
}

func (state *FrequentFlierAccount) transition(event interface{}) {
	switch e := event.(type) {
	case FrequentFlierAccountCreated:
		state.id = e.AccountId
		state.miles = e.OpeningMiles
		state.tierPoints = e.OpeningTierPoints
		state.status = StatusRed
	case StatusMatched:
		state.status = e.NewStatus
	case FlightTaken:
		state.miles += e.MilesAdded
		state.tierPoints += e.TierPointAdded
	case PromotedToGoldStatus:
		state.status = StatusGold
	}
}

func (state *FrequentFlierAccount) trackChange(event interface{}) {
	state.changes = append(state.changes, event)
	state.transition(event)
}

func (state *FrequentFlierAccount) RecordFlightTaken(miles int, tierPoints int) {
	state.trackChange(FlightTaken{MilesAdded: miles, TierPointAdded: tierPoints})

	if state.tierPoints > 20 && state.status != StatusGold {
		state.trackChange(PromotedToGoldStatus{})
	}
}
