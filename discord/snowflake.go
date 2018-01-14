package discord

import "strconv"

// Snowflake is the ID type used by discord
type Snowflake uint64

func (snowflake Snowflake) String() string {
	return strconv.FormatUint(uint64(snowflake), 10)
}
