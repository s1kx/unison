package discord

import "strconv"

func discordIDStringToUint64(id string) uint64 {
	if id == "" {
		return 0
	}

	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(err)
	}

	return u
}

func uint64ToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}
