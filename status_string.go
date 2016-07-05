package main

import "fmt"

const _STATUS_NAME = "StatusRedStatusSilverStatusGold"

var _STATUS_INDEX = [...]int8{0, 9, 21, 31}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_STATUS_INDEX)-1) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _STATUS_NAME[_STATUS_INDEX[i]:_STATUS_INDEX[i+1]]
}
