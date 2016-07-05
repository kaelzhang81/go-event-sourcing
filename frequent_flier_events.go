package main

type Status int

const (
	StatusRed    Status = itoa
	StatusSilver Status = itoa
	StatusGold   Status = itoa
)

type FrequentFlierAccountCreated struct {
	AccountId        string
	OpeningMiles     int
	OpeningTierPoint int
}

type StatusMatched struct {
	NewStatus Status
}

type FlightTaken struct {
	MilesAdded     int
	TierPointAdded int
}

type PromotedToGoldStatus struct{}
