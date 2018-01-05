package state

import (
	"strconv"
)

// Type state type
//type Type uint8

// ToStr convert state.Type into string
func ToStr(state uint8) string {
	return strconv.Itoa(int(state))
}

// IsReserved check if given state is reserved. These cannot be used for custom states.
// Custom states starts at 11 and ends at 255. custom state range [11, 255].
func IsReserved(state uint8) bool {
	return state >= 0 && state <= 10
}

// Different reserved states
const (
	// MissingState is used when no state was specified
	MissingState uint8 = iota // 0
	// Normal is the default bot state
	Normal // 1

	// Pause is used when the bot should not react to events
	Pause // 2

	// Debug should gives more detailed discord feedback
	Debug // 3

	// Silence Respond to nothing
	Silence // 4

	// Reserved5 ..
	Reserved5

	// Reserved6 ..
	Reserved6

	// Reserved7 ..
	Reserved7

	// Reserved8 ..
	Reserved8

	// Reserved9 ..
	Reserved9

	// Reserved10 ..
	Reserved10
)
