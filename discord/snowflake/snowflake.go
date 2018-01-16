package snowflake

import "strconv"

// ID is a database identifier used by go.
// See https://snowflake.net
type ID uint64

// String returns the decimal representation of the ID.
func (f ID) String() string {
	return strconv.FormatUint(uint64(f), 10)
}

// NewID creates a new Snowflake ID from a uint64
func NewID(i uint64) ID {
	return ID(i)
}

// ParseID interprets a string with a decimal number.
// Note that in contrast to ParseInt, this function assumes the given string is
// always valid and thus will panic rather than return an error.
func ParseID(v string) ID {
	id, err := ParseUint(v, 10)
	if err != nil {
		panic(err)
	}
	return id
}

// ParseUint converts a string and given base to a Snowflake
func ParseUint(v string, base int) (ID, error) {
	if v == "" {
		return ID(0), nil
	}

	id, err := strconv.ParseUint(v, base, 64)
	return ID(id), err
}
