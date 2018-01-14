package discord

import "strconv"

// Snowflake is the ID type used by discord
type Snowflake uint64

// SnowflakeToStr converts a snowflake into a string
func SnowflakeToStr(id Snowflake) string {
	return strconv.FormatUint(uint64(id), 10)
}
