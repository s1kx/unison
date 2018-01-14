package snowflake

import "strconv"

// ID is a database identifier used by go.
// See https://snowflake.net
type ID int64

// String returns the decimal representation of the ID.
func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}

func NewID(i int64) ID {
	return ID(i)
}

// ParseID interprets a string with a decimal number.
// Note that in contrast to ParseInt, this function assumes the given string is
// always valid and thus will panic rather than return an error.
func ParseID(v string) ID {
	id, err := ParseInt(v, 10)
	if err != nil {
		panic(err)
	}
	return id
}

func ParseInt(v string, base int) (ID, error) {
	if v == "" {
		return 0, nil
	}

	id, err := strconv.ParseInt(v, base, 64)
	return ID(id), err
}
