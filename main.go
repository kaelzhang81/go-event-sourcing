package main

import (
	"fmt"
)

func main() {
	history := []interface{}{
		FrequentFlierAccountCreated{AccountId: "1234567", OpeningMiles: 10000, OpeningTierPoints: 0},
		StatusMatched{NewStatus: StatusSilver},
		FlightTaken{MilesAdded: 2525, TierPointAdded: 5},
		FlightTaken{MilesAdded: 2512, TierPointAdded: 5},
		FlightTaken{MilesAdded: 5600, TierPointAdded: 5},
		FlightTaken{MilesAdded: 3000, TierPointAdded: 3},
	}

	aggregate := NewFrequentFlierAccountFromHistory(history)
	fmt.Println("before RecordFlightTaken")
	fmt.Println(aggregate)

	aggregate.RecordFlightTaken(1000, 3)
	fmt.Println("after RecordFlightTaken")
	fmt.Println(aggregate)
}
